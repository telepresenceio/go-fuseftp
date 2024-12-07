# go-fuseftp: An FTP based Fuse implementation written in Go

[![Unit tests](https://github.com/telepresenceio/go-fuseftp/actions/workflows/unit_tests.yaml/badge.svg)](https://github.com/telepresenceio/go-fuseftp/actions/workflows/unit_tests.yaml)
[![Coverage Status](https://coveralls.io/repos/github/telepresenceio/go-fuseftp/badge.svg?branch=master)](https://coveralls.io/github/telepresenceio/go-fuseftp?branch=master)
[![Go ReportCard](https://goreportcard.com/badge/telepresenceio/go-fuseftp)](http://goreportcard.com/report/telepresenceio/go-fuseftp)
[![Go Reference](https://pkg.go.dev/badge/github.com/telepresenceio/go-fuseftp.svg)](https://pkg.go.dev/github.com/telepresenceio/go-fuseftp)
[![Go version](https://img.shields.io/github/go-mod/go-version/telepresenceio/go-fuseftp)](https://golang.org/doc/devel/release.html)

## Install ##

```
go get -u github.com/telepresenceio/go-fuseftp
```

## Architecture

This is a FUSE (Filesystem in USEr Space) implementation for FTP, intended to be used as a library. It builds upon:
- [github.com/winfsp/cgofuse](https://github.com/winfsp/cgofuse), a cross-platform FUSE library for Go
- [github.com/jlaffaye/ftp](https://github.com/jlaffaye/ftp), an FTP client library for Go

The tests uses an FTP server based on [github.com/fclairamb/ftpserverlib](https://github.com/fclairamb/ftpserverlib),
a Golang FTP Server Libray

## Building

The `cgofuse` library relies on CGO, and that the headers for libfuse is installed. On Ubuntu, you'd typically do
```console
$ sudo apt-get install -y libfuse-dev
```
Once installed, the package can be unit tested using:
```console
$ make test
```
The gRPC server binary is built using:
```console
$ make fuseftp
```
