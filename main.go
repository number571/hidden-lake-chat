package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/number571/hidden-lake/build"
	"github.com/number571/hl-chat/internal/app"
)

const (
	appVersion     = "v0.0.1"
	defaultPath    = "."
	defaultNetwork = build.CDefaultNetwork
)

var (
	v *bool
	h *bool
	p *string
	n *string
)

func init() {
	v = flag.Bool("version", false, "print version of application")
	flag.BoolVar(v, "v", *v, "alias for -version")

	h = flag.Bool("help", false, "print information about application")
	flag.BoolVar(h, "h", *h, "alias for -help")

	p = flag.String("path", defaultPath, "set path to config, database files")
	flag.StringVar(p, "p", *p, "alias for -path")

	n = flag.String("network", defaultNetwork, "set network key from build")
	flag.StringVar(n, "n", *n, "alias for -network")

	flag.Parse()
}

func main() {
	if *v {
		fmt.Println(appVersion)
		return
	}

	if *h {
		flag.Usage()
		return
	}

	app, err := app.InitApp(*p, *n)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	if err := app.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
