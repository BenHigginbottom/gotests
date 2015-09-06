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
	xmlFile := os.Open("/home/ben/Media/Code/go/gotests/test.xml")
	xmlData, _ := ioutil.ReadAll(xmlFile)
	
	xml.Unmarshal(xmlData, x)
	fmt.Printf("foo: %s\n", x.Foo)

}