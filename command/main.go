package main

import (
	"errors"
	"fmt"
	"os"

	"command/modules/photo"
	"command/types/cmm"
)

func main() {
	if err := cmd(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func cmd(args []string) error {
	if len(args) < 1 {
		return errors.New("no sub-command found")
	}

	cmds := []cmm.Subcommand{
		photo.NewPhoto(),
	}

	subcommand := os.Args[1]
	for _, cmd := range cmds {
		if cmd.Name() == subcommand {
			if err := cmd.Parse(os.Args[2:]); err != nil {
				return fmt.Errorf("parse command error: %s", err.Error())
			}
			return cmd.Run()
		}
	}

	return fmt.Errorf("unknown subcommand: %s", subcommand)
}
