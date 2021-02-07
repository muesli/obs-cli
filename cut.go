package main

import (

	"strconv"



	obsws "github.com/christopher-dG/go-obs-websocket"
	"github.com/spf13/cobra"
)

var transCmd = &cobra.Command{
	Use:   "trans",
	Short: "transitions with name ",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 1 {
			return trans(args[0],0)
		}
		if len(args) == 2 {
			i, _ :=strconv.Atoi(args[1])
			return trans(args[0],i)
		}
		return trans("",0)
	},
}
// Change the active transition before switching scenes.
// Defaults to the active transition.
// Required: No.
//WithTransition map[string]interface{} `json:"with-transition"`
// Name of the transition.
// Required: Yes.
//WithTransitionName string `json:"with-transition.name"`
// Transition duration (in milliseconds).
// Required: No.
//WithTransitionDuration int `json:"with-transition.duration"`
//_request               `json:",squash"`
//response               chan TransitionToProgramResponse



func trans(name string,duration int) error {
	req := obsws.NewTransitionToProgramRequest(nil,name,duration)
	return req.Send(*client)
}

func init() {
	rootCmd.AddCommand(transCmd)
}
