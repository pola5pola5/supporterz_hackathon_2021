package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	f := flag.String("imgpath", "./test2.jpg", "input image_file")
	flag.Parse()
	filepath := *f
	file, err := os.Open(filepath)
	if err != nil {
		log.Println("[-]Couldn't opne file: ", filepath)
	}
	defer file.Close()

	// encode
	binary, _ := ioutil.ReadAll(file)
	base64Data := base64.StdEncoding.EncodeToString(binary)
	fmt.Print(base64Data)
}
