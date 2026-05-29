---
id: task-20260529-history-navigation-ui
title: Capture, recall, re-share, and navigate cluster history
type: feature
priority: normal
effort: epic
owner: ""
created: 2026-05-29
---

## Problem

Operators need temporal visibility into cluster events and state changes to
debug intermittent issues and understand system behavior over time. The same
data also needs to be navigable in the browser and selectively re-shared without
turning pulsed into a metrics platform or a control plane.

This epic owns the data lifecycle after tags, identity, and trust metadata exist:

- **Write**: intentionally capture selected observations and events to local
  history
- **Recall**: query what a node observed at a point in time or across a tagged
  window
- **Re-share**: decide which retained observations are safe and useful to pass
  along based on local care/share policy, tags, and trust metadata
- **Navigate**: make captured history usable through trends, timeline controls,
  node focus, zoom/density, and keyboard navigation

Large-cluster dashboard pain belongs here too: high core counts, dense metric
rows, refresh-driven scroll jumps, and lack of node selection all make current
and historical data harder to inspect.

## Done when

### Write
- [ ] Bounded in-memory ring captures recent local cluster events
- [ ] Optional history streaming writes selected observations to a flat file
  format such as line-delimited JSON or CSV
- [ ] Capture policy can filter by tag, role, trust ring, event type, or TTL
- [ ] Events include timestamp, source node ID, tags, trust metadata, and
  signature status when available
- [ ] Events preserve whether an observation was known locally, sensed locally,
  directly seen, heard second-hand, or re-shared

### Recall
- [ ] Recall API can return observed state at a point in time
- [ ] Recall API can filter by tag, role, trust ring, source, or event type
- [ ] Recall API can distinguish direct local observations from heard or
  re-shared observations
- [ ] Dashboard renders compact recent trends for key metrics where history is
  available
- [ ] Timeline UI can show tagged events and scrub or step through state
  transitions

### Re-share
- [ ] Retained observations have a clear policy for whether they may be
  re-shared through gossip or future APIs
- [ ] Re-sharing respects tags, trust ring, source, signature status, and local
  operator policy
- [ ] Re-sharing honors rebuff/backoff signals from neighbors that do not care
  to receive a category of data
- [ ] Re-shared data remains distinguishable from direct local observations
- [ ] Defaults preserve today's local-only lightweight behavior

### Navigate
- [ ] Node rows/cards support selection and focused inspection without a full
  page navigation
- [ ] Large clusters have a usable zoom/density model for high core counts and
  many nodes
- [ ] Refreshes preserve scroll position, selected node, focus target, and user
  density/zoom preferences
- [ ] Keyboard users can move between nodes, open focused details, and operate
  history controls

## Related

- Replaces the separate planning tickets for historical sparklines,
  keyboard-friendly dashboard navigation, and isolated history playback UI
- Absorbs the previous local history retention and recent event ring planning
- Depends on the tag, identity, HMAC, and rings-of-trust API model in
  `tagging_trust_api_overhaul.md`
- Uses observation categories and share policy from
  `observation_and_share_policy.md`
