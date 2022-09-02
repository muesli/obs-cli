package main

import (
	"errors"
	"fmt"

	// "github.com/andreykaipov/goobs/api/requests/sources"
	"github.com/muesli/coral"
)

var (
	labelCmd = &coral.Command{
		Use:   "label",
		Short: "manage text labels",
		Long:  `The label command manages text labels`,
		RunE:  nil,
	}

	textCmd = &coral.Command{
		Use:   "text",
		Short: "Changes a text label",
		RunE: func(cmd *coral.Command, args []string) error {
			if len(args) < 2 {
				return errors.New("text requires a source and the new text")
			}
			return changeLabel(args[0], args[1])
		},
	}
)

func changeLabel(source string, text string) error {
	// p := sources.GetTextFreetype2PropertiesParams{
	// 	Source: source,
	// }

	// resp, err := client.Sources.GetTextFreetype2Properties(&p)
	// if err != nil {
	// 	return err
	// }

	// r := sources.SetTextFreetype2PropertiesParams{
	// 	Source:      source,
	// 	Color1:      resp.Color1,
	// 	Color2:      resp.Color2,
	// 	CustomWidth: resp.CustomWidth,
	// 	DropShadow:  &resp.DropShadow,
	// 	Font:        resp.Font,
	// 	FromFile:    &resp.FromFile,
	// 	LogMode:     &resp.LogMode,
	// 	Outline:     &resp.Outline,
	// 	Text:        text,
	// 	TextFile:    resp.TextFile,
	// 	WordWrap:    &resp.WordWrap,
	// }

	// _, err = client.Sources.SetTextFreetype2Properties(&r)
	// return err
	return fmt.Errorf("TODO: refactor for v5 protocol changes")
}

func init() {
	labelCmd.AddCommand(textCmd)
	rootCmd.AddCommand(labelCmd)
}
