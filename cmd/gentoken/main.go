package main

import (
	"fmt"
	"log"

	"github.com/kamatama41/md-to-prd-sample/client"
)

func main() {
	tok, err := client.GetTokenFromWeb()
	if err != nil {
		log.Fatalf("Failed to get access token: %v", err)
	}

	// Test
	srv, err := client.NewService(tok)
	if err != nil {
		log.Fatalf("Failed to retrieve Drive client: %v", err)
	}

	r, err := srv.Files.List().PageSize(10).
		Fields("nextPageToken, files(id, name)").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve files: %v", err)
	}
	fmt.Println("Files:")
	if len(r.Files) == 0 {
		fmt.Println("No files found.")
	} else {
		for _, i := range r.Files {
			fmt.Printf("%s (%s)\n", i.Name, i.Id)
		}
	}
}
