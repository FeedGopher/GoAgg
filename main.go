/*
	Base package for the ATOM server.

	Server periodically queries the ATOM sources of its users, stores updates

*/
package main

import (
	"fmt"
	"log"
	"os"
	"net/http"

	"github.com/clbanning/mxj"
	// "database/sql"
	// "github.com/lib/pq"
)

var testFeed = "https://www.blogger.com/feeds/8100407163665430627/posts/default"

func main() {
	xml := fetchFeedXML(testFeed)

	file, err := os.Create("output.json")
	if err != nil {
		fmt.Println("Error creating test output file.", err)
	}
	writeXMLAsJSON(xml, file) // should be writing to a db or server response
}

func fetchFeedXML(feedURI string) (mxj.Map) {
	response, err := http.Get(feedURI)
	if err != nil {
		fmt.Println("ATOM GET error:", err)
		log.Fatal(err)
	}
	defer response.Body.Close()
	// fmt.Printf("response body: %s", response.Body)

	xmlMap, err := mxj.NewMapXmlReader(response.Body, true)
	if err != nil {
		fmt.Println("Error creating map from XML reader", err)
	}

	return xmlMap
}

func writeXMLAsJSON(xmlMap mxj.Map, destination *os.File) () {
	xmlMap.JsonIndentWriter(destination, "", "\t", true)
	return
}

/*func fetchFeed(feedURI string) (io.Reader) {
	response, err := http.Get(feedURI)
	if err != nil {
		fmt.Println("ATOM GET error:", err)
		log.Fatal(err)
	}
	defer response.Body.Close()
	// fmt.Printf("response body: %s", response.Body)

	// We're getting a bug where feedToXML is trying to parse the closed response.
	// A couple solutions come to mind: create a new reader: response.Body like bufio.NewReader(response.Body),
	// copy the data: (ioutils.ReadAll(body)), or combine the functions.
	// I imagine that this is a common problem that has a solid solution

	return response.Body
}

func feedToXML(feed io.Reader) (mxj.Map) {
	// Use mxj package to translate XML structured bytes to a map[string]interface{}
	fmt.Printf("converting to xml %S", feed)
	_ = "breakpoint"
	xmlMap, err := mxj.NewMapXmlReader(feed, true)
	if err != nil {
		fmt.Println("Error creating map from XML reader", err)
	}
	return xmlMap
}*/
