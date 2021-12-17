package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/dustin/go-humanize"
	"github.com/spf13/cobra"
)

var (
	recordingCmd = &cobra.Command{
		Use:   "recording",
		Short: "manage recordings",
		Long:  `The recording command manages recordings`,
		RunE:  nil,
	}

	startStopRecordingCmd = &cobra.Command{
		Use:   "toggle",
		Short: "Toggle recording",
		RunE: func(cmd *cobra.Command, args []string) error {
			return startStopRecording()
		},
	}

	startRecordingCmd = &cobra.Command{
		Use:   "start",
		Short: "Starts recording",
		RunE: func(cmd *cobra.Command, args []string) error {
			return startRecording()
		},
	}

	stopRecordingCmd = &cobra.Command{
		Use:   "stop",
		Short: "Stops recording",
		RunE: func(cmd *cobra.Command, args []string) error {
			return stopRecording()
		},
	}

	pauseRecordingCmd = &cobra.Command{
		Use:   "pause",
		Short: "manage paused state",
	}

	enablePauseRecordingCmd = &cobra.Command{
		Use:   "enable",
		Short: "Pause recording",
		RunE: func(cmd *cobra.Command, args []string) error {
			return pauseRecording()
		},
	}

	resumePauseRecordingCmd = &cobra.Command{
		Use:   "resume",
		Short: "Resume recording",
		RunE: func(cmd *cobra.Command, args []string) error {
			return resumeRecording()
		},
	}

	togglePauseRecordingCmd = &cobra.Command{
		Use:   "toggle",
		Short: "Pause/resume recording",
		RunE: func(cmd *cobra.Command, args []string) error {
			return pauseResumeRecording()
		},
	}

	recordingStatusCmd = &cobra.Command{
		Use:   "status",
		Short: "Reports recording status",
		RunE: func(cmd *cobra.Command, args []string) error {
			return recordingStatus()
		},
	}
)

func startStopRecording() error {
	_, err := client.Recording.StartStopRecording()
	return err
}

func startRecording() error {
	_, err := client.Recording.StartRecording()
	return err
}

func stopRecording() error {
	_, err := client.Recording.StopRecording()
	return err
}

func pauseRecording() error {
	_, err := client.Recording.PauseRecording()
	return err
}

func resumeRecording() error {
	_, err := client.Recording.ResumeRecording()
	return err
}

func pauseResumeRecording() error {
	r, err := client.Recording.GetRecordingStatus()
	if err != nil {
		return err
	}
	if !*r.IsRecording {
		return fmt.Errorf("recording is not running")
	}

	if *r.IsRecordingPaused {
		return resumeRecording()
	}
	return pauseRecording()
}

func recordingStatus() error {
	r, err := client.Recording.GetRecordingStatus()
	if err != nil {
		return err
	}

	fmt.Printf("Recording: %s\n", strconv.FormatBool(*r.IsRecording))
	if !*r.IsRecording {
		return nil
	}

	fmt.Printf("Paused: %s\n", strconv.FormatBool(*r.IsRecordingPaused))
	fmt.Printf("File: %s\n", r.RecordingFilename)
	fmt.Printf("Timecode: %s\n", r.RecordTimecode)

	st, err := os.Stat(r.RecordingFilename)
	if err != nil {
		return err
	}
	fmt.Printf("Filesize: %s\n", humanize.Bytes(uint64(st.Size())))

	return nil
}

func init() {
	pauseRecordingCmd.AddCommand(enablePauseRecordingCmd)
	pauseRecordingCmd.AddCommand(resumePauseRecordingCmd)
	pauseRecordingCmd.AddCommand(togglePauseRecordingCmd)

	recordingCmd.AddCommand(startStopRecordingCmd)
	recordingCmd.AddCommand(startRecordingCmd)
	recordingCmd.AddCommand(stopRecordingCmd)
	recordingCmd.AddCommand(pauseRecordingCmd)
	recordingCmd.AddCommand(recordingStatusCmd)

	rootCmd.AddCommand(recordingCmd)
}
