# Western Skies Contracting — Redesign Project
## Stage 1 Design Document: Scaffold & Config Loading

**Purpose:** This document is a complete implementation plan for Stage 1 of the Western Skies Contracting website rebuild. It is intended for a coding agent. Do not deviate from the described structure without flagging it first. Write clean, idiomatic Go. No premature abstraction.

---

## Project Context

This is a brochure website for Western Skies Contracting, a trades contractor based in Hamilton, Montana (roofing, framing, siding, fencing). It is being built as a handcrafted replacement for their current Duda-based site.

**Tech stack:**
- Go (standard library leaning)
- `templ` for HTML templating
- htmx for lightweight interactivity (not needed in Stage 1)
- Tailwind CSS (not needed in Stage 1)
- `github.com/BurntSushi/toml` for config parsing

**Deployment target:** Single binary behind Caddy on a Fedora VPS via Docker Compose. Per the developer's standard infrastructure pattern.

**Stage 1 goal:** The server starts, loads site configuration from a TOML file, and renders a single unstyled home route that confirms config loaded correctly. Nothing more.

---

## Directory Structure

Create the following structure. Do not create files not listed here.

```
western-skies/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   └── config/
│       └── config.go
├── content/
│   └── site.toml
├── templates/
│   └── layout.templ
├── go.mod
└── .gitignore
```

**Notes:**
- `static/` directory is intentionally omitted — added in a later stage
- No Dockerfile yet — added in a later stage
- `templates/` holds all `.templ` files. The generated `_templ.go` files will land alongside them after `templ generate`

---

## go.mod

Module path: `github.com/fireflymt/western-skies`

Required dependencies:
- `github.com/BurntSushi/toml` — latest stable
- `github.com/a-h/templ` — latest stable

Run `go mod tidy` after creating this file.

---

## content/site.toml

This file holds all site-specific content that may change per client. It must be valid TOML. Populate it with real Western Skies data as follows:

```toml
[business]
name        = "Western Skies Contracting"
tagline     = "A Local Building Contractor You Can Trust"
email       = ""  # leave blank — not publicly listed on their current site
address     = "Hamilton, Montana"

  [business.hours]
  weekdays = "Mon – Sat: 7:00 am – 7:00 pm"
  sunday   = "Closed"

[[phones]]
label  = "Office Manager"
name   = "Matt Detweiler"
number = "406-546-1018"

[[phones]]
label  = "Owner"
name   = "John Plocher"
number = "406-381-8391"

[[services]]
slug        = "roofing"
name        = "Roofing Services"
description = "High-quality roofing that protects your home and stands up to Montana weather."

[[services]]
slug        = "framing"
name        = "Framing Services"
description = "Custom frameworks built with accuracy and skill for residential and commercial structures."

[[services]]
slug        = "siding"
name        = "Siding Services"
description = "Long-lasting, weatherproof siding that improves insulation and curb appeal."

[[services]]
slug        = "fencing"
name        = "Fencing Services"
description = "Durable fencing built for security, privacy, and Montana ranch life."
```

---

## internal/config/config.go

Define the following structs to match the TOML shape exactly. Use the `toml` struct tags from `BurntSushi/toml`.

```
Package: config

Structs:

SiteConfig
  Business  BusinessConfig
  Phones    []PhoneEntry
  Services  []ServiceEntry

BusinessConfig
  Name     string
  Tagline  string
  Email    string
  Address  string
  Hours    HoursConfig

HoursConfig
  Weekdays string
  Sunday   string

PhoneEntry
  Label   string
  Name    string
  Number  string

ServiceEntry
  Slug        string
  Name        string
  Description string
```

**Function to implement:**

```
func Load(path string) (*SiteConfig, error)
```

- Opens the file at `path`
- Decodes it into a `SiteConfig` using `toml.NewDecoder`
- Returns a descriptive error if the file cannot be opened or parsed
- Does not panic — let the caller decide how to handle the error

---

## cmd/server/main.go

**Behavior:**
1. Determine config path: check for a `SITE_CONFIG` environment variable first; fall back to `"content/site.toml"`
2. Call `config.Load(path)` — if it returns an error, log the error and `os.Exit(1)` with a clear message
3. Log to stdout: `"Loaded: <business name>"` — this is the Stage 1 smoke test
4. Parse templates using `templ` (see layout.templ section below)
5. Register one route: `GET /` — handler renders `layout.templ` passing the loaded config
6. Listen on `:8080` (hardcoded for now — configurable in a later stage)
7. Log: `"Serving on :8080"`

Use only the standard library for HTTP (`net/http`). No router library yet.

---

## templates/layout.templ

A minimal HTML shell. No CSS classes, no Tailwind, no styling. Just valid semantic HTML to confirm rendering works.

**Component signature:**
```
templ Layout(cfg config.SiteConfig)
```

**Renders:**
```html
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8"/>
  <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
  <title>{cfg.Business.Name}</title>
</head>
<body>
  <h1>{cfg.Business.Name}</h1>
  <p>{cfg.Business.Tagline}</p>
  <p>{cfg.Business.Address}</p>
  <!-- phones -->
  <ul>
    for each phone in cfg.Phones:
      <li>{phone.Name} ({phone.Label}): {phone.Number}</li>
  </ul>
  <!-- services -->
  <ul>
    for each service in cfg.Services:
      <li><strong>{service.Name}</strong> — {service.Description}</li>
  </ul>
</body>
</html>
```

Use proper `templ` syntax for loops and interpolation. Do not use `html/template` directly.

---

## .gitignore

```
# Go
/western-skies  # compiled binary
*.exe

# Templ generated files (regenerated on build)
*_templ.go

# Env / secrets
.env
*.env

# Editor
.DS_Store
.idea/
.vscode/
```

---

## Build & Run Instructions

The coding agent should confirm these commands work before considering Stage 1 complete:

```bash
# 1. Generate templ files
templ generate

# 2. Tidy dependencies
go mod tidy

# 3. Run
go run ./cmd/server

# Expected output:
# Loaded: Western Skies Contracting
# Serving on :8080

# 4. Smoke test
curl http://localhost:8080
# Expected: HTML page containing "Western Skies Contracting"
```

---

## Definition of Done — Stage 1

- [ ] `go run ./cmd/server` starts without errors
- [ ] Stdout shows `Loaded: Western Skies Contracting`
- [ ] `GET /` returns valid HTML containing the business name, tagline, address, both phone numbers, and all four services
- [ ] No hardcoded strings in Go — all content comes from `site.toml` via the config struct
- [ ] `go vet ./...` passes clean
- [ ] No unused imports

---

## What Stage 1 Does NOT Include

Do not implement these — they belong to later stages:

- Tailwind CSS or any styling
- htmx
- Static file serving
- Multiple routes/pages
- Contact form
- Gallery
- Testimonials
- Docker / deployment config
- Any database or file scanning

---

## Stage 2 Preview (for context only — do not implement)

Stage 2 will add Tailwind CSS via CDN (for development speed), a proper nav component, and the core page routes: Home, Services, Contact. The layout established in Stage 1 will be refactored into a named `Base` layout component that all pages extend.
