// +build ignore

// simple command to generate CGOFLAGS for a given python VM
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/pkg/errors"
)

func main() {
	vm := flag.String("vm", "python", "path to a python VM")
	flag.Parse()

	if *vm == "" {
		log.Fatalf("need a python VM")
	}

	cfg, err := getPythonConfig(*vm)
	if err != nil {
		log.Fatalf("could not infer python configuration: %v", err)
	}

	oname := "cgoflags_unix.go"
	switch runtime.GOOS {
	case "windows":
		oname = "cgoflags_windows.go"
	}

	err = ioutil.WriteFile(oname, []byte(fmt.Sprintf(tmpl,
		cfg.cflags,
		cfg.ldflags,
		runtime.GOOS,
	)), 0644)
	if err != nil {
		log.Fatalf("could not write %q: %v", oname, err)
	}
}

// getPythonConfig returns the needed python configuration for the given
// python VM (python, python2, python3, pypy, etc...)
func getPythonConfig(vm string) (pyconfig, error) {
	code := `import sys
import distutils.sysconfig as ds
import json
print(ds.get_config_vars())
print(ds.get_python_lib())
print(json.dumps({
	"version": sys.version_info.major,
	"prefix":  ds.get_config_var("prefix"),
	"incdir":  ds.get_python_inc(),
	"libdir":  ds.get_config_var("LIBDIR"),
	"libdest":  ds.get_config_var("LIBDEST"),
	"libpy":   ds.get_config_var("LIBRARY"),
	"shlibs":  ds.get_config_var("SHLIBS"),
	"syslibs": ds.get_config_var("SYSLIBS"),
	"shlinks": ds.get_config_var("LINKFORSHARED"),
	"DLLs": ds.get_config_var("DLLLIBRARY"),
}))
`

	var cfg pyconfig
	bin, err := exec.LookPath(vm)
	if err != nil {
		return cfg, errors.Wrapf(err, "could not locate python vm %q", vm)
	}

	buf := new(bytes.Buffer)
	cmd := exec.Command(bin, "-c", code)
	cmd.Stdin = os.Stdin
	cmd.Stdout = buf
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return cfg, errors.Wrap(err, "could not run python-config script")
	}

	log.Printf("distutils:\n%s", buf.String())

	var raw struct {
		Version int    `json:"version"`
		IncDir  string `json:"incdir"`
		LibDir  string `json:"libdir"`
		LibPy   string `json:"libpy"`
		ShLibs  string `json:"shlibs"`
		SysLibs string `json:"syslibs"`
	}
	err = json.NewDecoder(buf).Decode(&raw)
	if err != nil {
		return cfg, errors.Wrapf(err, "could not decode JSON script output")
	}

	if strings.HasSuffix(raw.LibPy, ".a") {
		raw.LibPy = raw.LibPy[:len(raw.LibPy)-len(".a")]
	}
	if strings.HasPrefix(raw.LibPy, "lib") {
		raw.LibPy = raw.LibPy[len("lib"):]
	}

	cfg.version = raw.Version
	cfg.cflags = strings.Join([]string{
		"-I " + raw.IncDir,
	}, " ")
	cfg.ldflags = strings.Join([]string{
		"-L " + raw.LibDir,
		"-l " + raw.LibPy,
		raw.ShLibs,
		raw.SysLibs,
	}, " ")

	return cfg, nil
}

type pyconfig struct {
	version int
	cflags  string
	ldflags string
}

const tmpl = `// Automatically generated. Do not edit.

// +build %[3]s

package python

// #cgo %[3]s CFLAGS: %[1]s
// #cgo %[3]s LDFLAGS: %[2]s
//
// #include "go-python.h"
import "C"
`
