package schema

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var _schema_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x90\xb1\x4e\x03\x31\x0c\x86\xf7\x3c\x85\xb3\x15\xa9\x4f\xe0\x15\x96\x1b\x40\x20\xd4\x09\x31\x58\xc5\x3d\x22\x5d\x72\x25\x76\x86\x0a\xf5\xdd\x51\x72\x49\xe8\xf5\xa6\xd8\x7f\xfe\x7c\xbf\x1d\x39\x7e\xb3\x27\xf8\x35\x00\x00\x3f\x89\xe3\x05\xe1\x2d\x1f\x50\x14\x9f\x94\xd4\xcd\x01\xe1\xb9\x56\xe6\x6a\x8c\x5e\xce\x5c\x5d\xcb\xc3\x24\x1c\x77\xec\xc9\x4d\x08\xef\x1a\x5d\x18\xed\x03\xc2\x41\x38\xf6\x6b\xd9\x9d\x69\x64\x84\x21\xa8\xdd\xc3\xe4\xbc\xd3\xa5\xa9\x46\xe9\xe0\x96\x54\xd9\xc7\xc8\xa4\x7c\xd8\x26\xec\xe1\xe4\xa2\xe8\x0b\x79\xbe\xd1\x26\xba\x93\xda\x20\x0d\x9f\x9b\x8a\x76\x5f\x08\xc3\x93\x2d\xf5\x9a\x5d\xa4\x2d\xbe\xc8\x9b\x84\x8e\x7e\xa5\x91\x87\x70\x9a\x2b\xfe\x7f\xe1\xe5\x33\x6f\x96\x5e\x4d\x23\xd5\xaf\xb3\xd2\xf4\x38\xa7\xd0\x4c\x65\x46\x65\x2f\x08\x1f\xd9\x68\x3f\x3b\x37\xc7\x60\x0f\xb4\xe6\xfa\x17\x00\x00\xff\xff\x69\xb2\xd2\xf2\xc9\x01\x00\x00")

func schema_graphql() ([]byte, error) {
	return bindata_read(
		_schema_graphql,
		"schema.graphql",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"schema.graphql": schema_graphql,
}
// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"schema.graphql": &_bintree_t{schema_graphql, map[string]*_bintree_t{
	}},
}}
