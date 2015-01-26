package tmpl

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
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

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _layout_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x56\x41\x77\x22\x37\x0c\xbe\xf3\x2b\xbc\xde\x1e\x77\xc6\x01\xb2\x21\xc9\x0e\x1c\xba\xed\xad\x7d\xe9\xcb\xa3\x87\x1e\x35\xb6\x00\x67\x8d\x67\xd6\xf6\x10\x78\x3c\xfe\x7b\xe5\x99\x81\x18\xba\xa4\x3d\x35\x1c\xc6\x96\x3e\xc9\x9f\x24\x4b\xce\x7e\xaf\x70\xa1\x2d\x32\xfe\xfc\xf4\x34\xe7\x87\xc3\xa0\xf8\xf0\xcb\xd3\xd7\xf9\x5f\x7f\xfc\xca\x56\x61\x6d\x66\x83\xa2\xfb\x30\x56\xac\x10\x54\x5c\xd0\xb2\x04\x8f\x6c\xe5\x70\x31\xe5\xfb\x7d\xfe\x33\xed\xfe\x7c\xfe\xed\x70\xe0\xbd\x3a\xe8\x60\x70\xb6\xdf\x07\x5c\xd7\x06\x02\x79\x9f\x47\x09\x67\x3f\x1d\x0e\x85\xe8\xb4\x1d\x72\x8d\x01\x98\x5c\x81\xf3\x18\xa6\xbc\x09\x8b\xec\x9e\xa7\x2a\x0b\x6b\x9c\xf2\x8d\xc6\xd7\xba\x72\x81\x33\x59\xd9\x80\x96\xa0\xaf\x5a\x85\xd5\x54\xe1\x46\x4b\xcc\xda\xcd\x27\xa6\xad\x0e\x1a\x4c\xe6\x25\x18\x9c\x0e\x8f\x8e\x8c\xb6\xdf\x98\x43\x33\xe5\x7e\x45\x4e\x64\x13\x98\x26\x3f\xbc\x0f\x40\x41\x80\x47\xbd\x86\x25\x8a\x6d\x16\x15\x5f\x3e\x71\x16\x76\x35\x9d\x9b\x4a\xf9\x8c\x25\xee\x3a\x53\x21\x2c\x06\x65\x21\x2f\xab\x2a\xf8\xe0\xa0\x96\xca\xe6\xb2\x5a\x8b\x93\x40\x8c\xf3\x71\x3e\x14\xd2\xfb\x37\x59\xbe\xd6\x84\xf2\x9e\xf7\xac\xc2\xce\xa0\x5f\x21\x86\x23\xe3\x56\xd2\xad\x59\x59\xa9\x1d\xdb\xb3\x1a\x94\xd2\x76\x99\x85\xaa\x7e\x64\x9f\x6f\xea\xed\x17\x76\xe8\x00\xb9\xba\x21\x7d\x09\xf2\xdb\xd2\x55\x8d\x55\x99\xac\x4c\xe5\x1e\xd9\xc7\xb2\x2c\x13\xd0\xf0\xc7\x20\x29\xcb\x49\x39\x49\x70\xa3\x6b\x38\x98\x40\x8a\x1b\x5f\xc3\x3d\xd0\x5f\x82\xbb\xbd\x86\xbb\x1f\xdd\x8f\x12\xdc\xe7\x6b\xb8\xc9\xdd\xe4\x2e\xc1\xdd\x5d\xc3\xdd\x2e\x6e\x17\x09\x6e\x72\x0d\x37\x7e\x18\xa7\xfc\xee\xaf\xe1\x46\x30\x82\x04\xf7\x70\x0d\x37\xa4\x5f\x9a\xe7\x2b\xd5\x90\xf2\x06\x6f\xf0\x08\x2c\x44\x52\xe3\xc2\x4b\xa7\xeb\xc0\xbc\x93\x53\xbe\x0a\xa1\xf6\x8f\x42\xc8\x4a\x61\xfe\xf2\xbd\x41\xb7\x6b\x6f\x54\xb7\xcc\x46\xf9\x90\x7e\x2f\x9e\xcf\xc8\x45\x6b\xf6\x9e\x0f\x65\x5f\x7c\x2e\x4d\xd5\xa8\x85\x01\x87\xad\x23\x78\x81\xad\x30\xba\xf4\x42\x8d\xe9\x76\xde\xe6\xc3\x31\xad\xfe\xab\x4b\x07\xaf\x4b\x1d\x3a\x46\x1a\x56\x0d\xd8\x25\x59\x67\x41\xaf\x91\x1a\x03\xc5\x1a\x7c\x40\x27\xc8\x26\x15\xff\xc3\x7b\x21\x8e\xd3\xa4\x88\x17\xbc\x3f\xd0\xc2\x86\x49\x03\xde\x4f\x39\x2d\x4b\x70\xac\xfb\x64\xda\x6e\x90\x46\xc4\x71\xbb\xd0\x5b\x54\xb1\x15\xa8\x83\x2a\x6a\xf5\x88\xd6\x4b\x08\x3a\xb6\x69\x57\x08\x56\x28\x7d\x72\x16\xa7\x06\x10\x0d\x77\xd2\x9e\xeb\x7b\xb7\x91\xd2\x19\x26\xb2\x6b\x42\xa8\x6c\x3f\x0f\xba\x0d\xbf\x30\x0b\xd5\x72\x69\x90\x46\x93\x31\x50\x7b\x54\x9c\xc5\x91\xd2\x8b\xe3\xe1\x9d\xfc\x28\x06\xb7\x8c\xb3\xee\x63\x67\xcd\x19\x38\x0d\x19\x6e\x6b\xb0\x0a\xd5\x94\x2f\xc0\x44\x6c\x2b\x8d\xbc\x29\xc0\xd3\x51\x67\xd4\x62\x81\xc8\xe8\x48\xc6\xbb\xac\xb2\x66\xc7\x67\xf3\x8e\xce\x5b\x4a\x28\xef\x84\x7b\xc7\x34\x8e\xb7\xac\x75\xff\x7f\x41\x0b\xd1\xa5\xf2\x4c\x06\x17\x79\x2d\x1d\xa5\x84\xff\xf0\x99\x81\xba\xa6\x41\x2a\xb1\x10\x90\x54\x54\x50\x49\x2f\x0a\xac\xd5\x29\x77\x6f\x97\xa1\xab\xc7\xf1\x32\x9d\xea\x73\x46\xa6\x31\x09\x9b\x23\x94\x3e\x97\x25\x30\x9a\xed\xf7\x7a\xc1\xf0\x3b\xcb\xbf\x36\xce\xd1\xe3\xf4\x5c\x35\xf1\xc9\x6b\x09\x12\xd1\xbc\x5d\x78\x7a\x5b\x7b\x87\x20\x83\xde\x20\x45\x84\x56\x1d\x0e\x33\x8a\xbb\x0b\xb1\xc7\xcd\xe6\xed\x37\x86\x56\x50\x9f\x9e\xa7\xad\x31\x97\xf1\x16\x1f\xb2\x4c\xe4\xc4\xec\x14\x08\xcb\xb2\x53\x13\xbc\xa5\xa4\x10\x84\x99\x0d\x06\xff\xde\x1b\xe9\xbb\xfd\x3b\xa9\xda\x67\x7b\x70\xe6\x8e\xea\xd7\x76\x2d\xb5\x71\xfb\xdf\x41\x1f\xcb\xe0\xef\x00\x00\x00\xff\xff\xda\xad\x8f\xd8\x4f\x08\x00\x00")

