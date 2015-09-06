package main

import (
	"os"
	"fmt"
	"encoding/xml"
	"io/ioutil"
)

type XML struct {
	Foo []string `xml:"foo"`
}


func main() {
	x := XML{}
	xmlFile, err := os.Open("/home/ben/Media/Code/go/gotests/test.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	
	xml.Unmarshal(xmlData, x)
	fmt.Printf("foo: %s\n", x.Foo)

}