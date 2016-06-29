/*
	Base package for the ATOM server.

	Server periodically queries the ATOM sources of its users, stores updates

*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/clbanning/mxj"
	// "database/sql"
	// "github.com/lib/pq"
)

var testFeed = "https://www.blogger.com/feeds/8100407163665430627/posts/default"

func main() {
	body := fetchFeed(testFeed)
	xml := feedToXML(body)

	file, err := os.Create("output.json")
	if err != nil {
		fmt.Println("Error creating test output file.", err)
	}
	writeXMLAsJSON(xml, file) // should be writing to a db or server response
}

func fetchFeed(feedURI string) (io.Reader) {
	// feedURI instead
	response, err := http.Get(feedURI)
	if err != nil {
		fmt.Println("ATOM GET error:", err)
		log.Fatal(err)
	}
	defer response.Body.Close()
	// fmt.Printf("response body: %s", response.Body)

	return response.Body
}

func feedToXMLMap(feed io.Reader) (Map) {
	// Use mxj package to translate XML structured bytes to a map[string]interface{}
	xmlMap, err := mxj.NewMapXmlReader(response.Body)
	if err != nil {
		fmt.Println("Error creating map from XML reader", err)
	}
	return xmlMap
}

func writeXMLAsJSON(xmlMap Map, destination *File) () {
	xmlMap.JsonIndentWriter(destination, "", "\t", true)

	return
}