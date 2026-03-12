package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fireflymt/western-skies/internal/config"
	"github.com/fireflymt/western-skies/templates"
)

func main() {
	path := os.Getenv("SITE_CONFIG")
	if path == "" {
		path = "content/site.toml"
	}

	cfg, err := config.Load(path)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	fmt.Printf("Loaded: %s\n", cfg.Business.Name)

	// Load gallery page
	galleryPage, err := config.LoadGalleryPage("content/gallery.toml")
	if err != nil {
		log.Fatalf("Failed to load gallery page: %v", err)
	}
	fmt.Println("Loaded gallery page")

	// Load service pages
	servicePages := map[string]*config.ServicePageConfig{}
	for _, slug := range []string{"roofing", "framing", "siding", "fencing"} {
		page, err := config.LoadServicePage("content/" + slug + ".toml")
		if err != nil {
			log.Fatalf("Failed to load %s page: %v", slug, err)
		}
		servicePages[slug] = page
		fmt.Printf("Loaded service page: %s\n", slug)
	}

	// Static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Routes
	http.HandleFunc("/roofing-services", func(w http.ResponseWriter, r *http.Request) {
		templates.ServicePage(*cfg, *servicePages["roofing"], "/roofing-services").Render(r.Context(), w)
	})

	http.HandleFunc("/framing-services", func(w http.ResponseWriter, r *http.Request) {
		templates.ServicePage(*cfg, *servicePages["framing"], "/framing-services").Render(r.Context(), w)
	})

	http.HandleFunc("/siding-services", func(w http.ResponseWriter, r *http.Request) {
		templates.ServicePage(*cfg, *servicePages["siding"], "/siding-services").Render(r.Context(), w)
	})

	http.HandleFunc("/fencing-services", func(w http.ResponseWriter, r *http.Request) {
		templates.ServicePage(*cfg, *servicePages["fencing"], "/fencing-services").Render(r.Context(), w)
	})

	http.HandleFunc("/gallery", func(w http.ResponseWriter, r *http.Request) {
		templates.Gallery(*cfg, *galleryPage).Render(r.Context(), w)
	})

	http.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprint(w, "User-agent: *\nAllow: /\n\nSitemap: "+cfg.Business.URL+"/sitemap.xml\n")
	})

	http.HandleFunc("/sitemap.xml", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")
		base := cfg.Business.URL
		fmt.Fprintf(w, `<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  <url><loc>%s/</loc><priority>1.0</priority></url>
  <url><loc>%s/roofing-services</loc><priority>0.8</priority></url>
  <url><loc>%s/framing-services</loc><priority>0.8</priority></url>
  <url><loc>%s/siding-services</loc><priority>0.8</priority></url>
  <url><loc>%s/fencing-services</loc><priority>0.8</priority></url>
  <url><loc>%s/gallery</loc><priority>0.6</priority></url>
</urlset>`, base, base, base, base, base, base)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			w.WriteHeader(http.StatusNotFound)
			templates.NotFound(*cfg).Render(r.Context(), w)
			return
		}
		templates.Home(*cfg).Render(r.Context(), w)
	})

	fmt.Println("Serving on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
