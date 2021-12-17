package main

import (
	"errors"
	"fmt"
	"strings"

	scenecollections "github.com/andreykaipov/goobs/api/requests/scene_collections"
	"github.com/spf13/cobra"
)

var (
	sceneCollectionCmd = &cobra.Command{
		Use:   "scenecollection",
		Short: "manage scene collections",
		Long:  `The scenecollection command manages scene collections`,
		RunE:  nil,
	}

	listSceneCollectionCmd = &cobra.Command{
		Use:   "list",
		Short: "List all scene collections",
		RunE: func(cmd *cobra.Command, args []string) error {
			return listSceneCollections()
		},
	}

	getSceneCollectionCmd = &cobra.Command{
		Use:   "get",
		Short: "Get the current scene collection",
		RunE: func(cmd *cobra.Command, args []string) error {
			return getSceneCollection()
		},
	}

	setSceneCollectionCmd = &cobra.Command{
		Use:   "set",
		Short: "Set the current scene collection",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("set requires a scene collection name as argument")
			}
			return setSceneCollection(strings.Join(args, " "))
		},
	}
)

func listSceneCollections() error {
	r, err := client.SceneCollections.ListSceneCollections()
	if err != nil {
		return err
	}

	for _, v := range r.SceneCollections {
		fmt.Println(v.ScName)
	}
	return nil
}

func setSceneCollection(collection string) error {
	r := scenecollections.SetCurrentSceneCollectionParams{
		ScName: collection,
	}
	_, err := client.SceneCollections.SetCurrentSceneCollection(&r)
	return err
}

func getSceneCollection() error {
	r, err := client.SceneCollections.GetCurrentSceneCollection()
	if err != nil {
		return err
	}

	fmt.Println(r.ScName)
	return nil
}

func init() {
	sceneCollectionCmd.AddCommand(listSceneCollectionCmd)
	sceneCollectionCmd.AddCommand(setSceneCollectionCmd)
	sceneCollectionCmd.AddCommand(getSceneCollectionCmd)
	rootCmd.AddCommand(sceneCollectionCmd)
}
