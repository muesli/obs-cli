package main

import (
	obsws "github.com/muesli/go-obs-websocket"
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
	m := make(map[string]interface{})
	req := obsws.NewStartStreamingRequest(m, "", m, m, "", "", true, "", "")
	return req.Send(*client)
}

func stopStream() error {
	req := obsws.NewStopStreamingRequest()
	return req.Send(*client)
}

func init() {
	// rootCmd.AddCommand(startStreamCmd)
	rootCmd.AddCommand(stopStreamCmd)
}
