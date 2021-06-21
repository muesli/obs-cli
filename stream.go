package main

import (
	"github.com/andreykaipov/goobs/api/requests/streaming"
	"github.com/spf13/cobra"
)

var (
	startStreamCmd = &cobra.Command{
		Use:   "start-stream",
		Short: "Starts streaming",
		RunE: func(cmd *cobra.Command, args []string) error {
			return startStream()
		},
	}

	stopStreamCmd = &cobra.Command{
		Use:   "stop-stream",
		Short: "Stops streaming",
		RunE: func(cmd *cobra.Command, args []string) error {
			return stopStream()
		},
	}
)

func startStream() error {
	_, err := client.Streaming.StartStreaming(&streaming.StartStreamingParams{})
	return err
}

func stopStream() error {
	_, err := client.Streaming.StopStreaming()
	return err
}

func init() {
	rootCmd.AddCommand(startStreamCmd)
	rootCmd.AddCommand(stopStreamCmd)
}
