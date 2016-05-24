/*
	Base package for the ATOM server.

	Server periodically queries the ATOM sources of its users, stores updates

*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	// "encoding/json"

	// "github.com/clbanning/mxj"
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

	

	JSON writer

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


/*

// package level variables
// var dbConnection psqlConn

// Outside a function, every statement must begin with a keyword
type Feed struct (

)

type Vertex struct {
	X int
	Y int
}

func fetchFeed(uri string) (*feed) {
	response, err := http.Get(uri)
	if err != nil {
		fmt.Println('error fetching feed')
	}

	decoder := xml.NewDecoder(response)

	err := decoder.Decode()
	if err != nil {
		fmt.Println('error decoding feed')
	}

}

// https://godoc.org/github.com/clbanning/mxj#example-HandleXmlReader
func XMLtoJSON(xml *smth) (json interface {}) {
	return json.Marshal(xml)
}

func add(x, y int) int {
	return x + y
}

func subtract(x int, y int) (int, int) {
	primary = x - y
	secondary = y - x

	// empty, or naked, return returns named variables
	return
}

func pow(x, n, lim float64) float64 {
	// if statement runs the first argument, making declarations available in
	// the if block scope and any related else blocks
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	// if an initializer (value) is provided, types may be omitted.
	var jelly, pb, raisins = 2, true, false

	fmt.Println(add(42, 13))

	// := indicates implicit typing
	pos, neg := subtract(42, 13)
	fmt.Println(pos, neg)

	const f = "%T(%v)\n"
	fmt.Printf(f, neature, neature)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, sign, sign)
	fmt.Printf(f, glyph, glyph)

	v := Vertex{1, 2}
	p := &v
	// when accessing a field of a struct you do not need to use a pointer *
	p.X = 1e9 // (*p).X = 1e9
	fmt.Println(v)
}

*/
