package main

import (
	"errors"
	"fmt"

	"github.com/andreykaipov/goobs/api/requests/sceneitems"
	"github.com/andreykaipov/goobs/api/typedefs"
	"github.com/muesli/coral"
)

var (
	sceneItemCmd = &coral.Command{
		Use:   "sceneitem",
		Short: "manage scene items",
		Long:  `The sceneitem command manages a scene's items`,
		RunE:  nil,
	}

	listSceneItemsCmd = &coral.Command{
		Use:   "list",
		Short: "Lists all items of a scene",
		RunE: func(cmd *coral.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("list requires a scene")
			}
			return listSceneItems(args[0])
		},
	}

	toggleSceneItemCmd = &coral.Command{
		Use:   "toggle",
		Short: "Toggles visibility of a scene-item",
		RunE: func(cmd *coral.Command, args []string) error {
			if len(args) < 2 {
				return errors.New("toggle requires a scene and scene-item")
			}
			return toggleSceneItem(args[0], args[1:]...)
		},
	}

	showSceneItemCmd = &coral.Command{
		Use:   "show",
		Short: "Makes a scene-item visible",
		RunE: func(cmd *coral.Command, args []string) error {
			if len(args) < 2 {
				return errors.New("show requires a scene and scene-item(s)")
			}
			return setSceneItemVisible(true, args[0], args[1:]...)
		},
	}

	hideSceneItemCmd = &coral.Command{
		Use:   "hide",
		Short: "Hides a scene-item",
		RunE: func(cmd *coral.Command, args []string) error {
			if len(args) < 2 {
				return errors.New("hide requires a scene and scene-item(s)")
			}
			return setSceneItemVisible(false, args[0], args[1:]...)
		},
	}

	getSceneItemVisibilityCmd = &coral.Command{
		Use:   "visible",
		Short: "Show visibility status of a scene-item",
		RunE: func(cmd *coral.Command, args []string) error {
			if len(args) < 2 {
				return errors.New("visible requires a scene and scene-item")
			}
			return getSceneItemVisibility(args[0], args[1:]...)
		},
	}

	centerSceneItemCmd = &coral.Command{
		Use:   "center",
		Short: "Horizontally centers a scene-item",
		RunE: func(cmd *coral.Command, args []string) error {
			if len(args) < 2 {
				return errors.New("center requires a scene and scene-item")
			}
			return centerSceneItem(args[0], args[1:]...)
		},
	}
)

func getSceneItems(scene string) ([]*typedefs.SceneItem, error) {
	resp, err := client.SceneItems.GetSceneItemList(&sceneitems.GetSceneItemListParams{SceneName: scene})
	if err != nil {
		return nil, err
	}
	return resp.SceneItems, nil
}

func listSceneItems(scene string) error {
	items, err := getSceneItems(scene)
	if err != nil {
		return err
	}

	for _, v := range items {
		fmt.Println(v.SourceName)
	}

	return nil
}

func sceneItemWithName(sceneItems []*typedefs.SceneItem, sourceName string) (*typedefs.SceneItem, error) {
	for _, item := range sceneItems {
		if item.SourceName == sourceName {
			return item, nil
		}
	}
	return nil, fmt.Errorf("unable to locate scene item with name: %s", sourceName)
}

func getSceneItemId(sceneItems []*typedefs.SceneItem, sourceName string) (int, error) {
	item, err := sceneItemWithName(sceneItems, sourceName)
	if err != nil {
		return 0, err
	}
	return item.SceneItemID, nil
}

func setSceneItemVisible(visible bool, scene string, itemNames ...string) error {
	items, err := getSceneItems(scene)
	if err != nil {
		return err
	}

	for _, itemName := range itemNames {
		id, err := getSceneItemId(items, itemName)
		if err != nil {
			return err
		}

		p := sceneitems.SetSceneItemEnabledParams{
			SceneItemEnabled: &visible,
			SceneItemId:      float64(id),
			SceneName:        scene,
		}
		_, err = client.SceneItems.SetSceneItemEnabled(&p)
		if err != nil {
			return err
		}
	}

	return nil
}

func toggleSceneItem(scene string, itemNames ...string) error {
	items, err := getSceneItems(scene)
	if err != nil {
		return err
	}

	for _, itemName := range itemNames {
		item, err := sceneItemWithName(items, itemName)
		if err != nil {
			return err
		}

		enabled := !item.SceneItemEnabled
		p := sceneitems.SetSceneItemEnabledParams{
			SceneItemEnabled: &enabled,
			SceneItemId:      float64(item.SceneItemID),
			SceneName:        scene,
		}
		_, err = client.SceneItems.SetSceneItemEnabled(&p)
		if err != nil {
			return err
		}
	}

	return nil
}

func getSceneItemVisibility(scene string, itemNames ...string) error {
	items, err := getSceneItems(scene)
	if err != nil {
		return err
	}

	for _, itemName := range itemNames {
		item, err := sceneItemWithName(items, itemName)
		if err != nil {
			return err
		}

		fmt.Printf("%s: %t\n", item.SourceName, item.SceneItemEnabled)
	}

	return nil
}

func centerSceneItem(scene string, itemNames ...string) error {
	items, err := getSceneItems(scene)
	if err != nil {
		return err
	}

	for _, itemName := range itemNames {
		id, err := getSceneItemId(items, itemName)
		if err != nil {
			return err
		}
		p := sceneitems.GetSceneItemTransformParams{
			SceneItemId: float64(id),
			SceneName:   scene,
		}
		resp, err := client.SceneItems.GetSceneItemTransform(&p)
		if err != nil {
			return err
		}

		vresp, err := client.Config.GetVideoSettings()
		if err != nil {
			return err
		}

		transform := resp.SceneItemTransform
		transform.PositionX = float64(vresp.BaseWidth) / 2

		r := sceneitems.SetSceneItemTransformParams{
			SceneName:          scene,
			SceneItemId:        float64(id),
			SceneItemTransform: transform,
		}

		_, err = client.SceneItems.SetSceneItemTransform(&r)
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
	sceneItemCmd.AddCommand(getSceneItemVisibilityCmd)
	sceneItemCmd.AddCommand(listSceneItemsCmd)
	rootCmd.AddCommand(sceneItemCmd)
}
