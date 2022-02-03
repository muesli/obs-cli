package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/muesli/coral"
)

var countdownCmd = &coral.Command{
	Use:   "countdown",
	Short: "Triggers a countdown and continuously updates a label with the remaining time",
	RunE: func(cmd *coral.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("countdown requires a label and the countdown in seconds")
		}

		d, err := time.ParseDuration(args[1])
		if err != nil {
			return err
		}

		return countdown(args[0], d)
	},
}

func countdown(label string, duration time.Duration) error {
	until := time.Now().Add(duration).Add(time.Second)

	c := time.Tick(time.Second)
	for range c {
		rem := time.Until(until)
		if rem < 0 {
			rem = 0
		}
		if err := changeLabel(label, fmtDuration(rem)); err != nil {
			return err
		}

		if time.Now().After(until) {
			break
		}
	}

	return nil
}

func fmtDuration(d time.Duration) string {
	d = d.Round(time.Second)
	m := d % time.Hour / time.Minute
	s := d % time.Minute / time.Second
	return fmt.Sprintf("%02d:%02d", m, s)
}

func init() {
	labelCmd.AddCommand(countdownCmd)
}
