// Code generated by go-bindata.
// sources:
// cloud.json
// credential.json
// DO NOT EDIT!

package azure

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

var _cloudJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x5b\x5d\x6f\xdb\xc6\x12\x7d\xf7\xaf\x58\xf0\xe9\x5e\x20\x16\xb4\xfc\x54\xfc\x96\x6b\x45\xbc\x69\x8a\xb4\x05\x9b\xe4\xa1\x08\x8c\xb5\xb4\x76\x58\xc9\xa4\xca\x5d\x2a\x71\x02\xff\xf7\x82\xb4\x23\x59\x5c\x0e\x77\x86\x12\xeb\xfa\xa5\x0d\xa8\x11\xe7\xec\xd9\x39\x73\x66\x29\xfa\xfb\x09\x63\x4e\x26\x6e\xa4\x73\xc6\x1c\xf1\xad\x2c\xa4\xf3\xa2\xba\x24\xb3\x8d\x72\xce\xd8\x1f\x27\x8c\x31\xe6\x2c\xe4\xa6\xbe\xcc\x98\xf3\x97\x70\x4e\x18\xfb\x54\x07\x15\xf2\x3a\xcd\xb3\x5d\xdc\xf7\xfa\xbf\x8c\x39\xab\x7c\x2e\x74\x9a\x67\xd5\x4d\x3f\xa4\xc5\x75\x9a\xa5\xe2\xe1\x06\xdb\xaf\x55\x9f\xbd\x16\x4a\xb3\xf7\xc9\xee\xa3\x6f\x79\x26\x77\xf7\xab\x2f\x49\xa1\x74\xa9\x9c\x87\x0b\x9f\xea\xff\xdf\xbd\x38\x28\x1f\x73\x11\x19\x5d\x7c\xca\x37\xf9\x97\xd6\x74\xe7\x32\xd3\x85\x58\xd9\x56\x38\xbf\x0f\xa3\x2c\xf2\xcd\x6a\x95\x66\x79\xaa\xda\xb2\xbe\xcb\x0b\xfd\x99\x21\x73\x67\x55\x70\x0f\x00\xbf\xcb\xaf\xa2\x35\x7b\x92\x97\xf8\xec\xaa\x0a\xee\x91\xfd\xa3\x54\xba\x35\xc7\x0e\x07\x18\xd2\x02\xe3\x8b\x54\xba\x07\x8a\x73\xb1\x4a\xaf\xf2\x02\xa8\xb5\x1a\x00\x22\x31\x79\xdd\xfb\xf5\x6b\x24\xb4\x14\xf7\x7d\x4a\x42\x71\xff\x56\xca\x4b\x39\x67\xe7\xa9\xbe\x6d\xad\x71\x91\x89\x85\x60\x95\xb2\xba\x8b\xbc\x8e\xab\xa4\x45\x28\xb2\xbc\xc8\x33\x9d\x77\xa4\x7d\xd8\x60\x44\xe6\x87\xfd\xc5\x27\x4f\x44\xce\x7e\x15\xe5\x2a\x67\x89\x16\x5a\xb6\x81\xf8\x5f\x21\xbe\xa5\x2b\x56\x97\x7c\x27\x84\xcb\x3a\xb0\xae\x76\x82\xc6\x0b\xb9\x12\xd9\x02\x96\xf8\xeb\xb2\xc8\xd7\xd2\x2e\x6f\x79\x1f\x87\x4e\xfc\x4e\xea\xcf\xb2\xa8\x72\xb7\x2a\xbc\xae\x33\x44\xee\xaa\xd2\xa8\xa9\xcf\x45\xb1\x48\xaf\xae\xda\xd2\xbe\x7f\xcb\xaa\xcc\x9d\x29\xcb\x65\x95\x14\x9f\xee\xe7\x3c\x5b\xe4\x19\x90\xcd\xbe\xad\xe5\x92\xb8\xa5\x49\x9a\x5d\x8b\x75\x5e\xb4\x56\x53\x9d\xaf\x12\x08\x7b\xa5\x1e\xb7\x14\xa8\x6f\x56\xa1\xa2\x8a\x44\xa7\xff\x7f\x9e\x5d\xb3\xb7\x79\x76\x0d\x7a\xa3\x35\x33\x3d\xe9\x3b\xf9\xe5\x9e\x4a\xf6\x51\xac\x64\x6b\x45\xbd\x2a\x55\xa5\xcd\x14\xd1\x46\xc4\x8f\x50\x5a\x27\xf9\x90\xce\x75\x5e\xb4\x37\xea\x5d\xf6\xed\x0e\xe0\x20\x6c\x77\x81\xd2\xd1\x96\xb7\xf9\x0b\x96\x88\x54\x8b\x9b\x56\x34\x3f\x89\xb5\xc8\xec\x3c\xfc\x59\x85\xd1\x72\xff\xa2\xc4\xb2\x23\xa5\x55\x5c\x75\xca\x16\x7d\x6d\xc7\xc1\x34\x53\x5a\x64\x73\x79\xa1\x6f\xd7\xb2\x65\x2a\x94\x5f\xb5\x2c\x32\xb1\xba\x50\xcb\xb2\x2e\x78\x2d\xb2\x85\x28\x16\x17\xaf\xc6\xbb\xc4\x0b\xa9\xe6\x45\xba\xde\x2a\xa6\x2d\x66\x2e\xb4\xbc\xce\x8b\xdb\x7a\xf7\x4e\x95\x2c\xd2\xc7\x75\x35\x5f\x57\x77\xe7\xbb\x85\x8a\x1b\xe7\x8c\x8d\x47\x51\x38\xd9\x65\x49\xd5\xd2\x39\x63\xee\xb8\x95\x35\x18\x28\x47\x00\xe5\x87\x01\xe5\xa3\x28\x68\xe2\x8c\xa8\x38\x5d\x04\x4e\x97\x86\xd3\x6d\xe0\xf4\x46\x06\x4c\xee\x05\x44\x9c\x1e\x02\xa7\x47\xc3\xe9\x37\x70\x46\xc6\xa6\x4f\xa8\x28\x7d\x04\x4a\x9f\x86\x72\xd2\xdc\x75\xbf\x09\x33\x1c\x53\x61\x06\x08\x98\xc1\x61\x9b\x6e\xc2\xa4\xef\x79\x88\x80\x19\x1e\xb6\xe7\xae\xa9\x74\xf2\xa6\x47\x08\x98\xd1\x61\x9b\x1e\x84\x87\x6f\xfa\x04\x01\x73\xd2\x0d\x93\x9d\xb2\x79\x7e\xb3\x2e\xb5\x3c\x4d\x33\x2d\x33\x95\x6e\x24\xfb\xd1\xcc\x7b\x2c\xc2\x9b\xb8\xc4\x45\xbc\x44\x2c\xe2\xe5\x31\x17\xc1\xc3\x66\x61\x73\xf7\xf0\x65\x70\x8c\x8f\x71\x8b\x91\x3d\xfd\x6e\x70\x94\xcb\x59\x6c\xee\x5f\xb0\x1f\x53\xc4\x3a\xa6\xc0\x32\xa6\x48\xb7\x6e\x71\xc1\x80\x68\xd6\x53\x84\x59\x4f\x01\xb3\x86\x60\x36\xfb\xb6\x61\x82\x7c\x4c\x45\x89\xb0\xea\x29\x60\xd5\x10\xca\x66\xdb\x36\xdd\xc5\x25\xc3\x44\x78\xf5\x14\xf0\x6a\x08\x66\x53\x63\xa6\xbb\xf8\x64\x98\x18\x8d\x4d\x21\x8d\x61\x77\xbd\xc5\xad\xe9\x40\x31\xd5\xc9\x89\xe5\x89\xf0\x6b\x3a\x50\x4c\x81\x72\x62\x85\xda\xdb\x6b\x8f\xad\xc7\x94\x28\x27\xd6\x28\xa6\x81\x4e\xe8\x50\x2f\x36\xa8\xed\xdf\x0b\xdb\x83\xbb\x71\xff\xc9\x46\x8a\x83\xeb\xf6\x80\x3b\x44\x43\xc5\xa1\xf5\x7a\xa0\x1d\xa4\xb1\xe2\xe0\xfa\x3d\xe0\x0e\xd1\x60\x03\x1c\xdc\xa0\x4f\xe9\x36\xb5\x66\x76\x85\x1e\x52\xc3\x6a\xad\x8f\xd8\x86\x31\x06\x24\xe0\x3e\x72\x1b\xc6\x20\x90\x80\xfb\x28\x6e\x18\xa3\x40\x02\xee\xa3\xb9\x81\x0c\x03\x29\x3b\xde\x47\x77\xee\xd8\x28\xe3\x71\x4b\x1d\x53\x31\x27\x98\x51\x2c\x81\x46\xb1\xa4\xbf\xbf\x45\x54\xa0\x18\x66\x13\x88\x56\x08\xa8\xdd\xd9\x7c\x2a\x4e\xcc\x24\x96\x40\x93\x18\x84\x13\xe1\x69\x13\x2a\x50\xcc\x24\x96\x40\x93\x18\x04\xd4\xee\x66\x41\x48\x2e\x51\x5c\x8d\x52\x8b\xd4\xee\x0b\x74\x4e\x51\x07\x86\x04\x3c\x31\x60\xb7\xff\x28\xac\xa2\x0a\x15\x3c\x33\x60\x0b\xc0\xf4\x02\xce\xa9\x8f\x32\x12\xd4\xa9\x21\x01\x8f\x0d\x60\x9b\x42\xd8\x80\xeb\x92\x3b\x00\x72\x98\x49\x3a\x86\x99\xe4\x90\xa3\x03\xbd\xb5\x22\x01\x77\x0c\x33\x30\xe0\x21\x5a\x2c\x12\x6f\xc7\x2c\x03\xe3\x1d\xa4\xd5\x22\x01\x77\xcc\x32\x30\xe0\x21\x5a\x2e\x72\x92\x49\x3a\x26\x99\x8e\x12\xb6\x1f\x21\xfa\xf4\x08\xb4\xec\x7a\xe9\x6e\x18\xb7\xc0\x42\xee\xa5\xbc\x61\x5c\x03\x0b\xb9\x97\xf8\x06\x72\x0f\x2c\xe6\x5e\xfa\x1b\xca\x45\xb0\x1a\xec\x3a\x4e\x74\xd4\x33\xe2\x3c\xe1\x4e\x88\xc7\x89\x19\x62\x52\x9b\x01\x73\xda\x0c\x69\x78\x06\xb7\x9c\x58\xc4\x33\x04\xab\x33\x80\x50\x08\x64\xb3\x3b\x18\xcd\xc1\x23\x56\xed\x0c\x31\xf1\xcc\x80\x79\x07\x02\xd9\xec\x07\x46\x3b\x08\x89\x45\x3a\x43\xfc\xfe\x3a\x03\x7e\x7f\x85\x40\x1a\xef\x06\x98\x1d\x80\xda\x68\x67\x1c\xf1\xab\xfb\x5e\x10\xaa\x2e\x9b\xb2\xf7\x4c\xd5\x53\xfb\xeb\x8c\x2b\x0c\x52\x05\x20\x55\x7d\x25\x44\xdd\x77\x17\x03\xd3\xa5\xc2\xb4\x8a\x88\xba\xef\x3e\x06\xa6\x4f\x85\x69\x95\x11\xb9\x21\x4d\x30\x38\x27\x54\x9c\x76\x25\x91\x9b\x12\x0f\x51\xf5\x19\x92\x0b\xd4\xae\x25\x6a\x6f\x8a\x11\x56\x14\x03\x56\x14\x23\x0b\xd4\x1c\xa8\xbc\x09\x15\x26\xc2\x8c\x62\xc0\x8c\x20\x98\xcd\x02\x35\x87\xa8\x28\x24\x4a\x29\x46\x3c\x2c\x88\x81\x47\x05\x10\x4c\xa3\x3e\xcd\xb1\x89\x07\x1e\x51\x4b\x31\xc2\x37\x63\xc0\x37\x21\xa0\x46\x75\xba\xae\xe9\xef\xe3\x88\x28\xa6\x18\xf1\xc6\x5a\x0c\xbc\xb1\x06\x21\xf5\x8c\x16\xea\x9b\x26\xcf\x7d\x6a\x8d\x62\x1e\x12\xc7\xd0\x43\xe2\x18\xfb\xfc\xed\xf0\xe3\x49\x8c\x79\x48\x1c\x43\x0f\x89\x41\xa0\x76\x3d\x05\xd4\x91\x24\xc6\x3c\x26\x8e\xa1\xc7\xc4\x20\x52\x8c\xa4\xc6\x74\x56\x31\x9a\x82\x9e\x14\x83\x58\x31\xaa\x72\xc9\xa7\xbd\x38\xc1\xc8\x2a\x81\x74\x05\x81\xc5\x08\xcb\xdf\x1e\xf2\xb6\x2f\x85\x2f\xcb\x4b\x59\x64\x52\xd7\x2f\x84\xdf\x83\x77\x16\xf2\x4a\x94\x2b\x7d\xa1\xa4\x2e\xd7\xdb\xcb\xf5\x9f\xab\x2c\xa4\x92\xfa\xd1\x25\x66\xfc\x70\x7f\xc6\xdc\x87\xcf\xee\x1e\xf3\xe2\x6c\x64\xa1\xd2\x3c\x53\x17\x97\xb7\x17\x32\xdb\x3c\xbe\xed\x42\x6e\xf6\x5e\x65\xdf\xdd\x7c\xf7\xbd\x6a\xed\x7c\xe4\x8f\x76\x6f\xe7\xb5\xf1\x57\x47\xb0\xff\x5c\x4a\x2d\xfe\xbb\x1f\x28\xd6\x6b\xb5\x87\xbb\xbe\xba\x5b\x7d\xc5\xea\x46\x16\xed\x69\x1a\x91\x62\xa5\xb7\x71\xa7\x62\xde\x0c\x55\x5a\x14\xfa\x74\x8f\x58\x67\x3c\x72\x47\xdc\xb8\xe9\xe7\x5c\xe9\x2b\x31\xd7\xea\xfe\x76\xc1\x28\x74\x1e\x05\xdc\x35\x56\xba\x2e\x64\x55\x09\x0b\xe7\x8c\xe9\xa2\x94\x27\x2d\x71\x5d\xcc\x71\x2b\x73\xfc\x18\xcc\x71\x24\x73\xfc\xf9\x30\xe7\x5b\x99\xf3\x8f\xc1\x9c\x8f\x64\xce\x7f\x3e\xcc\x05\x56\xe6\x82\x63\x30\x17\x20\x99\x0b\x9e\x0f\x73\xa1\x95\xb9\xf0\x18\xcc\x85\x48\xe6\x42\x34\x73\xe1\xc8\x7d\xf9\x84\xcc\x05\xd6\x3e\x17\x1c\xa3\xcf\x35\xd3\x40\xcc\x05\x84\x3e\xf7\xe4\xcc\xb9\x56\xe6\xdc\x63\x30\xe7\x22\x99\x73\x9f\x0f\x73\x9e\x95\x39\xef\x18\xcc\x79\x48\xe6\xbc\xe7\xc3\x9c\xcd\x5b\x83\x63\x78\x6b\x33\x0d\xcc\x1c\xde\x5b\x9f\x9c\x39\x9b\xb7\x06\xc7\xf0\xd6\x66\x1a\x98\x39\xbc\xb7\x3e\x39\x73\x36\x6f\x0d\x8e\xe1\xad\xcd\x34\x30\x73\xcf\xc8\x5b\x23\x2b\x73\xd1\x41\x94\x45\x48\xca\xa2\xe3\x52\x16\xed\x51\x46\x23\x25\xb2\x1e\x49\x8d\x08\x12\x29\xcd\x6f\x43\xa4\x44\x84\xb3\xe8\x41\xa4\x3c\xfc\x6b\xf7\xd7\xe5\x77\x27\x77\x7f\x07\x00\x00\xff\xff\xc9\x72\xf7\xdd\xa7\x48\x00\x00")

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

	info := bindataFileInfo{name: "cloud.json", size: 18599, mode: os.FileMode(420), modTime: time.Unix(1453795200, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _credentialJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x92\xc1\x6a\xc4\x20\x10\x86\xef\x3e\xc5\x30\xe7\x7d\x82\xbd\x95\x9e\x7a\xde\xde\xca\xb2\x98\x64\x52\xa6\x35\x2a\x3a\x96\xb6\x21\xef\x5e\x34\x90\x86\x54\x2f\xdd\x8b\xc8\xff\xfd\xf2\x8d\x30\xb3\x02\x40\x1f\xdc\x07\x0f\x14\xf0\x0c\xa8\xbf\x53\x20\x3c\xe5\x78\xe0\xe8\x8d\xfe\xca\xe9\xc3\x9f\xf4\x36\xba\x30\x69\xc9\x70\x64\x32\xc3\x0a\xcb\x35\xe2\x19\x5e\x14\x00\xc0\x5c\x4e\x00\x7c\x8b\xce\xe6\xaa\x90\xd5\x56\x6e\xbc\xd6\x0b\xb2\x7a\xa2\x06\x32\xba\x23\x93\xd9\x73\x61\xf0\xb4\x63\x6c\x7d\x92\xf5\xdd\xa7\x60\x49\x97\x53\x5d\x1a\x53\x17\xfb\xc0\x5e\xd8\xd9\xaa\xba\x59\xd8\x06\xb8\xec\x1a\xff\x1d\xa3\x37\x4c\x8d\xbf\x57\xd0\xa6\x7e\x2c\xec\x4e\x69\xa4\x3e\x90\x34\xc5\x47\x7c\x94\x5f\x0e\xbc\x36\x80\x02\xb8\x96\x1d\x10\xfd\xfa\xbb\x01\xf8\x9e\x3a\x0a\x96\x84\x62\xee\x5d\xd5\xa2\x7e\x02\x00\x00\xff\xff\x6c\x74\x77\xde\x73\x02\x00\x00")

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

	info := bindataFileInfo{name: "credential.json", size: 627, mode: os.FileMode(420), modTime: time.Unix(1453795200, 0)}
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
	"cloud.json": cloudJson,
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
	"cloud.json": &bintree{cloudJson, map[string]*bintree{}},
	"credential.json": &bintree{credentialJson, map[string]*bintree{}},
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
