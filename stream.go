package main

import (
	"fmt"
	"strconv"

	"github.com/andreykaipov/goobs/api/requests/streaming"
	"github.com/muesli/coral"
)

var (
	streamCmd = &coral.Command{
		Use:              "stream",
		Short:            "manage streams",
		Long:             `The stream command manages streams`,
		RunE:             nil,
		PersistentPreRun: connectOBSCommand,
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
	_, err := client.Streaming.StartStopStreaming(&streaming.StartStopStreamingParams{})
	return err
}

func startStream() error {
	_, err := client.Streaming.StartStreaming(&streaming.StartStreamingParams{})
	return err
}

func stopStream() error {
	_, err := client.Streaming.StopStreaming()
	return err
}

func streamStatus() error {
	r, err := client.Streaming.GetStreamingStatus()
	if err != nil {
		return err
	}

	fmt.Printf("Streaming: %s\n", strconv.FormatBool(r.Streaming))
	if !r.Streaming {
		return nil
	}

	fmt.Printf("Timecode: %s\n", r.StreamTimecode)

	rs, err := client.Streaming.GetStreamSettings()
	if err != nil {
		return err
	}

	fmt.Printf("URL: %s\n", rs.Settings.Server)
	return nil
}

func init() {
	streamCmd.AddCommand(startStopStreamCmd)
	streamCmd.AddCommand(startStreamCmd)
	streamCmd.AddCommand(stopStreamCmd)
	streamCmd.AddCommand(streamStatusCmd)
	rootCmd.AddCommand(streamCmd)
}
