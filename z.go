package main

import (
	"archive/tar"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

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
		return
	}

	defer resp.Body.Close()

	jsonBody, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	var packageInfo NPMPackage

	err = json.Unmarshal(jsonBody, &packageInfo)

	if err != nil {
		fmt.Println(err)
		return
	}

	tarball := packageInfo.Versions["0.1.1"].Dist.Tarball

	r, err := http.Get(tarball)

	if err != nil {
		fmt.Println(err)
	}

	defer r.Body.Close()

	err = Untar(r.Body, getCachePath()+"/",packageName)

	if err != nil {
		fmt.Println(err)
	}

	return
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

	err = os.Mkdir(path, 0777)

	if err != nil {
		fmt.Println(err)
	}

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

// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package untar untars a tarball to disk.

// TODO(bradfitz): this was copied from x/build/cmd/buildlet/buildlet.go
// but there were some buildlet-specific bits in there, so the code is
// forked for now.  Unfork and add some opts arguments here, so the
// buildlet can use this code somehow.

// Untar reads the gzip-compressed tar file from r and writes it into dir.
func Untar(r io.Reader, dir string, packageName string) error {
	return untar(r, dir, packageName)
}

func untar(r io.Reader, dir string, packageName string) (err error) {
	t0 := time.Now()
	nFiles := 0
	madeDir := map[string]bool{}
	defer func() {
		td := time.Since(t0)
		if err == nil {
			log.Printf("extracted tarball into %s: %d files, %d dirs (%v)", dir, nFiles, len(madeDir), td)
		} else {
			log.Printf("error extracting tarball into %s after %d files, %d dirs, %v: %v", dir, nFiles, len(madeDir), td, err)
		}
	}()
	zr, err := gzip.NewReader(r)
	if err != nil {
		return fmt.Errorf("requires gzip-compressed body: %v", err)
	}
	tr := tar.NewReader(zr)
	loggedChtimesError := false
	for {
		f, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("tar reading error: %v", err)
			return fmt.Errorf("tar error: %v", err)
		}
		if !validRelPath(f.Name) {
			return fmt.Errorf("tar contained invalid name error %q", f.Name)
		}
		rel := filepath.FromSlash(f.Name)

		rel = strings.Replace(rel, "package/", packageName+"/", 1)

		abs := filepath.Join(dir, rel)

		fi := f.FileInfo()
		mode := fi.Mode()
		switch {
		case mode.IsRegular():
			// Make the directory. This is redundant because it should
			// already be made by a directory entry in the tar
			// beforehand. Thus, don't check for errors; the next
			// write will fail with the same error.
			dir := filepath.Dir(abs)
			if !madeDir[dir] {
				if err := os.MkdirAll(filepath.Dir(abs), 0755); err != nil {
					return err
				}
				madeDir[dir] = true
			}
			wf, err := os.OpenFile(abs, os.O_RDWR|os.O_CREATE|os.O_TRUNC, mode.Perm())
			if err != nil {
				return err
			}
			n, err := io.Copy(wf, tr)
			if closeErr := wf.Close(); closeErr != nil && err == nil {
				err = closeErr
			}
			if err != nil {
				return fmt.Errorf("error writing to %s: %v", abs, err)
			}
			if n != f.Size {
				return fmt.Errorf("only wrote %d bytes to %s; expected %d", n, abs, f.Size)
			}
			modTime := f.ModTime
			if modTime.After(t0) {
				// Clamp modtimes at system time. See
				// golang.org/issue/19062 when clock on
				// buildlet was behind the gitmirror server
				// doing the git-archive.
				modTime = t0
			}
			if !modTime.IsZero() {
				if err := os.Chtimes(abs, modTime, modTime); err != nil && !loggedChtimesError {
					// benign error. Gerrit doesn't even set the
					// modtime in these, and we don't end up relying
					// on it anywhere (the gomote push command relies
					// on digests only), so this is a little pointless
					// for now.
					log.Printf("error changing modtime: %v (further Chtimes errors suppressed)", err)
					loggedChtimesError = true // once is enough
				}
			}
			nFiles++
		case mode.IsDir():
			if err := os.MkdirAll(abs, 0755); err != nil {
				return err
			}
			madeDir[abs] = true
		default:
			return fmt.Errorf("tar file entry %s contained unsupported file type %v", f.Name, mode)
		}
	}
	return nil
}

func validRelativeDir(dir string) bool {
	if strings.Contains(dir, `\`) || path.IsAbs(dir) {
		return false
	}
	dir = path.Clean(dir)
	if strings.HasPrefix(dir, "../") || strings.HasSuffix(dir, "/..") || dir == ".." {
		return false
	}
	return true
}

func validRelPath(p string) bool {
	if p == "" || strings.Contains(p, `\`) || strings.HasPrefix(p, "/") || strings.Contains(p, "../") {
		return false
	}
	return true
}