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

var _schema_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8d\x31\xce\xc2\x30\x0c\x46\x77\x9f\xe2\xeb\xf6\xff\x12\x27\xc8\xcc\xc2\x00\x12\x42\x3d\x80\x55\x0c\x44\x6a\x0a\x38\xce\x50\xa1\xde\x1d\x35\x04\x04\xcd\x14\xe7\xc9\xef\x99\x28\x76\x17\x09\x8c\x07\x01\xc0\x3d\x89\x8e\x0e\xfb\xf9\x41\x26\x21\x19\x9b\xbf\x0e\x0e\xdb\x32\xd1\x44\x64\xe3\x4d\xca\xd6\x4b\x4c\x51\xf4\x4f\x02\xfb\xde\xe1\x60\xea\x87\x73\xf3\xef\xd0\x46\xd1\xcf\xfa\xdb\x2f\x46\xa7\xc2\x26\x6d\xed\xad\x70\xf2\x1a\x6d\xc7\x41\xbe\x58\xcf\x0b\xb4\xcc\xcf\x9f\x92\xf6\x47\x87\xcd\xba\xc9\xf3\x6f\x3b\xa3\x3a\x9f\x71\x75\x81\xa6\x67\x00\x00\x00\xff\xff\xd0\x74\x8a\x28\x1e\x01\x00\x00")

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
