package main

import (
	"errors"
	"strings"

	obsws "github.com/christopher-dG/go-obs-websocket"
	"github.com/spf13/cobra"
)

var previewSceneCmd = &cobra.Command{
	Use:   "preview-scene",
	Short: "Preview to a different OBS scene",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("preview-scene requires a scene name as argument")
		}
		return previewScene(strings.Join(args, " "))
	},
}

func previewScene(scene string) error {
	req := obsws.NewSetPreviewSceneRequest(scene)
	return req.Send(*client)
}

func init() {
	rootCmd.AddCommand(previewSceneCmd)
}
