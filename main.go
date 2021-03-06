package main

import (
	"fmt"
	"os"

	"github.com/nethings/internal-api/cmd"
)

var (
	buildVersion string
)

func main() {
	cmd.RootCmd.Version = buildVersion

	fmt.Println("1++++++++++++")

	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
