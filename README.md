# 🧞 Genie - CLI game scanner and launcher

![](genie.png) `genie` is a modern, cross-platform command-line tool for **managing and launching your PC games** — without needing to open the launchers manually. It scans your system for installed games, stores them in a local sqlite database, and lets you launch them directly from the terminal.

> NOTE
>
> Currently work in progress, and not all features might work as expected.

## Features

- Detect installed game launchers
- Index installed games
- Launch any game directly from the CLI
- Optionally bypass launcher GUIs (where supported)
- Local SQLite database for quick lookups
- Managing custom games (e.g. installed from CD)

## Installation

Install via `go install`:
```bash
go install github.com/solaire/genie@latest
```

## Usage

```sh
# See full list of commands
genie --help
genie {command} --help
```

### Platform operations

You can list indexed platforms along with the number of games per platform:
```sh
genie platform list
genie platform ls
```

### Scanner

By default, the scanner will try all supported platforms
```sh
# Scan for installed games
genie scan 
```

You can scan specific platforms (e.g. if you just installed a steam game).
```sh
# Scan for steam and gog games only
genie scan --platform steam,gog
```

### Running games

```sh
# Launch game
genie run {game} 
```

### Game operations

```sh
# List indexed games
genie game list
genie game ls
```

You can add games manually. The platform will be set to 'custom'
Useful for games installed outside of supported launchers (e.g. games from CDs)
```sh
genie game add --name {name} --binary {path/to/bin} --dir {path/to/bin}
```

Games added manually can be removed. 
Does not work for platform-indexed games (they would be re-indexed anyway).
```sh
# Remove custom game
genie game rm {game}
```

Game aliasing is supported. Aliases are unique
```sh
# Update game alias
genie game alias {game} {alias}
```

```sh
# Remove game alias
genie game alias {game}
```

## Database

All game metadata is stored in a local SQLite file at:
- Windows: `%USERPROFILE%\.local\share\genie\games.db`
- Linux:   `~/.local/share/genie/games.db`
- macOS:   `~/Library/Application Support/genie/games.db (planned)`

## Logs

For the purposes of troubleshooting, `genie` will create logs on the device. Logs are never transmitted anywhere, and log files older than 24 hours are deleted.

Logs are stored in the following directories:
- Windows: `%USERPROFILE%\.local\share\genie\logs\...`
- Linux:   `~/.local/share/genie/logs/...`
- macOS:   `~/Library/Application Support/genie/logs... (planned)`