package command

import (
    "flag"
    "fmt"
    "github.com/robn/spillway/db"
    "github.com/mitchellh/cli"
    "strings"
)

type UserCommand struct {
    Ui cli.Ui
}

func (c *UserCommand) Help() string {
    helpText := `
Usage: spillway user [options] name ...

  Tools to manipulate the user database

Options:

  -add  Add (create) users
`

    return strings.TrimSpace(helpText)
}

// XXX this would be better
//   spillway user add
//   spillway user remove
func (c *UserCommand) Run(args []string) int {
    cmdFlags := flag.NewFlagSet("user", flag.ContinueOnError)
    cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }

    var add bool
    cmdFlags.BoolVar(&add, "add", false, "add")

    if err := cmdFlags.Parse(args); err != nil {
        return 1
    }

    names := cmdFlags.Args()
    if len(names) == 0 {
        c.Ui.Error("At least one user name must be specified.")
        c.Ui.Error("")
        c.Ui.Error(c.Help())
        return 1
    }

    if add {
        var dbh *db.UserDB
        var err error
        if dbh, err = db.OpenUserDB(); err != nil {
            c.Ui.Error(fmt.Sprintf("open user database failed: %s", err))
            return 1
        }

        for _, name := range names {
            user := db.NewUser(name)
            if err := dbh.Save(&user); err != nil {
                c.Ui.Error(fmt.Sprintf("add user '%s' failed: %s", name, err))
            }
        }
    }

    return 0
}

func (c *UserCommand) Synopsis() string {
    return "Tools to manipulate the user database"
}
