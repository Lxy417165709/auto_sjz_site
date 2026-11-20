package httpapi

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Lxy417165709/auto_sjz_site/server/internal/content"
)

type sitePayload struct {
	Site    content.Site      `json:"site"`
	Pricing content.Pricing   `json:"pricing"`
	FAQ     []content.FAQItem `json:"faq"`
	Docs    []content.Doc     `json:"docs"`
	Meta    siteMeta          `json:"meta"`
}

type siteMeta struct {
	GeneratedAt string `json:"generatedAt"`
	ManagerURL  string `json:"managerUrl"`
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{
		"status": "ok",
		"time":   time.Now().UTC().Format(time.RFC3339),
	})
}

func (s *Server) handleSiteAll(w http.ResponseWriter, r *http.Request) {
	site, pricing, faq, docs, loaded := s.store.Snapshot()
	writeJSON(w, http.StatusOK, sitePayload{
		Site:    site,
		Pricing: pricing,
		FAQ:     faq,
		Docs:    docs,
		Meta: siteMeta{
			GeneratedAt: loaded.UTC().Format(time.RFC3339),
			ManagerURL:  s.cfg.ManagerURL,
		},
	})
}

func (s *Server) handlePricing(w http.ResponseWriter, r *http.Request) {
	_, pricing, _, _, _ := s.store.Snapshot()
	writeJSON(w, http.StatusOK, pricing)
}

func (s *Server) handleFAQ(w http.ResponseWriter, r *http.Request) {
	_, _, faq, _, _ := s.store.Snapshot()
	writeJSON(w, http.StatusOK, faq)
}

func (s *Server) handleDocs(w http.ResponseWriter, r *http.Request) {
	_, _, _, docs, _ := s.store.Snapshot()
	writeJSON(w, http.StatusOK, docs)
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	_ = enc.Encode(payload)
}

func writeError(w http.ResponseWriter, status int, code, message string) {
	writeJSON(w, status, map[string]any{
		"error": map[string]string{"code": code, "message": message},
	})
}
