---
id: task-20260517-prune-palettes
title: Support pluggable dashboard themes
type: feature
priority: low
effort: trip
creator: codex
owner: ""
created: 2026-05-17
revised: 2026-05-29
---

## Problem

Theme selection is valuable, and the dashboard has room to be more personal
without making the pulsed binary carry every visual idea forever. The current
palette list is bundled directly into the embedded dashboard template, which
makes every built-in theme part of the binary and part of the project
maintenance surface.

The better long-term shape is closer to old web customization: pulsed ships a
small, readable default, then lets operators opt into their own CSS for colors,
scale, density, breakpoints, and local taste.

## Design direction

- Keep a minimal built-in stylesheet so the dashboard works offline and remains
  readable with no configuration
- Treat a theme as a small key/value token set first, with CSS variables as one
  renderer of those tokens
- Expose stable CSS custom properties and class hooks as the browser theme API
- Allow a local CSS file or URL to be injected deliberately, such as through an
  environment variable or query/localStorage setting
- Consider a separate theme repository for curated CSS files that the browser can
  load when an operator opts in
- Do not fetch remote CSS automatically by default; remote themes are code-like
  browser input and should be explicit
- Keep dashboard data and behavior in pulsed; keep decorative palettes and visual
  experiments outside the binary where possible

## Theme token model

A pulsed theme can be expressed as semantic key/value tokens rather than raw CSS:

- `color.ok`, `color.warn`, `color.critical`, `color.info`, `color.accent`
- `color.bg`, `color.surface`, `color.surface_alt`, `color.border`
- `color.text`, `color.muted`
- `scale.font`, `scale.density`, `scale.gap`, `scale.radius`
- state aliases such as `fresh -> color.ok`, `stale -> color.warn`,
  `offline -> color.critical`

Color values should be able to carry RGBA. Scale tokens should stay separate
from colors, even if the transport is compact, so layout changes do not get
confused with palette changes.

The wire/storage format can start as readable JSON or CSS custom properties. A
compact representation such as MessagePack may make sense later if themes are
shared over gossip or stored as part of a broader key/value exchange, but it
should not be required for the first implementation.

## Queryable theme sharing

Themes are a useful first-class example of tagged resource sharing. A node may
advertise that it has a queryable theme package, and another node may request it
by tag or name. The serving node should only send it if local share policy says
the requester is trusted or worth obliging.

They are also a practical enumerability demo. A "theme server" node could simply
enumerate available theme packages and serve them to peers that ask. That is
silly in the right way: small enough to understand, but powerful enough to prove
the tag request, resource enumeration, and share-policy machinery.

A theme package should include enough metadata to be universal and temporal:

- resource kind, such as `theme`
- stable theme ID, human name, version, and optional variant
- tags, such as `dark`, `high-contrast`, `compact`, or `ansi16`
- created/updated timestamps and optional expiry
- token payload or reference URL
- content hash and optional signature metadata
- compatibility bounds for pulsed dashboard/theme API versions

If a node declines, it should be able to say not found, not shared, stale,
unsupported, or try later, plus any rebuff/backoff hint needed to avoid repeated
requests.

## Open questions

- Should custom CSS be configured server-side, browser-side, or both?
- Should remote theme URLs require an allowlist or be accepted as any explicit
  URL?
- Should pulsed serve a local `PULSED_THEME_CSS` file from disk, or should the
  browser load remote stylesheets directly?
- Which CSS variables and layout hooks are stable enough to document as the
  theme contract?
- Is the canonical theme format JSON-like key/value tokens, CSS variables, or
  both?
- If themes are ever shared through pulsed's key/value path, do they need a
  compact binary representation such as MessagePack?
- What makes a theme queryable: a local file, a signed package, a cached remote
  URL, a gossip-advertised resource, or some combination?
- What tag request API should a browser-serving node use to ask a peer for a
  theme, and what response states are needed?
- How does a node enumerate available themes by tag, version, compatibility,
  freshness, source, and trust/signature status before a peer requests one?
- Should the built-in palettes be moved to a separate `pulsed-themes` repo,
  reduced to a tiny set, or left as examples until external themes exist?

## Done when

- Dashboard has a documented CSS variable/hook contract for themes
- Theme tokens have a documented semantic key/value shape for colors, state
  aliases, and scale/density values
- Default built-in styling remains readable with no network access
- Operators can opt into custom CSS without rebuilding pulsed
- Theme packages can be marked queryable by tag/name and requested from another
  node when local share policy allows it
- Theme packages can be enumerated before retrieval, making a simple
  theme-serving node possible without special-case dashboard logic
- Theme package format includes identity, version/time metadata, compatibility,
  hash/signature metadata, and payload/reference fields
- Remote CSS loading, if supported, is explicit and documented with security and
  offline tradeoffs
- Built-in palettes are either reduced to defaults/examples or moved behind the
  same external theme mechanism
- README explains the theme mechanism without turning dashboard styling into a
  core product surface
