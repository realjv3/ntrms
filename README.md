# NTRMS — Ninja Turtle Risk Management System

A terminal UI for penny stock day traders to calculate position sizing and stop-loss using the Turtle Trading risk management methodology.

Enter a ticker, your account equity, and risk parameters — NTRMS fetches the 14-day ATR and live price from Twelve Data, then displays a live-updating dashboard with all your risk management variables.

```
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣠⣤⣤⣤⣤⣄⡀⠀⠀⠀⠀⢠⣤⣄⠀⣀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⣠⣴⠟⠛⠉⠁⠀⠀⠈⠉⠛⠻⣦⣄⠀⢸⡟⠙⣿⡟⣷⡀
⠀⠀⠀⠀⠀⢠⣾⠏⠁⣀⠀⠀⠀⠀⠀⠀⠀⢀⣀⡈⠻⣷⣼⣧⠀⢹⡇⣹⡇
⠀⠀⠀⠀⣰⡿⠟⠛⢛⣛⣛⡿⢶⣶⣶⡶⢿⣛⣛⡛⠛⠿⢿⣿⣷⣿⣣⡿⠁  ███╗   ██╗████████╗██████╗ ███╗   ███╗███████╗
⠀⠀⠀⠀⣿⠁⢀⣼⠟⣯⣝⣻⣦⣤⣤⣾⣟⣫⣭⠻⣷⡄⠈⣿⣨⣿⠋⠀⠀  ████╗  ██║╚══██╔══╝██╔══██╗████╗ ████║██╔════╝
⠀⠀⣠⡾⠻⢷⣬⣛⣿⡿⠟⠋⠁⠀⠀⠈⠉⠛⢿⣿⣋⣵⡾⠛⢿⣅⠀⠀⠀  ██╔██╗ ██║   ██║   ██████╔╝██╔████╔██║███████╗
⠀⣼⠟⠀⠀⠀⠉⠿⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⠿⠁⠀⠀⠀⠻⣧⠀⠀  ██║╚██╗██║   ██║   ██╔══██╗██║╚██╔╝██║╚════██║
⠰⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⠆⠀  ██║ ╚████║   ██║   ██║  ██║██║ ╚═╝ ██║███████║
⠀⢻⣦⠀⠀⠀⠀⠀⢴⣤⣤⣀⣀⠀⠀⣀⣠⣤⡾⢿⡆⠀⠀⠀⠀⣴⡟⠀⠀  ╚═╝  ╚═══╝   ╚═╝   ╚═╝  ╚═╝╚═╝     ╚═╝╚══════╝
⠀⠀⠙⢷⣤⣀⠀⠀⠀⠈⠉⠙⠛⠛⠛⠛⠉⠁⠀⠈⠁⠀⣀⣤⡾⠋⠀⠀⠀
⠀⠀⠀⠀⠈⠛⠷⢶⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⡶⠟⠋⠁⠀⠀⠀⠀⠀  ━━ Ninja Turtle Risk Management System ━━
⠀⠀⠀⠀⠀⠀⠀⠈⠛⢷⣤⣀⡀⠀⠀⢀⣠⣴⡾⠛⠁⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠙⠛⠛⠛⠛⠉⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
```

```
╭────────────────────────────────────────────────────╮
│  ⣿⣿  NTRMS ━━━ AAPL Dashboard                    │
│  Entry Price          $3.4200                      │
│  ATR (14d)            0.2500                       │
│  Account Equity       $10,000.00                   │
│  Risk %               2.0%                         │
│  Stop Multiplier      1.0×                         │
│                                                    │
│  Risk Unit (U)        $200.00                      │
│  Stop Distance        $0.2500                      │
│  Stop Price           $3.1700                      │
│  Position Size        800 shares                   │
│  Total Cost           $2,736.00                    │
│                                                    │
│  r: recalculate · q/ctrl-c: quit · refreshes 30s  │
╰────────────────────────────────────────────────────╯
```

## How It Works

The position sizing formula follows the classic Turtle Trading rules:

| Variable | Formula | Description |
|----------|---------|-------------|
| **Risk Unit (U)** | `Equity × Risk%` | Max dollars you're willing to lose on a trade |
| **Stop Distance** | `Multiplier × ATR` | How far below entry your stop sits |
| **Stop Price** | `Entry - Stop Distance` | The price that triggers your stop-loss |
| **Position Size** | `floor(U / Stop Distance)` | Number of shares to buy |
| **Total Cost** | `Position Size × Entry` | Capital required for the position |

ATR (Average True Range) measures daily volatility — the stop is placed at a multiple of ATR below your entry so it respects the stock's natural movement.

## Prerequisites

- Go 1.25+
- A free [Twelve Data](https://twelvedata.com/) API key (800 requests/day on the free tier)

## Setup

```sh
git clone <repo-url> && cd ntrms
go mod download
echo "TWELVEDATA_API_KEY=your_key_here" > .env
```

## Usage

```sh
go run .
```

### Input Screen

| Key | Action |
|-----|--------|
| `Tab` / `Shift+Tab` | Navigate between fields |
| `Enter` | Fetch data and calculate |
| `Ctrl+C` | Quit |

**Fields:**
- **Ticker** — stock symbol (e.g. `AAPL`, `TSLA`)
- **Account Equity** — your account size in dollars
- **Risk %** — percentage of equity to risk per trade (default: 2)
- **Stop Multiplier** — ATR multiple for stop distance (default: 1)
- **Entry Price** — leave blank to use the live market price

### Results Screen

| Key | Action |
|-----|--------|
| `r` | Return to inputs to recalculate |
| `q` / `Ctrl+C` | Quit |

Price refreshes automatically every 30 seconds.

## Running Tests

```sh
go test ./...
```

## Project Structure

```
ntrms/
  main.go                 entrypoint
  internal/
    calc/                 pure position-sizing math
    api/                  Twelve Data HTTP client
    ui/                   lipgloss styles + view rendering
    model/                Bubbletea state machine
```

## License

MIT
