package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

	fmt.Println("http://registry.npmjs.org/" + packageName)

	resp, err := http.Get("http://registry.npmjs.org/" + packageName)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	jsonBody, err := io.ReadAll(resp.Body)

	var packageInfo NPMPackage

	err = json.Unmarshal(jsonBody, &packageInfo)

	if err != nil {
		fmt.Println(err)
	}

	tarball := packageInfo.Versions["0.1.1"].Dist.Tarball

	DownloadFile(tarball, tarball)

	fmt.Println(tarball)
}

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

	os.Setenv("Z_PKG_CACHE_PATH", path)

	return path
}

type NPMPackage struct {
	ID       string `json:"_id"`
	Rev      string `json:"_rev"`
	Name     string `json:"name"`
	DistTags struct {
		Latest string `json:"latest"`
	} `json:"dist-tags"`
	Versions    map[string]PackageVersion `json:"versions"`
	Maintainers []struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"maintainers"`
	Description string `json:"description"`
	Author      struct {
		Name string `json:"name"`
	} `json:"author"`
	License        string `json:"license"`
	Readme         string `json:"readme"`
	ReadmeFilename string `json:"readmeFilename"`
}

type PackageVersion struct {
	Version          string `json:"version"`
	License          string `json:"license"`
	Main             string `json:"main"`
	Typings          string `json:"typings"`
	PeerDependencies struct {
	} `json:"peerDependencies"`
	Name   string `json:"name"`
	Author struct {
		Name string `json:"name"`
	} `json:"author"`
	Module    string `json:"module"`
	SizeLimit []struct {
		Path  string `json:"path"`
		Limit string `json:"limit"`
	} `json:"size-limit"`
	DevDependencies map[string]string `json:"devDependencies"`
	Dependencies    map[string]string `json:"dependencies"`
	Description     string            `json:"description"`
	LicenseText     string            `json:"licenseText"`
	ID              string            `json:"_id"`
	Dist            struct {
		Shasum       string `json:"shasum"`
		Integrity    string `json:"integrity"`
		Tarball      string `json:"tarball"`
		FileCount    int    `json:"fileCount"`
		UnpackedSize int    `json:"unpackedSize"`
		NpmSignature string `json:"npm-signature"`
	} `json:"dist"`
	NpmUser struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"_npmUser"`
	Directories struct {
	} `json:"directories"`
	Maintainers []struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"maintainers"`
	NpmOperationalInternal struct {
		Host string `json:"host"`
		Tmp  string `json:"tmp"`
	} `json:"_npmOperationalInternal"`
	HasShrinkwrap bool `json:"_hasShrinkwrap"`
}
