# Western Skies Contracting — Brand Guide

**Direction: Granite & Sage**
*For use by the coding agent implementing the site. All decisions here are final unless explicitly overridden by the developer.*

---

## 1. Design Philosophy

The brand draws from the physical environment of the Bitterroot Valley — blue-grey granite peaks, sage meadows, warm limestone outcroppings, and weathered concrete. It is **not** a generic contractor palette. Color is used sparingly. The logo is the loudest visual element; the site defers to it.

**Tone:** Quiet authority. Warm without being rustic. The kind of contractor who doesn't need to shout.

**One rule to carry through every component:** When in doubt, add space and remove color.

---

## 2. Color Tokens

Define these as CSS custom properties on `:root`. Use token names exactly as specified — the coding agent should never use raw hex values in component styles.

```css
:root {
  --color-ink:        #1E2120;  /* Near-black with a green undertone. Primary text, darkest surfaces. */
  --color-granite:    #3D4A52;  /* Blue-grey. Nav background, dark section backgrounds, primary CTA. */
  --color-sage:       #7A8C6E;  /* Muted green. Eyebrow labels, secondary accents. Used sparingly. */
  --color-sunburst:   #C49A3C;  /* Warm gold. Pulled from logo sunburst rays. Used in exactly ONE role per component — CTA hover, accent rules, or logo sub-label. Never competing with granite. */
  --color-limestone:  #F2EDE4;  /* Warm off-white. Primary body/page background. */
  --color-concrete:   #D4CEC6;  /* Warm grey. Borders, dividers, card backgrounds, muted surfaces. */
  --color-white:      #FAFAF8;  /* Near-pure white. Card surfaces on limestone backgrounds. */
  --color-body-text:  #3A3830;  /* Warm dark brown-grey. Body copy — softer than --color-ink. */
  --color-muted-text: #7A7468;  /* Mid-grey-brown. Captions, metadata, secondary labels. */
}
```

### Usage Rules

| Token | Permitted uses | Never use for |
|---|---|---|
| `--color-ink` | Primary headings, nav text, button text on light bg | Body copy (too harsh) |
| `--color-granite` | Nav bg, dark section bg, primary CTA bg | Accent or decorative roles |
| `--color-sage` | Eyebrow labels, section label text | Large filled areas, buttons |
| `--color-sunburst` | Bottom accent rules, logo sub-label, nav CTA bg, one highlight per section | Multiple competing uses in one view |
| `--color-limestone` | Page bg, section bg (warm), form bg | Dark sections |
| `--color-concrete` | Borders, dividers, hr elements, subtle card bg | Text (too low contrast) |
| `--color-white` | Card surfaces, input fields | Page background (too stark) |
| `--color-body-text` | All body copy, list items | Headings |
| `--color-muted-text` | Captions, timestamps, metadata | Body paragraphs |

### Dark Section Rule

When a section uses `--color-granite` as background:

- Headings → `--color-white`
- Body copy → `--color-concrete`
- Labels → `--color-sunburst`
- Borders → `rgba(255,255,255,0.08)`

---

## 3. Typography

### Font Families

```css
:root {
  --font-display:    'Cormorant Garamond', Georgia, serif;
  --font-ui:         'Barlow Condensed', system-ui, sans-serif;
  --font-body:       'Barlow', system-ui, sans-serif;
}
```

**Google Fonts import (place in `<head>`):**

```html
<link rel="preconnect" href="https://fonts.googleapis.com">
<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
<link href="https://fonts.googleapis.com/css2?family=Cormorant+Garamond:ital,wght@0,400;0,600;0,700;1,400;1,600;1,700&family=Barlow:wght@300;400;500&family=Barlow+Condensed:wght@400;600;700&display=swap" rel="stylesheet">
```

### Font Roles

