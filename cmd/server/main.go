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
		templates.ServicePage(*cfg, *servicePages["roofing"]).Render(r.Context(), w)
	})

	http.HandleFunc("/framing-services", func(w http.ResponseWriter, r *http.Request) {
		templates.ServicePage(*cfg, *servicePages["framing"]).Render(r.Context(), w)
	})

	http.HandleFunc("/siding-services", func(w http.ResponseWriter, r *http.Request) {
		templates.ServicePage(*cfg, *servicePages["siding"]).Render(r.Context(), w)
	})

	http.HandleFunc("/fencing-services", func(w http.ResponseWriter, r *http.Request) {
		templates.ServicePage(*cfg, *servicePages["fencing"]).Render(r.Context(), w)
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
