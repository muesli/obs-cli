package main

import (
	"fmt"
	"strconv"

	"github.com/muesli/coral"
)

var (
	recordingCmd = &coral.Command{
		Use:   "recording",
		Short: "manage recordings",
		Long:  `The recording command manages recordings`,
		RunE:  nil,
	}

	startStopRecordingCmd = &coral.Command{
		Use:   "toggle",
		Short: "Toggle recording",
		RunE: func(cmd *coral.Command, args []string) error {
			return startStopRecording()
		},
	}

	startRecordingCmd = &coral.Command{
		Use:   "start",
		Short: "Starts recording",
		RunE: func(cmd *coral.Command, args []string) error {
			return startRecording()
		},
	}

	stopRecordingCmd = &coral.Command{
		Use:   "stop",
		Short: "Stops recording",
		RunE: func(cmd *coral.Command, args []string) error {
			return stopRecording()
		},
	}

	pauseRecordingCmd = &coral.Command{
		Use:   "pause",
		Short: "manage paused state",
	}

	enablePauseRecordingCmd = &coral.Command{
		Use:   "enable",
		Short: "Pause recording",
		RunE: func(cmd *coral.Command, args []string) error {
			return pauseRecording()
		},
	}

	resumePauseRecordingCmd = &coral.Command{
		Use:   "resume",
		Short: "Resume recording",
		RunE: func(cmd *coral.Command, args []string) error {
			return resumeRecording()
		},
	}

	togglePauseRecordingCmd = &coral.Command{
		Use:   "toggle",
		Short: "Pause/resume recording",
		RunE: func(cmd *coral.Command, args []string) error {
			return pauseResumeRecording()
		},
	}

	recordingStatusCmd = &coral.Command{
		Use:   "status",
		Short: "Reports recording status",
		RunE: func(cmd *coral.Command, args []string) error {
			return recordingStatus()
		},
	}
)

func startStopRecording() error {
	_, err := client.Record.ToggleRecord()
	return err
}

func startRecording() error {
	_, err := client.Record.StartRecord()
	return err
}

func stopRecording() error {
	_, err := client.Record.StopRecord()
	return err
}

func pauseRecording() error {
	_, err := client.Record.PauseRecord()
	return err
}

func resumeRecording() error {
	_, err := client.Record.ResumeRecord()
	return err
}

func pauseResumeRecording() error {
	r, err := client.Record.GetRecordStatus()
	if err != nil {
		return err
	}
	if !r.OutputActive {
		return fmt.Errorf("recording is not running")
	}

	if r.OuputPaused {
		return resumeRecording()
	}
	return pauseRecording()
}

func recordingStatus() error {
	r, err := client.Record.GetRecordStatus()
	if err != nil {
		return err
	}

	fmt.Printf("Recording: %s\n", strconv.FormatBool(r.OutputActive))
	if !r.OutputActive {
		return nil
	}

	fmt.Printf("Paused: %s\n", strconv.FormatBool(r.OuputPaused))
	// TODO: see if recording filename is available from another API method
	// fmt.Printf("File: %s\n", r.RecordingFilename)
	fmt.Printf("Timecode: %s\n", r.OutputTimecode)

	// st, err := os.Stat(r.RecordingFilename)
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("Filesize: %s\n", humanize.Bytes(uint64(st.Size())))

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
