package command

import (
	"log"

	"github.com/philproctor/confluence-cli/client"
)

//OperationOptions holds all the options that apply to the specified operation
type OperationOptions struct {
	Title         string
	SpaceKey      string
	Filepath      string
	BodyOnly      bool
	StripImgs     bool
	AncestorTitle string
	AncestorID    int64
}

var options *OperationOptions

var restClient *client.ConfluenceClient

//Run the provided command with the options from the command line
func Run(command string, config *client.ConfluenceConfig, opts *OperationOptions) {

	options = opts

	switch command {
	case "add-page":
		prepareClient(config)
		addPage()

	case "find-page":
		prepareClient(config)
		findPage()

	case "help", "":
		printUsage()

	default:
		log.Fatal("Unknown Command: ", command)
	}

}

func prepareClient(config *client.ConfluenceConfig) {
	validateClientDetails(config)
	restClient = client.Client(config)
	processFlags()
}

func processFlags() {

}