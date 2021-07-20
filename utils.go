package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func getPackageJSON() PackageJSON {
	data, err := ioutil.ReadFile("./package.json")

	if err != nil {
		fmt.Println(err)
	}

	var packageJSON PackageJSON

	err = json.Unmarshal(data, &packageJSON)

	if err != nil {
		fmt.Println(err)
	}

	return packageJSON
}
