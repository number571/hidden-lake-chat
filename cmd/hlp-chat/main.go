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
	"github.com/number571/hidden-lake-chat/internal/settings"
	"github.com/number571/hidden-lake/build"
	"github.com/number571/hidden-lake/pkg/utils/flag"
	"github.com/number571/hidden-lake/pkg/utils/help"
)

var (
	gFlags = flag.NewFlagsBuilder(
		flag.NewFlagBuilder("-v", "--version").
			WithDescription("print version of service"),
		flag.NewFlagBuilder("-h", "--help").
			WithDescription("print information about service"),
		flag.NewFlagBuilder("-p", "--path").
			WithDescription("set path to config, database files").
			WithDefinedValue("."),
		flag.NewFlagBuilder("-n", "--network").
			WithDescription("set network key of connections from build").
			WithDefinedValue(build.CDefaultNetwork),
	).Build()
)

func main() {
	args := os.Args[1:]
	if ok := gFlags.Validate(args); !ok {
		panic("args invalid")
	}

	if gFlags.Get("-v").GetBoolValue(args) {
		fmt.Println(build.GetVersion())
		return
	}

	if gFlags.Get("-h").GetBoolValue(args) {
		help.Println(settings.GetAppName(), settings.CProjectDescription, gFlags)
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
