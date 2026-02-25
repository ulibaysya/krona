# krona

This is a REST backend app for furniture shop

## Usage

```
Usage:
        krona   run [config_path]
        krona   version
        krona   help [command]
```

## Config

Config is required. YAML file is specified via:
1. CLI
2. environment variable `KRONA_CONF`
3. /etc/krona/conf.yaml

You can see config example in [configs/conf.yaml](configs/conf.yaml)

## Warning 

There is a russian language mostly in tests and templates. It is used only for names, so it is ok to rename things in Engish
