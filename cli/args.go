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
	getCmd := parser.NewCommand("get", "Get current callbacks")

	changeFile := changeCmd.String("c", "conf", &argparse.Options{Required: true, Help: "file with config"})
	entryPoint := stateCmd.String("e", "entry_point", &argparse.Options{Required: true, Help: "Entry point"})
	stateFile := stateCmd.String("c", "conf", &argparse.Options{Required: true, Help: "file with source"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	if changeCmd.Happened() {
		fmt.Println(sendChange(*address, ReadDtos(ReadFile(*changeFile))))
	} else if stateCmd.Happened() {
		fmt.Println(sendState(*address, *entryPoint, ReadSource(ReadFile(*stateFile))))
	} else if infoCmd.Happened() {
		fmt.Println(sendInfo(*address))
	} else if getCmd.Happened() {
		fmt.Println(sendGet(*address))
	} else {
		err := fmt.Errorf("bad arguments, check usage")
		fmt.Print(parser.Usage(err))
	}
}
