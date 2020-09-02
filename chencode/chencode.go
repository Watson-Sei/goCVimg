package chencode

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func EncodeJPG(directory,filename string)  {
	fmt.Println(directory, filename)
	// jpgファイル開きます。
	of, err := os.Open(filepath.Join(directory, filename))
	if err != nil {
		log.Fatal(err)
	}
	defer of.Close()

	img, err := jpeg.Decode(of)
	// pngファイル作成します。
	cf, err := os.Create(filepath.Join(directory, strings.Split(filename,".")[0] + ".png"))
	if err != nil {
		log.Fatal(err)
	}
	defer cf.Close()

	err = png.Encode(cf, img)
	if err != nil {
		log.Fatal(err)
	}
}

func EncodePNG(directory, filename string)  {
	fmt.Println(directory, filename)
	// pngファイルを開きます。
	of, err := os.Open(filepath.Join(directory, filename))
	if err != nil {
		log.Fatal(err)
	}
	defer of.Close()

	img, err := png.Decode(of)
	// jpgファイル作成します。
	cf, err := os.Create(filepath.Join(directory, strings.Split(filename,".")[0] + ".jpg"))
	if err != nil {
		log.Fatal(err)
	}
	defer cf.Close()

	err = jpeg.Encode(cf, img, nil)
	if err != nil {
		log.Fatal(err)
	}
}