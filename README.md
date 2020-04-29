# Versioner

The Hall Arranger Schema versioning tool.

[![Documentation](https://godoc.org/github.com/vitanovs/versioner?status.svg)](http://godoc.org/github.com/vitanovs/versioner)
[![Go Report Card](https://goreportcard.com/badge/github.com/vitanovs/versioner)](https://goreportcard.com/report/github.com/vitanovs/versioner)

## Requirements

* [Go](https://golang.org/) v1.13.0 at least.

* [Docker](https://www.docker.com) v19.03.4 at least.

## Installation

In order to install `versioner`, follow these steps after
cloning the repository.

```sh
make install
```

To uninstall the binary, run

```sh
make uninstall
```

The `versioner` binary will be installed in your `$GOPATH/bin` directory.

## Usage

The `versioner` CLI tool provides comprehensive menu where all commands and options can be reviewed. To show the help menu, run

```sh
versioner --help
```

To print the menus of the sub commands use:

```sh
versioner [options] <command> [options] <sub-command> --help
```

## Docker

The `versioner` tool can also be generated as a Docker image. To build image use

```sh
make docker
```

### Docker Image Usage

The following is an example of how to use the `versioner` Docker image to run its commands

```sh
docker run --rm versioner --help
```

Use the following template to create new database:

```sh
docker run --rm versioner \
--endpoint <remote> \
--port <remote-port> \
--database <database> \
--username <username-credential> \
--password <password-credential> \
--sslmode <postgres-ssl-mode> \
database create \
--name <database-name>
```

To remove already existing database, replace `create` with `drop`.
In case your Postgres server runs on localhost and your OS of choice is macOS or Windows, replace `<remote>` with:

* `docker.for.mac.localhost` for macOS
* `docker.for.win.localhost` for Windows

### Docker Requirements

* [Go Docker Image](https://hub.docker.com/_/golang) v1.13.0 at least.

```sh
docker pull golang:1.13.0
```

* [BusyBox Docker Image](https://hub.docker.com/_/busybox) v1.31.1 at least.

```sh
docker pull busybox:1.31.1
```

## Contact

Direct questions or issue to hall.arranger@gmail.com or open GitHub issue right away.

## License

Copyright © 2020 Hall Arranger
