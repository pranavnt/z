package main

type PackageJSON struct {
	Name            string            `json:"name"`
	Version         string            `json:"version"`
	Private         bool              `json:"private"`
	Dependencies    map[string]string `json:"dependencies"`
	Scripts         map[string]string `json:"scripts"`
	DevDependencies map[string]string `json:"devDependencies"`
}

type NPMPackage struct {
	ID          string                    `json:"_id"`
	Rev         string                    `json:"_rev"`
	Name        string                    `json:"name"`
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
	Name      string `json:"name"`
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
