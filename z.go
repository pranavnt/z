package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	// app := mamba.New()
	// app.AddCommand("install {packageName}", addPackage)
	// app.Run(os7.Args)
	install()

	addPackage("notistack", "1.0.9")
}

func install() {

	if !doesLockfileExist() {
		os.Create("./z.toml")

		packageJSON := getPackageJSON()

		deps := getDeps(packageJSON.Dependencies)

		fmt.Println("deps: ", deps)

		for _, el := range deps {
			fmt.Println(el)
			arr := strings.Split(el, "|")
			fmt.Println(arr[0], arr[1])
			addPackage(arr[0], arr[1])
		}
	} else {
		// TODO: install with lockfile

	}
}

func getDeps(depMap map[string]string) []string {
	var deps []string

	for key, val := range depMap {
		deps = append(deps, key+"|"+strings.Trim(val, "^"))
	}

	return deps
}

func addPackage(name string, version string) map[string]string {
	full := name + "@" + version

	fmt.Println("http://registry.npmjs.org/" + name)

	resp, err := http.Get("http://registry.npmjs.org/" + name)

	if err != nil {
		log.Fatal(err)
		return
	}

	defer resp.Body.Close()

	jsonBody, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var pkgInfo NPMPackage

	err = json.Unmarshal(jsonBody, &pkgInfo)

	if err != nil {
		log.Fatal(err)
	}

	tarball := pkgInfo.Versions[version].Dist.Tarball

	// fmt.Println(pkgInfo.Versions)
	// fmt.Println(pkgInfo.Versions[version])
	fmt.Println(pkgInfo.Versions[version].Dist.Tarball)

	resp, err = http.Get(tarball)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	err = Untar(resp.Body, getCachePath()+"/", full)

	if err != nil {
		log.Fatal(err)
	}

	return pkgInfo.Versions[version].Dependencies
}
