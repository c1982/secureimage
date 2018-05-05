# secureimage
TR: Go web uygulamalarında upload edilen resim dosyalarının güvenilir olup, olmadığını kontrol eden küçük bir doğrulama paketidir.
Bu paket sadece gif, jpeg ve png dosya formatlarını doğrular.

EN: This is a small verification package that checks whether image files uploaded in Go web applications are reliable.
This package only supports gif, jpeg and png file formats.

## Install

```bash
go get github.com/c1982/secureimage
```

## Usage

```go
package main

import (
	"fmt"
	"os"
	"github.com/c1982/secureimage"
)

func main() {
	trusted, err := secureimage.Check("./uploads/tmp_test.jpg")

	if err != nil {
		panic(err)
	}

	if trusted {
		fmt.Println("file is trusted.")
	} else {
		fmt.Println("bad file")
	}
}
```
## Todos

- [x] Magic Bytes check
- [x] Validate image file format
- [ ] Clean exif data in jpeg format

## Credits

 * [Oğuzhan](https://github.com/c1982)

## License

The MIT License (MIT) - see LICENSE.md for more details