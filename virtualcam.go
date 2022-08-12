package main

import (
	"fmt"
	"strconv"

	"github.com/muesli/coral"
)

var (
	virtualCamCmd = &coral.Command{
		Use:   "virtualcam",
		Short: "manage virtual camera",
		Long:  `The virtualcam command manages the virtual camera`,
		RunE:  nil,
	}

	startStopVirtualCamCmd = &coral.Command{
		Use:   "toggle",
		Short: "Toggle virtual camera status",
		RunE: func(cmd *coral.Command, args []string) error {
			return starStopVirtualCam()
		},
	}

	startVirtualCamCmd = &coral.Command{
		Use:   "start",
		Short: "Starts virtual camera",
		RunE: func(cmd *coral.Command, args []string) error {
			return startVirtualCam()
		},
	}

	stopVirtualCamCmd = &coral.Command{
		Use:   "stop",
		Short: "Stops virtual camera",
		RunE: func(cmd *coral.Command, args []string) error {
			return stopVirtualCam()
		},
	}

	virtualCamStatusCmd = &coral.Command{
		Use:   "status",
		Short: "Reports virtual camera status",
		RunE: func(cmd *coral.Command, args []string) error {
			return virtualCamStatus()
		},
	}
)

func starStopVirtualCam() error {
	_, err := client.Outputs.ToggleVirtualCam()
	return err
}

func startVirtualCam() error {
	_, err := client.Outputs.StartVirtualCam()
	return err
}

func stopVirtualCam() error {
	_, err := client.Outputs.StopVirtualCam()
	return err
}

func virtualCamStatus() error {
	r, err := client.Outputs.GetVirtualCamStatus()
	if err != nil {
		return err
	}

	fmt.Printf("Virtual camera: %s\n", strconv.FormatBool(r.OutputActive))
	if !r.OutputActive {
		return nil
	}

	// TODO: see if virtual cam timecode is available with different API
	// fmt.Printf("Timecode: %s\n", r.Ou)
	return nil
}

func init() {
	virtualCamCmd.AddCommand(startStopVirtualCamCmd)
	virtualCamCmd.AddCommand(startVirtualCamCmd)
	virtualCamCmd.AddCommand(stopVirtualCamCmd)
	virtualCamCmd.AddCommand(virtualCamStatusCmd)
	rootCmd.AddCommand(virtualCamCmd)
}
