package main

import (
	"fmt"

	obsws "github.com/muesli/go-obs-websocket"
	"github.com/spf13/cobra"
)

var listSourcesCmd = &cobra.Command{
	Use:   "list-sources",
	Short: "Lists all sources",
	RunE: func(cmd *cobra.Command, args []string) error {
		return listSources()
	},
}

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
		req := obsws.NewGetSpecialSourcesRequest()
		resp, err := req.SendReceive(*client)
		if err != nil {
			return err
		}

		fmt.Println("Special Sources\n===============\n")
		fmt.Printf("Desktop1: %s\n", resp.Desktop1)
		fmt.Printf("Desktop2: %s\n", resp.Desktop2)
		fmt.Printf("Mic1: %s\n", resp.Mic1)
		fmt.Printf("Mic2: %s\n", resp.Mic2)
		fmt.Printf("Mic3: %s\n", resp.Mic3)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(listSourcesCmd)
}
