# obs-cli

[![Latest Release](https://img.shields.io/github/release/muesli/obs-cli.svg)](https://github.com/muesli/obs-cli/releases)
[![Build Status](https://github.com/muesli/obs-cli/workflows/build/badge.svg)](https://github.com/muesli/obs-cli/actions)
[![Go ReportCard](https://goreportcard.com/badge/muesli/obs-cli)](https://goreportcard.com/report/muesli/obs-cli)
[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://pkg.go.dev/github.com/muesli/obs-cli)

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

    go install github.com/muesli/obs-cli@latest

## Usage

All commands support the following flags:

- `--host`: which OBS instance to connect to
- `--port`: port to connect to
- `--password`: password used for authentication

### Streams

Change the streaming state:

```
obs-cli stream start
obs-cli stream stop
obs-cli stream toggle
```

Display streaming status:

```
obs-cli stream status
```

### Recordings

Change the recording state:

```
obs-cli recording start
obs-cli recording stop
obs-cli recording toggle
```

Display recording status:

```
obs-cli recording status
```

### Scenes

List all scene names:

```
obs-cli scene list
```

Switch program to a scene:

```
obs-cli scene current <scene>
```

Switch preview to a scene (studio mode must be enabled):

```
obs-cli scene preview <scene>
```

Switch program (studio mode disabled) or preview (studio mode enabled) to a scene:

```
obs-cli scene switch <scene>
```

### Labels

Change a FreeType text label:

```
obs-cli label text <label> <text>
```

Trigger a countdown and continuously update a label with the remaining time:

```
obs-cli label countdown <label> <duration>
```

### Scene Items

List all items of a scene:

```
obs-cli sceneitem list <scene>
```

Change the visibility of a scene-item:

```
obs-cli sceneitem show <scene> <item>
obs-cli sceneitem hide <scene> <item>
obs-cli sceneitem toggle <scene> <item>
```

Center a scene-item horizontally:

```
obs-cli sceneitem center <scene> <item>
```

### Sources

List special sources:

```
obs-cli source list
```

Toggle mute status of a source:

```
obs-cli source toggle-mute <source>
```

### Studio Mode

Enable or disable Studio Mode:

```
obs-cli studiomode enable
obs-cli studiomode disable
obs-cli studiomode toggle
```

Display studio mode status:

```
obs-cli studiomode status
```

Transition to program (when the studio mode is enabled):

```
obs-cli studiomode transition
```