func layout_html_bytes() ([]byte, error) {
	return bindata_read(
		_layout_html,
		"layout.html",
	)
}

func layout_html() (*asset, error) {
	bytes, err := layout_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "layout.html", size: 2127, mode: os.FileMode(436), modTime: time.Unix(1417495397, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _root_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x94\xdf\x4f\xdb\x30\x10\xc7\x9f\x9b\xbf\xe2\x96\xa7\x4d\x6a\x12\xc1\xc3\x1e\x20\x64\x9a\x34\xa9\x4c\x1a\xd2\x26\x2a\xf1\xec\x24\x97\xc4\xc3\xb1\x3d\xfb\x42\xe9\xaa\xfe\xef\x3b\xbb\x2d\xd0\x01\xd3\x78\x4a\x72\xbf\xfc\xb9\xef\x9d\xb3\xd9\xb4\xd8\x49\x8d\x90\x2e\x25\x29\x4c\xb7\x5b\x61\x2d\x39\xd1\xe0\x66\x83\xba\xdd\x6e\x93\xe4\x31\xe4\x4a\x48\x9d\x06\x53\xe9\x69\xad\xb0\x4a\x72\x3f\xd5\xb0\x49\x66\xa3\x70\xbd\xd4\x99\xc2\x8e\xce\xe0\x04\xc7\xf3\x64\xd6\x19\x4d\x99\x97\xbf\x91\x0d\x1f\xed\xfd\x79\xb2\x4d\xf2\xce\x18\x42\x17\x12\x08\xef\x29\x13\x4a\xf6\xfa\x0c\x1a\xd4\x6c\x0d\x01\x65\xb1\xaf\x9b\x94\xc3\x49\x75\x83\xaa\x31\x23\x96\x05\xbf\x27\x65\x2b\xef\xa0\x51\xc2\xfb\x8b\x94\x0f\x4d\xab\x64\x56\xda\xea\xf3\x9e\x15\xa4\x07\x01\x16\x5d\x67\xdc\x28\x34\x1b\x84\x6e\xa1\xc5\x7a\xea\x21\x04\x48\xdd\x83\x9f\x24\x21\x70\x00\x2c\x4c\x5e\x16\x76\x57\xe1\x2b\x81\x50\xca\xac\x7c\xf4\x1c\x62\x69\x40\xe0\xee\x33\x32\x19\x3f\x8e\x0a\x9b\x0e\x06\x89\x4e\xb8\x66\x90\x0d\xa7\xae\xc1\x93\x9b\x1a\x9a\x1c\xb6\xc0\xda\x29\xb6\x92\x34\xda\xc3\x7b\x25\x6f\x11\x5a\xc9\x7e\x59\x4f\xc4\xee\x15\xd6\x21\xc4\x7f\xc8\xe1\x81\x1c\xb5\xa8\x15\x7a\x58\x9b\x09\xc8\xcc\x23\x06\xde\x8b\xd1\x2a\x9c\x83\x47\x2e\x80\x24\xa4\xe2\x6c\xa9\x23\x44\x28\x1e\x28\x50\x34\x03\x5c\x2e\x97\xdf\xc1\xe1\xaf\x09\x3d\xc5\x9e\xaf\x7f\x7c\x03\xfe\x72\x6b\x18\x45\x8b\x50\xaf\xd9\xca\x67\x90\x74\x2f\xa2\x1c\x68\x77\x82\x94\x05\xab\x1c\xd5\x77\x05\x7f\x0d\xa7\xd5\x02\x89\x82\x20\xd7\x24\x1c\xa7\xf1\x2c\x4e\x5f\x9b\xc5\xd2\x40\x8f\xc4\x6a\xc4\xc8\x39\x07\xc8\xe6\x16\x98\x35\x88\x59\xe2\x58\x2d\x43\xbf\xbe\x2c\xf8\x15\x48\xf0\xf1\x14\x5d\x64\x6c\x68\x87\x06\x9e\xa1\x15\x3d\xe6\x70\x33\xa0\xe6\x71\x3e\x4c\xb6\x31\x4a\x61\x13\xa8\x25\xc1\x4a\x2a\x05\x7e\x30\x2b\x98\x6c\xc8\x77\xf8\x1a\xfb\x95\x71\xf8\x2f\xe0\xa3\xe5\x61\x49\x2d\xea\xcc\x9b\xc9\xb1\xc5\x3a\xf3\x93\x4f\x84\xc1\xf8\x70\x6c\x29\x60\x70\xd8\x5d\xa4\x03\x91\xf5\x67\x45\xd1\x4b\x1a\xa6\x3a\xe7\xf5\x2c\x76\x19\xbd\x13\x76\x28\x0e\x57\x27\xad\xb8\xed\x85\xa4\xcb\xa9\x2e\x0b\x51\xcd\xe3\x64\xc2\x80\x1b\x3e\x67\xf2\xb8\x93\xe4\x8d\x45\x0b\xe9\x3d\x8f\x39\xad\xe2\x33\xaa\x73\x8b\x2e\xd4\x67\x09\xa1\xe3\x15\x09\x5d\xec\x9c\xbc\x43\x87\xa5\xd0\xb8\x82\x0e\x45\xd8\x4f\xff\x97\x52\x4f\x54\xd9\xdd\xcc\x28\x4c\x1d\xf4\x9b\x95\xef\xb2\x0c\xae\x23\xc6\x22\x60\x40\x2d\xda\x1e\x21\xcb\x82\x4f\x8e\x7c\x9d\x5c\xf3\x88\xfe\x84\x37\xf2\x0b\x2b\x0b\x87\xd6\xf8\x67\x9e\x17\x3b\xcb\x63\x71\x5f\xb4\xa6\xf1\xd9\x7e\xf9\x7d\x6e\x75\x9f\xf2\xe5\xa4\x8b\xf4\xc8\xce\x90\x7b\xbc\x85\xf9\x62\x9a\x23\xb0\x67\x92\x1a\x4e\xcd\x8d\xeb\xff\x8b\x23\xad\x9e\x77\xf6\xb6\x0a\x9f\x78\xfb\x69\xf2\xb9\xbf\x3b\xa0\x47\x46\x2e\xcc\x63\x7a\x10\xfe\xf0\x67\xfd\x13\x00\x00\xff\xff\x7f\x58\xc1\x38\x7c\x05\x00\x00")

func root_html_bytes() ([]byte, error) {
	return bindata_read(
		_root_html,
		"root.html",
	)
}

func root_html() (*asset, error) {
	bytes, err := root_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "root.html", size: 1404, mode: os.FileMode(436), modTime: time.Unix(1422234964, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _trace_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x57\x51\x6f\xdb\x36\x10\x7e\xf7\xaf\xb8\x29\x01\x4c\x77\xb6\x9c\xa6\xc5\x80\x39\x8e\x5f\x96\x76\xcb\x80\xae\xc3\x92\x75\x0f\x5d\x51\xd0\x12\x65\xb1\xa1\x44\x81\xa4\x9c\xb8\x86\xfe\xfb\xee\x48\x49\xb6\x53\x07\xcb\x80\xe5\x41\x96\x78\xfc\x8e\x1f\xef\xbe\x3b\x32\xdb\x6d\x2a\x32\x59\x0a\x88\x6e\xa5\x53\x22\x6a\x9a\xed\x56\x66\x10\xdf\x1a\x9e\x88\xf8\xfa\x2a\xfe\x9d\x1b\x51\xba\xa6\xb1\x15\x2f\x61\xbb\xdd\x19\x6e\x70\xa0\x69\x60\x82\x83\xa2\x4c\xf1\xcd\x91\xe5\x60\x8a\x7f\xf1\x73\x78\x55\x79\x73\x3b\x77\x30\xd8\xad\xfb\x8e\xcb\x12\x97\x1d\xcc\xf3\x97\x8b\xdb\xa7\x5c\xcc\xa7\x68\x1d\x3c\x41\x6d\x9e\x9f\x2f\x6e\xea\xe5\xe4\x09\x8a\x88\x3d\x5f\xf4\x0b\xcf\x53\xb9\x5e\xf8\x27\xc8\xf4\x32\x3a\xf1\xb4\x26\x47\x96\x8c\x20\x51\xdc\xda\xcb\x28\xcc\x70\xb2\x10\x0a\x09\x47\x8b\xf9\xf4\xd0\x45\xae\xd7\xc2\xfc\x21\x6c\xb4\x18\x00\xf8\xd1\x16\x98\x68\xa5\x8d\x48\xaf\xe4\xba\x07\xb5\x13\x08\x56\xf2\x42\x1c\x1b\xb7\x89\xd1\x4a\x89\xf4\x73\xca\xdd\xde\x6a\x07\x3f\x83\xb9\x75\x1b\x25\xc0\x6d\x2a\x81\x04\xc5\x83\x9b\x26\xd6\x13\x88\xf9\x83\xb4\x50\x71\x97\x8f\xbb\x0f\x62\x0d\x5b\x5c\x02\x32\xa9\xd4\x0c\x4a\x5d\x8a\x0b\xfa\xb4\xce\xe8\x3b\x31\x83\xa5\xe2\xc9\x5d\x18\xc9\x79\x25\x26\x18\xd5\x54\x18\x59\xae\x66\x90\x18\x69\xab\x37\xe9\x4a\x58\xb4\x37\x9d\x47\x5a\xb0\xf5\xa8\x4b\x37\xc9\x78\x21\xd5\x66\x06\x96\x97\x76\x62\x11\x99\x5d\xf4\x36\x2b\xbf\xe2\x0a\x2f\xcf\xaa\x87\xd6\x41\x17\xc7\x89\xe2\x4b\xa1\xfe\x9b\x97\xf3\xce\xcb\x49\xe7\xe5\xbc\x65\xe4\xdd\x60\xa2\x4a\x9b\x69\x53\xcc\xc2\xab\xc2\x00\x32\x5c\x79\xfc\x0a\x1f\x23\xef\x6d\x52\xd8\xc9\xbf\x4c\x83\xe9\x0b\xb8\x7e\x03\x3f\xc2\x8b\xa9\x47\xdc\x8b\xe5\x9d\x74\xcf\x41\xdd\xf0\x8c\x1b\x09\xbc\x4c\xe1\xa7\xdc\xe8\x42\x74\x2e\xf4\x73\xd0\xef\x2b\x61\x78\x87\x28\xf4\xd7\xe7\x60\xde\x4a\x23\x32\xfd\xe0\x51\x14\xdc\x9d\xe2\x42\x44\x72\x21\x57\xb9\x9b\x9d\x53\xf8\xe1\x5e\xa6\x2e\x6f\xdf\x33\xa5\xb9\x9b\x29\x91\x39\x1f\xd0\xf9\xd4\xeb\xc9\x0b\x0b\x53\x5e\xb9\x7d\x65\x7d\xe1\x6b\x1e\x46\x49\x60\x6b\x6e\x00\x95\xc9\xe1\x92\x4a\xed\x83\xb4\x57\xf8\xd1\x34\x17\xc1\xe2\xd7\x40\xd3\x29\x8b\x90\x4c\xe9\xb0\xb6\x85\x89\x46\xb1\x1f\x67\x98\x83\x01\x4c\xa7\x20\x0a\x40\x23\x56\x8d\x43\x2d\xe5\x02\x64\x59\xd5\x0e\x98\x2c\xc9\x52\x97\xd2\xd9\x11\x38\x0d\x95\x7c\x10\xca\x86\x01\x1f\x55\x23\x5c\x6d\x4a\x0b\xd2\xc5\x03\xc8\xea\x32\x71\x52\x13\x86\x89\xe2\xcf\x80\xf2\xbb\x26\x22\x24\x9b\x1b\x54\x0d\x72\xa9\xb8\xb1\xe2\x2d\x6d\x98\x9d\xb2\xe1\x52\xa7\x9b\xe1\x28\xc6\x72\x61\xc3\x5e\x5b\xc3\x51\x90\x47\x58\x60\x07\x7e\x01\xad\x67\x1f\x25\xcf\xbd\xac\x8b\xb7\x98\xdb\x37\x7b\x74\x68\x0b\x38\xbc\x14\xb8\x2c\x9a\xfc\x37\x96\x10\xe8\xcc\xbf\x56\xda\x61\x9b\x92\x5c\xa9\x0d\xac\xb8\x59\xf2\x95\xf0\x9e\xb0\xfc\xb0\xc6\xc6\x20\xe2\x55\x0c\x51\xa7\xe8\x6b\x27\x8a\xcf\x2f\x5f\xbf\x7e\x15\xc1\x64\x01\xf4\xb2\xb7\xd5\xdd\xe2\x0c\xd1\xed\x76\x5b\xd2\x7e\x9b\xd7\xa5\x23\x4b\x5c\x70\x97\xe4\x6c\xca\xfe\x4e\xbf\x1f\x9d\x4e\x47\x1f\xcf\x3e\x8d\xb1\x08\x47\xed\x36\x7a\x7f\xdd\xa2\xbf\x50\x0b\x63\x7b\xe1\x23\x03\xb5\x53\x6f\xc0\x18\x76\x08\x96\xe4\xdc\xb8\x31\x66\x2c\x15\x0f\xed\xfc\x80\xa0\xf6\x45\x79\x1f\x9e\x74\xfd\x70\x18\x62\x1a\xcc\x5e\x99\x16\x67\x78\x07\x41\xa8\x96\x75\x33\x10\x1c\xe3\x69\x90\xb2\xe1\x9e\x84\xbb\x2c\x2d\xb1\x3b\xad\x8c\xae\xcb\x74\xe2\x8d\xc3\x71\xeb\x8d\x05\x16\xdf\x3a\x39\xa1\xde\x8a\x70\x92\x2f\x23\xb1\x7e\xf4\x33\x3f\xc5\xbe\xe7\x84\xf9\x14\x87\x96\x1a\x31\x42\x66\xe9\xab\xbe\x39\xb1\x51\x70\xb9\xff\xd7\x6a\xd8\x3f\x8f\x99\xad\x43\x9e\x47\x81\x05\x37\x2b\x59\xb2\x2d\x95\xdb\x0c\xd5\xfa\xc3\x68\x0c\xc6\x17\xe6\xd9\x18\x75\x5e\xd1\xcf\x52\x3b\xa7\x8b\xd9\x59\x73\xcc\x81\x8f\x28\xeb\xb3\xc6\x52\x4c\xc0\x98\xaa\xb0\x2e\x30\x07\x87\xc9\xea\x33\x34\x82\xa3\xbe\x12\x25\x91\xe6\x53\xbe\xbe\x05\x00\xd6\x74\x99\xea\xfb\x58\xe9\x84\x13\x24\xce\xb1\xe7\x50\xb8\x08\x12\xd7\x46\x5d\x1c\x03\x4d\xa7\x5c\x61\x7d\xb3\x5f\x6f\xde\xff\x16\x07\xa5\xcb\x6c\xc3\x02\xc8\xf0\x7b\xea\x1a\x63\xd4\xb3\x52\x63\x38\xef\x53\xb8\xff\xd7\x84\x41\xca\x90\x5d\xaf\x42\x7e\xac\x50\x22\x71\xd8\x5b\x1e\x9d\xc9\xa3\x18\x2f\x18\x58\x72\x2c\xc2\xa9\xf4\xe5\x9c\x61\x91\x4f\x55\x34\x86\xe3\x29\x8b\x3d\x17\x2f\x0f\x14\x1a\x56\x67\x88\x9c\xef\x51\xc4\x1f\xde\xf1\x3b\x11\x8e\x3a\x6a\x32\x3c\xc9\xfb\x72\x01\xa4\x51\x60\x4d\x83\x0f\xe6\x84\x2f\x95\x88\x91\x5e\xcf\x27\xfe\x62\x21\xd5\xc2\x96\x43\xd7\xfa\xb2\x02\x3b\x1b\xb6\xb4\x9c\xaf\x05\x70\xb8\xe7\x1b\xfa\xb2\x75\x55\x69\xd4\x9e\xcb\xf1\x08\x13\xdc\xe2\xf9\x17\x07\x40\x0b\xfb\x4b\xa0\x9f\x60\x5e\x6e\x20\xec\x1e\x23\xe9\xdb\x8a\x67\xd6\x12\x19\x03\x09\xbf\xb3\x54\x46\xac\x7b\x8a\x0c\x47\x5a\x6f\x37\x1f\x7e\xc6\x66\x91\x38\x94\x1f\x75\x53\x34\x84\xae\xd1\xe1\xae\xaf\x80\xdd\xe7\x12\x37\xaa\xb4\xbe\xa3\x8b\x03\xdd\x0e\x1e\x75\x26\x8c\xee\x01\xc5\x5b\x04\xe2\x75\xc7\x75\x1d\xd0\x37\x3d\xee\x40\x76\xad\x1d\x0b\x0f\x9f\xb8\x5b\x5d\x87\x93\xc3\xe3\x4f\xbf\x49\x22\xac\x16\xb4\x25\xf4\x4f\xb1\xee\x15\xca\x7a\x55\x9e\x32\x0a\xc4\xe8\x51\x35\xb0\x3d\xd5\x92\x56\xc2\x82\x97\xfb\xad\xb2\x03\x52\x5c\x58\x2b\x8e\xa1\x4c\x87\x7b\xb2\x3b\x5e\x42\xbe\xbd\xb4\x73\x9a\x56\x18\x3b\x1e\x87\x95\xf4\xff\xf0\x78\xba\xd6\xfa\x1e\xb6\xab\xb8\xb6\x42\x9a\xae\xad\xcb\xcc\xab\x19\xbe\xbb\xf4\x85\xd5\x12\x7a\xd4\xe3\xbb\xf3\xde\x1f\xe8\x74\xe0\xd7\xea\xe0\xa2\x1b\xae\xb1\xdb\x2d\x26\xbb\xa2\x1b\x07\xfe\x63\x40\xc3\x11\x74\xf7\x62\x04\xd7\x6a\x41\x37\xf8\xf6\x46\xfd\x4f\x00\x00\x00\xff\xff\x65\xae\x2a\xbb\x43\x0c\x00\x00")

func trace_html_bytes() ([]byte, error) {
	return bindata_read(
		_trace_html,
		"trace.html",
	)
}

func trace_html() (*asset, error) {
	bytes, err := trace_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "trace.html", size: 3139, mode: os.FileMode(436), modTime: time.Unix(1422235116, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _trace_inc_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x52\xbd\x6e\xf3\x30\x0c\xdc\xf3\x14\x82\xa6\xef\x1b\x62\xbf\x80\x62\xa0\x40\x96\xa0\x45\x97\x16\xdd\x65\x8b\x49\x05\x30\x72\x60\xd1\x43\x21\xe8\xdd\x4b\x53\xf1\x4f\xda\xa1\x8b\x25\x9e\x48\xde\x1d\xe9\x94\x1c\x9c\x7d\x00\xa5\xdf\x07\xdb\x81\xce\x79\x67\xd0\xab\x0e\x6d\x8c\x07\x4d\x82\x29\xef\x0e\x3a\xde\x6c\xd8\xa7\x54\xbd\xf1\x59\x9d\x8e\x72\xe6\xac\x9b\x9d\x52\x29\x11\x5c\x6f\x68\x89\xbb\x4c\xf0\x8b\x6d\x01\xb5\xba\xa7\x48\x82\x3f\x97\xb0\x7a\x0a\xa1\x27\x4b\xbe\x0f\x51\x9e\x0c\xd9\x16\x61\xe1\x93\x40\xbe\xfb\xae\x0f\x0e\x42\x04\x77\x8f\x23\x0d\xfe\x06\x4e\x18\xa7\x96\x83\x0d\x17\x50\xff\xce\x1e\x09\x86\x4d\xdb\xdf\x44\xff\x85\x69\xe2\x1a\x1a\x43\x9f\x0d\xbb\x78\x86\xaf\x9c\x4d\xcd\x81\x21\xc7\x00\x37\x57\xd5\x87\xc5\x11\x04\x76\x0d\x7f\x86\x99\x09\x82\x2b\x5a\x6b\x51\x52\x2c\x17\x70\x36\xe7\x20\x76\x8c\xc8\x0c\xe3\xd6\xf3\xd8\x96\xd2\x11\x67\x8f\x71\x6c\x95\xcc\x35\xfe\xb0\xb2\x24\x3f\x4e\xb4\xec\x45\x55\xcb\xd3\x2a\x67\xc4\xad\x96\xf5\x66\x6a\xf4\xcd\x6e\xd1\xb8\xee\x78\xdd\x0e\xe3\x45\xe1\xab\xbd\xc2\x54\xc2\x13\xe8\xc3\x45\x91\x27\x84\x83\xe6\x11\x9d\x8e\xd3\x7a\xf9\x52\x32\x4c\x5d\x32\xa4\x2f\xc6\x3f\x6a\x96\x1f\xe4\xa1\x4c\xe4\xcc\xe7\x77\x00\x00\x00\xff\xff\xe0\x81\x61\xb5\x7a\x02\x00\x00")

func trace_inc_html_bytes() ([]byte, error) {
	return bindata_read(
		_trace_inc_html,
		"trace.inc.html",
	)
}

func trace_inc_html() (*asset, error) {
	bytes, err := trace_inc_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "trace.inc.html", size: 634, mode: os.FileMode(436), modTime: time.Unix(1422234960, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _traces_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x8f\xbd\x8e\xc3\x20\x10\x84\x7b\x9e\x62\x45\x7f\xa0\xeb\xb1\xab\x6b\xae\x48\x15\xbf\xc0\xca\x5e\xc7\x48\x84\x20\xc0\xd5\x8a\x77\x0f\x3f\xb2\x22\xa5\x1b\x98\x6f\x77\x66\x99\x37\xda\xad\x27\x90\x8b\xcd\x8e\x64\x29\x4b\xc4\x95\x12\xfc\x00\x86\x90\x9b\x66\x26\xbf\x95\x22\xc4\x87\xbd\xa1\xf5\x15\x15\xe6\xf8\x9d\x07\x6f\x74\x95\x42\x98\xd3\xcd\x02\x80\x39\xa2\x7f\x10\xa8\x61\x56\x12\xc0\x38\xdb\xac\x2a\x10\x8e\x48\xfb\x24\x99\xcf\xe8\x96\x57\x67\x40\xdd\x03\x7a\xf5\xff\x37\x46\x4a\x91\x33\xf3\xf7\x9f\xd1\x58\x33\xfa\x8e\xd3\xc1\xea\x30\xa5\x49\xf6\x8e\x49\x8e\xdd\x2d\x3a\xd3\x33\x38\xcc\xed\xa4\x66\x49\x50\x3d\xbf\x0e\xe9\xd1\xce\xe8\x51\xe5\x3a\x6c\xfc\x5f\xaf\x77\x00\x00\x00\xff\xff\xb4\x9a\x48\x82\x12\x01\x00\x00")

func traces_html_bytes() ([]byte, error) {
	return bindata_read(
		_traces_html,
		"traces.html",
	)
}

func traces_html() (*asset, error) {
	bytes, err := traces_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "traces.html", size: 274, mode: os.FileMode(436), modTime: time.Unix(1417495397, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	"layout.html": layout_html,
	"root.html": root_html,
	"trace.html": trace_html,
	"trace.inc.html": trace_inc_html,
	"traces.html": traces_html,
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
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"layout.html": &_bintree_t{layout_html, map[string]*_bintree_t{
	}},
	"root.html": &_bintree_t{root_html, map[string]*_bintree_t{
	}},
	"trace.html": &_bintree_t{trace_html, map[string]*_bintree_t{
	}},
	"trace.inc.html": &_bintree_t{trace_inc_html, map[string]*_bintree_t{
	}},
	"traces.html": &_bintree_t{traces_html, map[string]*_bintree_t{
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

