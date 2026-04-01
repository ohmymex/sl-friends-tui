# sl-friends-tui

Terminal dashboard for monitoring [Second Life](https://secondlife.com) friends, groups, and Linden Dollar balance.

Go rewrite of [sl-friends](https://github.com/Jiab77/sl-friends) (bash). Single binary, cross-platform, no external dependencies.

## Features

- Friends list with online/offline status and optional internal names
- Groups list
- L$ balance display
- Real-time search/filter
- Desktop notifications when watched friends come online
- Configurable auto-refresh interval
- Works on Linux, macOS, and Windows

## Install

### Binary release

Download from [GitHub Releases](https://github.com/ohmymex/sl-friends-tui/releases).

### go install

```
go install github.com/ohmymex/sl-friends-tui/cmd/sl-friends@latest
```

### Build from source

```
git clone https://github.com/ohmymex/sl-friends-tui.git
cd sl-friends-tui
make build
```

## Configuration

Copy the example config and add your session token:

```
cp config.example.yaml config.yaml
```

```yaml
token: "your-session-token-here"
```

**Getting your session token:** Log in to secondlife.com, open browser devtools (F12), go to Application > Cookies, copy the `session-token` value.

Config file locations (checked in order):
- `./config.yaml`
- `~/.config/sl-friends/config.yaml`

All settings can also be set via CLI flags or environment variables (prefix `SLF_`).

## Usage

```
sl-friends                          # run with config file
sl-friends --token "your-token"     # pass token directly
sl-friends --filter offline         # show offline friends
sl-friends --refresh 10s            # custom refresh interval
sl-friends --notify alice.doe       # notify when alice comes online
SLF_TOKEN="token" sl-friends        # token via env var
```

Run `sl-friends --help` for all options.

## Keyboard Shortcuts

```
Tab     Switch focus between panels
j/k     Scroll down/up in focused panel
/       Search friends
Esc     Cancel search
f       Cycle filter (online/offline/all)
r       Force refresh
?       Toggle help
q       Quit
```

## Library

The SL client package can be imported independently:

```go
import "github.com/ohmymex/sl-friends-tui/pkg/sl"

client := sl.NewClient("your-token")
friends, err := client.FetchFriends(context.Background())
groups, err := client.FetchGroups(context.Background())
lindens, err := client.FetchLindens(context.Background())
```

## TODO

- [ ] Groups HTML selector needs to be confirmed with real SL data (currently uses a guessed selector)

## Credits

- Original bash version by [Jiab77](https://github.com/Jiab77/sl-friends)
- Go rewrite by [ohmymex](https://github.com/ohmymex)

## License

[WTFPL](LICENSE)
