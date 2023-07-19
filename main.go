package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"os"
)

func main() {
	parser := argparse.NewParser("sandbox-cli", "tool for in-time configuraion gVisor")
	fuckFlag := parser.Flag("f", "force", &argparse.Options{Required: false, Help: "Send fuck json"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	if *fuckFlag {
		message := Message{Data: "hi", Type: "log"}
		SendToSocker("127.0.0.1", 3333, message.ToString())
	}

	fmt.Println(*fuckFlag)
}
