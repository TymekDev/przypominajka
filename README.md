# przypominajka

przypominajka runs a Telegram bot to send notifications about birthdays,
namedays, and anniversaries.

Note: reminders are written in Polish. If you want to customize them, then
modify [formats.go](formats.go).

## Installation
Run `make` to compile przypominajka.
Run `make install` to install przypominajka and completions to `/usr/local/`.
Clean up with `make clean` and `make uninstall`, respectively.

To override `/usr/local/` PREFIX variable use `make -e PREFIX=/foo/bar/baz/`.

## Usage
```
przypominajka - a Telegram bot for sending event reminders

Description:
  przypominajka reads a YAML file with events and sends reminders about them.
  The reminders are sent out on the day of the event between 08:30 and 09:29
  system time (exact time depends on serve command startup time).

  Reminders are written in Polish. 

Example events.yaml:
  january:
    5:
      - name: "John"
        type: "birthday"
      - name: "Jane"
        surname: "Doe"
        type: "nameday"
  april:
    17:
      - names: ["John", "Jane"]
        surname: "Doe"
        type: "wedding anniversary"

Notes:
  - Name and names are mutually exclusive.
  - Names, if present, must have two elements.
  - Surname is optional.
  - Type has to be one of: birthday, nameday, wedding anniversary.

Usage:
  przypominajka [command]

Available Commands:
  help        Help about any command
  list        List events
  serve       Start Telegram bot

Flags:
      --events string   YAML file defining events (default "events.yaml")
  -h, --help            help for przypominajka
  -v, --version         version for przypominajka

Use "przypominajka [command] --help" for more information about a command.
```

### List
```
List events

Usage:
  przypominajka list [flags]

Flags:
  -h, --help   help for list

Global Flags:
      --events string   YAML file defining events (default "events.yaml")
```

### Serve
```
Start Telegram bot

Usage:
  przypominajka serve [flags]

Flags:
      --chat-id int    Telegram chat ID
  -h, --help           help for serve
      --token string   Telegram bot token

Global Flags:
      --events string   YAML file defining events (default "events.yaml")
```