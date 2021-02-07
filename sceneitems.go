package main

import (
	"errors"
	"fmt"

	obsws "github.com/christopher-dG/go-obs-websocket"
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
	req := obsws.NewGetSceneListRequest()
	resp, err := req.SendReceive(*client)
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
		req := obsws.NewGetSceneItemPropertiesRequest(scene, item)
		resp, err := req.SendReceive(*client)
		if err != nil {
			return err
		}

		chreq := obsws.NewSetSceneItemPropertiesRequest(
			scene,
			item,
			resp.PositionX,
			resp.PositionY,
			resp.PositionAlignment,
			resp.Rotation,
			resp.ScaleX,
			resp.ScaleY,
			resp.CropTop,
			resp.CropBottom,
			resp.CropLeft,
			resp.CropRight,
			visible,
			false,
			resp.BoundsType,
			resp.BoundsAlignment,
			resp.BoundsX,
			resp.BoundsY)
		err = chreq.Send(*client)
		if err != nil {
			return err
		}
	}

	return nil
}

func toggleSceneItem(scene string, items ...string) error {
	for _, item := range items {
		req := obsws.NewGetSceneItemPropertiesRequest(scene, item)
		resp, err := req.SendReceive(*client)
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
