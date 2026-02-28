# krona

This is a simple backend app for furniture shop

## Features (aka goals)

- HTML pages
- - Header and footer templates on each page: catalog list, redirect buttons to shopping cart, favourites, information, other pages
- - Root page: ads banners with attached redirect link, some reviews, popular products, slogans (/)
- - Furniture catalogs with ability to search by parameters, sort by price, addition date, hits and other factors (/catalogs, /catalogs/{catalog})
- - Products inside catalogs with attached parameters (/catalogs/{catalog}/{product ID})
- - Reviews collected with third party APIs (like maps or review aggregators) and inside app reviews from storage (/reviews)
- - Rooms list with redirect links to specific categories (/rooms/{room}) (e.g. /rooms/bedroom contains links to /catalogs/wardrobes?room=bedroom, /catalogs/mirrors?room=bedroom)
- - Blog pages (/blog, /blog/{article ID})
- REST api
- - Add new categories, products, blogs (POST)
- - Modify (PUT)
- - Receive (GET)
- - Delete (...)
- Storage implementation
- - Relational database (postgresql and sqlite)
- - Support for cache (valkey, memcached)
- Configuration
- - YAML
- - Environment variables
- User management
- - Different authentication methods
- - - Telegram
- - - Bare keys
- - - We should have another one (maybe classic password authentication, oauth2 provider, oidc, jwt), come up with decision later
- - Service should have simple cookies, so user doesn't have to register to make an order
- - Admin panel
- - - There should be 2 administrator levels
- - - - root can do anything: assign new helpers, add, edit and delete any object, pass the root status (not grant, there can't be several roots)
- - - - helpers that have limited abilities
- Different static content serving methods
- - Something like integration with webservers like nginx, caddy and other
- - Serving by ourselves (http.FileServer)

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
