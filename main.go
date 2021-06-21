package main

import (
	"fmt"
	"os"

	"github.com/andreykaipov/goobs"
	"github.com/spf13/cobra"
)

var (
	host     string
	password string
	port     uint32

	rootCmd = &cobra.Command{
		Use:   "obs-cli",
		Short: "obs-cli is a command-line remote control for OBS",
	}

	client *goobs.Client
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if client != nil {
		_ = client.Disconnect()
	}
}

func init() {
	cobra.OnInitialize(connectOBS)
	rootCmd.PersistentFlags().StringVar(&host, "host", "localhost", "host to connect to")
	rootCmd.PersistentFlags().StringVar(&password, "password", "", "password for connection")
	rootCmd.PersistentFlags().Uint32VarP(&port, "port", "p", 4444, "port to connect to")
}

func connectOBS() {
	var err error
	client, err = goobs.New(host+fmt.Sprintf(":%d", port), goobs.WithPassword(password))
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
