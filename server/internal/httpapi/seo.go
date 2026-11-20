package httpapi

import (
	"fmt"
	"net/http"
	"sort"
	"strings"
)

// handleRobots 固定放行全站，屏蔽 /api，并声明 sitemap 地址。
func (s *Server) handleRobots(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Cache-Control", "public, max-age=3600")
	fmt.Fprintf(w, "User-agent: *\nAllow: /\nDisallow: /api/\n\nSitemap: %s/sitemap.xml\n", s.cfg.SiteBaseURL)
}

// handleSitemap 用导航配置生成 sitemap.xml；deploy、contact 不在导航里，固定补上。
func (s *Server) handleSitemap(w http.ResponseWriter, r *http.Request) {
	site, _, _, _, loaded := s.store.Snapshot()
	lastmod := loaded.UTC().Format("2006-01-02")

	routes := map[string]bool{"/": true, "/deploy": true, "/contact": true, "/console": true}
	for _, item := range site.Nav {
		href := strings.TrimSpace(item.Href)
		if href == "" || strings.HasPrefix(href, "http") || strings.HasPrefix(href, "#") {
			continue
		}
		routes["/"+strings.TrimPrefix(href, "/")] = true
	}

	list := make([]string, 0, len(routes))
	for route := range routes {
		list = append(list, route)
	}
	sort.Slice(list, func(i, j int) bool {
		if list[i] == "/" {
			return true
		}
		if list[j] == "/" {
			return false
		}
		return list[i] < list[j]
	})

	w.Header().Set("Content-Type", "application/xml; charset=utf-8")
	w.Header().Set("Cache-Control", "public, max-age=3600")
	var b strings.Builder
	b.WriteString(`<?xml version="1.0" encoding="UTF-8"?>` + "\n")
	b.WriteString(`<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">` + "\n")
	for _, route := range list {
		priority := "0.6"
		switch route {
		case "/":
			priority = "1.0"
		case "/pricing", "/deploy", "/start", "/console":
			priority = "0.8"
		}
		fmt.Fprintf(&b, "  <url><loc>%s%s</loc><lastmod>%s</lastmod><priority>%s</priority></url>\n",
			s.cfg.SiteBaseURL, route, lastmod, priority)
	}
	b.WriteString("</urlset>\n")
	_, _ = w.Write([]byte(b.String()))
}
