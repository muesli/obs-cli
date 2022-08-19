package main

import (
	"fmt"
	"strconv"

	"github.com/andreykaipov/goobs/api/requests/stream"
	"github.com/muesli/coral"
)

var (
	streamCmd = &coral.Command{
		Use:   "stream",
		Short: "manage streams",
		Long:  `The stream command manages streams`,
		RunE:  nil,
	}

	startStopStreamCmd = &coral.Command{
		Use:   "toggle",
		Short: "Toggle streaming",
		RunE: func(cmd *coral.Command, args []string) error {
			return startStopStream()
		},
	}

	startStreamCmd = &coral.Command{
		Use:   "start",
		Short: "Starts streaming",
		RunE: func(cmd *coral.Command, args []string) error {
			return startStream()
		},
	}

	stopStreamCmd = &coral.Command{
		Use:   "stop",
		Short: "Stops streaming",
		RunE: func(cmd *coral.Command, args []string) error {
			return stopStream()
		},
	}

	streamStatusCmd = &coral.Command{
		Use:   "status",
		Short: "Reports streaming status",
		RunE: func(cmd *coral.Command, args []string) error {
			return streamStatus()
		},
	}
)

func startStopStream() error {
	_, err := client.Stream.ToggleStream(&stream.ToggleStreamParams{})
	return err
}

func startStream() error {
	_, err := client.Stream.StartStream(&stream.StartStreamParams{})
	return err
}

func stopStream() error {
	_, err := client.Stream.StopStream()
	return err
}

func streamStatus() error {
	r, err := client.Stream.GetStreamStatus()
	if err != nil {
		return err
	}

	fmt.Printf("Streaming: %s\n", strconv.FormatBool(r.OutputActive))
	if !r.OutputActive {
		return nil
	}

	fmt.Printf("Timecode: %s\n", r.OutputTimecode)

	rs, err := client.Config.GetStreamServiceSettings()
	if err != nil {
		return err
	}

	fmt.Printf("URL: %s\n", rs.StreamServiceSettings.Server)
	return nil
}

func init() {
	streamCmd.AddCommand(startStopStreamCmd)
	streamCmd.AddCommand(startStreamCmd)
	streamCmd.AddCommand(stopStreamCmd)
	streamCmd.AddCommand(streamStatusCmd)
	rootCmd.AddCommand(streamCmd)
}
