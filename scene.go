package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/andreykaipov/goobs/api/requests/scenes"
	"github.com/andreykaipov/goobs/api/requests/studio_mode"
	"github.com/spf13/cobra"
)

var (
	sceneCmd = &cobra.Command{
		Use:   "scene",
		Short: "manage scenes",
		Long:  `The scene command manages scenes`,
		RunE:  nil,
	}

	currentSceneCmd = &cobra.Command{
		Use:   "current",
		Short: "Switch program to a different scene",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("current requires a scene name as argument")
			}
			return setCurrentScene(strings.Join(args, " "))
		},
	}

	listSceneCmd = &cobra.Command{
		Use:   "list",
		Short: "List all scene names",
		RunE: func(cmd *cobra.Command, args []string) error {
			return listScenes()
		},
	}

	previewSceneCmd = &cobra.Command{
		Use:   "preview",
		Short: "Switch preview to a different scene (studio mode must be enabled)",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("preview requires a scene name as argument")
			}
			return setPreviewScene(strings.Join(args, " "))
		},
	}

	switchSceneCmd = &cobra.Command{
		Use:   "switch",
		Short: "Switch program or preview in studio mode to a different scene",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("switch requires a scene name as argument")
			}
			return switchScene(strings.Join(args, " "))
		},
	}
)

func listScenes() error {
	r, err := client.Scenes.GetSceneList()
	if err != nil {
		return err
	}

	for _, v := range r.Scenes {
		fmt.Println(v.Name)
	}
	return nil
}

func setCurrentScene(scene string) error {
	r := scenes.SetCurrentSceneParams{
		SceneName: scene,
	}
	_, err := client.Scenes.SetCurrentScene(&r)
	return err
}

func setPreviewScene(scene string) error {
	r := studiomode.SetPreviewSceneParams{
		SceneName: scene,
	}
	_, err := client.StudioMode.SetPreviewScene(&r)
	return err
}

func switchScene(scene string) error {
	isStudioModeEnabled, err := IsStudioModeEnabled()
	if err != nil {
		return err
	}

	if isStudioModeEnabled {
		return setPreviewScene(scene)
	}
	return setCurrentScene(scene)
}

func init() {
	sceneCmd.AddCommand(currentSceneCmd)
	sceneCmd.AddCommand(listSceneCmd)
	sceneCmd.AddCommand(previewSceneCmd)
	sceneCmd.AddCommand(switchSceneCmd)
	rootCmd.AddCommand(sceneCmd)
}
