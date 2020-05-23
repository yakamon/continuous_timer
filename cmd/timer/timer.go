package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	cmdName, interval, unit, unitStr := ParseArgs()

	message := fmt.Sprintf("You will be notified every %d %s.", interval, unitStr)
	TerminalNotifierCommand(cmdName, message, "Bosso").Run()

	var cumDur time.Duration

	for {
		time.Sleep(interval * unit)
		cumDur += interval

		message = fmt.Sprintf("%d %s have elapsed.", cumDur, unitStr)
		TerminalNotifierCommand(cmdName, message, "Glass").Run()
	}
}

// ParseArgs parses command line args.
func ParseArgs() (cmdName string, interval time.Duration, unit time.Duration, unitStr string) {
	cmd := os.Args[0]
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <interval> <s [seconds,sec] | m [minutes,min] | h [hours]>\n", cmd)
		os.Exit(0)
	}

	if i, err := strconv.Atoi(os.Args[1]); err != nil {
		fmt.Printf("Interval should be an interger: %v\n", i)
		os.Exit(1)
	} else {
		interval = time.Duration(i)
	}

	switch u := os.Args[2]; u {
	case "s", "sec", "seconds":
		unit = time.Second
		unitStr = "seconds"
	case "m", "min", "minutes":
		unit = time.Minute
		unitStr = "minutes"
	case "h", "hours":
		unit = time.Hour
		unitStr = "hours"
	default:
		fmt.Printf("Time unit should be either seconds, minutes, or hours: %v\n", u)
		os.Exit(1)
	}

	return cmdName, interval, unit, unitStr
}

// TerminalNotifierCommand initializes "terminal-notifier" command.
func TerminalNotifierCommand(cmdName, message, sound string) *exec.Cmd {
	args := []string{
		"-title", cmdName,
		"-message", message,
		"-sound", sound,
	}
	return exec.Command("terminal-notifier", args...)
}
