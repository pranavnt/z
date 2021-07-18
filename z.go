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

	

	fmt.Println(string(packageName))
}

type NPMPackage struct {
	ID       string `json:"_id"`
	Rev      string `json:"_rev"`
	Name     string `json:"name"`
	DistTags struct {
		Latest string `json:"latest"`
	} `json:"dist-tags"`
	Versions struct {
		Zero10 struct {
			Version string `json:"version"`
			License string `json:"license"`
			Main    string `json:"main"`
			Typings string `json:"typings"`

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
		} `json:"0.1.0"`
		Zero11 struct {
			Version string `json:"version"`
			License string `json:"license"`
			Main    string `json:"main"`
			Typings string `json:"typings"`
			Engines struct {
				Node string `json:"node"`
			} `json:"engines"`
			Scripts struct {
				Start   string `json:"start"`
				Build   string `json:"build"`
				Test    string `json:"test"`
				Lint    string `json:"lint"`
				Prepare string `json:"prepare"`
				Size    string `json:"size"`
				Analyze string `json:"analyze"`
			} `json:"scripts"`
			PeerDependencies struct {
			} `json:"peerDependencies"`
			Husky struct {
				Hooks struct {
					PreCommit string `json:"pre-commit"`
				} `json:"hooks"`
			} `json:"husky"`
			Prettier struct {
				PrintWidth    int    `json:"printWidth"`
				Semi          bool   `json:"semi"`
				SingleQuote   bool   `json:"singleQuote"`
				TrailingComma string `json:"trailingComma"`
			} `json:"prettier"`
			Name   string `json:"name"`
			Author struct {
				Name string `json:"name"`
			} `json:"author"`
			Module    string `json:"module"`
			SizeLimit []struct {
				Path  string `json:"path"`
				Limit string `json:"limit"`
			} `json:"size-limit"`
			DevDependencies struct {
				SizeLimitPresetSmallLib string `json:"@size-limit/preset-small-lib"`
				Husky                   string `json:"husky"`
				SizeLimit               string `json:"size-limit"`
				Tsdx                    string `json:"tsdx"`
				Tslib                   string `json:"tslib"`
				Typescript              string `json:"typescript"`
			} `json:"devDependencies"`
			Dependencies struct {
				BabelPluginProposalClassProperties   string `json:"@babel/plugin-proposal-class-properties"`
				BabelPluginTransformTypescript       string `json:"@babel/plugin-transform-typescript"`
				KobraDevJsRegression                 string `json:"@kobra-dev/js-regression"`
				KobraDevMlSvm                        string `json:"@kobra-dev/ml-svm"`
				KobraDevMultivariateLinearRegression string `json:"@kobra-dev/multivariate-linear-regression"`
				TypesJest                            string `json:"@types/jest"`
				MlKnn                                string `json:"ml-knn"`
				MlRandomForest                       string `json:"ml-random-forest"`
			} `json:"dependencies"`
			Description string `json:"description"`
			LicenseText string `json:"licenseText"`
			ID          string `json:"_id"`
			Dist        struct {
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
		} `json:"0.1.1"`
	} `json:"versions"`
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
