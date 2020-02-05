package main

import (
	"errors"

	obsws "github.com/christopher-dG/go-obs-websocket"
	"github.com/spf13/cobra"
)

var toggleMuteCmd = &cobra.Command{
	Use:   "toggle-mute",
	Short: "Toggles mute",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("toggle-mute requires a source name as argument")
		}
		return toggleMute(args[0])
	},
}

func toggleMute(source string) error {
	req := obsws.NewToggleMuteRequest(source)
	return req.Send(*client)
}

func init() {
	rootCmd.AddCommand(toggleMuteCmd)
}
