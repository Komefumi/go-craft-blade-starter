package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var filenames []string

	root, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}

		if !info.IsDir() {
			filenames = append(filenames, path)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	for _, currentFilename := range filenames {
		content, err := ioutil.ReadFile(currentFilename)
		if err != nil {
			log.Fatal(err)
		}
		newContentString := strings.Replace(string(content), "craft-blade", "craft-blade", -1)
		err = os.WriteFile(currentFilename, []byte(newContentString), 0666)
		if err != nil {
			log.Fatal(err)
		}
	}
}
