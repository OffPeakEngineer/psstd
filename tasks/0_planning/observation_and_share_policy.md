---
id: task-20260518-selectable-metrics-sensors
title: Design observation categories and share policy
type: feature
priority: normal
effort: trip
creator: codex
owner: ""
created: 2026-05-18
revised: 2026-05-29
---

## Problem

pulsed currently treats metrics mostly as fields to render. That is too narrow
for a mesh where each node may know local facts, sense host conditions, observe
neighbors, hear second-hand reports, and choose what it cares to re-share.

The design needs to separate observation categories from display preferences:

- **What I know**: local identity, config, version, role, and self-declared
  status
- **What I feel**: local host metrics and sensors, such as CPU, memory,
  temperature, battery, UPS, HID, I2C, fan, and power
- **What I see**: direct neighbor observations and connectivity/topology facts
- **What I hear**: second-hand observations received through the mesh
- **What I care to share, and with whom**: local policy for forwarding,
  suppressing, summarizing, or declining observations

Trust and care are bundled in the sending policy but not owned globally. A node
can say what it knows and why it is sharing it. The receiver decides whether to
trust it, store it, display it, or forward it based on its own configuration.

What a node should share depends on its purpose and scope. A bare-metal pulsed
node in a Kubernetes cluster may care about CPU/RAM pressure, disk IO, and
network IO. A pod-level node on that same machine may care mostly about
Kubernetes health from inside the workload. A cluster orchestrator node may care
about workers as child nodes. Availability-zone or region nodes may care about
rollups and reachability, not per-core detail.

This creates a scoped mesh:

- **Internal nodes**: local loopback children that can share high-detail local
  observations cheaply
- **Body nodes**: external-facing representatives for a pod, host, orchestrator,
  zone, or region
- **Boundary policy**: trust/share rules decide which internal observations are
  summarized, forwarded, hidden, or exposed through body nodes

The mesh should remain distributed and expansive, but nodes need a rebuff/backoff
mechanism so they do not keep sending information their neighbors have already
said they do not want.

The same policy should support explicit requests. A node can ask a neighbor for
a tagged resource or observation, such as a theme package. The receiver decides
whether it trusts the requester enough, cares to serve the request, has a fresh
enough copy, or should decline with a rebuff/backoff hint.

Enumeration is part of the protocol, not an afterthought. Nodes should be able
to ask "what do you have?" before asking for a specific payload. That applies to
CPUs, sensors, direct observations, heard reports, historical slices, and even
themes.

## Telemetry taxonomy

Telemetry is already a well-trodden path. pulsed should borrow the familiar
shape instead of inventing sensor categories from scratch:

- **Measurement semantics**: use a Home Assistant-like model for device class,
  state class, unit, precision, availability, and expiry
- **Hardware attachment**: use an ESPHome-like model for buses and platforms,
  such as built-in OS metrics, hwmon/sysfs, I2C, SPI, UART/serial, 1-Wire, HID,
  USB, network, and bespoke adapters
- **Collection adapter**: use Go libraries or small OS adapters to enumerate
  what the supported OS can actually expose, without promising every ESPHome
  device class is available everywhere
- **Payload shape**: keep a stable internal observation schema so native OS
  metrics, hardware sensors, and second-hand mesh observations hang from the
  same structure

The lowest-common-denominator target is not "support every sensor." It is a
small, predictable envelope:

- stable observation ID and source node ID
- optional parent/child node path for nested devices, buses, services, or
  subcomponents
- category: know, feel, see, hear, or share-policy
- device class or measurement kind
- unit, value type, precision, and collection timestamp
- attachment/source kind, such as os, hwmon, i2c, spi, uart, hid, usb, network,
  or mesh
- availability, expiry/staleness, and confidence/signature metadata

## Planning questions

- Which observations are core and always emitted by a node?
- Which observations are optional capabilities discovered per-node?
- What is a node, and which nested things should be modeled as child nodes
  instead of tags, structures, or parameters?
- What is this node's purpose and scope: sensor, pod, host, orchestrator, zone,
  region, or another body boundary?
- How are observations categorized as know, feel, see, hear, or share-policy
  decisions?
- Should display selection be a browser preference, node-local config,
  cluster-wide convention, or a mix?
- How should the UI represent an observation that only some nodes can produce or
  only some peers care to receive?
- Which operating-system sensor APIs are reliable enough on Linux, macOS,
  Windows, and containers?
- Which Go libraries or native command integrations should be considered for
  sensor discovery and reading?
- How often should slow or expensive sensor reads run compared with CPU,
  memory, load, and gossip heartbeats?
- Which Home Assistant device classes and ESPHome sensor/bus categories map
  cleanly to pulsed's first-pass observation schema?
- Which sensor categories are observable through maintained Go libraries on
  Linux, macOS, Windows, and containers, and which should remain future adapter
  work?
- How should observation values be normalized, named, unit-tagged, versioned,
  signed, and source-attributed in the gossip payload?
- What does a node say when it declines or suppresses a category of data?
- What rebuff/backoff signal prevents peers from repeatedly sending unwanted
  observations?
- What does a node send when it is requesting a tag or resource from a peer?
- How does a node mark a resource as queryable, temporal, stale, superseded, or
  not available?
- What are the stable enumeration dimensions for observations and resources:
  when, who, where, what, and how/currentness?
- Which dimensions are indexed eagerly, and which remain scan/filter-only until
  a real use case requires an index?

## Done when

- Existing CPU, memory, load, age, and health metrics are mapped into explicit
  observation categories
- Optional metrics can be hidden, shown, shared, suppressed, or summarized
  without adding server API surface unless the design justifies it
- Sensor capability discovery is defined separately from sensor reading
- Unsupported metrics render as unavailable, not stale/offline
- Observation payloads include stable names, units, values, collection
  timestamps, source identity, observation category, and signature/trust
  metadata where available
- Observation payloads can attach to a node path, allowing hierarchical child
  nodes without turning the mesh into an arbitrary graph
- Observation schema has a documented mapping to familiar telemetry concepts:
  device class, state class, unit, bus/source kind, availability, and expiry
- First implementation identifies the supported OS/library adapters and the
  sensor classes that are intentionally deferred
- Share policy can express what a node will forward, what it will not forward,
  and which peers or trust rings it cares to share with
- Share policy can vary by node purpose/scope, including high-detail internal
  loopback sharing and summarized external/body-node sharing
- Receiver behavior is explicitly local: each node may trust, store, display,
  rebuff, or forward the same observation differently
- Rebuff/backoff behavior is defined so nodes do not talk past a neighbor's
  configured interest
- Request/response behavior is defined for queryable tagged resources, including
  accepted, declined, stale, not found, and try-later outcomes
- Enumeration behavior is defined before payload retrieval, including how a node
  lists available sensors, metrics, structures, themes, and historical slices
- The first implementation scope is small enough for one MR
- Follow-up tickets exist for OS-specific sensor support after research

## Related

- Uses identity, tags, HMAC, and rings of trust from
  `tagging_trust_api_overhaul.md`
- Feeds write, recall, and re-sharing behavior in
  `history_and_navigable_dashboard.md`
- Research references: Home Assistant sensor device classes, ESPHome sensor and
  bus/component docs, Go libraries such as gopsutil, procfs/sysfs readers, and
  periph.io
