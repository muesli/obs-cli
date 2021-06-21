package main

import (
	"github.com/andreykaipov/goobs/api/requests/streaming"
	"github.com/spf13/cobra"
)

var (
	streamCmd = &cobra.Command{
		Use:   "stream",
		Short: "manage streams",
		Long:  `The stream command manages streams`,
		RunE:  nil,
	}

	startStreamCmd = &cobra.Command{
		Use:   "start",
		Short: "Starts streaming",
		RunE: func(cmd *cobra.Command, args []string) error {
			return startStream()
		},
	}

	stopStreamCmd = &cobra.Command{
		Use:   "stop",
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
	streamCmd.AddCommand(startStreamCmd)
	streamCmd.AddCommand(stopStreamCmd)
	rootCmd.AddCommand(streamCmd)
}
