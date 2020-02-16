package main

import (
	"errors"
	"strings"

	obsws "github.com/muesli/go-obs-websocket"
	"github.com/spf13/cobra"
)

var switchSceneCmd = &cobra.Command{
	Use:   "switch-scene",
	Short: "Switch to a different OBS scene",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("switch-scene requires a scene name as argument")
		}
		return switchScene(strings.Join(args, " "))
	},
}

func switchScene(scene string) error {
	req := obsws.NewSetCurrentSceneRequest(scene)
	return req.Send(*client)
}

func init() {
	rootCmd.AddCommand(switchSceneCmd)
}
