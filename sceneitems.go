package main

import (
	"errors"
	"fmt"

	sceneitems "github.com/andreykaipov/goobs/api/requests/scene_items"
	"github.com/spf13/cobra"
)

var (
	listSceneItemsCmd = &cobra.Command{
		Use:   "list-sceneitems",
		Short: "Lists all scene-items of a source",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("list-sceneitems requires a scene")
			}
			return listSceneItems(args[0])
		},
	}

	toggleSceneItemCmd = &cobra.Command{
		Use:   "toggle-sceneitem",
		Short: "Toggles visibility of a scene-item",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return errors.New("toggle-sceneitem requires a scene and scene-item")
			}
			return toggleSceneItem(args[0], args[1:]...)
		},
	}

	showSceneItemCmd = &cobra.Command{
		Use:   "show-sceneitem",
		Short: "Makes a scene-item visible",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return errors.New("show-sceneitem requires a scene and scene-item(s)")
			}
			return setSceneItemVisible(true, args[0], args[1:]...)
		},
	}

	hideSceneItemCmd = &cobra.Command{
		Use:   "hide-sceneitem",
		Short: "Hides a scene-item",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return errors.New("hide-sceneitem requires a scene and scene-item(s)")
			}
			return setSceneItemVisible(false, args[0], args[1:]...)
		},
	}
)

func listSceneItems(scene string) error {
	resp, err := client.Scenes.GetSceneList()
	if err != nil {
		return err
	}

	for _, v := range resp.Scenes {
		if v["name"] != scene {
			continue
		}

		for _, s := range v["sources"].([]interface{}) {
			src := s.(map[string]interface{})
			fmt.Println(src["name"])
		}
	}

	return nil
}

func setSceneItemVisible(visible bool, scene string, items ...string) error {
	for _, item := range items {
		p := sceneitems.GetSceneItemPropertiesParams{
			Item:      item,
			SceneName: scene,
		}
		resp, err := client.SceneItems.GetSceneItemProperties(&p)
		if err != nil {
			return err
		}

		r := sceneitems.SetSceneItemPropertiesParams{
			SceneName: scene,
			Item:      item,
			Bounds:    resp.Bounds,
			Crop:      resp.Crop,
			Position:  resp.Position,
			Rotation:  resp.Rotation,
			Scale:     resp.Scale,
			Visible:   visible,
		}

		_, err = client.SceneItems.SetSceneItemProperties(&r)
		if err != nil {
			return err
		}
	}

	return nil
}

func toggleSceneItem(scene string, items ...string) error {
	for _, item := range items {
		p := sceneitems.GetSceneItemPropertiesParams{
			Item:      item,
			SceneName: scene,
		}
		resp, err := client.SceneItems.GetSceneItemProperties(&p)
		if err != nil {
			return err
		}

		err = setSceneItemVisible(!resp.Visible, scene, item)
		if err != nil {
			return err
		}
	}

	return nil
}

func init() {
	rootCmd.AddCommand(listSceneItemsCmd)
	rootCmd.AddCommand(toggleSceneItemCmd)
	rootCmd.AddCommand(showSceneItemCmd)
	rootCmd.AddCommand(hideSceneItemCmd)
}
