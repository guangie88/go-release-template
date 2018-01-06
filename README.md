# `go-release-template`

For testing out the full CI/CD development.

## Requirements

* [`go`](https://golang.org/dl/) + [`glide`](https://glide.sh/) **OR**
* [`docker`](https://www.docker.com/get-docker) +
  [`docker-compose`](https://docs.docker.com/compose/install/) that is able to
  run Compose v2 file

## How to Build

### Native Build

Install the go dependencies for this project into `vendor/`:

```bash
glide install
```

Build and you are done:

```bash
go build
```

The compiled executable will be located at the repository root directory, and is
named `go-release-template`.

For fully statically linked executable, use the following command instead:

```bash
CGO_ENABLED=0 go build
```

### Docker Build

This alternative method is recommended if you do not wish to up `go`, `glide` or
the environment variable `GOPATH`, or prefer not to have this repository files
not within the `GOPATH`.

Run the following for full compilation:

```bash
docker-compose run -u $UID:`id -g` all
```

Note that `` -u $UID:`id -g` `` is optional, but useful to perform the
compilation under your current user's UID and GID.

The compiled executable will be located at the repository root directory, and is
named `go-release-template`. The executable is always fully statically linked.

The following commands are available to run for `docker-compose`:

* `all`
  * Performs `glide install`, followed by `go build`.
  * e.g. `` docker-compose run -u $UID:`id -g` all ``
* `install`
  * Performs only `glide install`.
  * e.g. `` docker-compose run -u $UID:`id -g` install ``
* `build`
  * Performs only `go build`.
  * e.g. `` docker-compose run -u $UID:`id -g` build ``
* `clean`
  * Performs `go clean`.
  * e.g. `` docker-compose run -u $UID:`id -g` clean ``
* `update`
  * Performs `glide update`. Not recommended since it updates the dependencies.
  * e.g. `` docker-compose run -u $UID:`id -g` update ``

## How to Run

```bash
./go-release-template
```

Running the above should print the current executable path in a JSON format.
