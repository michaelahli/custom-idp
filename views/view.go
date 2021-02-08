// Code generated by go-bindata. DO NOT EDIT.
// sources:
// views/authc/consent.html
// views/authc/login.html
// views/fe/after_login.html
// views/fe/homepage.html

package views


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
	info  fileInfoEx
}

type fileInfoEx interface {
	os.FileInfo
	MD5Checksum() string
}

type bindataFileInfo struct {
	name        string
	size        int64
	mode        os.FileMode
	modTime     time.Time
	md5checksum string
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
func (fi bindataFileInfo) MD5Checksum() string {
	return fi.md5checksum
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bindataViewsAuthcConsentHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x56\x6d\x6f\xdb\x36\x10\xfe\x9e\x5f\xc1\x72\xd8\xa7\x95\x96\x9d\xb7\x76\xaa\x64\x20\xcd\xd6\x02\xeb\x5a\xb4\xcb\xd6\x6d\x18\x86\x82\x22\x4f\x12\x6b\x8a\xd4\xc8\x93\x5f\xe1\xff\x3e\x50\xb2\x5c\x5b\x71\xda\x38\x40\x24\xde\x3d\x3c\xde\xf3\xf0\x8e\x54\xf2\x44\x5a\x81\xab\x1a\x48\x89\x95\x9e\x9e\x25\xe1\x41\x34\x37\x45\x4a\xc1\xd0\x60\x00\x2e\xa7\x67\x84\x10\x92\x54\x80\x9c\x88\x92\x3b\x0f\x98\xd2\x06\x73\xf6\x9c\x1e\xba\x0c\xaf\x20\xa5\x73\x05\x8b\xda\x3a\xa4\x44\x58\x83\x60\x30\xa5\x0b\x25\xb1\x4c\x25\xcc\x95\x00\xd6\x0e\x9e\x12\x65\x14\x2a\xae\x99\x17\x5c\x43\x3a\x79\x4a\x7c\xe9\x94\x99\x31\xb4\x2c\x57\x98\x1a\x7b\x22\xb4\x04\x2f\x9c\xaa\x51\x59\x73\x10\xfd\x04\x90\x37\x58\x5a\x77\x80\xf9\xbb\xf1\x4d\xde\x03\x51\xa1\x86\xe9\x4d\x8b\x51\x6b\x20\x37\x86\xdc\xd4\x75\x12\x75\xf6\xb3\xb3\x0e\xf5\x84\x31\xf2\xd2\x5a\xf4\xe8\x78\x4d\x84\x75\x40\x6e\xef\xee\x08\x63\xbb\x28\x5a\x99\x19\x71\xa0\x53\xea\x71\xa5\xc1\x97\x00\x48\x49\xe9\x20\x4f\x69\x89\x58\xfb\x38\x8a\x84\x34\x9f\xfd\x48\x68\xdb\xc8\x5c\x73\x07\x23\x61\xab\x88\x7f\xe6\xcb\x48\xab\xcc\x47\xb8\x50\x88\xe0\x58\xd6\x2f\x13\x5d\x8e\xae\x47\xe3\x48\x78\x1f\xed\x6d\xa3\x4a\x99\x91\xf0\x9e\x12\x65\x10\x0a\xa7\x70\x95\x52\x5f\xf2\xab\xc9\x39\x7b\x7f\xf5\xb6\x78\x6b\x26\x9f\x5f\xbe\x1b\x4f\xb8\x7f\x59\xfc\x31\x5e\x5f\x8f\x3f\xcc\x2e\x3f\x2c\xd5\x5f\xf6\xf9\xf5\x0f\x0b\xfd\x8a\x97\x6f\x9c\xff\x90\x5f\x3c\x13\x3f\x0a\x67\xaf\x26\xcf\xfe\x5c\x7f\xbc\x7b\xff\xfe\xe3\x04\x7f\x7a\xb3\x2e\x67\xcd\xb9\xfa\xe5\xfc\xd5\xc7\xe2\xd7\xeb\x67\x8b\xd7\xe3\x8b\xbb\xd7\xe6\xdd\x4d\x9a\x52\x22\x9c\xf5\xde\x3a\x55\x28\x93\x52\x6e\xac\x59\x55\xb6\xf1\x94\x44\x3b\xf6\x2d\xe7\xee\x3d\xfc\x42\xd9\x3c\xdd\x8f\x32\x2b\x57\x64\xb3\x1f\xb6\x00\x50\x45\x89\x31\x99\x8c\xc7\xdf\xbf\xd8\x7b\xb6\x67\x5f\x9b\x23\x95\xaf\x35\x5f\xc5\x84\x55\x9e\xe5\x1a\x96\x99\x5d\xbe\x38\x8d\x08\xde\x63\x57\x3f\x87\x71\xad\x0a\x13\x13\x01\x06\xc1\x1d\x63\x5a\x17\x53\x08\x95\x3f\x0d\xa8\xb9\x94\xca\x14\x0c\x6d\x1d\x93\xcb\x71\xbd\x3c\xed\xce\x2c\xa2\xad\x4e\x21\x32\x2e\x66\x85\xb3\x8d\x91\x4c\x58\x6d\x5d\x4c\xbe\xcb\xaf\xc2\xdf\x49\x0d\x46\xb9\x75\x15\x13\xd6\x78\x30\x38\x10\xa3\x6d\x9b\xa1\x7e\xe1\x57\xf1\x25\xdb\x39\x2f\x2e\x1e\x4a\x31\x26\x93\xab\xa1\xab\xe2\xae\x50\x26\x26\xbc\x41\x7b\x98\xce\xe9\x6c\x46\xa2\x04\x31\xcb\xec\x72\x90\x57\x6e\x0d\xb2\xc5\x6e\x77\x2f\xc7\xe3\x47\x44\xea\x47\xe8\xac\x1e\x44\xab\xad\x57\xa1\xbd\xe3\xd0\x59\x1c\xd5\x1c\x06\x7a\xda\x25\xf3\x6a\xdd\x32\xca\xac\x93\x6d\xf3\x0c\x78\xf5\xa5\x76\xcc\xeb\x58\x8c\x7b\x3a\xb5\x34\xbc\x5a\x43\x4c\x26\xd7\x87\xce\x47\xb1\x88\x73\x2b\x1a\x3f\xe0\xb2\x66\xca\x48\x58\xc6\xe4\xfc\xdb\xd1\x94\xa9\x1b\xfc\x27\x1c\xc4\x29\x85\x8a\x2b\x4d\xff\x1d\x44\xeb\x76\x6b\x5f\x6a\x6c\x72\xaf\xd4\x7a\x39\x02\x80\xb9\xa0\x01\x73\x5c\xaa\xc6\xc7\x64\xfc\x35\xa8\x86\xfc\x24\xf2\x31\xb9\xd6\xdc\xfb\x85\x75\xf2\x5b\xe9\xde\x17\x7c\x97\x03\xda\xfa\xc1\x04\x06\xb0\x07\x29\xed\x3a\x28\x89\x76\x87\xd2\x59\x12\x75\xb7\x56\xd2\x1e\x2a\x42\x73\xef\x53\x8a\xb0\x44\xd6\x35\x39\x0d\x98\xc0\xa9\xf7\x1d\xf2\xa3\xa4\x02\x2c\xad\x4c\x69\x6d\x3d\x52\xc2\x45\x28\xc8\x94\x46\xe1\x52\x01\x83\x4a\xf0\x60\x88\x7a\x78\x77\x08\x6e\x36\x2a\x27\xc6\x22\x19\xdd\x76\xf6\xdb\x92\x6b\x0d\xa6\x80\xed\x17\x1d\x13\xa9\xe6\x27\xd2\x21\x5c\x83\xc3\xee\x3f\x93\xdc\x14\xe0\x28\x71\x56\x87\x8b\x2c\xd8\xe8\xf4\x48\x92\x24\x9b\x6e\x36\x64\xf4\xb3\x73\xd6\xfd\x1e\xae\x2c\xb2\xdd\x26\x51\x36\x04\xb9\x63\xc3\x7e\xca\x6d\x77\x27\x92\xc3\xbc\x22\xa9\xe6\x3d\x0f\x30\x72\xbb\x53\xb4\x25\xf5\x10\xa1\xa4\x9c\xf4\x5c\xca\x0b\x52\x65\xec\xe2\xf0\x2c\x60\xc6\xba\x8a\x6b\xba\xbf\x66\x5b\xd1\x92\xa8\x9c\xec\x2e\x90\x7a\x9a\x44\x75\xbf\xa6\x0b\xa4\xc9\xe8\x37\xf8\xaf\x01\x8f\x20\xef\x84\xad\xc1\x3f\xa0\x5c\xb7\x59\xe1\x34\x1a\x0a\xd3\x56\xe6\x7d\x18\x6b\xed\x94\x74\x25\xdb\x9f\x63\x74\xf7\xad\x50\x38\x6e\xf0\x93\x0f\x2b\x52\x32\xe7\xba\x81\x94\x6e\x36\xa3\xed\x96\x12\x25\xf7\xaf\xed\x2c\x90\x83\x05\x35\xcf\x40\x9f\x58\xb0\xb5\x53\x92\x5b\xd7\x07\x98\xb6\x8f\x24\x6a\x3d\xd3\xaf\x29\x7f\xc0\xa4\x4b\xb8\x54\x52\x82\xe9\xd3\xdd\x95\xdd\x27\xd1\x6f\xc7\x61\xd2\xf7\xb7\xaa\xff\xe6\xc9\x1a\x44\x6b\xfa\x54\x33\x34\x24\x43\xc3\x74\xd1\x3e\x6a\xa7\x2a\xee\x56\xed\x7b\xa6\xad\x98\xf5\x5a\xf9\x26\xab\x14\x7e\xd9\x44\x48\xa2\x2e\xd0\x71\xca\x49\x14\xb8\xb7\x8d\x17\x3a\x6e\x1a\x1a\xb0\xfd\x9c\xfc\x3f\x00\x00\xff\xff\x19\x8d\x14\x4d\x5f\x0a\x00\x00")

func bindataViewsAuthcConsentHtmlBytes() ([]byte, error) {
	return bindataRead(
		_bindataViewsAuthcConsentHtml,
		"views/authc/consent.html",
	)
}



func bindataViewsAuthcConsentHtml() (*asset, error) {
	bytes, err := bindataViewsAuthcConsentHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "views/authc/consent.html",
		size: 2655,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1612769779, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataViewsAuthcLoginHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x57\xfb\x6f\xdb\x36\x10\xfe\x3d\x7f\x05\xcb\x61\xc3\x86\x95\x96\x9d\x67\xa7\x4a\xc6\xd6\xae\x2d\xb6\x3e\x90\x2e\x5b\xb7\x61\x18\x0a\x8a\x3c\x4b\xac\xf9\xd0\x48\x2a\xb6\x63\xe4\x7f\x1f\xa8\x87\x23\xcb\x4a\x53\x15\xa8\x25\xde\xc7\xe3\x77\x1f\xef\x8e\x4c\xf2\x88\x1b\xe6\x37\x25\xa0\xc2\x2b\x39\x3f\x4a\xc2\x0f\x92\x54\xe7\x29\x06\x8d\xc3\x00\x50\x3e\x3f\x42\x08\xa1\x44\x81\xa7\x88\x15\xd4\x3a\xf0\x29\xae\xfc\x82\x3c\xc1\x7d\x93\xa6\x0a\x52\x7c\x2d\x60\x55\x1a\xeb\x31\x62\x46\x7b\xd0\x3e\xc5\x2b\xc1\x7d\x91\x72\xb8\x16\x0c\x48\xfd\xf1\x18\x09\x2d\xbc\xa0\x92\x38\x46\x25\xa4\xb3\xc7\xc8\x15\x56\xe8\x25\xf1\x86\x2c\x84\x4f\xb5\x19\x71\xcd\xc1\x31\x2b\x4a\x2f\x8c\xee\x79\x1f\x01\xd2\xca\x17\xc6\xf6\x30\x7f\x57\xae\x5a\x74\x40\x2f\xbc\x84\xf9\x95\xc8\x35\xfa\x45\x27\x51\xf3\x79\x74\xd4\x18\x1f\x11\x82\x9e\x19\xe3\x9d\xb7\xb4\x44\xcc\x58\x40\xcf\xaf\xae\x10\x21\xed\x64\x29\xf4\x12\x59\x90\x29\x76\x7e\x23\xc1\x15\x00\x1e\xa3\xc2\xc2\x22\xc5\x85\xf7\xa5\x8b\xa3\x88\x71\xfd\xc9\x4d\x98\x34\x15\x5f\x48\x6a\x61\xc2\x8c\x8a\xe8\x27\xba\x8e\xa4\xc8\x5c\xe4\x57\xc2\x7b\xb0\x24\xeb\x96\x89\x4e\x27\xe7\x93\x69\xc4\x9c\x8b\x76\x63\x13\x25\xf4\x84\x39\x87\x91\xd0\x1e\x72\x2b\xfc\x26\xc5\xae\xa0\x67\xb3\x63\x72\x79\xf6\x36\x7f\xab\x67\x9f\x9e\xbd\x9b\xce\xa8\x7b\x96\xff\x31\xbd\x39\x9f\xbe\x5f\x9e\xbe\x5f\x8b\xbf\xcc\x93\xf3\xef\x57\xf2\x25\x2d\x5e\x5b\xf7\x7e\x71\x72\xc1\x7e\x60\xd6\x9c\xcd\x2e\xfe\xbc\xf9\x70\x75\x79\xf9\x61\xe6\x7f\x7e\x7d\x53\x2c\xab\x63\xf1\xeb\xf1\xcb\x0f\xf9\x9b\xf3\x8b\xd5\xab\xe9\xc9\xd5\x2b\xfd\xee\xa7\x34\xc5\x88\x59\xe3\x9c\xb1\x22\x17\x3a\xc5\x54\x1b\xbd\x51\xa6\x72\x18\x45\x6d\xf4\x75\xcc\xcd\x7b\x78\x26\x19\x27\xa5\xa4\x0c\x0a\x23\x39\x58\x22\x54\x8e\xb6\x3b\x6b\x78\x16\x46\x7b\xe2\xc4\x0d\xc4\x68\x36\x99\x1d\x9f\x59\x50\x4f\xf7\x00\x1e\xd6\x9e\x50\xcd\x0a\x63\x63\xa4\x04\xe7\x12\xf6\x01\x64\x05\xd9\x52\x78\x52\x39\xb0\xc4\x81\x04\xe6\x63\xa4\x8d\x1e\xc2\x94\xb9\x79\x18\xe3\x1e\x82\x7c\xc6\x7c\x7b\xb4\x7b\xfd\x51\x01\x17\x14\x7d\xab\x84\x6e\xf2\x39\x46\x17\xe7\x4f\xca\xf5\x77\x83\xe8\x47\xf4\x21\x72\x28\xd1\x40\xa6\x93\xc9\xa1\x48\xb7\x63\x24\x42\xa9\x3e\xde\x7d\x65\x86\x6f\x06\x8e\x0b\x10\x79\xe1\x63\x34\x9b\x4e\xbf\x1e\x0d\x63\x64\x0e\x17\xae\x94\x74\x13\xd7\x5a\x2d\x24\xac\x33\xb3\x7e\x3a\x8e\x08\xd6\x43\x7d\xc3\x28\xa1\x52\xe4\x3a\x46\x0c\xb4\x07\xbb\x8f\xa9\x4d\x44\x78\x50\x6e\x1c\x50\x52\xce\x85\xce\x89\x37\x65\x8c\x4e\xa7\xe5\x7a\xdc\x9c\x19\xef\x8d\x1a\x43\x64\x94\x2d\x73\x6b\x2a\xcd\x09\x33\x32\xa4\xd5\x57\x8b\xb3\xf0\x6f\x54\x83\xc9\xc2\x58\x45\x9c\xc8\xb5\xd0\x03\x2d\xda\x9d\xdd\x97\x2f\x3c\x8a\xae\xbb\x6d\x3f\x39\xb9\x8f\x61\x8c\x66\x67\x43\x93\xa2\x36\x17\x3a\x46\xb4\xf2\xa6\xcf\x66\x94\xcc\x84\x15\xc0\x96\x99\x59\x8f\x95\xd4\xaa\xdd\xdb\xd3\xe9\xf4\x61\x47\xf5\x47\xe8\x85\xd6\xc8\x81\xb3\xd2\x38\x11\xfa\x69\x1c\x7a\x1a\xf5\xe2\x7a\x50\x0f\x99\x59\x87\xbc\xac\xe3\xc9\x8c\xe5\x75\xdb\x1a\x44\xd5\xe5\xd9\x7e\x54\xfb\x52\x1c\xa8\xd4\x6f\x0c\xe7\x7d\xe3\x97\x04\x11\x2f\x0c\xab\xdc\x20\x94\x1b\x22\x34\x87\x75\x8c\x8e\x1f\x74\x26\x74\x59\xf9\x7f\xc2\xb9\x97\x62\x50\x54\x48\xfc\xef\xc0\x59\xb3\x53\xbb\x2c\x23\xb3\x83\x2c\xeb\xc4\x08\x00\x62\x83\x02\xc4\x52\x2e\x2a\x17\xa3\xe9\xe7\xa0\x12\x16\xa3\xc8\x2f\xa0\x5a\x52\xe7\x56\xc6\xf2\x87\xd8\x1e\xaa\xdd\x52\xf0\xa6\xbc\x77\xfd\x01\xec\xde\x88\xda\xda\x49\xa2\xf6\x2c\x38\x4a\xa2\xe6\x8e\x90\xd4\xed\x84\x49\xea\x5c\x8a\xeb\xb6\xde\x94\x37\x0e\x98\x10\x52\x67\xeb\x85\x87\x91\x02\x5f\x18\x9e\xe2\xd2\x38\x8f\x11\x65\x21\x19\x53\x1c\x85\x13\x1c\xb4\x17\x8c\x86\x81\x48\x9a\x5c\xe8\xf6\xfc\xde\x6e\xc5\x02\x69\xe3\xd1\xe4\x4d\x18\x7d\x5e\x50\x29\x41\xe7\x70\x7b\xa7\x60\xc2\xc5\xf5\x08\x13\x44\x25\x58\xdf\xfc\x4f\x38\xd5\x39\x58\x8c\xac\x91\xe1\xc6\x10\xc6\xf0\x7c\x4f\x8d\x24\x9b\x6f\xb7\x68\xf2\xc2\x5a\x63\x7f\x0f\x97\x04\x74\x7b\x9b\x44\xd9\x10\x64\xf7\x07\x76\x53\x9e\x37\x97\x0f\xd4\xe7\x15\x71\x71\xdd\x45\x01\x9a\xdf\xb6\x62\xd6\x21\x8d\x87\x93\x14\xb3\x2e\x92\xe2\x04\xa9\x8c\x9c\xf4\x1b\x00\xd1\xc6\x2a\x2a\xf1\xfc\x52\x02\x75\x80\x82\xaa\x48\xe8\x24\x2a\x66\xed\x89\x5d\xe7\x0f\x6a\xf2\xa7\x10\x9c\x83\xc6\xed\x25\xa9\xd6\xf4\x23\xeb\xd6\xc3\xe8\x9a\xca\x0a\x52\xbc\xdd\x1e\x30\xe9\x6e\x4e\x92\x66\x20\xd1\xc2\xd8\x14\xd7\x7e\x5f\xd4\xb5\xd3\xf1\x73\x96\x18\x2d\x37\x78\x5e\x0f\x23\xca\xb9\x05\xe7\x92\xa8\x9e\x35\x42\xa7\xa9\x3c\x24\xf8\xa8\xb7\x7e\xbd\x77\x94\xdb\x19\xbd\x23\x35\xc5\x7b\x8b\x61\x64\xe1\xbf\x4a\x58\xe0\x75\x37\xaa\xdb\xc4\x3d\xdc\x2f\xbb\x62\x3a\xa0\xdf\x59\xee\x67\xbe\x2b\xc4\x3b\xf2\x07\xee\xc6\xf8\xdf\xcd\xdb\x0b\xe1\x6e\x6e\xc7\xbe\x5d\xb3\x97\xc6\xbb\xc3\x20\xa4\x40\x2f\x51\x93\x1e\xc7\xdd\x58\x9f\x6b\x37\xb1\xe3\x60\x41\x81\xca\xc0\x7e\x54\x77\x5b\xee\x6d\x05\x78\x8e\x7e\x6b\x4d\x48\x41\x2f\x65\xfb\x22\xdc\xe5\x6f\x92\x55\xde\x1b\xdd\xd1\xcb\xbc\x46\x99\xd7\xe1\x76\x13\x7e\x4a\x2b\x14\xb5\x9b\xfa\x3d\x93\x86\x2d\x71\xcb\xc6\x55\x99\x12\x1e\x37\xb7\xef\x90\xa8\x8d\x9b\xd6\x67\xd9\xb9\x53\x9e\x9c\x35\xc9\x5e\x97\xaf\xaa\x3c\x70\x3c\xff\x86\x99\x72\xf3\x14\x1d\x4f\x8f\x67\x49\x54\xee\xd7\x51\x12\x05\xc1\xc3\x35\xfe\x28\x89\x42\x27\x9a\x87\xc6\x54\xff\x51\xf3\x7f\x00\x00\x00\xff\xff\xbf\x17\xd0\x5f\xe5\x0c\x00\x00")

func bindataViewsAuthcLoginHtmlBytes() ([]byte, error) {
	return bindataRead(
		_bindataViewsAuthcLoginHtml,
		"views/authc/login.html",
	)
}



func bindataViewsAuthcLoginHtml() (*asset, error) {
	bytes, err := bindataViewsAuthcLoginHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "views/authc/login.html",
		size: 3301,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1612408792, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataViewsFeAfterloginHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x91\x41\x4f\xc4\x20\x10\x85\xef\xfd\x15\x23\x77\xe5\x6a\x36\xb3\x24\x1b\x77\x3d\x99\x68\x0c\x1e\x3c\xb2\xed\x6c\x4b\xa4\xd0\x00\x26\x36\x4d\xff\xbb\x01\x7a\xa8\xa9\xbd\xc0\xcc\x7b\x7c\x6f\x02\xe0\xdd\xf9\xf5\x49\x7e\xbe\x5d\xa0\x8b\xbd\x11\x15\xa6\x0d\x8c\xb2\xed\x91\x91\x65\x49\x20\xd5\x88\x0a\x00\x00\x7b\x8a\x0a\xea\x4e\xf9\x40\xf1\xc8\x3e\xe4\xf3\xfd\x23\x5b\xac\xa8\xa3\x21\x71\xba\x45\xf2\xf0\xe2\x5a\x6d\x91\x17\xa9\x42\x5e\x12\xf0\xea\x9a\xb1\x9c\x3e\x53\x54\xda\x80\xa7\x30\x38\x1b\xe8\x50\x22\xae\x7e\xc9\xfa\x36\xa5\xc8\x8d\xd1\xe2\x54\xd7\x14\x02\x48\xf7\x45\xf6\x00\xd3\xf4\x50\x84\xdc\xaf\xeb\x79\x46\x6e\xf4\x5f\x36\x1b\x20\xc7\x81\x36\x64\x5e\x93\xf3\x1f\xf7\x4e\x37\x4f\xa1\xdb\x19\xba\xb8\xbb\x53\x2f\x3f\x83\xf6\xd4\x6c\xb8\xac\x8f\x6b\x02\x79\xba\x2d\xf2\xf2\x38\xc8\xf3\x2f\xfc\x06\x00\x00\xff\xff\x0d\x25\xb8\x77\x95\x01\x00\x00")

func bindataViewsFeAfterloginHtmlBytes() ([]byte, error) {
	return bindataRead(
		_bindataViewsFeAfterloginHtml,
		"views/fe/after_login.html",
	)
}



func bindataViewsFeAfterloginHtml() (*asset, error) {
	bytes, err := bindataViewsFeAfterloginHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "views/fe/after_login.html",
		size: 405,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1611913868, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataViewsFeHomepageHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\x8f\x3b\x4f\xc4\x30\x10\x84\xfb\xfc\x8a\xc5\x3d\xb8\xa5\xd8\x58\x42\x3c\x84\xc4\xf1\x10\xba\x2b\x28\x17\xbc\x38\x2b\xdb\x9b\x53\x62\x17\x28\xca\x7f\x47\xc6\x54\xfb\x98\x19\x7d\x1a\xbc\xb8\x7b\xbd\x3d\x7e\xbc\xdd\xc3\x54\x72\x72\x03\xb6\x01\x89\x34\x8c\x86\xd5\xb4\x07\x93\x77\x03\x00\x00\x66\x2e\x04\x5f\x13\x2d\x2b\x97\xd1\x9c\x8e\x0f\x97\xd7\xe6\x5f\x2a\x52\x12\xbb\xc7\x39\xf3\x99\x02\xa3\xed\xf7\x80\xb6\xc7\xf1\x73\xf6\x3f\xdd\xfa\x94\x24\x82\x17\x58\x45\x05\xaa\x96\x1a\x21\x73\xa2\x58\x23\x29\x20\xc1\xb4\xf0\xf7\x68\xb6\xed\xea\x30\x07\xd1\xd3\xfb\x61\xdf\x8d\x4b\x6d\x07\xcf\x1a\x48\xe1\x85\x32\xdd\x9c\x93\x44\x5a\xe5\xb9\xa2\xa5\xc6\xe9\x00\xb4\x7f\x35\x7e\x03\x00\x00\xff\xff\xe8\xed\x68\x59\xd6\x00\x00\x00")

func bindataViewsFeHomepageHtmlBytes() ([]byte, error) {
	return bindataRead(
		_bindataViewsFeHomepageHtml,
		"views/fe/homepage.html",
	)
}



func bindataViewsFeHomepageHtml() (*asset, error) {
	bytes, err := bindataViewsFeHomepageHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "views/fe/homepage.html",
		size: 214,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1611906036, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}


//
// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
//
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
// nolint: deadcode
//
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

//
// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or could not be loaded.
//
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// AssetNames returns the names of the assets.
// nolint: deadcode
//
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

//
// _bindata is a table, holding each asset generator, mapped to its name.
//
var _bindata = map[string]func() (*asset, error){
	"views/authc/consent.html":  bindataViewsAuthcConsentHtml,
	"views/authc/login.html":    bindataViewsAuthcLoginHtml,
	"views/fe/after_login.html": bindataViewsFeAfterloginHtml,
	"views/fe/homepage.html":    bindataViewsFeHomepageHtml,
}

//
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
//
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, &os.PathError{
					Op: "open",
					Path: name,
					Err: os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op: "open",
			Path: name,
			Err: os.ErrNotExist,
		}
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

var _bintree = &bintree{Func: nil, Children: map[string]*bintree{
	"views": {Func: nil, Children: map[string]*bintree{
		"authc": {Func: nil, Children: map[string]*bintree{
			"consent.html": {Func: bindataViewsAuthcConsentHtml, Children: map[string]*bintree{}},
			"login.html": {Func: bindataViewsAuthcLoginHtml, Children: map[string]*bintree{}},
		}},
		"fe": {Func: nil, Children: map[string]*bintree{
			"after_login.html": {Func: bindataViewsFeAfterloginHtml, Children: map[string]*bintree{}},
			"homepage.html": {Func: bindataViewsFeHomepageHtml, Children: map[string]*bintree{}},
		}},
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
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
