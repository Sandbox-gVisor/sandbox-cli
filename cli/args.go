package cli

import (
	"errors"
	"fmt"
	"github.com/akamensky/argparse"
	"os"
	"sandbox-cli/models"
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
	changeCmd := parser.NewCommand("change", "Change callbacks")
	infoCmd := parser.NewCommand("info", "Show info")
	stateCmd := parser.NewCommand("state", "Change state")
	getCmd := parser.NewCommand("get", "Get current callbacks")
	deleteCmd := parser.NewCommand("delete", "Unregister callbacks")

	verboseFlag := getCmd.Flag("v", "verbose", &argparse.Options{Required: false, Help: "Verbose output"})

	changeFile := changeCmd.String("c", "conf", &argparse.Options{Required: true, Help: "file with config"})
	entryPoint := stateCmd.String("e", "entry_point", &argparse.Options{Required: true, Help: "Entry point"})
	stateFile := stateCmd.String("c", "conf", &argparse.Options{Required: true, Help: "file with source"})

	deleteAll := deleteCmd.Flag("u", "all", &argparse.Options{Required: false, Help: "Unregister all callbacks"})
	sysno := deleteCmd.Int("s", "sysno", &argparse.Options{Required: false, Help: "Callback sysno"})
	callbackType := deleteCmd.String("t", "type", &argparse.Options{Required: false, Help: "Callback type"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Printf("\n%s\n\n%s",
			models.MakeTextBoldAndColored("Bad arguments, check usage", models.RedColorText), parser.Usage(err))
		return
	}

	err = replaceAddressWithSavedInArgv(address)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	var responseHandler models.ResponseHandler = &models.DefaultResponseHandler{}
	var request *models.Request

	if changeCmd.Happened() {
		request = models.MakeChangeCallbacksRequest(*changeFile)
	} else if stateCmd.Happened() {
		request = models.MakeChangeStateRequest(*entryPoint, *stateFile)
	} else if infoCmd.Happened() {
		request = models.MakeHookInfoRequest()
		responseHandler = models.HooksInfoResponseHandler()
	} else if getCmd.Happened() {
		request = models.MakeGetCallbacksRequest()
		responseHandler = models.GetCallbackResponseHandler(*verboseFlag)
	} else if deleteCmd.Happened() {
		if *deleteAll {
			request = models.MakeDeleteCallbacksRequest("all", *sysno, *callbackType)
		} else {
			request = models.MakeDeleteCallbacksRequest("list", *sysno, *callbackType)
		}
	}

	response, err := models.SendRequest(*address, request)
	if err != nil {
		fmt.Printf("\nError: %s\n\n", models.MakeTextBoldAndColored(err.Error(), models.RedColorText))
	} else {
		responseHandler.Handle(response)
	}
}
