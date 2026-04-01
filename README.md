# sl-friends-tui

Terminal dashboard for monitoring [Second Life](https://secondlife.com) friends, groups, and Linden Dollar balance.

Go rewrite of [sl-friends](https://github.com/Jiab77/sl-friends) (bash). Single binary, cross-platform, no external dependencies.

## Features

- Friends list with online/offline status and internal names
- Groups list with member count
- L$ balance display
- Real-time search/filter
- Scrollable panels (j/k or arrow keys)
- Desktop notifications when watched friends come online
- ntfy.sh push notifications to your phone
- Configurable auto-refresh interval
- Debug mode for HTTP request logging
- Demo mode for testing without a token
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
sl-friends --debug 2> debug.log     # log HTTP requests to file
sl-friends --demo                   # run with dummy data (no token needed)
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

## ntfy.sh (Mobile Notifications)

Get push notifications on your phone when friends come online.

1. Install the [ntfy app](https://ntfy.sh) on your phone (Android/iOS)
2. Subscribe to a topic (pick something random, e.g. `sl-friends-a8f3x9`)
3. Add to your `config.yaml`:

```yaml
notify:
  enabled: true
  users: ["alice.doe"]
  ntfy:
    enabled: true
    topic: "sl-friends-a8f3x9"
```

Both desktop and mobile notifications fire simultaneously.

## TODO

- [ ] Confirm groups HTML selector with real SL data
- [ ] Account summary panel (membership plan, status, L$/USD balance, land holdings) from `/my/account/index.php`
- [ ] Profile data from `my.secondlife.com` (interests, SL birthdate, bio) — requires AJAX with CSRF token

## Credits

- Original bash version by [Jiab77](https://github.com/Jiab77/sl-friends)
- Go rewrite by [ohmymex](https://github.com/ohmymex)

## License

[WTFPL](LICENSE)
