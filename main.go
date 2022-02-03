package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/andreykaipov/goobs"
	"github.com/muesli/coral"
)

var (
	host     string
	password string
	port     uint32
	version  string

	rootCmd = &coral.Command{
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
	coral.OnInitialize(connectOBS)
	rootCmd.PersistentFlags().StringVar(&host, "host", "localhost", "host to connect to")
	rootCmd.PersistentFlags().StringVar(&password, "password", "", "password for connection")
	rootCmd.PersistentFlags().Uint32VarP(&port, "port", "p", 4444, "port to connect to")
}

func getUserAgent() string {
	userAgent := "obs-cli"
	if version != "" {
		userAgent += "/" + version
	}
	return userAgent
}

func connectOBS() {
	var err error
	client, err = goobs.New(
		host+fmt.Sprintf(":%d", port),
		goobs.WithPassword(password),
		goobs.WithRequestHeader(http.Header{"User-Agent": []string{getUserAgent()}}),
	)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
