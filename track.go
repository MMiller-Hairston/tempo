package tempo

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type TrackCommand struct{}

const (
	// storage  = "/opt/ds" TODO: Move to Config
	file_ext = "_time.log"
)

func (c *TrackCommand) Help() string {
	return "Track a time event"
}

func getToday() time.Time {
	return time.Now().Local()
}

func formatDate(t time.Time) string {
	return fmt.Sprintf("%d%d%d", t.Year(), t.Month(), t.Day())
}

func formatTime(t time.Time) string {
	return fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())
}

func (c *TrackCommand) Run(args []string) int {
	t := getToday()
	p := formatDate(t)
	tm := formatTime(t)

	f, err := os.OpenFile(storage+"/"+p+file_ext, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}

	defer f.Close()

	w := io.MultiWriter(os.Stdout, f)
	log.SetFlags(0)
	log.SetOutput(w)
	log.Printf("[%s] %s", tm, args[0])
	return 0
}

func (c *TrackCommand) Synopsis() string {
	return "Keep track of your day to day. Add an event to the timelog for recollection at a later time."
}
