package libwyag

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/google/subcommands"
)

// InitCmd represents a command go initialize an empty git repo
type InitCmd struct {
	path string
}

func (*InitCmd) Name() string     { return "init" }
func (*InitCmd) Synopsis() string { return "Initialize and empty git repository" }
func (*InitCmd) Usage() string {
	return `init [-path PATH=PWD]:
Initialize an empty git repository at PATH
`
}

func (i *InitCmd) SetFlags(f *flag.FlagSet) {
	// f.BoolVar(&p.capitalize, "capitalize", false, "capitalize output")
	pwd, err := os.Getwd()
	orDie(err)
	f.StringVar(&i.path, "path", pwd, "Path where the repository should be initialized")
}

func (i *InitCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
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