| Font | Role | Why |
|---|---|---|
| Cormorant Garamond | Display headlines (H1, H2) | Calligraphic editorial serif. Echoes the hand-lettered Western style of the logo without copying it. Carries authority without stiffness. |
| Barlow Condensed | UI labels, nav links, section eyebrows, buttons, captions | Tight, utilitarian, legible at small sizes. All-caps with letter-spacing for structure. |
| Barlow | Body copy, paragraph text, testimonials | Clean, neutral, pairs seamlessly with Condensed. Light weight (300) creates breathing room. |

### Type Scale

```css
:root {
  /* Display */
  --text-hero:    clamp(38px, 5vw, 60px);   /* H1 hero headline */
  --text-h2:      clamp(28px, 3.5vw, 44px); /* Section headlines */
  --text-h3:      clamp(20px, 2.5vw, 26px); /* Card/subsection headlines */

  /* UI */
  --text-label:   11px;   /* Section eyebrows, nav — always all-caps + tracked */
  --text-nav:     11px;   /* Nav links */
  --text-btn:     12px;   /* Button text — always all-caps + tracked */
  --text-caption: 10px;   /* Metadata, credits */

  /* Body */
  --text-body:    15px;   /* Standard paragraph */
  --text-body-lg: 17px;   /* Lead paragraph / intro copy */
  --text-small:   13px;   /* Secondary body, card descriptions */
}
```

### Type Style Rules

**Cormorant Garamond headlines:**

- Weight: 700 for H1, 600 for H2/H3
- Italic (`font-style: italic`) is used for the *second line* of a split headline to create contrast — e.g., "Built to last. / *Built for this land.*"
- Line height: 1.05–1.1 for large sizes, 1.15–1.2 for smaller
- Do not use all-caps for display type — the letterforms are designed for mixed case
- Letter spacing: default (0) — Cormorant is already optically spaced

**Barlow Condensed labels:**

- Always `text-transform: uppercase`
- Always `letter-spacing: 0.18em` to `0.26em` depending on size (smaller size = more tracking)
- Weight 700 for primary labels, 600 for secondary/nav
- Pair with a short horizontal rule (`::before` pseudo-element, 22–28px wide) when used as section eyebrows

**Barlow body:**

- Weight 300 for paragraphs
- Weight 400 for testimonials and slightly emphatic text
- Weight 500 for owner names, contact info
- Line height: 1.75 for paragraphs, 1.5 for tighter contexts

---

## 4. Spacing & Layout Tokens

```css
:root {
  --gutter:       clamp(24px, 5vw, 72px);       /* Horizontal page padding */
  --section-y:    clamp(64px, 8vw, 112px);       /* Vertical section padding */
  --card-pad:     clamp(24px, 3vw, 40px);        /* Card internal padding */
  --nav-height:   68px;
}
```

---

## 5. Component Patterns

### Section Eyebrow Label

Used above every H2 section headline. Always `--font-ui`, always `--color-sage`, always all-caps, always tracked.

```html
<div class="section-eyebrow">Who We Are</div>
```

```css
.section-eyebrow {
  font-family: var(--font-ui);
  font-size: var(--text-label);
  font-weight: 700;
  letter-spacing: 0.24em;
  text-transform: uppercase;
  color: var(--color-sage);
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 16px;
}
.section-eyebrow::before {
  content: '';
  display: block;
  width: 24px;
  height: 2px;
  background: var(--color-sage);
  flex-shrink: 0;
}
/* On dark (granite) backgrounds, use sunburst instead */
.dark .section-eyebrow,
.section-eyebrow--light {
  color: var(--color-sunburst);
}
.dark .section-eyebrow::before,
.section-eyebrow--light::before {
  background: var(--color-sunburst);
}
```

### Primary Button

```css
.btn-primary {
  font-family: var(--font-ui);
  font-size: var(--text-btn);
  font-weight: 700;
  letter-spacing: 0.16em;
  text-transform: uppercase;
  background: var(--color-granite);
  color: var(--color-white);
  padding: 13px 28px;
  border: 2px solid var(--color-granite);
  text-decoration: none;
  display: inline-block;
  transition: background 0.2s, border-color 0.2s, color 0.2s;
}
.btn-primary:hover {
  background: var(--color-sunburst);
  border-color: var(--color-sunburst);
  color: var(--color-ink);
}
```

