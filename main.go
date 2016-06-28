/*
	Base package for the ATOM server.

	Server periodically queries the ATOM sources of its users, stores updates

*/
package main

import (
	"fmt"
	"log"
	"net/http"
	// "bufio"
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

	// ByteBody, err := ioutil.ReadAll(response.Body)

	// // Unmarshall XML-encoded response body to a json-like map.
	// // mBody, err := mxj.NewMapReader(response.Body)
	// if err != nil {
	// 	fmt.Println("Feed response-to-XML error:", err)
	// 	log.Fatal(err)
	// }
	// fmt.Printf("mxj'd response body: %s", ByteBody)

	xmlMap, err := mxj.NewMapXmlReader(response.Body)
	if err != nil {
		fmt.Println("Error creating map from XML reader", err)
	}
	fmt.Printf("mxj XML map: %S", xmlMap)

	f, err := os.Create("output.json")	
	xmlMap.JsonIndentWriter(f, "", "\t", true)

	// JSON writer

	// 	// func to handle Map value from XML Reader
	// 	func maphandler(m mxj.Map) bool {

	// 		// marshal Map as JSON
	// 		jsonVal, err := m.Json()
	// 		if err != nil {
	// 			fmt.Println("JSON marshalling failed.", err)
	// 			log.Fatal(err)
	// 			return false // stop further processing of XML Reader
	// 		}

	// 		// write JSON somewhere
	// 		json, err = jsonWriter.Write(jsonVal)
	// 		if err != nil {
	// 			fmt.Println("Writing marshalled JSON failed.", err)
	// 			log.Fatal(err)
	// 			return false // stop further processing of XML Reader
	// 		}

	// 		return true // continue - get next XML from Reader
	// 	}

	// 	// func to handle error from unmarshaling XML Reader
	// 	func errhandler(errVal error) bool {
	// 		fmt.Println("Error caught by HandleXmlReader", errVal.Error());
	// 		return true // continue
	// 	}

	// 	err := mxj.HandleXmlReader(xmlReader, maphandler, errhandler)
	// 	if err != nil {
	// 		fmt.Println("Hamdle XML reader failed", err)
	// 	}

}
