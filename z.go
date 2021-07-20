package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	// app := mamba.New()
	// app.AddCommand("install {packageName}", addPackage)
	// app.Run(os7.Args)
	addPackage("kobra.js", "0.1.1")
	fmt.Println(getPackageJSON().Dependencies)
}

func install(depMap map[string]string) []string {
	var deps []string

	for key, val := range depMap {
		deps = append(deps, key+"|"+val)
	}

	return deps
}

func addPackage(name string, version string) {
	full := name + "@" + version

	fmt.Println("http://registry.npmjs.org/" + name)

	resp, err := http.Get("http://registry.npmjs.org/" + name)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	jsonBody, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	var pkgInfo NPMPackage

	err = json.Unmarshal(jsonBody, &pkgInfo)

	if err != nil {
		fmt.Println(err)
		return
	}

	tarball := pkgInfo.Versions[version].Dist.Tarball

	r, err := http.Get(tarball)

	if err != nil {
		fmt.Println(err)
	}

	defer r.Body.Close()

	err = Untar(r.Body, getCachePath()+"/", full)

	if err != nil {
		fmt.Println(err)
	}
}
