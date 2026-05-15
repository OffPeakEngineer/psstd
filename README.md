# psstd

[![CI](https://github.com/OffPeakEngineer/psstd/actions/workflows/ci.yml/badge.svg)](https://github.com/OffPeakEngineer/psstd/actions/workflows/ci.yml) [![Release](https://github.com/OffPeakEngineer/psstd/actions/workflows/release.yml/badge.svg)](https://github.com/OffPeakEngineer/psstd/actions/workflows/release.yml)

**psstd is a resilient cluster htop.** Run it on a few machines, open any node in a browser, and watch the whole cluster from a server-rendered dashboard. If the node serving your browser gets busy, it can send the next refresh to a quieter peer.

The result is intentionally simple: every node can serve the UI, every node shares fresh load metrics with its peers, and the browser can be hot-potatoed around the cluster without a central coordinator.

## Features

- **Cluster htop view**: CPU, memory, load, freshness, and offline status for every known node.
- **Hot-potato refresh**: a busy node can bake a lower-load peer into the next browser refresh.
- **Peer rebasing**: node links point at that node's own HTTP address, so the browser moves to the selected peer.
- **Zero-config LAN discovery**: nodes discover peers with mDNS.
- **Seed support**: provide explicit peers with `PSSTD_SEEDS` when discovery is not enough.
- **Local-first resilience**: each node keeps enough recent state to keep rendering during peer churn.

## Quick Start

```bash
go build ./
./psstd
```

Open:

```text
http://localhost:8080
```

Run the same binary on additional LAN machines and they should discover each other automatically. If the advertised browser URL needs to differ from the listen address, set `PSSTD_ADVERTISE_HTTP`.

## Configuration

```bash
export PSSTD_HTTP=":9000"                         # HTTP listen address, default :8080
export PSSTD_ADVERTISE_HTTP="http://10.0.1.25:9000" # browser-reachable URL for this node
export PSSTD_GOSSIP=":7947"                       # peer sync listen address, default :7946
export PSSTD_SEEDS="10.0.1.20:7946,10.0.1.21:7946" # explicit peer sync addresses
export PSSTD_DB="./data"                          # local state directory
export PSSTD_WEB="true"                           # set false for sync-only nodes
./psstd
```

## Discovery

| Environment | How nodes find each other |
|---|---|
| LAN / bare metal | mDNS service `_psstd._tcp` |
| Static hosts | `PSSTD_SEEDS` |
| Mixed setup | mDNS discoveries and explicit seeds are merged |
| Single node | Renders solo until peers appear |

## Health Check

```bash
curl http://localhost:8080/healthz
```

## Notes

psstd uses peer-to-peer state sharing and a small local store internally, but those are implementation details for the dashboard. It is not intended to be a general-purpose distributed database or key-value API.

## Requirements

- Go 1.20+
