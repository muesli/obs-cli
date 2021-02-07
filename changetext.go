package main

import (
	"errors"

	obsws "github.com/christopher-dG/go-obs-websocket"
	"github.com/spf13/cobra"
)

var changeTextCmd = &cobra.Command{
	Use:   "change-text",
	Short: "Changes a text label",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("change-text requires a source and the new text")
		}
		return changeLabel(args[0], args[1])
	},
}

func changeLabel(source string, text string) error {
	req := obsws.NewGetTextFreetype2PropertiesRequest(source)
	resp, err := req.SendReceive(*client)
	if err != nil {
		return err
	}

	chreq := obsws.NewSetTextFreetype2PropertiesRequest(
		source,
		resp.Color1,
		resp.Color2,
		resp.CustomWidth,
		resp.DropShadow,
		resp.Font,
		resp.FontFace,
		resp.FontFlags,
		resp.FontSize,
		resp.FontStyle,
		resp.FromFile,
		resp.LogMode,
		resp.Outline,
		text,
		resp.TextFile,
		resp.WordWrap)
	return chreq.Send(*client)
}

func init() {
	rootCmd.AddCommand(changeTextCmd)
}
