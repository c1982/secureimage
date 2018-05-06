package secureimage

import (
	"bytes"
	"errors"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

const (
	attentionPlease = 512
)

var (
	errUnknowExtension = errors.New("unknown file extension")
	errImageFileFormat = errors.New("invalid image format")
	signs              map[string][][]byte
	extensions         []string
)

func init() {

	signs = map[string][][]byte{
		".gif":  [][]byte{[]byte("GIF87a"), []byte("GIF89a")},
		".png":  [][]byte{[]byte("\x89\x50\x4E\x47\x0D\x0A\x1A\x0A")},
		".jpg":  [][]byte{[]byte("\xFF\xD8\xFF")},
		".jpeg": [][]byte{[]byte("\xFF\xD8\xFF")},
	}

	extensions = []string{".jpg", ".jpeg", ".gif", ".png"}
}

//Check you can check trusted image file.
func Check(imagefile string) (trusted bool, err error) {

	trusted = false
	f, err := os.Open(imagefile)

	if err != nil {
		return
	}

	defer f.Close()
	data, err := ioutil.ReadAll(f)

	ext := path.Ext(imagefile)
	ext = strings.ToLower(ext)

	if !knowedExtensions(ext) {
		return trusted, errUnknowExtension
	}

	if !matchMime(data, ext) {
		return trusted, errImageFileFormat
	}

	return checkIt(f, ext)
}

func checkIt(f *os.File, ext string) (isok bool, err error) {

	f.Seek(0, 0)

	switch ext {
	case ".png":
		_, err = png.Decode(f)
	case ".jpg", ".jpeg":
		_, err = jpeg.Decode(f)
	case ".gif":
		_, err = gif.Decode(f)
	default:
		err = errUnknowExtension
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func matchMime(data []byte, ext string) bool {

	if len(data) > attentionPlease {
		data = data[:attentionPlease]
	}

	prefixes := signs[ext]

	i := firstByteNonWhiteSpace(data)
	data = data[i:]

	for i := 0; i < len(prefixes); i++ {
		if bytes.HasPrefix(data, prefixes[i]) {
			return true
		}
	}

	return false
}

func firstByteNonWhiteSpace(data []byte) (nonws int) {

	nonws = 0
	for ; nonws < len(data) && isWhiteSpace(data[nonws]); nonws++ {
	}

	return
}

func isWhiteSpace(b byte) bool {

	switch b {
	case '\t', '\n', '\r', '\v', '\f', ' ':
		return true
	}

	return false
}

func knowedExtensions(ext string) bool {

	for i := 0; i < len(extensions); i++ {
		if ext == extensions[i] {
			return true
		}
	}

	return false
}
