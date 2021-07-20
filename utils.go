package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func getCachePath() string {
	path := os.Getenv("Z_PKG_CACHE_PATH")

	if path != "" {
		return path
	}

	dir, err := os.UserHomeDir()

	if err != nil {
		fmt.Println(err)
	}

	path = dir

	path += "/.z"

	err = os.Mkdir(path, 0777)

	if err != nil {
		fmt.Println(err)
	}

	os.Setenv("Z_PKG_CACHE_PATH", path)

	return path
}

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
