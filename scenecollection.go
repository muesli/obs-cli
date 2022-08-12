package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/andreykaipov/goobs/api/requests/config"
	"github.com/muesli/coral"
)

var (
	sceneCollectionCmd = &coral.Command{
		Use:   "scenecollection",
		Short: "manage scene collections",
		Long:  `The scenecollection command manages scene collections`,
		RunE:  nil,
	}

	listSceneCollectionCmd = &coral.Command{
		Use:   "list",
		Short: "List all scene collections",
		RunE: func(cmd *coral.Command, args []string) error {
			return listSceneCollections()
		},
	}

	getSceneCollectionCmd = &coral.Command{
		Use:   "get",
		Short: "Get the current scene collection",
		RunE: func(cmd *coral.Command, args []string) error {
			return getSceneCollection()
		},
	}

	setSceneCollectionCmd = &coral.Command{
		Use:   "set",
		Short: "Set the current scene collection",
		RunE: func(cmd *coral.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("set requires a scene collection name as argument")
			}
			return setSceneCollection(strings.Join(args, " "))
		},
	}
)

func listSceneCollections() error {
	r, err := client.Config.GetSceneCollectionList()
	if err != nil {
		return err
	}

	for _, v := range r.SceneCollections {
		fmt.Println(v)
	}
	return nil
}

func setSceneCollection(collection string) error {
	r := config.SetCurrentSceneCollectionParams{
		SceneCollectionName: collection,
	}
	_, err := client.Config.SetCurrentSceneCollection(&r)
	return err
}

func getSceneCollection() error {
	r, err := client.Config.GetSceneCollectionList()
	if err != nil {
		return err
	}

	fmt.Println(r.CurrentSceneCollectionName)
	return nil
}

func init() {
	sceneCollectionCmd.AddCommand(listSceneCollectionCmd)
	sceneCollectionCmd.AddCommand(setSceneCollectionCmd)
	sceneCollectionCmd.AddCommand(getSceneCollectionCmd)
	rootCmd.AddCommand(sceneCollectionCmd)
}
