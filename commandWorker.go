package main

import (
	"fmt"
	"bytes"
	"os"
	"os/exec"
)


//initCommandString generates the necessary flags for the notify-send command
func initCommandString() string {
	return fmt.Sprintf("notify-send --urgency=critical --expire-time=3000 ")
}


//requires absolute path to current working directory
func setIcon(curr_cmd string, icon_path string) string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("failed to generate icon path")
		return ""
	}
	return fmt.Sprintf(" %s -i %s/%s ", curr_cmd, pwd, icon_path)
}	

func setTitle(curr_cmd, title string) string {
	return fmt.Sprintf(" %s \" %s \"", curr_cmd, title)
}

func setBody(curr_cmd, body string) string {
	return fmt.Sprintf(" %s \" %s \"", curr_cmd, body)
}



// runShellCommand executes the given arguments through the /bin/sh environment.
func runShellCommand(args string) error {
	var out bytes.Buffer    //save command output
	var stderr bytes.Buffer //save command errors
	cmd := exec.Command("/bin/sh", "-c", args)
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()

	return err
}


