package main

import (
    "github.com/robn/spillway/command"
    "github.com/mitchellh/cli"
    "os"
)

var Commands map[string]cli.CommandFactory

func init() {
    ui := &cli.BasicUi{Writer: os.Stdout}

    Commands = map[string]cli.CommandFactory{
        "user": func() (cli.Command, error) {
            return &command.UserCommand{
                Ui: ui,
            }, nil
        },
    }
}
