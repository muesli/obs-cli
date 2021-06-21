package main

import (
	"errors"
	"strings"

	"github.com/andreykaipov/goobs/api/requests/scenes"
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
	r := scenes.SetCurrentSceneParams{
		SceneName: scene,
	}
	_, err := client.Scenes.SetCurrentScene(&r)
	return err
}

func init() {
	rootCmd.AddCommand(switchSceneCmd)
}
