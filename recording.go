package main

import (
	"github.com/spf13/cobra"
)

var (
	recordingCmd = &cobra.Command{
		Use:   "recording",
		Short: "manage recordings",
		Long:  `The recording command manages recordings`,
		RunE:  nil,
	}

	startStopRecordingCmd = &cobra.Command{
		Use:   "toggle",
		Short: "Toggle recording",
		RunE: func(cmd *cobra.Command, args []string) error {
			return starStopRecording()
		},
	}

	startRecordingCmd = &cobra.Command{
		Use:   "start",
		Short: "Starts recording",
		RunE: func(cmd *cobra.Command, args []string) error {
			return startRecording()
		},
	}

	stopRecordingCmd = &cobra.Command{
		Use:   "stop",
		Short: "Stops recording",
		RunE: func(cmd *cobra.Command, args []string) error {
			return stopRecording()
		},
	}
)

func starStopRecording() error {
	_, err := client.Recording.StartStopRecording()
	return err
}

func startRecording() error {
	_, err := client.Recording.StartRecording()
	return err
}

func stopRecording() error {
	_, err := client.Recording.StopRecording()
	return err
}

func init() {
	recordingCmd.AddCommand(startStopRecordingCmd)
	recordingCmd.AddCommand(startRecordingCmd)
	recordingCmd.AddCommand(stopRecordingCmd)
	rootCmd.AddCommand(recordingCmd)
}
