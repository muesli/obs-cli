# obs-cli

[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://godoc.org/github.com/muesli/obs-cli)
[![Go ReportCard](http://goreportcard.com/badge/muesli/obs-cli)](http://goreportcard.com/report/muesli/obs-cli)

OBS-cli is a command-line remote control for OBS

## Installation

Make sure you have a working Go environment (Go 1.11 or higher is required).
See the [install instructions](http://golang.org/doc/install.html).

To install obs-cli, simply run:

    go get github.com/muesli/obs-cli

## Usage

List special sources:

```bash
obs-cli list-sources
```

Stop streaming:

```bash
obs-cli stop-stream
```

Switch to a scene:

```bash
obs-cli switch-scene <scene>
```

Toggle mute status of a source:

```bash
obs-cli toggle-mute <source>
```
