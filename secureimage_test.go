package secureimage

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestCheckWhiteSpace(t *testing.T) {

	var list = []struct {
		b byte
	}{
		{'\t'},
		{'\n'},
		{'\x0c'},
		{'\r'},
		{' '},
		{'\v'},
		{'\f'},
	}

	for _, l := range list {
		if !isWhiteSpace(l.b) {
			t.Error("it is not expected:", l.b)
			break
		}
	}
}

func TestFirstByteNonSpace(t *testing.T) {

	str1 := `
	
	GIF89a;
	zksdla`

	str2 := `firstbyte=0`

	str3 := "   	  0"

	data1 := []byte(str1)
	data2 := []byte(str2)
	data3 := []byte(str3)

	var list = []struct {
		data []byte
		seek int
	}{
		{data1, 4},
		{data2, 0},
		{data3, 6},
	}

	for _, l := range list {

		if i := firstByteNonWhiteSpace(l.data); i != l.seek {
			t.Error("invalid seek:", i)
		}
	}
}

func TestMatchMimePrefix(t *testing.T) {

	jpgData, _ := ioutil.ReadFile("./testdata/test.jpg")
	jpg2Data, _ := ioutil.ReadFile("./testdata/not_cleaned.jpg")
	pngData, _ := ioutil.ReadFile("./testdata/test.png")

	var list = []struct {
		ext string
		b   []byte
	}{
		{".gif", []byte("GIF89a;<script>alert(1)</script>")},
		{".gif", []byte("GIF89aõ^@ ^@÷ÿ^@2kÝ<91>0$°Ýîùl-þHÚ£e<9e>ùÙØú×îè0<9a>")},
		{".jpg", jpgData},
		{".jpg", jpg2Data},
		{".png", pngData},
	}

	for _, l := range list {
		if !matchMime(l.b, l.ext) {
			t.Error("invalid match prefix:", fmt.Sprintf("%c", l.b))
		}
	}
}

func TestKnowExtension(t *testing.T) {

	if !knowedExtensions(".jpg") {
		t.Error("extension not found")
	}
}

func TestCheck(t *testing.T) {

	_, err := Check("./testdata/test.jpg")

	if err != nil {
		t.Error(err)
	}
}
