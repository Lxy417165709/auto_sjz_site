package httpapi

import (
	"errors"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/Lxy417165709/auto_sjz_site/server/internal/config"
	"github.com/Lxy417165709/auto_sjz_site/server/internal/content"
)

// Server 聚合官网路由：/api 提供站点数据，其余路径交给前端静态产物（SPA 回退到 index.html）。
type Server struct {
	cfg     config.Config
	store   *content.Store
	log     *slog.Logger
	webDir  string
	limiter *ipLimiter
}

// New 构建 http.Handler，并在启动时校验前端产物是否存在，避免部署出空白站。
func New(cfg config.Config, store *content.Store, log *slog.Logger) (http.Handler, error) {
	absWeb, err := filepath.Abs(cfg.WebDir)
	if err != nil {
		return nil, err
	}
	info, err := os.Stat(absWeb)
	if err != nil {
		return nil, errors.New("前端产物不存在：" + absWeb + "，请先执行前端构建（npm run build）或设置 WEB_DIR")
	}
	if !info.IsDir() {
		return nil, errors.New("WEB_DIR 不是目录：" + absWeb)
	}
	if _, err := os.Stat(filepath.Join(absWeb, "index.html")); err != nil {
		return nil, errors.New("前端产物缺少 index.html：" + absWeb)
	}

	s := &Server{cfg: cfg, store: store, log: log, webDir: absWeb, limiter: newIPLimiter(5, 10*time.Minute)}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/health", s.handleHealth)
	mux.HandleFunc("GET /api/site", s.handleSiteAll)
	mux.HandleFunc("GET /api/site/pricing", s.handlePricing)
	mux.HandleFunc("GET /api/site/faq", s.handleFAQ)
	mux.HandleFunc("GET /api/site/docs", s.handleDocs)
	mux.HandleFunc("POST /api/lead", s.handleLead)
	mux.HandleFunc("GET /robots.txt", s.handleRobots)
	mux.HandleFunc("GET /sitemap.xml", s.handleSitemap)
	mux.HandleFunc("GET /", s.handleStatic)

	return s.recoverer(s.logger(s.secHeaders(mux))), nil
}

func (s *Server) handleStatic(w http.ResponseWriter, r *http.Request) {
	clean := path.Clean("/" + r.URL.Path)
	if strings.Contains(clean, "..") {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	rel := strings.TrimPrefix(clean, "/")
	if rel == "" {
		rel = "index.html"
	}

	full := filepath.Join(s.webDir, filepath.FromSlash(rel))
	info, err := os.Stat(full)
	if err == nil && info.IsDir() {
		// 静态子应用（如 /console-demo/）需要目录 index，且不能用 ServeFile
		// 直接指向 index.html，否则会 301 回目录、再回退到官网首页。
		indexPath := filepath.Join(full, "index.html")
		f, openErr := os.Open(indexPath)
		if openErr == nil {
			defer f.Close()
			fi, statErr := f.Stat()
			if statErr == nil && !fi.IsDir() {
				w.Header().Set("Cache-Control", "public, max-age=300")
				http.ServeContent(w, r, "index.html", fi.ModTime(), f)
				return
			}
		}
	}
	if err == nil && !info.IsDir() {
		if strings.HasPrefix(rel, "assets/") {
			w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
		} else {
			w.Header().Set("Cache-Control", "public, max-age=300")
		}
		http.ServeFile(w, r, full)
		return
	}

	// 前端使用 history 路由，未知的无扩展名路径回退到 index.html。
	if !strings.Contains(path.Base(rel), ".") {
		w.Header().Set("Cache-Control", "no-cache")
		http.ServeFile(w, r, filepath.Join(s.webDir, "index.html"))
		return
	}

	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		s.log.Error("读取静态资源失败", "path", rel, "err", err)
	}
	http.NotFound(w, r)
}

func (s *Server) secHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := w.Header()
		h.Set("X-Content-Type-Options", "nosniff")
		h.Set("X-Frame-Options", "SAMEORIGIN")
		h.Set("Referrer-Policy", "strict-origin-when-cross-origin")
		h.Set("Permissions-Policy", "geolocation=(), microphone=(), camera=()")
		if strings.HasPrefix(r.URL.Path, "/api/") {
			h.Set("Cache-Control", "no-store")
		}
		next.ServeHTTP(w, r)
	})
}

type statusWriter struct {
	http.ResponseWriter
	status int
}

func (w *statusWriter) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

func (s *Server) logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		sw := &statusWriter{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(sw, r)
		if strings.HasPrefix(r.URL.Path, "/assets/") {
			return
		}
		s.log.Info("request",
			"method", r.Method,
			"path", r.URL.Path,
			"status", sw.status,
			"duration_ms", time.Since(start).Milliseconds(),
		)
	})
}

func (s *Server) recoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				s.log.Error("panic", "path", r.URL.Path, "recover", rec)
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
