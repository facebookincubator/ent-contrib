// Package internal Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// template/enums.tmpl
// template/service.tmpl
package internal

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templateEnumsTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x91\x41\x6f\xd4\x30\x10\x85\xef\xf9\x15\x23\x8b\x43\x22\xb1\x16\x82\x1b\x88\x1b\xdb\x55\xa5\x82\x56\xa2\x9c\x2a\x54\x79\x93\x49\xb1\x9a\xd8\xc6\x99\x2c\xaa\xcc\xfc\x77\x64\xa7\x4d\xea\x92\x45\xa5\x39\x45\xf6\xbc\xf7\xbd\x79\x0e\x01\x1a\x6c\xb5\x41\x10\x68\xc6\x7e\x10\xc0\x5c\x00\x00\x84\xb0\x81\x57\xde\x5a\x82\xf7\x1f\x41\xc2\x66\x3e\x06\xaf\xcc\x0d\x82\x3c\xd3\xd8\x35\x9f\x95\x93\xdb\xa8\x7b\x90\xcd\xd2\xe8\x76\x79\xe7\x30\xc9\xf7\x87\x34\xfd\x09\x87\xda\x6b\x47\xd6\xcb\x1d\xd2\xf6\x61\x62\xb3\xa6\xfd\xa2\xfa\xa4\x75\x5e\x1b\x9a\x92\xc8\xad\xa1\x28\x90\xe9\x4e\x5c\x8b\x85\x12\xfd\xd2\xe9\x5f\x5e\xee\x10\x39\xe7\x0d\x9a\xb4\xc9\x64\x74\xa6\x3b\x94\x3b\x7b\xde\x3b\xeb\x69\xaf\xe8\x87\x9c\x06\x16\x32\xac\xa5\xa2\x8b\x5a\x0d\x29\xd5\x60\xd4\x2d\xae\xa5\x5a\x53\x65\x01\xd0\xd0\xcc\xba\xf7\x93\xfb\xc3\x57\xf2\x63\x4d\xa9\xa4\xcc\xa2\x1d\x4d\x0d\x64\xf7\xde\x92\x0d\x21\xdb\x46\xee\x6c\x22\x32\x43\x89\xf1\x5d\xf4\x6c\xbb\x00\x81\xb9\x7a\x74\xf7\xb8\x0c\x66\x08\x33\x27\x7e\xba\x85\xe3\x6b\xb0\xb7\x31\xe5\x49\xd6\xf5\x51\x75\x23\x5e\x85\x00\x3f\x47\xd5\xe9\xf6\x0e\xc4\x40\x5e\x9b\x9b\x41\x80\xb8\xb4\xdf\x9c\x43\x2f\x98\xcb\xe9\xb0\xc4\xaa\xfa\xfe\x21\x7a\xe6\xa8\xf8\x79\xa4\xd1\x9b\xa7\x24\xf8\x7d\x9f\x95\xb9\x3c\x56\x99\x88\x8b\xff\x93\xbf\x59\xe4\x5c\x3c\x6d\x74\x6b\xe8\xe4\x8e\x59\x9d\x79\x65\xd5\x3f\x8a\x7e\x51\x9d\x46\xf5\x78\xa5\x0d\xbd\x7b\x5b\xe2\x33\xaa\x5a\x43\x33\x97\xa7\x9e\xe3\xc2\xfe\x42\x2f\xa6\x2e\x9f\x51\xa6\x10\x45\x7e\x1f\x02\xa0\x69\x98\x8b\xe9\x07\x98\xff\x04\x00\x00\xff\xff\x37\x7d\x23\x2a\x30\x04\x00\x00")

func templateEnumsTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateEnumsTmpl,
		"template/enums.tmpl",
	)
}

func templateEnumsTmpl() (*asset, error) {
	bytes, err := templateEnumsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/enums.tmpl", size: 1072, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateServiceTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x90\xc1\x4e\xc3\x30\x0c\x86\xef\x7d\x8a\x5f\x3d\x20\x98\x68\x72\xdf\x75\x1b\xd3\x2e\x1b\x12\xe3\x01\xba\xd4\x54\x15\xad\x53\xa5\x2e\xd3\x14\xf2\xee\x28\xcd\x04\x07\x3a\x24\x4e\xad\x7e\x7f\x76\x3e\xdb\xfb\x02\x7a\x51\x5b\xb9\xf4\xb4\x04\xb1\xd4\x56\x35\x56\x1b\xcb\xe2\x9a\x93\x26\x96\xde\x59\xb1\xda\x74\x95\x9e\xfe\x4c\x51\x13\x17\x11\x74\xbd\x51\x03\xb9\x8f\xc6\xd0\x96\x98\x5c\x29\xd6\x2d\x34\x8a\x10\x32\xad\xb1\xb2\x15\xa1\x4e\x39\x55\x38\x5d\x30\xd3\x8e\xf5\x01\xfb\xc3\x11\x9b\xf5\xee\xa8\xb2\xbe\x34\xef\x65\x4d\xf0\x1e\xea\xa9\x69\x49\x6d\xed\x73\x8a\xf6\x65\x47\x08\x21\x8b\x83\x63\xf5\x25\x3d\xab\xb6\xf6\x5a\x41\xd3\xf5\x2d\x75\xc4\x32\xcc\x03\x31\x20\x97\xc5\x35\x6f\x4c\x18\xc4\x8d\x46\xe0\x33\x00\x30\x6d\x43\x2c\x58\x44\x74\xc3\x72\xd5\x50\xbb\x2a\xa6\xf9\x6a\xaa\xe6\xf8\x44\x33\x05\x21\x4c\x4d\xaf\xfc\x6d\x41\xd5\x5f\x16\x69\x91\x3d\x9d\xe7\x4d\x1c\xc9\xe8\x78\x40\x09\xa6\xf3\xbc\x6d\xf6\x36\xb2\xb9\x39\xe1\xfe\xbf\xfa\x0f\x09\xfd\xad\x92\xae\x91\x84\x70\x37\xcb\x24\xe4\xe7\x68\xcb\xeb\xf7\x71\xca\x43\x5c\xd6\x7b\x08\x75\x7d\x5b\x0a\x21\x27\x1e\xbb\x21\x87\x42\x08\x5f\x01\x00\x00\xff\xff\x81\x8f\x94\xcb\x7d\x02\x00\x00")

func templateServiceTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateServiceTmpl,
		"template/service.tmpl",
	)
}

func templateServiceTmpl() (*asset, error) {
	bytes, err := templateServiceTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/service.tmpl", size: 637, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
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
var _bindata = map[string]func() (*asset, error){
	"template/enums.tmpl":   templateEnumsTmpl,
	"template/service.tmpl": templateServiceTmpl,
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
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"template": &bintree{nil, map[string]*bintree{
		"enums.tmpl":   &bintree{templateEnumsTmpl, map[string]*bintree{}},
		"service.tmpl": &bintree{templateServiceTmpl, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
