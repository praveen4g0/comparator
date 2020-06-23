# Comparator library in Go

#### Problem statement
Create a comparator library that can be used to compare 2 API responses (HTTP/HTTPS)

1. The input to your code should be two files with millions of request URLs.

### Prerequisites

This example requires the following softwares to run.
  * [Go](https://golang.org/)
  * [Gauge](https://docs.gauge.org/getting_started/installing-gauge.html)
  * Gauge Go plugin
    * can be installed using `gauge --install go`
  * Clone this repository in [GOPATH](https://github.com/golang/go/wiki/GOPATH).

#### Installing Gauge on Linux

* *Install using DNF Package Manager*

```sudo dnf install gauge```

* *Install using Curl*

Install Gauge to /usr/local/bin by running

```curl -SsL https://downloads.gauge.org/stable | sh```

Or install Gauge to a [custom path] using

```curl -SsL https://downloads.gauge.org/stable | sh -s -- --location-[custom path]```
### Alternative Gauge Installation methods

* [Gauge](https://docs.gauge.org/getting_started/installing-gauge.html)


 * Gauge Go plugin
  * can be installed using
  ```
  gauge install go
  ```
* Gauge html plugin

* can be installed using
```
gauge install html-report
```

## Running comparator tests
```
gauge run --log-level=debug --verbose  specs/comparator.spec
```

## Framwork structure organization
* Utilites are under `pkg` directory
* `env` contains gauge configuarations
* `specs` contains test scenraios written in `markdown` syntax
* `stepImpl` contains step definations for test scerarios under .spec file
* `testdata` contains testdata required by the framework to run
* On Execution of gauge command it generates 2 directories 	`logs` and `reports`

## Docker Run
* Build docker file
```
docker build . -t <username>/comparator:v0.1
```
* Docker run
```
docker run --rm -it <username>/comparator:v0.1 gauge run --log-level=debug --verbose specs/comparator.spec 
```
