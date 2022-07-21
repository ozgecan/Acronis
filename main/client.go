package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func downloadFile(filePath string) {

	// Build fileName from fullPath
	fileURL, err := url.Parse(filePath)
	if err != nil {
		log.Fatal(err)
	}
	path := fileURL.Path
	segments := strings.Split(path, "/")
	fileName := segments[len(segments)-1]
	fmt.Print(segments)

	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	// Put content on file
	resp, err := client.Get(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Print(strings.Split(string(bodyBytes), "")[3])
	if strings.Split(string(bodyBytes), "")[3] == "A" {
		// Create blank file
		file, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}
		size, err := io.Copy(file, resp.Body)
		if err != nil {
			fmt.Errorf("file cannot download %s", fileName)
		}
		defer file.Close()
		fmt.Printf("Downloaded a file %s with size %d", fileName, size)
	} else {
		fmt.Print("File didn't download because don't have A")
	}

}
func main() {
	downloadFile("http://127.0.0.1:8080/file4.txt")
}
