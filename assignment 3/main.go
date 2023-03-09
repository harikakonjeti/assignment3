package main

import (
	"url/Downloader/File"
	"url/Downloader/Web"
	"url/Downloader/Zip"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	downloader := Web.NewWebDownloader()

	r1, err := downloader.UrlDownload("http://www.yahoo.com/image_to_read.jpg")
	if err != nil {
		log.Fatal(err)
	}

	Filedownloader := File.NewFileDownloader()
	r2, err := Filedownloader.UrlDownload("/Users/konjeti.harika/Downloads/Signature.jpeg")
	if err != nil {
		log.Fatal(err)
	}

	zipper := Zip.New()

	outR, _ := zipper.Archive([]string{"image_to_read.jpg", "Signature.jpeg"}, r1, r2)
	fp, _ := os.Create("result.zip")
	fmt.Println(io.Copy(fp, outR))
}
