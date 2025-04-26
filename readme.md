# Overlock - Decentralized Control Planes Network

[![codecov](https://codecov.io/github/overlock-network/node/coverage.svg?branch=main)](https://codecov.io/github/overlock-network/node?branch=main)

[![Go Report Card](https://goreportcard.com/badge/github.com/overlock-network/node)](https://goreportcard.com/report/github.com/overlock-network/node)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

[Overlock](https://overlock.network) is a secure, transparent, and decentralized Control Planes marketplace that connects those who need ready to use Crossplane or ArgoCD Control Planes (tenants) with those that have required Control Plane capacity to lease (providers).

For a high-level overview of the Overlock protocol and network economics, check out the [whitepaper](https://overlock.network/docs/OverlockWhitepaper.pdf); a detailed protocol definition can be 
found in the [design documentation](https://docs.overlock.network); and the target workload definition spec is [here](https://docs.overlock.network/sdl).

# Branching and Versioning

The `main` branch contains new features and is under active development; the `mainnet/main` branch contains the current, stable release.

* **stable** releases will have even minor numbers ( `v0.0.28` ) and be cut from the `mainnet/main` branch.
* **unstable** releases will have odd minor numbers ( `v0.0.28` ) and be cut from the `main` branch.

## Overlock Suite

Overlock Suite is the reference implementation of the [Overlock Protocol](https://overlock.network/docs/OverlockWhitepaper.pdf). Overlock is an actively-developed prototype currently focused on the distributed marketplace functionality.

The Suite is composed of one binary, `overlock`, which contains a ([tendermint](https://github.com/tendermint/tendermint)-powered) blockchain node that
implements the decentralized exchange as well as client functionality to access the exchange and network data in general.

## Get Started with Overlock

The easiest way to get started with Overlock is by following the [Quick Start Guide](https://docs.overlock.network/guides/deploy) to get started. 

## Join the Community

- [Join Developer Chat](https://discord.gg/amQZEMFbTe)

# Supported platforms

Platform | Arch | Status
--- | --- | :---
Darwin | amd64 | ✅ **Supported**
Linux | amd64 | ✅ **Supported**
Windows | amd64 | ⚠️ **Experimental**

# Installing

[GoDownloader](https://github.com/goreleaser/godownloader):

```sh
curl -sSfL https://raw.githubusercontent.com/overlock-network/node/main/install.sh | sh
```

Or install a specific version with [GoDownloader](https://github.com/goreleaser/godownloader)

```sh
curl -sSfL https://raw.githubusercontent.com/overlock-network/node/main/install | sh -s -- 0.0.37
```

# Roadmap and contributing

Overlock is written in Golang and is Apache 2.0 licensed - contributions are welcomed whether that means providing feedback, testing existing and new feature or hacking on the source.

To become a contributor, please see the guide on [contributing](CONTRIBUTING.md)

## Development environment
[This doc](https://github.com/overlock-network/node/blob/main/docs/development-environment.md) guides through setting up local development environment

Overlock is developed and tested with [golang 1.24.0+](https://golang.org/). 
Building requires a working [golang](https://golang.org/) installation, a properly set `GOPATH`, and `$GOPATH/bin` present in `$PATH`.
It is also required to have C/C++ compiler installed (gcc/clang) as there are C dependencies in use (libusb/libhid)
Overlock build process and examples are heavily tied to Makefile.


## Building from Source
Command below will compile overlock executable and put it into `.cache/bin`
```shell
make # overlock is set as default target thus `make` is equal to `make overlock`
```
once binary compiled it exempts system-wide installed overlock within overlock repo

## Running

We use thin integration testing environments to simplify
the development and testing process. We currently have three environments:

* [Single node](start.sh): simple (no workloads) single node running locally.