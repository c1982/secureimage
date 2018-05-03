package secureimage

import (
	"bytes"
	"errors"
	"os"
	"path"
	"strings"
)

const (
	attentionPlease = 256
)

var (
	errUnknowExtension = errors.New("unknow file extension")
	signatures         = map[string][]byte{}
)

func init() {
	signatures[".gif1"] = []byte("GIF87a")
	signatures[".gif2"] = []byte("GIF89a")
	signatures[".png"] = []byte("\x89\x50\x4E\x47\x0D\x0A\x1A\x0A")
	signatures[".jpeg"] = []byte("\xFF\xD8\xFF")
}

//Check you can check trusted image file.
func Check(imagefile string) (isok bool, err error) {

	f, err := os.OpenFile(imagefile, os.O_RDWR, 0755)

	if err != nil {
		isok = false
		return
	}

	ext := path.Ext(imagefile)
	ext = strings.ToLower(ext)

	switch ext {
	case ".jpg":
	case ".jpeg":
	case ".gif":
	case ".png":
		return checkIt(f)
	default:
		return false, errUnknowExtension
	}

	return
}

func checkIt(f *os.File) (isok bool, err error) {

	return
}

func matchMime(data []byte, ext string) bool {

	prefix := signatures[ext]
	return bytes.HasPrefix(data, prefix)
}

func firstByteNonWhiteSpace(data []byte) (nonws int) {
	nonws = 0
	for ; nonws < len(data) && isWhiteSpace(data[nonws]); nonws++ {
	}

	return
}

func isWhiteSpace(b byte) bool {
	switch b {
	case '\t', '\n', '\x0c', '\r', ' ':
		return true
	}
	return false
}
