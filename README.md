# BurnerVPN

Disposable VPN servers. Create, use, delete.

## Features

- 🌍 **Multiple Regions** - Deploy VPN servers in any DigitalOcean region
- ⚡ **On-Demand** - Create and destroy servers instantly
- 💰 **Usage-Based** - Pay only for the time you use
- 📱 **QR Code** - Easy mobile configuration
- 🔒 **WireGuard** - Modern, fast, and secure VPN protocol

## Quick Start

### Prerequisites

- Go 1.21+
- DigitalOcean API token
- WireGuard client

### Installation

```bash
git clone https://github.com/spiron09/burnervpn.git
cd burnervpn
go build -o burnervpn
```

### Configuration

Create a `.env` file:

```env
DO_API_TOKEN=your_digitalocean_api_token
BURNERVPN_API_URL=http://localhost:8080
```

### Start the Server

```bash
go run server/main.go
```

## CLI Usage

```bash
# List available regions
burnervpn list

# Create a VPN session
burnervpn create <region>

# Check session usage
burnervpn usage <session-id>

# Delete a session
burnervpn delete <session-id>
```

## How It Works

1. **Create** - Provisions a WireGuard droplet in your chosen region
2. **Connect** - Use the generated config with any WireGuard client
3. **Delete** - Destroys the droplet when you're done

## Project Structure

```
burnervpn/
├── cmd/                 # CLI commands
├── internal/client/     # API client
└── server/
    ├── handlers/        # HTTP handlers
    ├── services/        # WireGuard service
    ├── store/           # Session storage
    └── models/          # Data models
```

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/regions` | List available regions |
| POST | `/sessions` | Create a VPN session |
| GET | `/sessions/{id}/usage` | Get session usage |
| DELETE | `/sessions/{id}` | Delete a session |
