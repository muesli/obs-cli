package main

import (
	"errors"
	"fmt"

	sceneitems "github.com/andreykaipov/goobs/api/requests/scene_items"
	"github.com/andreykaipov/goobs/api/typedefs"
	"github.com/spf13/cobra"
)

var (
	sceneItemCmd = &cobra.Command{
		Use:   "sceneitem",
		Short: "manage scene items",
		Long:  `The sceneitem command manages a scene's items`,
		RunE:  nil,
	}

	listSceneItemsCmd = &cobra.Command{
		Use:   "list",
		Short: "Lists all items of a scene",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("list requires a scene")
			}
			return listSceneItems(args[0])
		},
	}

	toggleSceneItemCmd = &cobra.Command{
		Use:   "toggle",
		Short: "Toggles visibility of a scene-item",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return errors.New("toggle requires a scene and scene-item")
			}
			return toggleSceneItem(args[0], args[1:]...)
		},
	}

	showSceneItemCmd = &cobra.Command{
		Use:   "show",
		Short: "Makes a scene-item visible",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return errors.New("show requires a scene and scene-item(s)")
			}
			return setSceneItemVisible(true, args[0], args[1:]...)
		},
	}

	hideSceneItemCmd = &cobra.Command{
		Use:   "hide",
		Short: "Hides a scene-item",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return errors.New("hide requires a scene and scene-item(s)")
			}
			return setSceneItemVisible(false, args[0], args[1:]...)
		},
	}

	centerSceneItemCmd = &cobra.Command{
		Use:   "center",
		Short: "Horizontally centers a scene-item",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return errors.New("center requires a scene and scene-item")
			}
			return centerSceneItem(args[0], args[1:]...)
		},
	}
)

func listSceneItems(scene string) error {
	resp, err := client.Scenes.GetSceneList()
	if err != nil {
		return err
	}

	for _, v := range resp.Scenes {
		if v.Name != scene {
			continue
		}

		for _, s := range v.Sources {
			fmt.Println(s.Name)
		}
	}

	return nil
}

func setSceneItemVisible(visible bool, scene string, items ...string) error {
	for _, item := range items {
		p := sceneitems.GetSceneItemPropertiesParams{
			Item:      &typedefs.Item{Name: item},
			SceneName: scene,
		}
		resp, err := client.SceneItems.GetSceneItemProperties(&p)
		if err != nil {
			return err
		}

		r := sceneitems.SetSceneItemPropertiesParams{
			SceneName: scene,
			Item:      &typedefs.Item{Name: item},
			Bounds:    resp.Bounds,
			Crop:      resp.Crop,
			Position:  resp.Position,
			Rotation:  resp.Rotation,
			Scale:     resp.Scale,
			Locked:    resp.Locked,
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
			Item:      &typedefs.Item{Name: item},
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

func centerSceneItem(scene string, items ...string) error {
	for _, item := range items {
		p := sceneitems.GetSceneItemPropertiesParams{
			Item:      &typedefs.Item{Name: item},
			SceneName: scene,
		}
		resp, err := client.SceneItems.GetSceneItemProperties(&p)
		if err != nil {
			return err
		}

		vresp, err := client.General.GetVideoInfo()
		if err != nil {
			return err
		}

		pos := resp.Position
		pos.X = float64(vresp.BaseWidth) / 2
		r := sceneitems.SetSceneItemPropertiesParams{
			SceneName: scene,
			Item:      &typedefs.Item{Name: item},
			Bounds:    resp.Bounds,
			Crop:      resp.Crop,
			Position:  pos,
			Rotation:  resp.Rotation,
			Scale:     resp.Scale,
			Locked:    resp.Locked,
			Visible:   resp.Visible,
		}

		_, err = client.SceneItems.SetSceneItemProperties(&r)
		if err != nil {
			return err
		}
	}

	return nil
}

func init() {
	sceneItemCmd.AddCommand(centerSceneItemCmd)
	sceneItemCmd.AddCommand(toggleSceneItemCmd)
	sceneItemCmd.AddCommand(showSceneItemCmd)
	sceneItemCmd.AddCommand(hideSceneItemCmd)
	sceneItemCmd.AddCommand(listSceneItemsCmd)
	rootCmd.AddCommand(sceneItemCmd)
}
