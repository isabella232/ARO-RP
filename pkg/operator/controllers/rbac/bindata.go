// Code generated for package rbac by go-bindata DO NOT EDIT. (@generated)
// sources:
// staticresources/clusterrole.yaml
// staticresources/clusterrolebinding.yaml
package rbac

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

// Mode return file modify time
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

var _clusterroleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x59\x5d\xaf\xdb\xbc\x0d\xbe\xcf\xaf\x30\xde\x5d\xbc\xc0\x80\x93\x62\xd8\xcd\x70\x76\xd9\x0e\xc3\x80\x61\x05\x8a\x6e\xf7\x8c\xcc\xd8\x6c\x64\x51\x95\xa8\x9c\x66\xbf\x7e\x90\x2c\x39\x76\xbe\xec\xe3\xac\xbd\x8a\x45\x51\x7c\x48\x89\xe2\x87\xf2\x87\xea\x23\xd7\x58\x35\x68\xd0\x81\x60\x5d\xed\x4e\x55\x0b\xea\xf0\xa1\x41\x53\x93\x57\x7c\x44\x77\x52\xa0\x5a\xfc\x6b\xf5\xe9\x73\xf5\xaf\xcf\x5f\xab\xbf\x7d\xfa\xc7\xd7\xed\x06\x2c\xfd\x07\x9d\x27\x36\xaf\x95\xdb\x81\xda\x42\x90\x96\x1d\xfd\x17\x84\xd8\x6c\x0f\x7f\xf1\x5b\xe2\x0f\xc7\x3f\x6d\x0e\x64\xea\xd7\xea\xa3\x0e\x5e\xd0\x7d\x61\x8d\x9b\x0e\x05\x6a\x10\x78\xdd\x54\x95\x72\x98\x16\x7c\xa5\x0e\xbd\x40\x67\x5f\x2b\x13\xb4\xde\x54\x95\x81\x0e\x5f\x2b\x7f\xf2\x82\xdd\x2b\x38\x7e\xf1\x0e\x37\x2e\x68\xf4\xaf\x9b\x97\x0a\x2c\xfd\xdd\x71\xb0\x3e\x0a\x79\xa9\x7e\xfb\x6d\x53\x55\x0e\x3d\x07\xa7\x30\xd3\x14\x77\x96\x0d\x1a\xf1\x02\x12\x3c\xfa\x4d\x55\x1d\xd1\xed\xf2\x74\x83\x92\x7e\x35\x79\x59\x2a\xd0\xec\xa9\xe9\xc0\xfa\x34\x44\x53\x5b\x26\x23\x79\x74\xc4\xf2\xa9\xa9\x23\x71\x60\x1a\xec\xc7\xd1\x12\x6f\x41\x95\x21\xd7\xf9\xcb\xc6\x0d\xf4\x82\x46\x8e\xac\x43\x87\x4a\x03\x75\xb7\xa7\x32\x95\xeb\xe1\x43\xb0\xb3\x1a\x24\xcf\x38\xb4\x9a\x54\xda\x4a\xc5\x46\x1c\x6b\x8d\xae\x4c\xf5\x56\x7c\x0f\x2c\xd0\x93\x3c\xba\x23\x29\x04\xa5\x38\x14\xad\x33\xed\xd1\x2e\xc5\x8f\x37\x10\xd5\x2e\xdb\xaf\xa8\xed\x07\xcd\xcd\xb5\xc4\xab\xe5\x50\x77\xe4\xa3\x33\x39\x6c\xc8\x8b\x1b\x3b\xd1\xb5\xe0\x2e\x08\x08\x99\xe6\x0d\x77\x2d\xf3\xa1\x3f\x97\xd0\x2f\xea\x8d\x39\x82\xa6\xfa\x21\xcf\x0a\x1b\xc1\x12\xfe\x10\x34\x51\x4f\x7f\x57\x39\x15\xbc\x70\x57\x88\x35\xee\xc9\xd0\x73\xa0\x8b\xf6\x04\x2c\x3d\x77\x82\x60\x93\x5f\x5f\xfb\x7c\x76\x26\x87\x47\xf2\xc3\x06\xd7\x80\x1d\x1b\x8f\xd9\x79\x6a\xb4\x9a\x4f\xdd\x70\x05\xb2\x3b\x0e\xf3\xf1\x0a\xe2\x3e\xe8\x4c\x58\xa9\xde\x96\x2d\x1a\xdf\xd2\x5e\x6e\x6e\xc1\x59\x89\xfe\xb4\x7f\x21\xd2\x52\x37\x77\x3c\x23\x59\xf5\x71\x72\xad\xea\x41\x5a\x34\x92\x03\xc1\x5d\x5f\x11\x3e\xa0\x89\xe7\x89\x6f\x17\x40\x29\x1c\xe3\x6d\xc1\x97\xc1\xfd\x5a\xae\x47\xbd\xf7\x61\xf7\x0d\x95\x80\x52\xe8\xfd\x19\x63\x32\x99\xa2\xf8\x64\xee\xf6\xa2\x77\x2b\xb6\x68\x6f\x1d\x6b\xdc\x91\xa9\xc9\x34\xfe\x92\x9e\xbd\xf7\x92\xa3\x4c\x2d\x4e\x1f\xef\x51\xab\x0c\x6f\x6c\xd9\x2f\xd9\x96\x91\xb5\x0e\xbd\x38\x52\xcf\x84\xab\x20\xec\x15\x68\x32\xcd\x35\x52\x52\x89\x8d\x80\xb6\x5c\x17\xce\x67\x9c\xbd\x40\x2d\x3b\xf8\x29\xe2\x4b\xd5\x81\x6a\xc9\xe0\xd3\x8a\xec\x12\xf9\x1a\xd5\xb1\xf9\xc6\xbb\x1e\x2b\x7f\xac\x91\x1e\x48\xd7\x33\x06\x26\x9e\x73\xd0\xcb\x84\x9f\x0d\xb8\x34\xea\x29\x74\x42\xfb\x18\x94\xf0\x41\xda\x1c\x31\x51\x63\x92\x33\x7e\x0f\xe8\xd7\xe6\x0b\xa5\x39\xd4\xca\x61\x1d\xe3\x21\xe8\x39\x0f\x19\x18\xfd\x93\xb0\xe9\x10\x66\xd0\x72\xae\x2e\x7e\x38\x8d\xda\x17\x07\x38\x78\x2f\xdb\x58\xa1\xb3\x9b\x10\x8f\x7d\x09\xee\x0b\xb6\x1f\x62\x58\x6d\x7c\xfe\xda\x23\x48\x70\xd8\x0c\xb5\x22\x75\x50\x4a\x53\x32\x7b\x07\x5e\x5c\x50\x91\xa5\xd0\x62\x1c\x28\xab\x0d\xca\x1b\xbb\x43\x3f\xe0\xa8\x6a\xfe\xcc\xea\xb4\x21\x7b\xb8\x75\x1c\x83\xd5\x30\xf8\x41\x59\x82\x57\x2d\xd6\x61\xfd\xf5\xca\x66\xcd\x9d\x60\xcf\xa5\x34\xd5\xfc\x66\x34\x43\x3d\xd9\x94\x58\xb8\x39\x03\x5a\x73\xa3\xc9\x1c\x26\x73\x57\x04\xc3\xd9\x15\x2f\xb7\xf6\x04\x9d\xf6\xd0\xd9\xc7\xd9\xe0\xa1\x2d\xec\x6a\x32\x8f\xd3\xa8\x46\x78\xdc\xad\x3c\x00\x18\x7a\xb6\xbb\xd2\x87\xb6\x45\xaf\xaf\x15\xfb\x66\xe7\x3e\x44\xe9\x85\xd6\x88\x1e\x2a\xec\x6b\xb9\x63\xcf\x5c\x21\x7a\xaf\xf9\x2d\xd7\xb3\xdb\xe1\x0a\xde\x35\x22\x72\x47\xd7\xed\xa0\xb8\x34\xb1\x23\x39\x69\x3c\xa2\xfe\x7f\xb4\x13\xe9\x1a\xce\x78\xf5\xf8\xaa\xa6\x4f\x71\x08\xdd\x4f\x07\x4c\x28\x03\xf6\x5c\x7c\x7f\xa7\x5c\x81\x66\x64\x51\x1e\x2d\xae\xac\xd2\xa2\xdc\x11\x9d\xb6\x25\x0a\xcd\x07\x87\x21\x35\x26\x01\xd6\x05\xb3\x3a\x20\x65\x37\x5c\x0a\x5e\x1b\xef\x50\xb1\x5b\x9b\x8f\xa3\x83\x2a\x43\x5b\x65\xd4\xfe\x26\x40\x0e\xd1\x2f\x20\x02\xaa\x8d\x5d\xc9\xcb\xd3\x6d\x67\xae\x8d\x66\x4c\xcb\x5c\x2d\x82\x96\x56\xb5\xa8\x0e\x93\xca\x6a\x3a\x58\xdd\xfa\x65\x01\x93\x3b\x37\x7f\xe0\x02\x64\xd0\xb9\x60\x84\x3a\x1c\x3b\xc0\xb9\xa5\x1d\x53\x0f\x61\x87\x1a\x65\x4c\x9a\xe0\x5a\x66\x7d\x83\xbc\xd6\x24\x14\xd0\x7f\xbe\x5d\x5a\x81\xc3\x34\xdd\xb2\x3f\x67\xd3\xbe\xf7\xce\x0d\xc9\x3a\x40\x47\xea\x7e\xc8\x1e\x3d\x4a\xf1\x43\x47\xbd\x96\x4c\xcd\xcc\xd3\x84\x17\x76\xe9\xf6\x0f\x15\x48\xa6\xe4\xea\x65\x90\xb0\xd6\x36\x36\x24\xec\x62\x1f\xa0\xd8\x21\xfb\xad\xe2\xee\x46\xd5\xa5\xd1\x49\x07\x26\xc6\x8f\xc1\xd4\xbc\x76\xd8\xe7\x0e\xa5\xc5\x50\x8a\x9f\x33\x21\xb5\xac\xbd\xee\xfd\x2b\xcb\x64\xa1\xb4\x60\x38\xf1\xac\x8d\x29\xf9\x0e\x2f\x6b\x64\x26\x35\x19\xa6\x60\x94\x49\x96\x35\xa9\x52\x78\x25\x07\x0a\x3b\x53\x1e\x61\x0c\xca\xe4\x29\xf2\x39\x35\x97\x84\xbe\xc2\x66\x0f\xf4\x1c\x60\x3c\xdb\x7b\xee\x95\xc3\xb1\xd2\x30\x14\xad\x37\x4b\xd8\xd1\xde\xac\x51\x84\x6b\xbc\xab\x42\x89\x32\x83\x0a\x2b\x00\x16\x6e\xe9\xad\x66\xe1\xa2\xd7\xf1\x57\x39\x6f\xd2\x19\x28\x4f\xde\x80\xf5\x2d\xcb\xe5\x53\xf1\xb9\x69\x40\x51\xf5\x28\x5f\x46\xc6\x58\xd3\x25\x4d\x26\x4e\x56\x76\xff\x42\x52\x0c\xa7\x17\x6d\x4e\x24\x9d\xd9\x26\x17\x31\x4e\x4d\xda\x84\x4c\xba\x11\x26\xca\x95\x9b\xf6\x25\x65\xcb\x2e\x20\x07\xfa\x1d\xdc\x7c\x97\x15\x3c\x79\x66\xfe\x61\xe0\x51\x20\xa0\xb9\xc9\xb4\xf1\x45\xce\xf8\x93\x2e\x8e\x8c\x17\xd0\xda\x6a\x30\xd3\x4e\xab\x49\xc0\x53\xdc\x91\x48\x1f\x76\x5e\x39\xb2\x4f\x04\x52\x0b\xea\x10\x03\xf5\x76\x99\x5d\x99\xbd\x03\x43\xfb\x99\x8e\xf9\x1a\x2a\xfa\xd0\xe9\xe6\x7f\x06\x35\x79\x17\x92\x19\xbb\x50\x37\x25\x78\xc5\xb4\x84\x2a\xc4\x0a\xfc\xb9\x8b\x9c\xfb\xd4\x99\x6b\x96\xb9\x6e\x3f\x06\xdc\x31\xe9\x3d\x82\x57\x2a\x9f\xfe\xc3\x99\x7d\x5d\xb0\x9a\xb0\x2e\xaf\x9a\x97\xff\xfe\x2c\x3e\xa2\x25\x58\xef\x05\x79\x60\xda\xdd\x3f\x11\x7f\xe5\x43\xee\x23\xfd\x38\xc8\x5c\x3d\x9c\x78\x56\xca\xcf\xef\x0a\xef\x6f\x69\xd6\x60\xf5\xb1\xf6\x51\x56\x2d\xed\xee\x53\x39\xad\xdc\xd8\xb9\x6d\x03\xd3\x20\x68\xcd\xe3\xa4\x56\xd6\xa6\xd4\xf3\x23\x46\x70\x2f\x0e\x68\xf5\xc3\x42\xc9\x79\xdb\x9c\x55\xee\x5a\x9e\xff\x70\x2d\x29\x72\x54\x56\x5c\xcc\xe4\x9c\x78\x63\x6a\xad\x8a\x33\x9a\x29\x4f\xb5\xa3\x21\xbb\x29\x4f\xe7\xb2\x3d\xaf\xbd\x56\xf7\xdc\x15\xae\xd4\xaa\xfc\xd1\x3c\xf7\x24\xec\xf8\x80\xae\x30\xa7\x2c\x66\x4a\x6e\x7a\x4c\x5d\xab\x57\x30\x38\xf7\x4e\x6d\x1d\xef\xa9\x44\x83\xb4\x60\x25\x58\xf0\x38\x77\x23\x47\xf9\x99\x52\x25\x26\xa5\x46\x8a\x8b\x7f\x1e\x6e\x64\xc9\x80\xa7\x0e\xac\xbd\xdd\x24\x5e\x49\x7e\x6b\xd1\x21\xec\x38\xc8\xcc\xf3\x02\xd9\xd2\xfb\x2e\xd1\xdf\xb0\xf9\x92\xd7\xff\xfb\xcb\x3f\x33\xf7\xef\x7f\xfc\xfd\x7a\xf9\xff\x02\x00\x00\xff\xff\x8d\xeb\xe9\x4e\x7b\x22\x00\x00")

