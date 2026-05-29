# Contributing

pulsed is intentionally lightweight: a local-first cluster view, not a general
metrics platform or control plane.

Prefer changes that keep these properties intact:

- read-only observation by default
- peer-to-peer operation without a central coordinator
- useful behavior with no external service dependencies
- explicit opt-in for features that expose more data or trust more peers
- small implementation slices backed by tests

When adding concepts, keep the public language clear but map each metaphor to a
plain implementation field. For example, "what I see" should eventually map to
an observation category, source path, and timestamp.
