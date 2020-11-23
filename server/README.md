[![Github Actions Badge](https://github.com/abtris/fiber-demo/workflows/Go/badge.svg)](https://github.com/abtris/fiber-demo/actions)

# Fiber Hello World Demo

> An Express-inspired web framework written in Go.

[Fiber](https://gofiber.io/) is Golang alternative to [Express](https://expressjs.com/), [Flask](https://flask.palletsprojects.com/) or [Sinatra](http://sinatrarb.com/) with similar API.

- Demo contains
  - static assets for css (`style.css`)
  - bootstrap linked (cdn links in template)
  - 404 message
  - rate limiting
  - basic logger
  - prometheus integration (`/metrics`)
  - request ID (for connect with metrics, traces)
  - support for env PORT (for example used in Heroku)
  - basic test for 200, 404 routes

## How to use it

```
$ make

Usage:
  make <target>

Targets:
  help        Display this help
  deps        Check && download dependencies
  build       Build docker image
  watch       Watch file changes and build
  run         Run
```

## Docs

- https://docs.gofiber.io/
- https://github.com/ansrivas/fiberprometheus

