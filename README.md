# weibo alert
monitor weibo user to send a notice

```
weibo_alert -h

NAME:
   weibo_alert - fight the loneliness!

USAGE:
   weibo_alert [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --config value  configuration file with json (default: "config-sample.json")
   --help, -h      show help (default: false)
   --version, -v   print the version (default: false)
```

## logger-config-sample.json

```json
[
  {
    "enable": true,
    "filter_type": "console",
    "filter_name": "iterm2",
    "min_level": "INFO"
  },
  {
    "enable": true,
    "filter_type": "file",
    "filter_name": "handler",
    "file_name": "logs/handler.log",
    "min_level": "DEBG",
    "max_line": "2",
    "max_size": "20M"
  },
  {
    "enable": true,
    "filter_type": "file",
    "filter_name": "urgent",
    "file_name": "urgent.log",
    "min_level": "EROR",
    "max_size": "20M"
  }
]
```

## config-sample.json

```json
{
"weibo_id": [
  {
    "uid": "3911702320",
    "container_id": ""
  }
],
  "notifier": {
    "url": "https://sc.ftqq.com/",
    "key": ""
  }
}
```