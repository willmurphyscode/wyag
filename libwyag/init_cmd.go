package libwyag

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"github.com/google/subcommands"
)

// InitCmd represents a command go initialize an empty git repo
type InitCmd struct {
	capitalize bool
}

func (*InitCmd) Name() string     { return "print" }
func (*InitCmd) Synopsis() string { return "Print args to stdout." }
func (*InitCmd) Usage() string {
	return `print [-capitalize] <some text>:
  Print args to stdout.
`
}

func (p *InitCmd) SetFlags(f *flag.FlagSet) {
	f.BoolVar(&p.capitalize, "capitalize", false, "capitalize output")
}

func (p *InitCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	for _, arg := range f.Args() {
		if p.capitalize {
			arg = strings.ToUpper(arg)
		}
		fmt.Printf("%s ", arg)
	}
	fmt.Println()
	return subcommands.ExitSuccess
}
