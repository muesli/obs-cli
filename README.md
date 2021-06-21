# obs-cli

[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://godoc.org/github.com/muesli/obs-cli)
[![Go ReportCard](https://goreportcard.com/badge/muesli/obs-cli)](https://goreportcard.com/report/muesli/obs-cli)

OBS-cli is a command-line remote control for OBS. It requires the
[obs-websocket](https://github.com/Palakis/obs-websocket) plugin to be installed on your system.

## Installation

### Packages & Binaries

On Arch Linux you can simply install the package from the AUR:

    yay -S obs-cli

Or download a binary from the [releases](https://github.com/muesli/obs-cli/releases)
page. Linux (including ARM) binaries are available, as well as Debian and RPM
packages.

### Build From Source

Alternatively you can also build `obs-cli` from source. Make sure you have a
working Go environment (Go 1.11 or higher is required). See the
[install instructions](https://golang.org/doc/install.html).

To install obs-cli, simply run:

    go get github.com/muesli/obs-cli

## Usage

All commands support the following flags:

- `--host`: which OBS instance to connect to
- `--port`: port to connect to
- `--password`: password used for authentication

### Streams

Start streaming:

```bash
obs-cli stream start
```

Stop streaming:

```bash
obs-cli stream stop
```

### Recordings

Start recording:

```bash
obs-cli recording start
```

Stop recording:

```bash
obs-cli recording stop
```

Toggle recording:

```bash
obs-cli recording toggle
```

### Scenes

Switch to a scene:

```bash
obs-cli scene switch <scene>
```

### Labels

Change a FreeType text label:

```bash
obs-cli label text <label> <text>
```

### Scene Items

List all items of a scene:

```bash
obs-cli sceneitem list <scene>
```

Make a scene-item visible:

```bash
obs-cli sceneitem show <scene> <item>
```

Hide a scene-item:

```bash
obs-cli sceneitem hide <scene> <item>
```

Toggle visibility of a scene-item:

```bash
obs-cli sceneitem toggle <scene> <item>
```

Center a scene-item horizontally:

```bash
obs-cli sceneitem center <scene> <item>
```

### Sources

List special sources:

```bash
obs-cli source list
```

Toggle mute status of a source:

```bash
obs-cli source toggle-mute <source>
```