### Ghost Button (on limestone/white backgrounds)

```css
.btn-ghost {
  font-family: var(--font-ui);
  font-size: var(--text-btn);
  font-weight: 600;
  letter-spacing: 0.16em;
  text-transform: uppercase;
  background: transparent;
  color: var(--color-granite);
  padding: 13px 28px;
  border: 2px solid var(--color-concrete);
  text-decoration: none;
  display: inline-block;
  transition: border-color 0.2s, color 0.2s;
}
.btn-ghost:hover {
  border-color: var(--color-granite);
}
```

### Accent Rule

A horizontal line used at the bottom of hero sections and as a section divider. The gradient flows from sunburst into sage, left to right, fading to transparent.

```css
.accent-rule {
  height: 3px;
  background: linear-gradient(
    90deg,
    var(--color-sunburst) 0%,
    var(--color-sage) 55%,
    transparent 100%
  );
}
```

### Nav

- Background: `--color-granite`
- Bottom border: `2px solid var(--color-sunburst)`
- Height: `--nav-height` (68px)
- Links: `--font-ui`, 11px, 600 weight, all-caps, `letter-spacing: 0.15em`, color `rgba(255,255,255,0.65)`
- Link hover: `--color-sunburst`
- CTA button in nav: `--color-sunburst` background, `--color-ink` text

---

## 6. Logo Usage

The logo (`Logo_-_Black-1920w.webp`) is a black-ink mark on transparent background.

| Context | Treatment |
|---|---|
| On limestone/white backgrounds | Use as-is (black ink on light) |
| On granite/dark backgrounds | Apply `filter: invert(1)` in CSS to render white |
| Minimum size | 120px wide |
| Clear space | Equal to the cap-height of the "W" on all sides |
| Never | Recolor, stretch, add drop shadows, or place on busy photo without an overlay |

In the nav, pair the logo mark with a text lockup:

- Line 1: Business name — `--font-display`, 15px, 700, `--color-white`
- Line 2: Sub-label — `--font-ui`, 9px, 400, `letter-spacing: 0.20em`, `--color-sunburst`
  - Content: "Hamilton, Montana · Est. 2018"

---

## 7. Grain Texture

A subtle noise grain is applied to limestone-background sections to add tactile warmth and prevent the off-white from reading as "empty." Apply as a pseudo-element overlay.

```css
.has-grain::before {
  content: '';
  position: absolute;
  inset: 0;
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='200' height='200'%3E%3Cfilter id='n'%3E%3CfeTurbulence type='fractalNoise' baseFrequency='.75' numOctaves='4' stitchTiles='stitch'/%3E%3C/filter%3E%3Crect width='200' height='200' filter='url(%23n)' opacity='.4'/%3E%3C/svg%3E");
  background-size: 200px 200px;
  opacity: 0.035;
  pointer-events: none;
  z-index: 0;
}
```

---

## 8. What This Brand Is Not

These are anti-patterns. Do not introduce them.

- No barn red, orange, or rust tones — that reads as a different brand
- No gradient hero backgrounds simulating sky or dusk — photography does that job
- No rounded corners on buttons or cards — this brand is structural, not friendly-rounded
- No drop shadows on text
- No emoji or icon fonts — use SVG only if icons are needed
- No centered body copy — left-aligned only
- No all-caps Cormorant Garamond — the display font is mixed case always

---

## 9. Templ / Go Implementation Notes

- All CSS tokens go in a `static/css/tokens.css` file, imported globally
- Tailwind, if used, should be configured to reference these tokens via `theme.extend` — do not override with arbitrary Tailwind values
- The `--font-display`, `--font-ui`, `--font-body` variables map directly to Tailwind's `fontFamily` extension if desired
- Google Fonts link tag belongs in the `<head>` of the base layout templ component
- The `.has-grain` pseudo-element pattern requires `position: relative` on the parent — ensure this is set in layout components
