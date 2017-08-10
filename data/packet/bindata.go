// Code generated by go-bindata.
// sources:
// cloud.json
// credential.json
// DO NOT EDIT!

package packet

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
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _cloudJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x97\xdf\x4f\xe2\x4a\x14\xc7\xdf\xf9\x2b\x4e\xe6\xe9\xde\xa4\x36\x74\xfa\x03\x6e\xdf\x2a\x92\x9b\x35\xab\x31\xa2\x6b\xcc\xc6\x90\xa1\x1c\x5d\x16\x68\xeb\xcc\xc0\xca\x1a\xfe\xf7\xcd\x54\x6c\xa1\x16\x0a\xda\x98\x6d\xd2\x17\xa5\xa7\x67\xce\xb7\xdf\xcf\x74\x7a\x72\x9e\x1b\x00\x24\x60\x53\x24\x2e\x90\x88\xf9\x63\x94\x44\x53\x31\x0c\xe6\x82\xb8\xf0\xbd\x01\x00\x40\x86\x38\x8f\xc3\x00\xe4\x91\x91\x06\xc0\x5d\x9c\xc4\xf1\x61\x14\x06\x69\xde\x73\xfc\x17\x80\x4c\x42\x9f\xc9\x51\x18\xa8\xaa\xe7\xb7\x1d\x38\x43\xc9\x43\x0d\xce\x6f\x35\xb8\xee\x79\xab\x52\x49\x01\x95\x75\xdd\x83\x2e\x13\x32\xbd\xf5\x3b\x0c\x30\xad\x1c\x87\xf0\x17\x37\xc8\xea\xf2\x2e\xfe\xbf\xd4\xb6\xeb\xf6\x58\x00\xa7\xa1\x40\x0d\x3a\xde\x2e\xd9\x1b\x2c\x90\x15\x3f\xfd\x03\x64\xbd\xa9\x90\xc8\x87\x6c\xaa\xc1\xf9\xd7\x3c\xc9\xee\x35\xdc\x74\x7b\x57\x3b\x25\xd9\x54\x1c\x20\x79\x15\x8e\x17\xa1\x06\xa7\x17\x79\x72\x5e\xef\x8b\x07\x5d\xaf\x40\x30\xe0\x32\x2b\x98\xec\xf2\x28\x10\x92\x05\x3e\xf6\xe5\x22\xc2\x9c\xcd\xc6\x27\x89\x3c\x60\x93\xbe\x18\xcf\x94\xe2\x80\x71\x9c\xa2\x64\x93\x7e\x33\xd5\x1c\xa2\xf0\xf9\x28\x4a\x1e\xf9\xf6\xa2\x0b\x4d\x17\x2c\x78\x6a\x3b\xe0\x58\x83\x91\x84\x4e\xc8\x51\x68\xd0\xfe\xff\x18\x4e\x4e\x2e\x4d\xb8\xf4\xce\xd2\xf5\x3e\x93\xf8\x10\xf2\x85\x5a\x7c\xcc\x38\xaa\x57\x8a\x4d\xd6\xee\x47\x4a\xdb\x4a\x01\xb0\x29\x71\xa1\x9d\xea\x8f\xc4\x58\x05\x9a\xb9\x34\xb7\x5b\x30\x76\x5b\x30\x72\x2d\x98\xb4\x44\x0f\x26\xcd\x9a\x30\xe8\xa1\x2e\xe8\x6e\x17\xd4\x05\x9a\x63\x83\xda\xce\x8b\x0f\x0b\xba\x9d\xce\x7b\xbc\xd0\xac\x19\x6a\x3b\x59\x37\xb4\x7d\xf0\xa6\x98\xbb\xed\x98\x2e\x18\xce\x5b\x3b\x06\x6d\xbf\xda\x79\x87\x15\xc3\xc9\x58\x31\xe8\x9b\xb7\x2b\xd9\x98\xe4\xec\x8c\x67\x03\xe4\x01\xca\xf8\xdc\x3c\xbf\x7e\x4c\xef\xd9\x6c\x22\xfb\x02\xe5\x2c\x4a\xc2\xea\x10\x86\x43\x14\x28\xd7\x42\x00\xe4\x5b\xc7\xe8\x29\x4a\xab\xc8\x72\x1d\x15\x99\x23\x17\xea\x0b\xdc\x1f\x2c\xfa\x18\xcc\xd7\x8b\xa9\x4f\xf6\xfa\x11\x4f\x4b\xa6\xeb\x94\x63\x43\xb7\x74\x3b\x31\x9b\xc7\x34\xce\x80\x7f\x06\x28\xd9\xbf\x9b\x89\x2c\x8a\xc4\xc6\xd3\xc6\xd1\xd4\xf3\x91\x40\x3e\x47\x9e\x2f\x93\xc9\x64\x13\x99\xe4\x1d\x31\x3f\x9b\x2a\x24\xe3\xf2\x68\x03\x27\x69\xea\x8e\x4e\xff\xcb\x66\xfe\x08\x85\xbc\x67\xbe\x14\x2f\xe5\x6c\xdd\x21\x6b\x09\xcb\x8c\xd3\x88\xa3\xda\xff\x21\x71\x41\xf2\x19\x36\x72\xf2\x76\x91\x73\x0a\xc9\x39\x65\x90\x73\xf6\x24\xe7\x54\x85\x9c\xad\x1b\x05\xe4\x6c\xdd\xf8\x38\xb9\xac\xcc\x36\x72\xb6\x6e\x54\x87\x1c\x2d\x24\x47\xcb\x20\x47\xf7\x24\x47\xab\x43\xce\x2c\x24\x67\x96\x41\xce\xdc\x93\x9c\x59\x1d\x72\x56\x21\x39\xab\x0c\x72\xd6\x9e\xe4\xac\xea\x90\x2b\xea\xad\x76\x19\xbd\x35\x2b\xb3\x9d\x5c\x65\x7a\xab\x5d\xd8\x5b\xed\x32\x7a\x6b\x56\x66\x3b\xb9\x0a\xf5\xd6\x56\x21\xb9\xd6\x87\x90\xb5\xf6\x44\xd6\x2a\x17\x59\x6b\x03\xd9\x61\x50\x5a\x7a\xb3\x00\xca\x9b\x8c\x83\xa0\x64\x57\x6f\x83\xd2\xd2\x9b\x9f\x03\x65\xf5\xeb\x2e\x99\x51\x1e\x59\x3d\x0f\xd4\xf3\xc0\x5f\x4a\xae\x9e\x07\xea\x79\xe0\xf3\xc9\xd5\xf3\x40\x3d\x0f\xd4\xf3\x40\x55\xc8\xd5\xf3\x40\x3d\x0f\xd4\xf3\xc0\x47\xe7\x81\xc6\xeb\xd5\xb2\xb1\xfc\x13\x00\x00\xff\xff\xac\x2e\xe2\xde\xfe\x1d\x00\x00")

