package main

import (
	"context"
	"flag"
	"fmt"
	outer "github.com/mar-heaven/gocli-learn"
	internal "github.com/mar-heaven/gocli-learn/internal/gocli"
	"os"
	"strings"

	"github.com/google/subcommands"
)

type printCmd struct {
	capitalize bool
}

func (*printCmd) Name() string { return "print" }
func (*printCmd) Synopsis() string {
	// test package
	internalRsp := internal.ReturnInternal()
	outRsp := outer.ReturnOuter()
	fmt.Println(internalRsp)
	fmt.Println(outRsp)
	return "Print args to stdout."
}
func (*printCmd) Usage() string {
	return `print [-capitalize] <some text>:
  Print args to stdout.
`
}

func (p *printCmd) SetFlags(f *flag.FlagSet) {
	f.BoolVar(&p.capitalize, "capitalize", false, "capitalize output")
}

func (p *printCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	for _, arg := range f.Args() {
		if p.capitalize {
			arg = strings.ToUpper(arg)
		}
		fmt.Printf("%s ", arg)
	}
	fmt.Println()
	return subcommands.ExitSuccess
}

func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&printCmd{}, "")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
