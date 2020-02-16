package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	obsws "github.com/muesli/go-obs-websocket"
	"github.com/spf13/cobra"
)

var (
	host string
	port uint32

	rootCmd = &cobra.Command{
		Use:   "obs-cli",
		Short: "obs-cli is a command-line remote control for OBS",
	}

	client *obsws.Client
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if client != nil {
		client.Disconnect()
	}
}

func init() {
	cobra.OnInitialize(connectOBS)
	rootCmd.PersistentFlags().StringVar(&host, "host", "localhost", "host to connect to")
	rootCmd.PersistentFlags().Uint32VarP(&port, "port", "p", 4444, "port to connect to")
}

func connectOBS() {
	// disable obsws logging
	obsws.Logger = log.New(ioutil.Discard, "", log.LstdFlags)

	client = &obsws.Client{Host: host, Port: int(port)}
	if err := client.Connect(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Set the amount of time we can wait for a response.
	obsws.SetReceiveTimeout(time.Second * 2)
}
