# Versioner

[![go-doc](https://godoc.org/github.com/vitanovs/versioner?status.svg)](https://pkg.go.dev/github.com/vitanovs/versioner?tab=overview)
[![go-report](https://goreportcard.com/badge/github.com/vitanovs/versioner)](https://goreportcard.com/report/github.com/vitanovs/versioner)
[![license](https://img.shields.io/badge/license-BSD%202--Clause-orange.svg)](https://github.com/vitanovs/versioner/blob/master/LICENSE)

![versioner-logo](https://github.com/vitanovs/versioner/blob/master/doc/images/logo.png)

PostgreSQL schema versioning tool.

## Requirements

* [Go](https://golang.org/) v1.18 at least.

* [GNU Make](https://www.gnu.org/software/make/) v3.81 at least.

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

## Autocompletion

Enabling the autocompletion capabilities happens by exposing `PROG` environment variable and
sourcing dedicated completion script, located in `autocompletion/` directory of the project.

* BASH

    ```sh
    PROG=./bin/versioner source ./autocomplete/bash_autocomplete
    ```

* ZSH

    ```sh
    PROG=./bin/versioner _CLI_ZSH_AUTOCOMPLETE_HACK=1 source ./autocomplete/zsh_autocomplete
    ```

For more detailed overview of the autocompletion scripts see the [official documentation](https://github.com/urfave/cli/blob/master/docs/v2/manual.md#bash-completion).


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

Direct questions or issue to stoyan.a.vitanov@gmail.com or open GitHub issue right away.

## License

Copyright © 2020 Versioner
