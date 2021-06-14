package main

import (
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
	/*
		resp, err := client.Streaming.GetStreamSettings()
		if err != nil {
			return err
		}

		r := streaming.StartStreamingParams{}
		r.Stream.Settings = resp.Settings
	*/

	_, err := client.Streaming.StartStreaming(nil)
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
