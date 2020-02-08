package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/danilovalente/golangspell-core/usecase"
	"github.com/spf13/cobra"
)

func init() {
	RunCommandFunctions["init"] = runInitCommand
}

//Module name to initialize with 'Go Modules'
var Module string

//AppName used to define the application's directory and the default value to the config variable with the same name
var AppName string

func runInitCommand(cmd *cobra.Command, args []string) {
	err := usecase.RenderInitTemplate(args)
	if err != nil {
		fmt.Printf("An error occurred while trying to create the application. Error: %s\n", err.Error())
		return
	}
	execCmd := exec.Command("go", "mod", "init", args[0])
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr
	err = execCmd.Run()
	if err != nil {
		fmt.Printf("An error occurred while trying to create the application. Error: %s\n", err.Error())
	}
	fmt.Println("Application created!")
}
