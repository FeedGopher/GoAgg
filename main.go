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

func main() {
	response, err := http.Get("https://www.blogger.com/feeds/8100407163665430627/posts/default")
	if err != nil {
		fmt.Println("ATOM GET error:", err)
		log.Fatal(err)
	}
	defer response.Body.Close()
	fmt.Printf("response body: %s", response.Body)

	// Use mxj package to translate XML structured bytes to a map[string]interface{}
	xmlMap, err := mxj.NewMapXmlReader(response.Body)
	if err != nil {
		fmt.Println("Error creating map from XML reader", err)
	}

	// Output map[string]interface{...} to file as JSON.
	f, err := os.Create("output.json")	
	if err != nil {
		fmt.Println("Error creating test output file.", err)
	}
	xmlMap.JsonIndentWriter(f, "", "\t", true)
}
