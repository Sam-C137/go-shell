package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	user2 "os/user"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		pref, err := getShellPrefix()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(pref)
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
		}
		if err = run(strings.TrimSuffix(text, "\n")); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
		}
	}
}

func run(input string) error {
	fields := strings.Fields(input)
	if len(fields) < 1 {
		return nil
	}

	switch fields[0] {
	case "cd":
		if len(fields) < 2 {
			return errors.New("directory must be specified")
		}
		os.Chdir(fields[1])
		return nil
	case "exit":
		os.Exit(-1)
	default:
		var args []string
		if len(fields) > 1 {
			args = fields[1:]
		}

		cmd := exec.Command(fields[0], args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		return cmd.Run()
	}

	return nil
}

func getShellPrefix() (string, error) {
	var out string
	dir, err := os.Getwd()
	if err != nil {
		return out, err
	}

	hostname, err := os.Hostname()
	if err != nil {
		return out, err
	}

	user, err := user2.Current()
	if err != nil {
		return out, err
	}

	out = fmt.Sprintf("%s@%s %s > ", user.Username, hostname, dir)
	return out, nil
}
