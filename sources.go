package main

import (
	"errors"
	"fmt"

	"github.com/andreykaipov/goobs/api/requests/sources"
	"github.com/spf13/cobra"
)

var (
	sourceCmd = &cobra.Command{
		Use:   "source",
		Short: "manage sources",
		Long:  `The source command manages sources`,
		RunE:  nil,
	}

	listSourcesCmd = &cobra.Command{
		Use:   "list",
		Short: "Lists all sources",
		RunE: func(cmd *cobra.Command, args []string) error {
			return listSources()
		},
	}

	toggleMuteCmd = &cobra.Command{
		Use:   "toggle-mute",
		Short: "Toggles mute",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("toggle-mute requires a source name as argument")
			}
			return toggleMute(args[0])
		},
	}
)

func listSources() error {
	/*
		{
			req := obsws.NewGetSourcesListRequest()
			resp, err := req.SendReceive(client)
			if err != nil {
				return err
			}

			fmt.Println("Sources\n=======\n")
			for _, v := range resp.Sources {
				spew.Dump(v)
			}
			fmt.Println()
		}
	*/

	{
		resp, err := client.Sources.GetSpecialSources()
		if err != nil {
			return err
		}

		fmt.Println("Special Sources")
		fmt.Println("===============")
		fmt.Printf("Desktop1: %s\n", resp.Desktop1)
		fmt.Printf("Desktop2: %s\n", resp.Desktop2)
		fmt.Printf("Mic1: %s\n", resp.Mic1)
		fmt.Printf("Mic2: %s\n", resp.Mic2)
		fmt.Printf("Mic3: %s\n", resp.Mic3)
	}

	return nil
}

func toggleMute(source string) error {
	p := sources.ToggleMuteParams{
		Source: source,
	}

	_, err := client.Sources.ToggleMute(&p)
	return err
}

func init() {
	sourceCmd.AddCommand(listSourcesCmd)
	sourceCmd.AddCommand(toggleMuteCmd)
	rootCmd.AddCommand(sourceCmd)
}
