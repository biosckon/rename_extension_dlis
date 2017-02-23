package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	files, err := ioutil.ReadDir(".")

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fn := strings.Split(file.Name(), ".")
		if "DLI" == fn[1] {
			newfn := strings.Join([]string{fn[0], ".dlis"}, "")
			err = os.Rename(file.Name(), newfn)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