func cloudJsonBytes() ([]byte, error) {
	return bindataRead(
		_cloudJson,
		"cloud.json",
	)
}

func cloudJson() (*asset, error) {
	bytes, err := cloudJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cloud.json", size: 7678, mode: os.FileMode(420), modTime: time.Unix(1453795200, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _credentialJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8f\x41\xca\x83\x30\x10\x46\xf7\x39\xc5\x30\x6b\x4f\xe0\xee\x5f\xca\xbf\x71\x5f\x44\xa2\x19\x4b\x34\x26\x21\x89\xa5\x22\xde\xbd\x64\x84\xb6\xa0\x74\x13\xc2\x7b\x1f\x3c\x66\x13\x00\xe8\x83\x7b\x68\x45\x01\x4b\x40\x2f\xfb\x89\x12\x16\x99\x2b\x1d\xbd\x91\x6b\xc6\xf5\x19\xb7\x83\x0b\xb3\x4c\xd9\x0e\x9a\x8c\x3a\x24\x7f\x23\x96\x70\x13\x00\x00\x1b\xbf\x00\x38\x46\x67\xf3\x54\x7a\xdd\x4e\xb4\xf2\x98\x85\x95\x33\x5d\x0a\x23\x3b\x32\xd9\xfc\xd5\x15\xfc\x7f\x1b\x6d\xfd\xc2\xdd\x44\xcf\x84\x4c\xf7\xe2\x3a\xe7\x83\x1b\xa9\x4f\xad\x56\xe7\xe2\x95\x7b\x47\xeb\x43\x42\xa5\x7e\x77\x05\x40\xc3\x87\x27\x79\xff\x9c\x8d\xd3\xd2\x51\xb0\x94\x28\xe6\x5d\x23\xf6\x57\x00\x00\x00\xff\xff\x64\xa4\xd7\xf8\x69\x01\x00\x00")

func credentialJsonBytes() ([]byte, error) {
	return bindataRead(
		_credentialJson,
		"credential.json",
	)
}

func credentialJson() (*asset, error) {
	bytes, err := credentialJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "credential.json", size: 361, mode: os.FileMode(420), modTime: time.Unix(1453795200, 0)}
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
	"cloud.json":      cloudJson,
	"credential.json": credentialJson,
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
	"cloud.json":      {cloudJson, map[string]*bintree{}},
	"credential.json": {credentialJson, map[string]*bintree{}},
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
