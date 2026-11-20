package config

import (
	"log/slog"
	"os"
	"strconv"
	"strings"
	"time"
)

// Config 是官网服务的运行时配置，全部通过环境变量注入，保证容器部署时零配置文件。
type Config struct {
	// Port 监听端口，默认 8080（与 Manager 的 5001 / 5003 不冲突）。
	Port string
	// WebDir 前端构建产物目录，默认 ../web/dist（在 server/ 目录下执行时）。
	WebDir string
	// ContentDir 站点文案目录（JSON），默认 content。
	ContentDir string
	// LeadStorePath 试用/联系表单落盘文件，默认 data/leads.jsonl。
	LeadStorePath string
	// LeadWebhookURL 可选，配置后表单提交会额外转发到该地址。
	LeadWebhookURL string
	// SiteBaseURL 用于生成 sitemap.xml 与 canonical，例如 https://auto-sjz.example.com。
	SiteBaseURL string
	// ManagerURL 官网上的「登录 / 购买」跳转地址，由 Manager 提供，官网自身不接订单。
	ManagerURL string
	LogLevel   slog.Level
	// ShutdownTimeout 优雅停机等待时长。
	ShutdownTimeout time.Duration
}

// Load 读取环境变量并填充默认值。
func Load() Config {
	return Config{
		Port:            env("PORT", "8080"),
		WebDir:          env("WEB_DIR", "../web/dist"),
		ContentDir:      env("CONTENT_DIR", "content"),
		LeadStorePath:   env("LEAD_STORE_PATH", "data/leads.jsonl"),
		LeadWebhookURL:  strings.TrimSpace(os.Getenv("LEAD_WEBHOOK_URL")),
		SiteBaseURL:     strings.TrimRight(env("SITE_BASE_URL", "http://localhost:8080"), "/"),
		ManagerURL:      strings.TrimRight(env("MANAGER_URL", ""), "/"),
		LogLevel:        parseLevel(env("LOG_LEVEL", "info")),
		ShutdownTimeout: 10 * time.Second,
	}
}

func env(key, fallback string) string {
	if v := strings.TrimSpace(os.Getenv(key)); v != "" {
		return v
	}
	return fallback
}

func parseLevel(raw string) slog.Level {
	switch strings.ToLower(raw) {
	case "debug":
		return slog.LevelDebug
	case "warn", "warning":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

// PortNumber 返回端口号，便于日志与测试使用。
func (c Config) PortNumber() int {
	n, err := strconv.Atoi(c.Port)
	if err != nil {
		return 8080
	}
	return n
}
