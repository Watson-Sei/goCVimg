package main

import (
	"ChImage/chencode"
	"log"

	//"ChImage/chencode"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	//	"log"
//	"os"
//	"path/filepath"
)

var direcotry = flag.String("dir", "", "ディレクトリパスを設定")
var extenstion = flag.String("extend", "", "変更したい拡張子を設定")

func main() {
	flag.Parse()
	switch *extenstion {
	case "jpg":
		fmt.Println("jpgをpngに変換します。")
		err := filepath.Walk(*direcotry,
			func(path string, info os.FileInfo, err error) error {
				if filepath.Ext(path) == ".jpg" {
					fmt.Println(filepath.Dir(path), filepath.Base(path))
					chencode.EncodeJPG(filepath.Dir(path), filepath.Base(path))
				}
				return nil
			})
		if err != nil {
			log.Fatal(err)
		}
	case "png":
		fmt.Println("pngをjpgに変換します。")
		err := filepath.Walk(*direcotry,
			func(path string, info os.FileInfo, err error) error {
				if filepath.Ext(path) == ".png" {
					fmt.Println(filepath.Dir(path), filepath.Base(path))
					chencode.EncodePNG(filepath.Dir(path), filepath.Base(path))
				}
				return nil
			})
		if err != nil {
			log.Fatal(err)
		}
	default:
		fmt.Println("他の拡張子は利用できません。")
	}
}