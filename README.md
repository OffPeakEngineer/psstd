# 🔗 psstd

[![CI](https://github.com/OffPeakEngineer/psstd/actions/workflows/ci.yml/badge.svg)](https://github.com/OffPeakEngineer/psstd/actions/workflows/ci.yml) [![Release](https://github.com/OffPeakEngineer/psstd/actions/workflows/release.yml/badge.svg)](https://github.com/OffPeakEngineer/psstd/actions/workflows/release.yml)

**Peer State Store That's Decentralized** — A lightweight, zero-config distributed key-value store for LAN clusters and Kubernetes. Each node gossips its state to peers using Memberlist; state is persisted locally in Pebble.

> TL;DR: Drop a binary on multiple machines, they auto-discover via mDNS, and form a cluster. Query `/api/kv/{key}` on any node to read/write shared state.

## ✨ Features

- **Zero-config on LAN**: mDNS auto-discovery — no manual peer registration needed
- **Kubernetes-ready**: Helm chart included; uses headless service DNS
- **Gossip protocol**: CockroachDB Memberlist for efficient state propagation
- **Persistent**: Pebble-backed storage per node; optional in-memory mode
- **Web dashboard**: Real-time stats, peer inventory, and key-value browser
- **Simple API**: REST endpoints for get/set/delete operations
- **High availability**: Survives node churn; solo nodes rejoin when cluster returns

## 📦 Installation

### Kubernetes (Helm)
```bash
helm install psstd ./helm/psstd
```

### Bare Metal / LAN (auto-discover via mDNS)
```bash
# Build the binary
go build ./

# Run with web UI enabled (default)
./psstd

# Or disable the web UI
PSSTD_WEB=false ./psstd

# View the dashboard in your browser
open http://localhost:8080
```

### Custom Configuration
```bash
export PSSTD_DB="/custom/db/path"      # database directory (default: ./data)
export PSSTD_HTTP=":9000"              # HTTP listen address (default: :8080)
export PSSTD_ADVERTISE_HTTP="http://192.168.1.10:9000" # browser-reachable node URL
export PSSTD_GOSSIP=":7947"            # Gossip listen port (default: :7946)
export PSSTD_SEEDS="192.168.1.10:7946" # comma-separated seed peers
export PSSTD_WEB="true"                # enable/disable web UI (default: true)
./psstd
```

## 🌐 Peer Discovery

| Environment | How it works |
|---|---|
| **LAN / bare metal** | mDNS `_psstd._tcp` — zero-config auto-discovery |
| **Kubernetes** | PSSTD_SEEDS headless DNS (configured by Helm) |
| **Hybrid** | mDNS + seeds are merged; duplicates ignored |
| **Single node** | Runs solo, joins cluster when others appear |

## 🚀 Usage

### HTTP API

**Get a key:**
```bash
curl http://localhost:8080/api/kv/mykey
```

**Set a key:**
```bash
curl -X PUT http://localhost:8080/api/kv/mykey -d "myvalue"
```

**Delete a key:**
```bash
curl -X DELETE http://localhost:8080/api/kv/mykey
```

**Health check:**
```bash
curl http://localhost:8080/healthz
```

**Web Dashboard:**
```
http://localhost:8080
```

## 🏗️ Architecture

- **Gossip Layer**: Each node broadcasts state changes to the cluster using CockroachDB Memberlist (tuned for LAN performance)
- **Storage**: Local Pebble database; node can be restarted without losing data
- **Discovery**: mDNS on LAN, explicit seeds in Kubernetes
- **Consensus**: Eventually consistent — no strict consistency guarantees (CRDTs or your app layer should handle conflicts)

## 📋 Requirements

- Go 1.20+
