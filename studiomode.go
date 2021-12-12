package main

import (
	"fmt"
	"strconv"

	studiomode "github.com/andreykaipov/goobs/api/requests/studio_mode"
	"github.com/muesli/coral"
)

var (
	studioModeCmd = &coral.Command{
		Use:              "studiomode",
		Short:            "manage studio mode",
		Long:             `The studiomode command manages the studio mode`,
		RunE:             nil,
		PersistentPreRun: connectOBSCommand,
	}

	disableStudioModeCmd = &coral.Command{
		Use:   "disable",
		Short: "Disables the studio mode",
		RunE: func(cmd *coral.Command, args []string) error {
			return disableStudioMode()
		},
	}

	enableStudioModeCmd = &coral.Command{
		Use:   "enable",
		Short: "Enables the studio mode",
		RunE: func(cmd *coral.Command, args []string) error {
			return enableStudioMode()
		},
	}

	studioModeStatusCmd = &coral.Command{
		Use:   "status",
		Short: "Reports studio mode status",
		RunE: func(cmd *coral.Command, args []string) error {
			return studioModeStatus()
		},
	}

	toggleStudioModeCmd = &coral.Command{
		Use:   "toggle",
		Short: "Toggles the studio mode (enable/disable)",
		RunE: func(cmd *coral.Command, args []string) error {
			return toggleStudioMode()
		},
	}

	transitionToProgramCmd = &coral.Command{
		Use:   "transition",
		Short: "Transition to program",
		RunE: func(cmd *coral.Command, args []string) error {
			return transitionToProgram()
		},
	}
)

func disableStudioMode() error {
	_, err := client.StudioMode.DisableStudioMode()
	return err
}

func enableStudioMode() error {
	_, err := client.StudioMode.EnableStudioMode()
	return err
}

// Determine if the studio mode is currently enabled in OBS.
func IsStudioModeEnabled() (bool, error) {
	r, err := client.StudioMode.GetStudioModeStatus()
	return r.StudioMode, err
}

func studioModeStatus() error {
	isStudioModeEnabled, err := IsStudioModeEnabled()
	if err != nil {
		return err
	}

	fmt.Printf("Studio Mode: %s\n", strconv.FormatBool(isStudioModeEnabled))
	return nil
}

func toggleStudioMode() error {
	_, err := client.StudioMode.ToggleStudioMode()
	return err
}

func transitionToProgram() error {
	_, err := client.StudioMode.TransitionToProgram(&studiomode.TransitionToProgramParams{})
	return err
}

func init() {
	studioModeCmd.AddCommand(disableStudioModeCmd)
	studioModeCmd.AddCommand(enableStudioModeCmd)
	studioModeCmd.AddCommand(studioModeStatusCmd)
	studioModeCmd.AddCommand(toggleStudioModeCmd)
	studioModeCmd.AddCommand(transitionToProgramCmd)
	rootCmd.AddCommand(studioModeCmd)
}
