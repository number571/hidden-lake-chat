package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

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
[ -p, --path ] = set path to database file
[ -n, --network ] = set network key from build`
)

var (
	gFlags = flag.NewFlagsBuilder(
		flag.NewFlagBuilder("-v", "--version"),
		flag.NewFlagBuilder("-h", "--help"),
		flag.NewFlagBuilder("-p", "--path").WithDefinedValue("hidden-lake-chat.db"),
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

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	closed := make(chan struct{})
	defer func() {
		cancel()
		<-closed
	}()

	go func() {
		defer func() { closed <- struct{}{} }()
		if err := app.Run(ctx); err != nil && !errors.Is(err, context.Canceled) {
			log.Fatal(err)
		}
	}()

	<-shutdown
}
