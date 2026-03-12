package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type SiteConfig struct {
	Business    BusinessConfig    `toml:"business"`
	SEO         SEOConfig         `toml:"seo"`
	Phones      []PhoneEntry      `toml:"phones"`
	Services    []ServiceEntry    `toml:"services"`
	Testimonials []TestimonialEntry `toml:"testimonials"`
	Features    []FeatureEntry    `toml:"features"`
	ServiceArea ServiceAreaConfig `toml:"service_area"`
}

type BusinessConfig struct {
	Name    string      `toml:"name"`
	Tagline string      `toml:"tagline"`
	Email   string      `toml:"email"`
	Address string      `toml:"address"`
	URL     string      `toml:"url"`
	Hours   HoursConfig `toml:"hours"`
}

type HoursConfig struct {
	Weekdays string `toml:"weekdays"`
	Sunday   string `toml:"sunday"`
}

type SEOConfig struct {
	Title       string `toml:"title"`
	Description string `toml:"description"`
}

type PhoneEntry struct {
	Label  string `toml:"label"`
	Name   string `toml:"name"`
	Number string `toml:"number"`
}

type ServiceEntry struct {
	Slug        string `toml:"slug"`
	Name        string `toml:"name"`
	Headline    string `toml:"headline"`
	Description string `toml:"description"`
	Image       string `toml:"image"`
}

type TestimonialEntry struct {
	Quote  string `toml:"quote"`
	Author string `toml:"author"`
}

type FeatureEntry struct {
	Label string `toml:"label"`
	Copy  string `toml:"copy"`
}

type ServiceAreaConfig struct {
	Description string   `toml:"description"`
	Towns       []string `toml:"towns"`
}

// ServicePageConfig represents a single service page (roofing, framing, etc.)
type ServicePageConfig struct {
	SEO          SEOConfig              `toml:"seo"`
	Header       ServicePageHeader      `toml:"header"`
	Lead         []TextBlock            `toml:"lead"`
	Services     []ServiceDetail        `toml:"services"`
	Materials    MaterialsSection       `toml:"materials"`
	Pricing      TextBlock              `toml:"pricing"`
	Note         NoteSection            `toml:"note"`
	CTA          CTASection             `toml:"cta"`
	Testimonials []TestimonialEntry     `toml:"testimonials"`
	Suppliers    []SupplierLink         `toml:"suppliers"`
}

type ServicePageHeader struct {
	Eyebrow  string `toml:"eyebrow"`
	H1       string `toml:"h1"`
	HeroImage string `toml:"hero_image"`
}

type TextBlock struct {
	Text string `toml:"text"`
	Copy string `toml:"copy"`
}

type ServiceDetail struct {
	Name  string   `toml:"name"`
	Copy  string   `toml:"copy"`
	Image string   `toml:"image"`
	Items []string `toml:"items"`
}

type MaterialsSection struct {
	Intro string         `toml:"intro"`
	Items []ServiceDetail `toml:"items"`
}

type NoteSection struct {
	Heading string `toml:"heading"`
	Copy    string `toml:"copy"`
	Image   string `toml:"image"`
}

type CTASection struct {
	Heading string `toml:"heading"`
	Copy    string `toml:"copy"`
	Button  string `toml:"button"`
}

type SupplierLink struct {
	Name string `toml:"name"`
	URL  string `toml:"url"`
}

func Load(path string) (*SiteConfig, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("opening config %s: %w", path, err)
	}
	defer f.Close()

	var cfg SiteConfig
	if _, err := toml.NewDecoder(f).Decode(&cfg); err != nil {
		return nil, fmt.Errorf("parsing config %s: %w", path, err)
	}

	return &cfg, nil
}

func LoadServicePage(path string) (*ServicePageConfig, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("opening service page config %s: %w", path, err)
	}
	defer f.Close()

	var page ServicePageConfig
	if _, err := toml.NewDecoder(f).Decode(&page); err != nil {
		return nil, fmt.Errorf("parsing service page config %s: %w", path, err)
	}

	return &page, nil
}
