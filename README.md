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

```
obs-cli stream start
```

Stop streaming:

```
obs-cli stream stop
```

Toggle streaming:

```
obs-cli stream toggle
```

Display streaming status:

```
obs-cli stream status
```

### Recordings

Start recording:

```
obs-cli recording start
```

Stop recording:

```
obs-cli recording stop
```

Toggle recording:

```
obs-cli recording toggle
```

Display recording status:

```
obs-cli recording status
```

### Scenes

Switch to a scene:

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

Make a scene-item visible:

```
obs-cli sceneitem show <scene> <item>
```

Hide a scene-item:

```
obs-cli sceneitem hide <scene> <item>
```

Toggle visibility of a scene-item:

```
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
