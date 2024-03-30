package internal

import (
	"errors"
	"fmt"
	"github.com/akamensky/argparse"
	"os"
	"sandbox-cli/internal/commands"
	"sandbox-cli/internal/communication"
	models2 "sandbox-cli/internal/pretty_output"
)

const AddressArgv = "CLI_ADDRESS"

func replaceAddressWithSavedInArgv(address *string) error {
	if *address == "" {
		argvAddress := os.Getenv(AddressArgv)
		if argvAddress == "" {
			return errors.New("gvisor sock address required")
		}
		*address = argvAddress
	}

	return nil
}

func ParseCli() {
	parser := argparse.NewParser("sandbox-cli", "tool for in-time configuraion gVisor")

	address := parser.String("a", "address", &argparse.Options{Required: false, Help: "Socket address"})
	infoCmd := parser.NewCommand("man", "Show man for accessors")
	stateCmd := parser.NewCommand("state", "Change state")
	getCmd := parser.NewCommand("get", "Get current callbacks")
	deleteCmd := parser.NewCommand("delete", "Unregister callbacks")

	verboseFlag := getCmd.Flag("v", "verbose", &argparse.Options{Required: false, Help: "Verbose output"})

	stateFile := stateCmd.String("c", "conf", &argparse.Options{Required: true, Help: "file with source"})

	deleteAll := deleteCmd.Flag("u", "all", &argparse.Options{Required: false, Help: "Unregister all callbacks"})
	sysno := deleteCmd.Int("s", "sysno", &argparse.Options{Required: false, Help: "Callback sysno"})
	callbackType := deleteCmd.String("t", "type", &argparse.Options{Required: false, Help: "Callback type"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Printf("\n%s\n\n%s",
			models2.MakeTextBoldAndColored("Bad arguments, check usage", models2.RedColorText), parser.Usage(err))
		return
	}

	err = replaceAddressWithSavedInArgv(address)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	var responseHandler models2.ResponseFormatter = &models2.DefaultResponseFormatter{}
	var request *communication.Request

	if stateCmd.Happened() {
		request = commands.MakeChangeStateRequest(*stateFile)
	} else if infoCmd.Happened() {
		request = commands.MakeAccessorsInfoRequest()
		responseHandler = commands.AccessorInfoResponseHandler()
	} else if getCmd.Happened() {
		request = commands.MakeGetCallbacksRequest()
		responseHandler = commands.GetCallbackResponseHandler(*verboseFlag)
	} else if deleteCmd.Happened() {
		if *deleteAll {
			request = commands.MakeDeleteCallbacksRequest("all", *sysno, *callbackType)
		} else {
			request = commands.MakeDeleteCallbacksRequest("list", *sysno, *callbackType)
		}
	}

	response, err := communication.SendRequest(*address, request)
	if err != nil {
		fmt.Printf("\nError: %s\n\n", models2.MakeTextBoldAndColored(err.Error(), models2.RedColorText))
	} else {
		fmt.Printf("%v\n", responseHandler.Format(response))
	}
}
