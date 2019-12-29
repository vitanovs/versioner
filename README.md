# Versioner

The Hall Arranger Schema versioning tool.

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

### Docker Requirements

* [Golang Docker Image](https://hub.docker.com/_/golang) v1.13.0 at least.

```sh
docker pull golang:1.13.0
```

* [Busybox Docker Image](https://hub.docker.com/_/busybox) v1.31.1 at least.

```sh
docker pull busybox:1.31.1
```


## Contact

Direct questions or issue to hall.arranger@gmail.com or open GitHub issue right away.

## License

Copyright © 2020 Hall Arranger
