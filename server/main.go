package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Lxy417165709/auto_sjz_site/server/internal/config"
	"github.com/Lxy417165709/auto_sjz_site/server/internal/content"
	"github.com/Lxy417165709/auto_sjz_site/server/internal/httpapi"
)

func main() {
	cfg := config.Load()
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: cfg.LogLevel}))

	store, err := content.Load(cfg.ContentDir)
	if err != nil {
		logger.Error("加载站点内容失败", "err", err, "content_dir", cfg.ContentDir)
		os.Exit(1)
	}

	handler, err := httpapi.New(cfg, store, logger)
	if err != nil {
		logger.Error("初始化 HTTP 服务失败", "err", err, "web_dir", cfg.WebDir)
		os.Exit(1)
	}

	srv := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           handler,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	errCh := make(chan error, 1)
	go func() {
		logger.Info("官网启动", "addr", srv.Addr, "web_dir", cfg.WebDir, "base_url", cfg.SiteBaseURL)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	select {
	case err := <-errCh:
		logger.Error("HTTP 服务异常退出", "err", err)
		os.Exit(1)
	case <-ctx.Done():
		logger.Info("收到退出信号，开始优雅停机")
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Error("优雅停机失败", "err", err)
	}
}
