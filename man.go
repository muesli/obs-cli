//go:build mango
// +build mango

package main

import (
	"fmt"

	"github.com/muesli/coral"
	mcoral "github.com/muesli/mango-coral"
	"github.com/muesli/roff"
)

var (
	manCmd = &coral.Command{
		Use:    "man [directory]",
		Short:  "Generates manpages",
		Long:   `Generates manpages in the given directory`,
		Hidden: true,
		RunE: func(cmd *coral.Command, args []string) error {
			return generateManPages()
		},
	}
)

func init() {
	rootCmd.AddCommand(manCmd)
}

func generateManPages() error {
	manPage, err := mcoral.NewManPage(1, rootCmd)
	if err != nil {
		return err
	}

	manPage = manPage.WithSection("Copyright", "(C) 2022 Christian Muehlhaeuser.\n"+
		"Released under MIT license.")

	fmt.Println(manPage.Build(roff.NewDocument()))
	return nil
}
