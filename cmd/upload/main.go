package main

import (
	"log"
	"os"

	"github.com/kamatama41/md-to-prd-sample/client"

	"google.golang.org/api/drive/v3"
)

var pdfFile = "work/sample.pdf"

func main() {
	tok, err := client.GetTokenFromFile()
	if err != nil {
		log.Fatalf("Failed to get token from file: %v", err)
	}
	srv, err := client.NewService(tok)
	if err != nil {
		log.Fatalf("Failed to retrieve Drive client: %v", err)
	}

	f, err := os.Open(pdfFile)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer f.Close()
	uploadedFile, err := srv.Files.Create(&drive.File{Name: f.Name()}).Media(f).Do()
	if err != nil {
		log.Fatalf("Failed to upload file: %v", err)
	}
	println(uploadedFile.Id)
}
