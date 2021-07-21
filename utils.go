package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func getCachePath() string {
	path := os.Getenv("Z_PKG_CACHE_PATH")

	if path != "" {
		return path
	}

	dir, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)
	}

	path = dir

	path += "/.z"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, 0777)
	}

	if err != nil {
		log.Fatal(err)
	}

	os.Setenv("Z_PKG_CACHE_PATH", path)

	return path
}

func doesLockfileExist() bool {
	if _, err := os.Stat("./z.toml"); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func getPackageJSON() PackageJSON {
	data, err := ioutil.ReadFile("./package.json")

	if err != nil {
		log.Fatal(err)
	}

	var packageJSON PackageJSON

	err = json.Unmarshal(data, &packageJSON)

	if err != nil {
		log.Fatal(err)
	}

	return packageJSON
}
