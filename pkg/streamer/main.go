package main

import (
	"fmt"
	"github.com/bojand/ghz/printer"
	"github.com/bojand/ghz/runner"
	"os"
)

func main() {
	report, err := runner.Run(
		"GHZExample.Greeter.SayHello",
		"localhost:50052",
		runner.WithProtoFile("./internal/proto_db/proto_db.proto", []string{}),
		runner.WithDataFromFile("./pkg/streamer/data.json"),
		runner.WithInsecure(true),
	)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	p := printer.ReportPrinter{
		Out:    os.Stdout,
		Report: report,
	}

	if err2 := p.Print("pretty"); err2 != nil {
		fmt.Println(err2)
	}
}
