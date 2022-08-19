package main

import (
	"fmt"
	"strconv"

	"github.com/andreykaipov/goobs/api/requests/transitions"
	"github.com/andreykaipov/goobs/api/requests/ui"
	"github.com/muesli/coral"
)

var (
	studioModeCmd = &coral.Command{
		Use:   "studiomode",
		Short: "manage studio mode",
		Long:  `The studiomode command manages the studio mode`,
		RunE:  nil,
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

func setStudioModeEnabled(enabled bool) error {
	_, err := client.Ui.SetStudioModeEnabled(&ui.SetStudioModeEnabledParams{StudioModeEnabled: &enabled})
	return err
}

func disableStudioMode() error {
	return setStudioModeEnabled(false)
}

func enableStudioMode() error {
	return setStudioModeEnabled(true)
}

// Determine if the studio mode is currently enabled in OBS.
func IsStudioModeEnabled() (bool, error) {
	r, err := client.Ui.GetStudioModeEnabled()
	return r.StudioModeEnabled, err
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
	enabled, err := IsStudioModeEnabled()
	if err != nil {
		return err
	}
	err = setStudioModeEnabled(!enabled)
	return err
}

func transitionToProgram() error {
	_, err := client.Transitions.TriggerStudioModeTransition(&transitions.TriggerStudioModeTransitionParams{})
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
