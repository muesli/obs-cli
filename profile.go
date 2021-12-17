package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/andreykaipov/goobs/api/requests/profiles"
	"github.com/spf13/cobra"
)

var (
	profileCmd = &cobra.Command{
		Use:   "profile",
		Short: "manage profiles",
		Long:  `The profile command manages profiles`,
		RunE:  nil,
	}

	listProfileCmd = &cobra.Command{
		Use:   "list",
		Short: "List all profiles",
		RunE: func(cmd *cobra.Command, args []string) error {
			return listProfiles()
		},
	}

	getProfileCmd = &cobra.Command{
		Use:   "get",
		Short: "Get the current profile",
		RunE: func(cmd *cobra.Command, args []string) error {
			return getProfile()
		},
	}

	setProfileCmd = &cobra.Command{
		Use:   "set",
		Short: "Set the current profile",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("set requires a profile name as argument")
			}
			return setProfile(strings.Join(args, " "))
		},
	}
)

func listProfiles() error {
	r, err := client.Profiles.ListProfiles()
	if err != nil {
		return err
	}

	for _, v := range r.Profiles {
		fmt.Println(v.ProfileName)
	}
	return nil
}

func setProfile(profile string) error {
	r := profiles.SetCurrentProfileParams{
		ProfileName: profile,
	}
	_, err := client.Profiles.SetCurrentProfile(&r)
	return err
}

func getProfile() error {
	r, err := client.Profiles.GetCurrentProfile()
	if err != nil {
		return err
	}

	fmt.Println(r.ProfileName)
	return nil
}

func init() {
	profileCmd.AddCommand(listProfileCmd)
	profileCmd.AddCommand(setProfileCmd)
	profileCmd.AddCommand(getProfileCmd)
	rootCmd.AddCommand(profileCmd)
}
