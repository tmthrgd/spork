// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package assets

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// FileSystem contains project assets.
var FileSystem = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2019, 11, 21, 3, 29, 52, 238418804, time.UTC),
		},
		"/control.css": &vfsgen۰CompressedFileInfo{
			name:             "control.css",
			modTime:          time.Date(2019, 11, 21, 3, 44, 37, 649839690, time.UTC),
			uncompressedSize: 587,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x91\xc1\x6e\xdb\x30\x0c\x86\xcf\xd6\x53\x08\xd8\xb5\x32\xe6\x22\x05\x3a\x19\x7b\x92\x61\x07\x45\x62\x6c\xae\x32\x69\xd0\x74\x16\x77\xe8\xbb\x0f\x92\x13\x2c\x6b\x4e\x86\xc5\xff\x27\x3f\xfe\x1c\x75\xca\x4f\xf6\xc8\x69\xb3\x7f\x4c\x33\x21\xb9\x11\x70\x18\xd5\xdb\xd7\x97\xf3\xd8\x9b\x0f\x63\x6e\xc5\x84\xcb\x9c\xc3\xe6\xed\x29\xc3\xa5\x37\x4d\xf9\xb8\x84\x02\x51\x91\xc9\xdb\xc8\x79\x9d\xa8\x37\xcd\xaf\x75\x51\x3c\x6d\x2e\x32\x29\x90\x7a\x1b\x81\x14\xa4\x37\xa6\x39\x31\xa9\x5b\xf0\x1d\xbc\xed\xbe\xcd\x97\xf2\xa4\x70\x51\x17\x32\x0e\x74\x2f\x8c\x9c\x59\xbc\xfd\x72\x38\x1c\x2a\xc4\xd8\x15\x84\x3b\xfb\x41\x60\xba\x56\x9e\xec\x5c\xe1\x83\x0c\x48\xee\xc8\xaa\x3c\x79\xdb\xdd\x04\x6d\xe1\x10\xce\xcb\xa7\x0e\x5d\xdb\x3d\xbf\x3c\x68\x42\x51\xed\xd3\xf7\xe1\x57\xc2\x04\x91\x25\xec\x9b\x12\x13\x7c\xb2\xb5\x21\x2a\x9e\xa1\x98\x8f\x2c\x09\xc4\x29\xcf\xde\x3e\xcf\x17\xbb\x70\xc6\x64\xe3\xc8\x91\x73\x50\xe8\x4d\x33\x87\x94\x90\x86\x5d\xd2\x95\x18\xfe\xeb\x85\x34\xaf\x7a\x3b\xc7\x6f\x4c\x3a\x7a\xfb\xda\x56\xd4\x87\x2d\xbf\x56\x6f\x75\xfc\xd0\x6d\x86\xef\x12\x68\x80\x9f\xde\xbb\x89\xdf\x5d\xfd\x71\x2a\x21\xbe\x55\xb2\x10\xdf\x06\xe1\x95\x92\xbb\xe6\x9b\xcb\xad\x07\x81\x6d\x47\x58\x98\x86\xbb\x28\x2b\xdf\x3e\xa1\x05\x11\x96\x7f\xd9\xd8\x09\x12\xae\xd3\x19\x39\x83\x0a\xa4\xde\x7c\xfc\x0d\x00\x00\xff\xff\x6a\xda\x7f\x05\x4b\x02\x00\x00"),
		},
		"/controls.js": &vfsgen۰CompressedFileInfo{
			name:             "controls.js",
			modTime:          time.Date(2019, 11, 21, 3, 30, 4, 516492963, time.UTC),
			uncompressedSize: 965,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x53\xc1\x6e\xdb\x3a\x10\x3c\x4b\x5f\xb1\x0f\x70\x42\x12\x4f\xa5\x83\x1e\x7a\xa8\xa0\x02\x45\x1a\xa0\x05\xd2\x4b\x9b\xdc\x4d\x50\x2b\x8b\x35\x4d\xaa\xe4\xca\x68\xa0\xf2\xdf\x0b\x4a\x8e\x9d\x04\x69\x2f\x36\xc1\xe5\xce\xce\xec\x8c\xcc\x7e\xf0\x81\x60\x82\x1e\x55\x8b\x21\x56\xd0\x21\xe9\xfe\xae\x47\x77\x3c\x5e\x2b\xd2\x7d\x05\x83\x55\x0f\xd6\x44\xba\x1f\x5a\x45\x18\x21\x41\x17\xfc\x1e\x98\x5c\xf7\x68\x07\x0c\x51\xfe\x88\xac\x2e\xcb\xd6\xeb\x71\x8f\x8e\xe4\xcf\x11\xc3\xc3\x77\xb4\xa8\xc9\x07\xce\xa4\xf6\x8e\x82\xb7\x91\x09\xa9\xda\xf6\xe6\x80\x8e\x6e\x4d\x24\x74\x18\x38\xd3\xd6\xe8\x1d\xab\x00\xa1\xf9\x00\x53\x59\x98\x0e\xf8\x7f\x1c\x25\xa9\xb0\x45\x02\xe3\x22\x29\xa7\xd1\x77\xf0\xf9\xee\xeb\xed\x47\xa7\x7b\x1f\x6e\x2c\xe6\x41\x42\xe4\x86\x22\x20\x8d\xc1\xd5\x65\x91\xca\xb2\xb0\x48\x40\x3d\x3a\x68\xce\x72\xea\x05\x95\x91\xdf\x6e\x2d\x32\x30\x0e\x1e\xf1\x65\xab\x48\x45\xa4\x05\xe9\xd8\x18\x30\x0e\x47\x36\x45\x71\x42\xe1\xf9\x5a\xd4\x65\xbe\x5c\x46\xce\x0f\xa5\xdf\xc1\xe5\xe5\x72\x24\xfc\x45\x5c\xc8\x0c\xc3\xfd\xee\x04\x51\x9c\xa6\x69\xab\x62\xcc\xda\xe5\xc2\x85\x33\xa5\xc9\x1c\x90\x55\xe0\x77\x92\x82\xd9\x73\x01\x4d\xd3\x00\xa3\x30\x22\x13\x75\x6e\x4f\xf3\x5f\x3a\x0a\x44\x39\x04\xcc\x2b\xfc\x84\x9d\x1a\x2d\xf1\x5c\x9d\x49\x9e\x96\x26\xfb\x80\x5d\x75\x76\x16\xd2\x91\x53\xfe\x11\x52\x67\x5b\xf9\xd9\x61\x51\x97\xa9\x82\x4e\xd9\x88\xe2\x5f\x36\x1e\xbc\x1d\xf7\xf8\xaa\x89\xc6\x0d\x23\x3d\x31\xf1\xc4\x84\x0c\x59\x84\x06\x36\xab\xe9\x74\x77\x50\x76\xc4\x74\xb1\xc9\xab\x5c\x88\x6f\xd6\x8f\x11\x59\x2f\x43\xd6\xab\x09\x9d\xf6\x2d\xde\x7f\xfb\x72\xed\xf7\x83\x77\xe8\x88\x3f\x47\x10\x69\xf3\x5c\x63\x59\x14\x8b\xcc\x57\x73\xfc\x42\xa5\xf6\x2e\x12\x44\xef\xb6\xd0\xc0\x5f\x25\xe7\x7a\x76\xa1\x7c\xf1\x0d\x70\x3e\xc1\x2c\xad\x02\x8b\x6e\x4b\x3d\x24\xf1\x24\xbf\x73\x69\x89\x54\x46\x98\x73\x71\xed\x1d\xa1\xa3\x65\x17\xf3\x83\x04\x7c\x35\xf1\x63\xff\x1a\xde\x5d\x09\xf8\x0d\x57\xe9\xfd\x6a\xe2\xec\x8a\xc1\xff\xf0\x58\xbb\xc8\x35\x21\xa3\x35\x1a\xf9\x9b\xb7\x22\x89\xcd\x1c\x86\x24\xea\x3f\x01\x00\x00\xff\xff\x1c\xaf\x48\xe1\xc5\x03\x00\x00"),
		},
		"/error.css": &vfsgen۰FileInfo{
			name:    "error.css",
			modTime: time.Date(2019, 11, 14, 2, 27, 13, 978845000, time.UTC),
			content: []byte("\x68\x74\x6d\x6c\x2c\x20\x62\x6f\x64\x79\x20\x7b\x0a\x09\x6d\x69\x6e\x2d\x68\x65\x69\x67\x68\x74\x3a\x20\x31\x30\x30\x76\x68\x3b\x0a\x7d\x0a\x0a\x62\x6f\x64\x79\x20\x7b\x0a\x09\x64\x69\x73\x70\x6c\x61\x79\x3a\x20\x66\x6c\x65\x78\x3b\x0a\x09\x66\x6c\x65\x78\x2d\x64\x69\x72\x65\x63\x74\x69\x6f\x6e\x3a\x20\x63\x6f\x6c\x75\x6d\x6e\x3b\x0a\x09\x6a\x75\x73\x74\x69\x66\x79\x2d\x63\x6f\x6e\x74\x65\x6e\x74\x3a\x20\x63\x65\x6e\x74\x65\x72\x3b\x0a\x7d"),
		},
		"/favicon.ico": &vfsgen۰CompressedFileInfo{
			name:             "favicon.ico",
			modTime:          time.Date(2019, 11, 21, 2, 5, 57, 400926370, time.UTC),
			uncompressedSize: 198,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\xcb\xb1\x0d\xc2\x30\x00\x05\xd1\x67\x17\xd4\xae\xa8\x4d\x47\xeb\xda\x05\x6c\xc1\x52\x0c\xc3\x1e\xac\x80\xc8\x28\x8e\x9c\x44\x72\x94\x36\x27\x9d\xae\xf8\xfa\x04\x41\x4a\x51\xe7\x83\x2b\xee\x48\xc8\xd6\xbd\xf3\x34\x88\x5b\x6f\xdf\xb6\x74\xaa\xfc\x2b\xbf\xa3\x85\x77\xe6\x92\x79\xed\x2d\xb4\xc7\xf0\xec\x7f\x0e\x00\x00\xff\xff\xbc\xfa\x7c\xaf\xc6\x00\x00\x00"),
		},
		"/helpers.js": &vfsgen۰CompressedFileInfo{
			name:             "helpers.js",
			modTime:          time.Date(2019, 11, 21, 3, 39, 5, 984807786, time.UTC),
			uncompressedSize: 894,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x92\x41\x6b\xdc\x3e\x10\xc5\xcf\xeb\x4f\x31\x87\x3f\x48\x86\xa0\xbd\xc7\xf8\x0f\x21\x2c\x94\x1e\xda\xc2\xba\xf7\x55\xe5\xe7\xb5\xa9\x2c\x39\xd2\xb8\xdd\x10\xfc\xdd\x8b\x65\xd7\xd9\xb4\xd9\x1e\x25\xcd\xbc\xf9\xcd\xd3\xc3\x65\xf0\x81\xc9\x78\x17\x99\x5a\xe8\x1a\x21\x52\x49\x2f\xd9\xee\xc1\x18\x0c\x7c\x4f\x82\x71\xe1\xfd\x60\x75\xe7\xc4\x5d\x36\x15\x59\xb6\x14\x23\x04\x1f\x0e\x16\x3d\x95\x54\x7b\x33\xf6\x70\xac\x9e\x46\x84\xe7\x23\x2c\x0c\xfb\x20\x85\x4a\x45\x22\xdf\x9a\x8c\x85\x0e\x87\xf9\x92\x4a\x92\x39\x95\xff\xcf\xa3\x36\x29\x35\xcf\x7a\xf4\x8e\xe1\x98\x4a\x12\xa2\x48\x03\x57\xc8\x66\x74\x86\x3b\xef\xa8\x01\x9b\xb6\x6a\xe1\x64\x40\x1c\xf2\x59\xa2\x6b\x28\x1d\x94\xff\x9e\xce\xbb\xd7\x49\x32\x2f\xb2\xdd\x44\xb0\x11\xcb\x8b\x77\xd1\x5b\x2c\x6c\x52\x24\x31\x0a\x78\x1a\x11\x99\x1a\xdd\x59\xd4\xf7\x24\xee\x28\x69\x17\x59\xb6\xbb\xc9\x77\x7a\xa0\x0f\x55\xf5\x65\xb1\x82\xbc\x31\x63\x08\x73\xf3\x7f\x2f\x09\x25\xb2\xe6\x31\x4e\x6f\x8f\x15\x2e\x3c\xa9\xd3\x8c\x74\x7b\xb7\x47\xcd\xa6\x95\x48\x9b\xfc\x13\x37\x5d\x26\x5a\x24\xd4\xdb\xa4\xee\x1d\x4a\x28\xf6\x47\x0e\x9d\x3b\xcb\x7c\x3a\xbd\xef\xf5\x60\xf5\xb3\xed\x22\x7f\x1d\x6a\xcd\x88\xd2\xbb\x1e\x31\xea\xf3\x2b\x1c\x13\xe6\xc8\x38\xfc\xa4\xc3\x0f\x38\x3e\xfa\x31\x18\x48\xb1\xff\xdd\xba\x1f\x97\x5e\xb1\x20\x46\xe5\x9d\x1f\xe0\xa8\xbc\x8a\xc3\xf6\xb2\xaa\x53\x49\x7d\x3c\xaf\xf9\xf8\xf3\x33\x77\x5b\x99\xfc\x78\xfc\xfc\x49\x0d\x3a\x44\xc8\x3e\x9e\x55\xad\x59\xe7\xe9\xbf\x37\x41\xac\x69\xc3\x26\xf6\xd6\xd0\x2b\xe6\xbf\xec\xbc\x9d\xcc\xaa\xc5\xf5\xb6\x6b\x6e\x0a\x1a\x9d\xfe\x66\x41\xec\x69\xde\x1c\x8e\x1a\x1f\x36\x0f\x69\x35\x42\x89\x85\x70\xfa\x15\x00\x00\xff\xff\xbf\x33\xfb\xf3\x7e\x03\x00\x00"),
		},
		"/normalize.min.css": &vfsgen۰CompressedFileInfo{
			name:             "normalize.min.css",
			modTime:          time.Date(2019, 11, 17, 3, 51, 44, 749372000, time.UTC),
			uncompressedSize: 1815,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x55\x4b\x6f\xe3\x36\x10\xfe\x2b\xea\x21\x97\x85\xa8\xc8\x29\xb2\x59\x50\xe8\xa5\xb7\x02\xed\x6d\x6f\x0b\x1f\x48\x71\x2c\x4d\x4d\x72\x08\x72\xe4\x47\xb4\xfa\xef\x85\x1e\x76\xac\x24\x46\xb1\x17\xc3\xe4\xbc\xbe\x99\xef\x1b\xea\xf1\xcb\x6f\x99\xa7\xe8\x94\xc5\x57\x28\xea\x94\xb2\xc3\xb7\xa2\x2c\x36\xd9\xcf\xec\x9f\xbf\xbe\x67\x7f\x63\x0d\x3e\x41\xf6\x33\x6b\x90\xdb\x4e\x17\x35\xb9\x47\x0f\x35\x59\x95\x1e\xd7\x71\x5f\x1e\x5b\x76\xb6\xb7\xe8\x41\xb4\x80\x4d\xcb\x72\x53\x6c\x9e\x2b\x71\x04\xbd\x47\x16\x0c\x27\x16\x09\x5f\x41\x28\xf3\x6f\x97\x58\x6e\xca\xf2\x61\xd0\x64\xce\xbd\x53\xb1\x41\x2f\xcb\xc1\x29\xf4\xbd\xc1\x14\xac\x3a\x4b\x6d\xa9\xde\x0f\xed\xa6\xdf\x91\x9f\x23\xe5\x13\xb8\x6a\x71\x2e\xbe\xbe\x80\xcb\xca\xa1\x8d\xbd\xa6\xd3\x68\x46\xdf\xc8\x9a\x3c\x83\x67\xa1\xe9\x54\x2d\x20\xca\x8a\x0e\x10\x77\x96\x8e\xf2\x80\x09\xb5\x85\x21\x44\x98\x93\xee\x94\x43\x7b\x96\x8e\x3c\xa5\xa0\x6a\xc8\xaf\xff\xaa\xb7\xa2\x1b\x70\x83\xea\xb5\xaa\xf7\x4d\xa4\xce\x1b\x51\x93\xa5\x28\x39\x2a\x9f\x82\x8a\xe0\x79\x50\x5a\xc7\x1f\x8c\x6c\x61\xdb\x6b\x8a\x06\xa2\xd0\xc4\x4c\x4e\x7a\xf2\x50\x4d\xad\x1b\xa8\x29\x2a\x46\xf2\xb2\xf3\x06\xe2\x38\xa8\xfb\x96\xcc\x10\x33\x98\x41\xe7\x89\x23\xf9\x66\xc6\x7b\x9c\x5b\xd2\x64\x0d\xc4\xa1\x26\x03\xf9\x5e\x9b\x3c\x29\x17\x7e\xb1\xa1\xe4\x94\xb5\x37\x93\xfd\x56\x3e\x0c\xa9\xd3\x79\xea\xc2\xcd\xed\xcb\xf3\x43\x75\x4b\x68\x59\x05\x4a\x38\x21\x8d\x60\x15\xe3\x01\xaa\x03\x44\xc6\x5a\x59\xa1\x2c\x36\x5e\x6a\x95\x60\x0c\x19\xb3\xf5\xcb\x10\x44\xf1\xf4\x3c\xd6\xec\x42\xcf\x14\xa4\x28\xc6\x13\xba\xe6\x32\xab\xc4\x67\x0b\xd3\xa8\x06\xdd\x31\x93\xcf\xd1\x87\x8e\x73\x0a\x3c\x8e\x3c\xe4\x09\x2c\xd4\x9c\x8f\xd3\x52\x11\xd4\xaa\x57\xf4\x2d\x44\xe4\xdb\xfe\xca\x72\x0d\x7b\xd2\xe1\x55\x64\xb7\x25\xfa\x0f\xda\x58\xac\x73\xc5\x7e\xe2\x67\x62\x7a\x47\x71\x66\x73\xf8\xc1\xe7\x00\x7f\xcc\x7e\xdb\x7c\x3e\x45\x48\xc0\x97\x43\xea\xb4\x43\xde\xe6\xb3\x4b\x7f\x59\x00\x15\x02\xa8\xa8\x7c\x0d\x72\xb6\xac\x33\x49\x29\x1c\xbd\x8a\x1d\xd5\x5d\x12\xe8\x3d\xc4\x55\xee\xbb\xe6\xa5\xda\x27\xf6\x39\xf1\x47\xc3\xc7\xb1\x57\x41\x19\x33\xee\x4f\xf9\x0e\xd3\x5b\x68\x44\xdf\xac\x01\x7d\x6a\xbb\xa0\x79\x67\x5c\xa0\xac\x6f\x7b\xea\x78\xa4\x49\x6e\xc2\x69\xd1\x7b\xf6\xe7\xe4\xf8\x1d\x4e\x3c\xec\x10\xac\x49\xc0\xfd\x05\x5c\xf1\xfb\x33\xb8\xac\x78\x99\x7e\xbf\x4e\x92\xb2\xd0\x80\x37\xb7\x2f\xc0\x75\xfd\x4e\xd5\xbc\xa7\x17\x7d\x5c\x9e\x15\x56\xda\x42\xe5\xd4\x49\x1c\xd1\x70\x3b\xab\xe5\xda\x7f\x75\x6c\x91\x41\x4c\x2b\x23\xe7\x07\x6e\x08\x91\x9a\x08\x29\xf5\xf7\xa4\x7e\xd5\xe5\x55\x4d\xaa\x63\x5a\x46\x59\xb7\x50\xef\x35\x9d\xae\x52\x51\x06\x69\x7b\x07\xf2\x7b\x1e\x7c\xe7\x34\xc4\x91\xdc\x45\x44\x13\x81\x22\x05\xf4\x62\x51\xea\x1d\x47\xea\x78\xed\xd8\x2f\xcb\x70\x03\x2d\x81\x8a\x75\xbb\xfd\x4c\xa1\x63\x4b\x13\x01\xd5\x42\x92\xa0\xdd\x2e\x01\x4b\xf1\x14\x4e\xeb\xf0\xb7\x9a\xf3\xc5\xcd\x7b\xf6\x59\xe6\x69\x87\xde\x62\x76\x68\x41\x74\xc1\x92\x32\xe2\xff\x16\x66\xda\xf1\x0b\xa1\x83\x01\x56\x68\xd3\xbb\xef\x45\xea\x9c\x53\xf1\x7c\xbd\xb5\x98\x58\x20\x83\x1b\x18\x5c\xb0\x8a\xe1\x6a\x9a\xb7\xb9\x45\x63\xc0\x6f\xd7\xb7\xff\x05\x00\x00\xff\xff\xde\x27\x49\xb4\x17\x07\x00\x00"),
		},
		"/playlist.css": &vfsgen۰CompressedFileInfo{
			name:             "playlist.css",
			modTime:          time.Date(2019, 11, 21, 3, 45, 47, 761269204, time.UTC),
			uncompressedSize: 272,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xce\xc1\x6a\xc3\x30\x0c\xc6\xf1\x73\xf5\x14\x82\x9d\x53\xd6\x61\x76\x70\x9f\xc6\xad\xb5\x54\x20\x5b\x41\x56\x33\xb2\x91\x77\x1f\x4e\x58\xb6\xab\xfd\xfb\xa3\xef\xa6\x79\xc1\x6f\x38\x4d\x29\x67\xae\x63\xc4\x60\x54\xf0\xcd\xa8\x5c\x01\x4e\x1f\x5a\x7d\x68\xfc\x45\x11\x2f\xe7\xf7\xfd\xed\xae\xa2\x16\xf1\x25\x84\x70\x85\x15\xe0\x71\xe9\xfd\x3f\x19\xb6\x78\x05\x50\xe9\x3f\xc2\xcd\x87\xe6\x8b\xd0\x30\x69\x63\x67\xad\x11\xf5\xe9\x8d\x33\x6d\x2c\x93\x27\x96\x86\x53\xd7\x25\xd9\xc8\x35\xe2\xeb\x71\xfe\x93\x78\x7c\x78\xc4\xaa\x56\x92\x6c\x85\x30\xa6\x8e\xf7\x29\xc7\x92\x73\xba\x3b\xcf\x74\xcc\xf9\x2d\x6f\x2a\x79\x07\x64\xa6\xf6\x57\x62\xa1\xcc\xcf\x32\xb3\x0a\xb9\x51\x47\x3f\x01\x00\x00\xff\xff\x59\x84\x86\x5d\x10\x01\x00\x00"),
		},
		"/playlist.js": &vfsgen۰CompressedFileInfo{
			name:             "playlist.js",
			modTime:          time.Date(2019, 11, 21, 3, 30, 25, 371618928, time.UTC),
			uncompressedSize: 847,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\xbf\x6e\xdc\x3c\x10\xc4\x6b\xe9\x29\xd6\xf8\x0a\x52\x80\xc0\xaf\x8f\xa0\x00\xf9\x63\x20\x01\x9c\x34\xbe\xa4\x09\x52\xd0\xd4\xea\xc8\x98\xe2\x2a\xe4\xea\x82\xc3\x59\xef\x1e\x50\xd2\xdd\xd9\x80\x8b\x74\xe4\x88\x3b\xb3\xf3\x93\x1b\x46\x8a\x0c\x27\xb0\xa8\x3b\x8c\xa9\x86\x1e\xd9\xd8\x9d\xc5\xb0\x1d\x3f\x68\x36\xb6\x86\xd1\xeb\xa3\x77\x89\xbf\x8d\x9d\x66\x4c\x30\x43\x1f\x69\x00\xa1\xfe\xb7\xe8\x47\x8c\x49\xfd\x4a\xa2\x29\x4b\x43\x21\xf1\xe5\x35\xb4\xd0\x91\x99\x06\x0c\xac\x7e\x4f\x18\x8f\xf7\xe8\xd1\x30\x45\x29\xd4\xf9\x8d\xa8\x9a\xb2\x3c\x5f\x94\xee\xba\xdb\x03\x06\xbe\x73\x89\x31\x60\x94\xc2\x78\x67\x1e\x45\x0d\x08\xed\x5b\x38\x95\x85\xeb\x41\xde\x48\x54\xac\xe3\x1e\x19\x5c\x48\xac\x83\x41\xea\xe1\xd3\xee\xcb\xdd\xbb\x60\x2c\xc5\x5b\x8f\x39\xb3\xaa\xf2\x40\x11\x91\xa7\x18\x9a\xb2\x98\xcb\xb2\x40\x35\x46\xcc\x09\x1f\xb1\xd7\x93\x67\x59\x35\x65\xb1\x34\xbd\x78\x2a\x1b\xb1\xaf\xaf\x50\x60\xae\x14\x5b\x0c\xf2\x55\x36\x55\x53\xce\x35\xf4\xda\x27\xac\x2e\x00\x1c\xe3\x90\xa0\xbd\x80\x50\x7b\xe4\x6d\xa9\xf4\xfe\xb8\xd3\xfb\xaf\x7a\x40\x29\xbc\x7b\xd1\x7e\x83\x2b\xe5\x09\x46\xca\xb1\x5b\xe5\xd5\x93\x7c\xf7\x2f\x3c\x41\x69\xc3\xee\x80\xd9\x79\x81\x45\xbe\x5b\x39\x90\xef\x94\xcb\x1e\x42\x34\xdb\xd5\x78\x9d\x52\x66\xad\x22\x0e\x74\x40\x29\x9e\x0d\x67\x5c\xd7\x3a\xd0\xae\xad\x7e\x8c\x94\x7e\x6e\xd6\x37\x59\x79\x05\x72\x96\xb7\x28\x33\xc5\x88\x81\x73\xe2\xa2\x5e\x13\x75\xd7\x3d\x8f\x5b\x1d\x3d\x19\xcd\x8e\x82\xb2\x3a\x59\x68\xdb\x16\xc4\x7f\x67\x8b\x35\x68\xdd\x28\x99\x48\xde\xef\x68\xdb\x6a\xf9\xa9\x8e\xa6\xb4\x41\xbe\x77\x0f\xde\x85\x3d\x3c\x3d\x2d\x9f\x73\xdd\xf3\x84\x5a\x0f\x9f\x03\xd3\x77\x87\x7f\x64\xf6\x2c\x1e\x3c\x99\xc7\x37\x20\x12\xeb\xc8\xa2\xce\x92\x0b\xde\x05\x7c\xa9\x3d\xa0\xd5\x07\x47\x31\xab\x03\x11\xdb\x45\x9e\x57\x58\x73\xd5\xfc\x0d\x00\x00\xff\xff\x25\x7c\xa8\xcf\x4f\x03\x00\x00"),
		},
		"/skeleton.min.css": &vfsgen۰CompressedFileInfo{
			name:             "skeleton.min.css",
			modTime:          time.Date(2019, 11, 17, 3, 44, 33, 525117000, time.UTC),
			uncompressedSize: 5879,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x58\x4f\x8f\xa3\x3a\x12\xff\x2a\x51\xb7\x5a\xda\x95\x0c\xe2\x4f\xa0\xd3\xa0\x1d\xed\xec\x4c\xb7\xf6\xb4\x87\xbd\x3e\xcd\xc1\x40\xd1\x58\x6d\x30\xcf\x98\x0e\x99\x88\xef\xfe\x64\x0c\x04\x1b\xc8\xcc\xed\x25\x87\xc4\x55\xae\x4a\xd5\xaf\x7e\x2e\x2a\xb6\x53\x56\x09\x4c\x2a\xe0\xd7\x9a\x35\x44\x10\x56\x45\x1c\x28\x16\xe4\x13\xe2\x33\xc9\x44\x11\xb9\x8e\xf3\x14\x97\xb8\xb3\xd4\xf2\x25\x74\xea\x2e\x2e\x31\x7f\x27\x55\xe4\x1c\x70\x2b\x58\x5c\xe3\x2c\x23\xd5\x7b\xe4\x1c\x3c\xa9\x4c\x58\x67\x35\xe4\xa7\x94\x24\x8c\x67\xc0\xad\x84\x75\xbd\x9d\x32\xda\x96\x15\x1a\x3f\x9b\xeb\xc2\x7d\x4e\x19\x16\x11\x85\x5c\xec\x18\xff\xbb\x84\x8c\xe0\xc3\x3f\x4a\x52\x8d\x71\x1c\x1d\xa7\xee\xfe\x79\x5d\x24\xa0\xe4\xa7\xe0\xe9\x16\x4f\xbf\x61\x18\x04\x3b\x86\xce\xd3\x3a\x46\x95\xa7\x25\x23\x8b\x8e\xb3\x3e\xca\x09\x6f\x84\x95\x16\x84\x66\xf3\xde\xa5\x50\xb3\x73\x7a\x9b\x55\x30\xbb\xbe\x7d\x9f\x20\x38\xda\xe1\xfc\x7a\x7e\xea\x6d\x71\x66\x26\x48\xbe\xed\xcf\x2f\xb9\xa3\xe0\x60\x7a\xf1\xbc\xa7\xde\xce\x59\xcb\x0d\xb9\xef\xd8\x9a\xf7\x9c\x7c\x9a\xa6\xfe\x8b\xee\xbe\x21\x9d\x19\xe2\x49\x8a\xe1\x13\x2a\x43\x11\x84\xba\x77\x20\xef\x85\x30\xf6\x84\x81\xee\xbe\x22\x2b\x08\x9e\x25\xb8\x62\xe5\xfd\xe4\x19\xde\xe9\x46\x08\x2f\xae\x01\xce\x19\xe8\x2a\xc5\x91\xc5\x66\x55\x2c\x51\x10\x9e\x8d\x7b\x77\x00\x13\x67\xa6\xb6\x35\xfa\x3e\x33\x2d\xe9\xae\xc0\x34\xd7\x77\x0d\xc8\xb1\x3c\x6f\x40\x58\xc9\xc5\xd2\x98\xb0\x21\xd5\x29\x77\x32\x98\x71\xb3\xb8\x71\x04\x6d\x4a\x75\x3f\xee\xb3\x11\xea\xcd\x62\xc1\x24\xb4\x23\xd7\x7d\x79\xa1\xe6\x60\xc1\x38\xb4\x2d\xd6\xcd\xfd\xa3\xbd\x93\xd2\x82\x98\x68\x5b\x6c\x9c\x47\x7f\x2f\xa9\x1b\x7f\xd1\xa6\x54\xf7\x13\x78\xba\xf1\x82\x62\x68\x47\xae\x3b\x08\x9d\xbd\x94\x96\xa7\x01\xed\xc8\x0d\x5f\x2f\x7b\x49\x2d\x4e\x0d\xda\x16\xeb\x9e\x9e\x75\xe2\x89\xcd\xa4\xc4\x4e\x4a\xa7\x70\x37\x25\xba\x83\x8f\x71\x36\x97\xde\x5e\xcc\x93\xb2\x24\xbe\x76\x04\xd1\x1d\xdd\xef\xf2\x68\x75\x5e\xd1\x3d\xe5\xef\xc2\x6f\x9c\x6e\x33\xd0\x85\x6a\xcd\xae\xbe\x10\x25\xbd\xe6\xac\x12\xf2\xd1\x06\x51\xe8\xd9\xc1\x53\x9f\xb0\xec\xb2\x10\xba\x76\x00\x65\x4c\x89\x74\x36\xf0\x23\x72\xed\x30\x1e\xf4\x67\xb5\x3e\x3a\x8e\x5a\xe7\xb8\x24\xf4\x12\xfd\x1f\x53\x38\xe3\x0b\xfa\xaf\xec\x77\x82\xa4\xf8\x7f\xd0\x02\x7a\x98\x97\x07\xb9\x7e\xb8\xa9\xd1\x57\x4e\x30\x45\x0d\xae\x1a\xab\x01\x4e\xf2\x38\x65\x94\xf1\xe8\xd1\xf3\xbc\xbe\x70\x51\xe1\xa1\xc2\x47\xc5\x11\x15\x01\x2a\xc2\x29\x0b\xc1\xea\xc8\x99\x5a\x67\xc2\x84\x60\x65\xe4\x71\x28\xb5\xd0\x7c\xc7\xe9\x0b\x77\x91\xcd\x91\xaf\x92\xf1\x62\x0a\x42\x00\xb7\x9a\x1a\xa7\xf2\xf1\x6c\xd9\x2e\x87\xb2\x2f\xbc\x85\x9d\x6f\x87\x1b\x96\xc1\x9e\xa9\xbf\x34\x5d\x1b\xfa\x7b\x76\xc7\x85\x9d\x67\x6f\x04\xeb\x6f\xfc\xa4\x73\x1a\x6c\x03\xad\x68\xa7\xb5\xed\x96\x69\x30\x98\x86\x7a\xbd\xd7\xa6\xa1\x69\xea\xec\x4f\x30\x1a\xde\xc1\x0a\xc9\xa3\xed\xad\x21\x1a\xd0\xd5\xd3\xf7\x57\x49\x29\x40\xb6\xa2\xed\xfb\x5a\xa3\x45\x8f\xaf\x23\x85\xdc\xd7\xaf\xaf\xdf\xff\xd3\xe3\xa8\x60\x9f\xc0\x27\xa9\xf3\xf6\xd5\xf9\xf6\xda\xdb\x49\x2b\x04\xab\xd0\xf8\x41\xaa\xba\x15\x7f\x88\x4b\x0d\xff\x52\x92\x1f\x4b\x11\x87\x06\x84\x26\x69\xda\xa4\x24\xe2\xc7\x35\x23\x4d\x4d\xf1\x25\x22\xd5\x00\x5a\x42\x59\xfa\x11\x8f\xd0\xf9\xa7\xba\x5b\xcc\xa1\xbe\x9c\x43\xc7\x20\x82\x20\x88\x05\x74\xc2\xc2\x94\xbc\x57\x51\x0a\x95\x00\x1e\x2f\x52\x73\xeb\x4e\x23\x73\xe8\x38\x5a\x59\x06\xdf\x46\x5d\x06\x1e\x29\xb7\x82\xe3\xaa\xc9\x19\x2f\xa3\xb6\xae\x81\xa7\xb8\x01\xa5\xc8\x20\x65\x1c\x0f\x33\x75\xc5\x2a\x88\xcf\x05\x11\x30\x78\x80\xa8\x62\x67\x8e\xeb\x38\xc1\xe9\xc7\x3b\x67\x6d\x95\x59\x2a\xda\xc1\x59\x8d\x39\x54\x72\x14\x1e\xe6\x5f\x8e\x33\xd2\x36\xd1\x71\x98\xac\xa5\x24\x72\xeb\xee\xd0\x30\x4a\xb2\xc3\x63\x92\x24\x71\xda\xf2\x86\xf1\xa8\x66\x64\xc8\x6c\x67\xfc\x56\x50\x47\x39\x4b\xdb\x06\x4d\xab\xa1\x5a\x48\x53\x69\x9a\x75\xa9\xc6\x5d\x1b\x8a\x95\x85\xaa\xe4\xda\x60\x94\xaf\xf6\x8f\x75\x5e\x1b\x4c\x0a\x8d\x5b\xbe\xef\x4f\x08\x8d\x92\xd3\xe9\x14\xb3\x56\xc8\xd2\xc9\xf1\x4e\xc5\x35\x7e\x58\x35\x27\x25\xe6\x17\xb4\x2d\x5d\xa7\x73\x67\x87\x8a\xff\xce\x86\x31\x5e\x63\xc7\x14\xf8\xdb\xdb\xdb\xba\xee\x8f\xbe\xff\xcd\x7f\x73\x8c\x8c\x94\x70\x27\x15\xbd\x94\xa6\x72\x59\xd9\x6d\xc3\x7b\x76\xbf\xc4\x63\x97\x06\xbf\xf4\xb7\x89\xde\x1e\x49\x7e\xe9\x6d\x1b\xea\x5d\x0a\x6d\xfa\xbb\x5f\x18\xd5\xd9\x8c\xc2\x8c\xed\x6e\xf1\x03\x50\x62\x42\xb5\xb6\x55\xb5\x65\x02\x5c\x13\xd5\xb8\x69\xce\x8c\x67\x7a\x7b\x03\xcc\xd3\x42\x13\x09\xa0\xc6\xba\xd3\x5b\x62\xcb\xe9\x0f\xd4\x00\x85\x54\x20\xa9\xc4\x1c\xf0\x75\xab\x15\x86\x75\x77\x70\x87\x3f\xe5\xab\xc4\xf2\x3c\xdf\x68\x28\xdf\x5d\xf9\xde\xec\x3d\x9d\xd5\x14\x38\x63\x67\xd5\xcd\xb6\xdb\xcc\xdf\x8b\xc8\x0c\x85\x75\x86\xe4\x83\x08\x0b\xd7\x35\x60\x8e\xab\xa1\xe3\x56\x10\x5b\x25\xfb\xb9\x12\x1a\xeb\x7e\xf6\x22\x1f\xb9\x23\xa8\x61\x70\x03\x75\x78\xf6\x85\x8b\xf5\x38\x15\x85\xf5\x46\xfe\x6b\x2a\x8e\x28\xac\x15\x33\x16\x1b\xf4\x55\x88\xac\x15\x12\x97\x2d\x69\xb7\xd5\x47\x25\x46\xa3\x54\x71\x67\x5c\x4c\x09\xab\xe5\x75\x4d\x8a\xb1\x39\xdd\x9a\x2b\xc5\x09\x50\x44\xe1\x1d\xaa\x6c\x7e\x2e\xab\x07\xb2\x3e\x29\xaa\x21\xc7\x78\xba\xf6\x39\x01\x9a\x35\x20\xae\xf3\x13\x7b\x22\x9c\x9a\x70\x9c\x25\x8e\x69\x01\xe9\x47\xc2\x3a\x7d\x4a\xc0\x19\x61\xe6\x48\xa0\xc2\xfa\x62\x0f\x1f\xd6\x30\x5f\x6f\xce\x0c\xcb\xf9\x7c\x1d\xe0\xd1\x71\xfa\x96\x5e\x29\x69\x84\xd5\x88\x0b\x85\x28\x25\x3c\xa5\x70\x20\x55\x43\x32\xe8\x99\xa6\xcb\x20\x25\x25\xa6\x37\x25\x6a\xe9\x94\xd6\x78\xd3\x10\x6b\x23\x13\xa3\x07\x46\x11\xa3\x87\x56\x6e\x3d\x0c\x06\x87\x96\x8e\x73\xd5\x38\x68\x1d\x9c\xc3\xf8\xc5\x9f\xc3\x1b\x86\x95\x17\xe7\xa9\xa7\xe4\xaa\x83\x3c\xcc\xb4\x29\xcb\x60\xc6\x73\x98\xfd\x0e\x2a\xb7\xf9\xae\xce\xf6\x56\xbe\xee\xcf\x24\xd1\xe3\x9b\x2b\xdf\x1b\x6d\xe2\xd5\x95\xef\x75\x9b\xe8\x6b\x0e\x5f\x86\x50\x74\x56\x4c\x81\xc9\x50\xc7\xd4\xb4\xdf\xae\x39\xf4\x22\x43\xa2\x98\x53\x70\x3d\xd9\xba\xe4\xb1\x5b\x0c\x6f\xe3\x15\xe1\xd8\x6e\x54\xf2\x66\x50\xbd\xc8\xb4\x0b\x3a\x51\x68\x57\x73\x7a\x6d\xe4\x66\x8a\x97\x7b\x6f\xab\x79\x2b\x1f\x78\xe1\x18\x73\xec\x56\x11\x26\x62\x2b\xaa\xae\x3a\xb4\x61\xa1\x66\xea\x01\xa0\x3f\x5b\x26\x00\x65\x14\xe5\xe4\xbd\xe5\x80\xe4\x3c\x89\x18\x45\x35\xaa\x39\x20\x81\x13\x0a\x68\x26\xc9\xfc\x37\x4c\x39\xb0\x5b\x2b\x6f\x29\x55\x67\x67\x79\xdf\xb5\x33\x09\xb6\x56\x89\xbb\xa5\xc9\xed\x76\xf7\xbe\x59\x2d\x4d\x06\x30\xae\xea\xc6\x76\xf8\x3e\x6b\x24\xa2\xd7\xdb\x55\x6e\x5f\xf0\xe5\x9f\x05\xff\xc6\xc5\x29\x7e\x5f\xf1\x40\x3f\xfa\xd3\x52\xda\xac\x6b\x7b\xbb\xb6\x8d\x70\x2e\x80\x23\x9b\xb3\xf3\xf4\xb5\xb5\xd2\xfc\x2a\x37\x40\x25\xa2\x87\x87\x78\xa2\xe0\x00\x5f\x9c\x52\xc0\x3c\x4a\x98\x28\xfa\xbf\x02\x00\x00\xff\xff\xe1\xa3\xd0\x37\xf7\x16\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/control.css"].(os.FileInfo),
		fs["/controls.js"].(os.FileInfo),
		fs["/error.css"].(os.FileInfo),
		fs["/favicon.ico"].(os.FileInfo),
		fs["/helpers.js"].(os.FileInfo),
		fs["/normalize.min.css"].(os.FileInfo),
		fs["/playlist.css"].(os.FileInfo),
		fs["/playlist.js"].(os.FileInfo),
		fs["/skeleton.min.css"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰FileInfo:
		return &vfsgen۰File{
			vfsgen۰FileInfo: f,
			Reader:          bytes.NewReader(f.content),
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰FileInfo is a static definition of an uncompressed file (because it's not worth gzip compressing).
type vfsgen۰FileInfo struct {
	name    string
	modTime time.Time
	content []byte
}

func (f *vfsgen۰FileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰FileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰FileInfo) NotWorthGzipCompressing() {}

func (f *vfsgen۰FileInfo) Name() string       { return f.name }
func (f *vfsgen۰FileInfo) Size() int64        { return int64(len(f.content)) }
func (f *vfsgen۰FileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰FileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰FileInfo) IsDir() bool        { return false }
func (f *vfsgen۰FileInfo) Sys() interface{}   { return nil }

// vfsgen۰File is an opened file instance.
type vfsgen۰File struct {
	*vfsgen۰FileInfo
	*bytes.Reader
}

func (f *vfsgen۰File) Close() error {
	return nil
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
