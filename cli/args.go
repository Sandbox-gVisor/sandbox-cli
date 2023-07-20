package cli

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
)

func ParseCli() {
	parser := argparse.NewParser("sandbox-cli", "tool for in-time configuraion gVisor")
	address := parser.String("a", "address", &argparse.Options{Required: true, Help: "Socket address"})
	changeCmd := parser.NewCommand("change", "Change callbacks")
	infoCmd := parser.NewCommand("info", "Show info")
	stateCmd := parser.NewCommand("state", "Change state")

	changeFile := changeCmd.String("c", "conf", &argparse.Options{Required: true, Help: "file with config"})
	stateFile := stateCmd.String("c", "conf", &argparse.Options{Required: true, Help: "file with config"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	if changeCmd.Happened() {
		sendChange(*address, ReadDtos(ReadFile(*changeFile)))
	} else if stateCmd.Happened() {
		sendState(*address, ReadDtos(ReadFile(*stateFile)))
	} else if infoCmd.Happened() {
		sendInfo(*address)
	} else {
		err := fmt.Errorf("Bad arguments, check usage")
		fmt.Print(parser.Usage(err))
	}
}
