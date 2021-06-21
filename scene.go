package main

import (
	"errors"
	"strings"

	"github.com/andreykaipov/goobs/api/requests/scenes"
	"github.com/spf13/cobra"
)

var (
	sceneCmd = &cobra.Command{
		Use:   "scene",
		Short: "manage scenes",
		Long:  `The scene command manages scenes`,
		RunE:  nil,
	}

	switchSceneCmd = &cobra.Command{
		Use:   "switch",
		Short: "Switch to a different scene",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("switch requires a scene name as argument")
			}
			return switchScene(strings.Join(args, " "))
		},
	}
)

func switchScene(scene string) error {
	r := scenes.SetCurrentSceneParams{
		SceneName: scene,
	}
	_, err := client.Scenes.SetCurrentScene(&r)
	return err
}

func init() {
	sceneCmd.AddCommand(switchSceneCmd)
	rootCmd.AddCommand(sceneCmd)
}
