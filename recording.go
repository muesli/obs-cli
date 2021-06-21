package main

import (
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
	rootCmd.AddCommand(startStopRecordingCmd)
	rootCmd.AddCommand(startRecordingCmd)
	rootCmd.AddCommand(stopRecordingCmd)
}
