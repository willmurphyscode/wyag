package libwyag

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/subcommands"
)

// InitCmd represents a command go initialize an empty git repo
type InitCmd struct {
	path string
}

func (*InitCmd) Name() string     { return "init" }
func (*InitCmd) Synopsis() string { return "Initialize and empty git repository" }
func (*InitCmd) Usage() string {
	return `init [PATH=PWD]:
Initialize an empty git repository at PATH
`
}

func (i *InitCmd) SetFlags(f *flag.FlagSet) {
}

func (i *InitCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {

	pwd, err := os.Getwd()
	orDie(err)
	args := f.Args()
	if len(args) == 0 {
		i.path = pwd
	} else {
		i.path = filepath.Join(pwd, args[0])

	}
	fmt.Println(i.path)
	fmt.Println()
	return subcommands.ExitSuccess
}

func orDie(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
