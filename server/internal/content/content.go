package content

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// Store 保存站点全部文案。官网不接数据库，文案以 JSON 文件为准，
// 由 Go 服务在启动时加载并通过 /api 暴露，前端只负责渲染。
type Store struct {
	Site    Site
	Pricing Pricing
	FAQ     []FAQItem
	Docs    []Doc
	mu      sync.RWMutex
	loaded  time.Time
}

// Site 站点级信息：品牌、导航、CTA、联系方式、页脚。
type Site struct {
	Name        string      `json:"name"`
	Tagline     string      `json:"tagline"`
	Description string      `json:"description"`
	Nav         []NavItem   `json:"nav"`
	CTA         Link        `json:"cta"`
	Secondary   Link        `json:"secondaryCta"`
	Contact     Contact     `json:"contact"`
	Footer      []FooterCol `json:"footer"`
	ICP         string      `json:"icp"`
	Copyright   string      `json:"copyright"`
}

type NavItem struct {
	Label string `json:"label"`
	Href  string `json:"href"`
}

type Link struct {
	Label string `json:"label"`
	Href  string `json:"href"`
}

type Contact struct {
	Email   string `json:"email"`
	WeChat  string `json:"wechat"`
	Working string `json:"working"`
}

type FooterCol struct {
	Title string `json:"title"`
	Links []Link `json:"links"`
}

// Pricing 描述「按账号实例收费」的规则与示例，价格数值允许为空（未定价时前端显示洽谈）。
type Pricing struct {
	Headline     string        `json:"headline"`
	Summary      string        `json:"summary"`
	UnitLabel    string        `json:"unitLabel"`
	UnitPriceCny *float64      `json:"unitPriceCny"`
	Currency     string        `json:"currency"`
	BillingRules []string      `json:"billingRules"`
	Calculator   Calculator    `json:"calculator"`
	Includes     []FeatureItem `json:"includes"`
	Notes        []string      `json:"notes"`
}

type Calculator struct {
	DefaultDevices         int    `json:"defaultDevices"`
	DefaultAccountsPerUnit int    `json:"defaultAccountsPerDevice"`
	Hint                   string `json:"hint"`
}

type FeatureItem struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

type FAQItem struct {
	ID       string `json:"id"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
	Group    string `json:"group"`
}

type Doc struct {
	Slug     string    `json:"slug"`
	Audience string    `json:"audience,omitempty"`
	Title    string    `json:"title"`
	Summary  string    `json:"summary"`
	Updated  string    `json:"updated"`
	Sections []Section `json:"sections"`
}

type Section struct {
	Heading    string   `json:"heading"`
	Paragraphs []string `json:"paragraphs"`
	Bullets    []string `json:"bullets"`
	Steps      []Step   `json:"steps,omitempty"`
	Figures    []Figure `json:"figures,omitempty"`
}

type Step struct {
	Title string     `json:"title"`
	Items []StepItem `json:"items"`
}

type StepItem struct {
	Text      string `json:"text"`
	Href      string `json:"href,omitempty"`
	LinkLabel string `json:"linkLabel,omitempty"`
}

type Figure struct {
	Src     string `json:"src"`
	Alt     string `json:"alt"`
	Caption string `json:"caption"`
	Width   int    `json:"width"`
	Height  int    `json:"height"`
}

// Load 从目录读取四个 JSON 文件；任一文件缺失或格式错误都会直接失败，避免线上出现空白站。
func Load(dir string) (*Store, error) {
	s := &Store{}
	if err := readJSON(filepath.Join(dir, "site.json"), &s.Site); err != nil {
		return nil, err
	}
	if err := readJSON(filepath.Join(dir, "pricing.json"), &s.Pricing); err != nil {
		return nil, err
	}
	if err := readJSON(filepath.Join(dir, "faq.json"), &s.FAQ); err != nil {
		return nil, err
	}
	if err := readJSON(filepath.Join(dir, "docs.json"), &s.Docs); err != nil {
		return nil, err
	}
	if s.Site.Name == "" {
		return nil, fmt.Errorf("site.json 缺少 name 字段")
	}
	s.loaded = time.Now()
	return s, nil
}

// Reload 支持运维改完文案后热加载，无需重启容器。
func (s *Store) Reload(dir string) error {
	fresh, err := Load(dir)
	if err != nil {
		return err
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Site, s.Pricing, s.FAQ, s.Docs, s.loaded = fresh.Site, fresh.Pricing, fresh.FAQ, fresh.Docs, fresh.loaded
	return nil
}

func (s *Store) Snapshot() (Site, Pricing, []FAQItem, []Doc, time.Time) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.Site, s.Pricing, s.FAQ, s.Docs, s.loaded
}

func readJSON(path string, dst any) error {
	raw, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("读取 %s: %w", path, err)
	}
	if err := json.Unmarshal(raw, dst); err != nil {
		return fmt.Errorf("解析 %s: %w", path, err)
	}
	return nil
}
