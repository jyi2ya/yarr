# yarr

**yarr** (yet another rss reader) is a web-based feed aggregator which can be used as a personal self-hosted server.

The app is a single binary that connects to an external Postgres database for data persistence. This is a fork of [the original](https://github.com/nkanaev/yarr) that has been greatly simplified to remove features and dependencies for running as a locally installed application and also changed from using sqlite to using postgres for persistence. The reason for these changes is to focus on running as a simple microservice instead of a full-featured local application.

![screenshot](docs/promo.png)

## usage

**yarr** requires a Postgres database to use. Once you have that set up the easiest way to use **yarr** is to run it in Docker:

```sh
docker run -it -e YARR_DB=your_postgres_uri_here -p 7070:7070 jgkawell/yarr:latest
```

Then go to [http://localhost:7070](http://localhost:7070) or [http://server_ip_address:7070](http://server_ip_address:7070) to start using it.

This fork of **yarr** is specifically designed to run as a microservice in orchestration. That could be Docker Compose or Kubernetes and the only requirement is the `YARR_DB` environment variable pointing to an existing Postgres database.

See more:

* [Building from source code](docs/build.md)
* [Fever API support](docs/fever.md)

## credits

[Feather](https://feathericons.com/) for icons.
