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
	xmlFile := os.Open("/home/ben/Media/Code/go/gotests/test.xml")
	xmldata, _ := ioutil.ReadAll(xmlFile)
	
	xml.Unmarshal(xmlData, x)
	fmt.Printf("foo: %s\n", x.Foo)

}