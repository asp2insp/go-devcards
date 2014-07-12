package main

import (
	"fmt"
	"io/ioutil"
)

// bindata_read reads the given file from disk. It returns
// an error on failure.
func bindata_read(path, name string) ([]byte, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset %s at %s: %v", name, path, err)
	}
	return buf, err
}


// data_devcards_tmpl reads file data from disk.
// It panics if something went wrong in the process.
func data_devcards_tmpl() ([]byte, error) {
	return bindata_read(
		"/Users/josiahgaskin/Projects/gopath/src/github.com/ca-geo/go-devcards/data/devcards.tmpl",
		"data/devcards.tmpl",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	if f, ok := _bindata[name]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string] func() ([]byte, error) {
	"data/devcards.tmpl": data_devcards_tmpl,

}
