package httpapi

import (
	"bytes"
	"encoding/json"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"
)

type leadRequest struct {
	Name     string `json:"name"`
	Company  string `json:"company"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Devices  int    `json:"devices"`
	Message  string `json:"message"`
	Source   string `json:"source"`
	Honeypot string `json:"website"` // 机器人填这个字段，真人留空
}

type leadRecord struct {
	Time    string `json:"time"`
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Devices int    `json:"devices"`
	Message string `json:"message"`
	Source  string `json:"source"`
	IP      string `json:"ip"`
	UA      string `json:"ua"`
}

var emailRe = regexp.MustCompile(`^[^@\s]+@[^@\s]+\.[^@\s]+$`)

// handleLead 接收试用/联系表单。官网不写业务库：落 JSONL 文件，可选转发到 LEAD_WEBHOOK_URL。
func (s *Server) handleLead(w http.ResponseWriter, r *http.Request) {
	ip := clientIP(r)
	if !s.limiter.allow(ip) {
		writeError(w, http.StatusTooManyRequests, "rate_limited", "提交过于频繁，请稍后再试")
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 16<<10)
	var req leadRequest
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid_body", "请求格式不正确")
		return
	}
	if req.Honeypot != "" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	req.Name = strings.TrimSpace(req.Name)
	req.Email = strings.TrimSpace(req.Email)
	req.Phone = strings.TrimSpace(req.Phone)
	req.Message = strings.TrimSpace(req.Message)
	if req.Name == "" {
		writeError(w, http.StatusBadRequest, "missing_name", "请填写称呼")
		return
	}
	if req.Email == "" && req.Phone == "" {
		writeError(w, http.StatusBadRequest, "missing_contact", "请至少留下邮箱或电话")
		return
	}
	if req.Email != "" && !emailRe.MatchString(req.Email) {
		writeError(w, http.StatusBadRequest, "invalid_email", "邮箱格式不正确")
		return
	}
	if req.Devices < 0 || req.Devices > 100000 {
		writeError(w, http.StatusBadRequest, "invalid_devices", "设备数量不合法")
		return
	}
	if len([]rune(req.Message)) > 1000 {
		writeError(w, http.StatusBadRequest, "message_too_long", "留言请控制在 1000 字以内")
		return
	}

	rec := leadRecord{
		Time:    time.Now().UTC().Format(time.RFC3339),
		Name:    req.Name,
		Company: strings.TrimSpace(req.Company),
		Email:   req.Email,
		Phone:   req.Phone,
		Devices: req.Devices,
		Message: req.Message,
		Source:  strings.TrimSpace(req.Source),
		IP:      ip,
		UA:      r.UserAgent(),
	}

	if err := appendJSONL(s.cfg.LeadStorePath, rec); err != nil {
		s.log.Error("写入表单失败", "err", err, "path", s.cfg.LeadStorePath)
		writeError(w, http.StatusInternalServerError, "store_failed", "提交失败，请直接邮件联系我们")
		return
	}

	if s.cfg.LeadWebhookURL != "" {
		go s.forward(rec)
	}

	s.log.Info("收到试用申请", "name", rec.Name, "devices", rec.Devices, "source", rec.Source)
	writeJSON(w, http.StatusOK, map[string]any{"ok": true})
}

func (s *Server) forward(rec leadRecord) {
	body, err := json.Marshal(rec)
	if err != nil {
		s.log.Error("序列化 webhook 失败", "err", err)
		return
	}
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Post(s.cfg.LeadWebhookURL, "application/json", bytes.NewReader(body))
	if err != nil {
		s.log.Warn("webhook 转发失败", "err", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		s.log.Warn("webhook 返回非 2xx", "status", resp.StatusCode)
	}
}

func appendJSONL(path string, rec leadRecord) error {
	if dir := filepath.Dir(path); dir != "" && dir != "." {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return err
		}
	}
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()
	raw, err := json.Marshal(rec)
	if err != nil {
		return err
	}
	_, err = f.Write(append(raw, '\n'))
	return err
}

func clientIP(r *http.Request) string {
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		if i := strings.Index(xff, ","); i > 0 {
			return strings.TrimSpace(xff[:i])
		}
		return strings.TrimSpace(xff)
	}
	if ip, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
		return ip
	}
	return r.RemoteAddr
}

// ipLimiter 是进程内的极简限流器，用于挡住表单被脚本刷爆；多副本部署时由网关统一限流。
type ipLimiter struct {
	mu     sync.Mutex
	hits   map[string][]time.Time
	max    int
	window time.Duration
}

func newIPLimiter(max int, window time.Duration) *ipLimiter {
	return &ipLimiter{hits: make(map[string][]time.Time), max: max, window: window}
}

func (l *ipLimiter) allow(key string) bool {
	now := time.Now()
	l.mu.Lock()
	defer l.mu.Unlock()
	kept := l.hits[key][:0]
	for _, t := range l.hits[key] {
		if now.Sub(t) < l.window {
			kept = append(kept, t)
		}
	}
	if len(kept) >= l.max {
		l.hits[key] = kept
		return false
	}
	l.hits[key] = append(kept, now)
	return true
}
