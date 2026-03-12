# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Brochure website for Western Skies Contracting, a trades contractor (roofing, framing, siding, fencing) in Hamilton, Montana. Handcrafted replacement for their current Duda-based site. Production URL: `https://www.westernskiescontractingmt.com`

## Tech Stack

- **Go** (standard library for HTTP — no router library)
- **templ** for HTML templating (not `html/template`)
- **htmx** for lightweight interactivity
- **Tailwind CSS** — configured to use brand tokens via `theme.extend`, never arbitrary values
- **TOML** config via `github.com/BurntSushi/toml`
- **Deployment:** Single binary behind Caddy on Fedora VPS via Docker Compose

Module path: `github.com/fireflymt/western-skies`

## Build & Run Commands

```bash
templ generate         # Generate templ files (must run before build)
go mod tidy            # Tidy dependencies
go run ./cmd/server    # Run the server on :8080
go vet ./...           # Lint
```

Config path: `SITE_CONFIG` env var, defaulting to `content/site.toml`.

## Architecture

```
cmd/server/main.go        — Entry point: loads config, registers routes, starts HTTP server
internal/config/config.go  — TOML config parsing (SiteConfig, BusinessConfig, etc.)
content/site.toml          — All site-specific content (business info, phones, services)
templates/                 — All .templ files; generated *_templ.go files land here
static/images/             — Site photos and logo
static/css/tokens.css      — CSS custom properties (brand colors, typography, spacing)
```

### Key Design Decisions

- **All content comes from `content/site.toml`** — no hardcoded strings in Go code or templates.
- **Config structs mirror TOML shape exactly** — `SiteConfig` → `BusinessConfig` (with `HoursConfig`), `[]PhoneEntry`, `[]ServiceEntry`. Use `toml` struct tags.
- **`config.Load(path string) (*SiteConfig, error)`** — decodes with `toml.NewDecoder`, returns descriptive errors, never panics.
- **Templates use templ syntax** — components receive config structs. Use templ loops/interpolation, not Go's `html/template`.

### Staged Development

The project follows a staged implementation plan (`western-skies-stage1-plan.md`). Each stage builds on the previous. Do not implement features from later stages unless asked.

## Planning Documents

- `western-skies-stage1-plan.md` — Stage 1 implementation plan (scaffold, config, single route)
- `western-skies-brand-guide.md` — Complete brand guide with color tokens, typography, component patterns
- `western-skies-copy.md` — Final site copy for all pages with SEO metadata

## Brand & Styling Rules

All brand details are in `western-skies-brand-guide.md`. Key constraints:

- **Use CSS custom property token names** (`--color-granite`, `--color-sage`, etc.) — never raw hex values in component styles
- **Three font families via Google Fonts:** Cormorant Garamond (display headlines), Barlow Condensed (UI/labels/nav/buttons), Barlow (body copy)
- **Never:** rounded corners on buttons/cards, drop shadows on text, centered body copy, all-caps Cormorant Garamond, emoji/icon fonts, barn red/orange/rust tones, gradient hero backgrounds
- **Dark sections** (granite bg): headings → `--color-white`, body → `--color-concrete`, labels → `--color-sunburst`
- **Section eyebrow labels** above every H2: `--font-ui`, `--color-sage`, all-caps, tracked, with `::before` rule
- **Logo** (`Logo+-+Black-1920w.webp`): use `filter: invert(1)` on dark backgrounds; minimum 120px wide
- **Guiding principle:** "When in doubt, add space and remove color."

## Copy & Voice

All page copy is in `western-skies-copy.md`. Voice rules:

- Short sentences (under 20 words). No marketing language ("top-notch," "solutions," "seamless").
- First person plural ("We build"). Earn trust through specificity, not superlatives.
- Faith present but not performative — one scripture reference per page max.
- Montana is context, not decoration — reference it for practical reasons (snow loads, wind).

Pages: Home `/`, Roofing `/roofing-services`, Framing `/framing-services`, Siding `/siding-services`, Fencing `/fencing-services`, Gallery `/gallery`, Testimonials `/testimonials`, Contact `/contact-us`, 404.

## Conventions

- Clean, idiomatic Go. No premature abstraction.
- `go vet ./...` must pass clean — no unused imports.
- Generated `*_templ.go` files are gitignored — always regenerate with `templ generate`.
- Keep `.env` files out of version control.
