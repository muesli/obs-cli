package main

import (
	obsws "github.com/muesli/go-obs-websocket"
	"github.com/spf13/cobra"
)

var (
	startStopRecordingCmd = &cobra.Command{
		Use:   "toggle-recording",
		Short: "Toggle recording",
		RunE: func(cmd *cobra.Command, args []string) error {
			return starStopRecording()
		},
	}

	startRecordingCmd = &cobra.Command{
		Use:   "start-recording",
		Short: "Starts recording",
		RunE: func(cmd *cobra.Command, args []string) error {
			return startRecording()
		},
	}

	stopRecordingCmd = &cobra.Command{
		Use:   "stop-recording",
		Short: "Stops recording",
		RunE: func(cmd *cobra.Command, args []string) error {
			return stopRecording()
		},
	}
)

func starStopRecording() error {
	req := obsws.NewStartStopRecordingRequest()
	return req.Send(*client)
}

func startRecording() error {
	req := obsws.NewStartRecordingRequest()
	return req.Send(*client)
}

func stopRecording() error {
	req := obsws.NewStopRecordingRequest()
	return req.Send(*client)
}

func init() {
	rootCmd.AddCommand(startStopRecordingCmd)
	rootCmd.AddCommand(startRecordingCmd)
	rootCmd.AddCommand(stopRecordingCmd)
}
