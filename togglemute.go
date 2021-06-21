package main

import (
	"errors"

	"github.com/andreykaipov/goobs/api/requests/sources"
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
	p := sources.ToggleMuteParams{
		Source: source,
	}

	_, err := client.Sources.ToggleMute(&p)
	return err
}

func init() {
	rootCmd.AddCommand(toggleMuteCmd)
}