func clusterroleYamlBytes() ([]byte, error) {
	return bindataRead(
		_clusterroleYaml,
		"clusterrole.yaml",
	)
}

func clusterroleYaml() (*asset, error) {
	bytes, err := clusterroleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "clusterrole.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _clusterrolebindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\xcd\xbd\x0a\xc2\x40\x0c\x07\xf0\xfd\x9e\x22\x2f\xd0\x8a\x9b\xdc\xa8\x83\x7b\x41\xf7\xb4\x8d\x1a\xdb\x26\x47\x92\x13\xf4\xe9\x45\x70\x93\x3a\xff\x3f\x7e\x58\xf8\x4c\xe6\xac\x92\xc1\x7a\x1c\x5a\xac\x71\x53\xe3\x17\x06\xab\xb4\xd3\xce\x5b\xd6\xcd\x63\x9b\x26\x96\x31\xc3\x61\xae\x1e\x64\x9d\xce\xb4\x67\x19\x59\xae\x69\xa1\xc0\x11\x03\x73\x02\x10\x5c\x28\x83\x3f\x3d\x68\xc9\x68\xda\xb8\x51\x32\x9d\xa9\xa3\xcb\x27\xc7\xc2\x47\xd3\x5a\xfe\x58\x09\xe0\x87\x5a\x7b\xf6\xda\xdf\x69\x08\xcf\xa9\xf9\x8e\x4e\x4e\xb6\xd6\x7e\x07\x00\x00\xff\xff\xc4\xb6\x1b\x05\xeb\x00\x00\x00")

func clusterrolebindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_clusterrolebindingYaml,
		"clusterrolebinding.yaml",
	)
}

func clusterrolebindingYaml() (*asset, error) {
	bytes, err := clusterrolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "clusterrolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"clusterrole.yaml":        clusterroleYaml,
	"clusterrolebinding.yaml": clusterrolebindingYaml,
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
	"clusterrole.yaml":        {clusterroleYaml, map[string]*bintree{}},
	"clusterrolebinding.yaml": {clusterrolebindingYaml, map[string]*bintree{}},
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