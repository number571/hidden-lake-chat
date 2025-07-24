package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/number571/hidden-lake-chat/internal/app"
	"github.com/number571/hidden-lake/build"
	"github.com/number571/hidden-lake/pkg/utils/flag"
)

const (
	appVersion = "v0.0.1"
	helpInfo   = `<Hidden Lake Chat>
Description: anonymous console group chat
Arguments:
[ -v, --version ] = print version of service
[ -h, --help ] = print information about service
[ -p, --path ] = set path to config, database files
[ -n, --network ] = set network key from build`
)

var (
	gFlags = flag.NewFlagsBuilder(
		flag.NewFlagBuilder("-v", "--version"),
		flag.NewFlagBuilder("-h", "--help"),
		flag.NewFlagBuilder("-p", "--path").WithDefinedValue("."),
		flag.NewFlagBuilder("-n", "--network").WithDefinedValue(build.CDefaultNetwork),
	).Build()
)

func main() {
	args := os.Args[1:]
	if ok := gFlags.Validate(args); !ok {
		panic("args invalid")
	}

	if gFlags.Get("-v").GetBoolValue(args) {
		fmt.Println(appVersion)
		return
	}

	if gFlags.Get("-h").GetBoolValue(args) {
		fmt.Println(helpInfo)
		return
	}

	app, err := app.InitApp(args, gFlags)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	if err := app.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
