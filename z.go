package main

import (
	"fmt"
	"os"
	"github.com/pranavnt/mamba"
)

func main() {
	app := mamba.New()

	app.AddCommand("install {packageName}", installPackage)

	app.Run(os.Args)
}

func installPackage(params mamba.Dict) {
	packageName := params["packageName"]
	
	fmt.Println(packageName)
}

