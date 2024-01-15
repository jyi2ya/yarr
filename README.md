# yarr

**yarr** (yet another rss reader) is a web-based feed aggregator which can be used both
as a desktop application and a personal self-hosted server.

The app is a single binary that connects to an external Postgres database for data persistence.

![screenshot](docs/promo.png)

## usage

**yarr** requires a Postgres database to use. Once you have that set up the easiest way to use **yarr** is to run it in Docker:

```sh
docker run -it -e YARR_DB=your_postgres_uri_here -p 7070:7070 jgkawell/yarr:latest
```

Then go to `localhost:7070` or `server_ip_address:7070` to start using it.

This fork of **yarr** is specifically designed to run as a microservice in orchestration. That could be Docker Compose or Kubernetes and the only requirement is the `YARR_DB` environment variable pointing an existing Postgres database.

See more:

* [Building from source code](docs/build.md)
* [Fever API support](docs/fever.md)

## credits

[Feather](https://feathericons.com/) for icons.
