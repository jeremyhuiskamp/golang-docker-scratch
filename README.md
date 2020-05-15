# golang-docker-scratch

[![Build](https://github.com/jeremyhuiskamp/golang-docker-scratch/workflows/build/badge.svg)](https://github.com/jeremyhuiskamp/golang-docker-scratch/actions?query=workflow%3Abuild)

A recipe for go binaries in a scratch docker container with
up-to-date tls certs and timezone data.

Golang binaries pair well with the docker scratch image, but
sometimes they need some help.  In particular, it is common to
need a TLS certificate database and timezone data on the
file system.

Many tutorials suggest checking these files into your source
control, but this puts a maintenance burden on you to keep them
current.

This recipe leverages docker multistage builds to source the
data from the latest version of the alpine linux package repository
without needing the rest of alpine in the final image.

## Instructions

```
docker build -t golang-scratch-test . && docker run -it --rm golang-scratch-test
```
