# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Brochure website for Western Skies Contracting, a trades contractor (roofing, framing, siding, fencing) in Hamilton, Montana. Handcrafted replacement for their current Duda-based site.

## Tech Stack

- **Go** (standard library for HTTP — no router library)
- **templ** for HTML templating (not `html/template`)
- **htmx** for lightweight interactivity
- **Tailwind CSS** for styling
- **TOML** config via `github.com/BurntSushi/toml`
- **Deployment:** Single binary behind Caddy on Fedora VPS via Docker Compose

Module path: `github.com/fireflymt/western-skies`

## Build & Run Commands

```bash
# Generate templ files (must run before build)
templ generate

# Tidy dependencies
go mod tidy

# Run the server
go run ./cmd/server

# Lint
go vet ./...
```

The server listens on `:8080`. Config path is read from `SITE_CONFIG` env var, defaulting to `content/site.toml`.

## Architecture

### Directory Structure

```
cmd/server/main.go       — Entry point: loads config, registers routes, starts HTTP server
internal/config/config.go — TOML config parsing (SiteConfig, BusinessConfig, etc.)
content/site.toml         — All site-specific content (business info, phones, services)
templates/                — All .templ files; generated *_templ.go files land here
```

### Key Design Decisions

- **All content comes from `content/site.toml`** — no hardcoded strings in Go code or templates. Business name, tagline, phones, services, hours are all in the TOML config.
- **Config structs mirror TOML shape exactly** — `SiteConfig` → `BusinessConfig` (with `HoursConfig`), `[]PhoneEntry`, `[]ServiceEntry`. Use `toml` struct tags.
- **`config.Load(path string) (*SiteConfig, error)`** — opens file, decodes with `toml.NewDecoder`, returns descriptive errors. Never panics.
- **Templates use templ syntax** — components receive `config.SiteConfig` as a parameter. Use templ loops and interpolation, not Go's `html/template`.

### Staged Development

The project follows a staged implementation plan (see `western-skies-stage1-plan.md`). Each stage builds on the previous. Do not implement features from later stages unless asked.

## Conventions

- Write clean, idiomatic Go. No premature abstraction.
- No unused imports — `go vet ./...` must pass clean.
- Generated `*_templ.go` files are gitignored — always regenerate with `templ generate`.
- Keep `.env` files out of version control.
