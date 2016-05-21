package main

import (
	"fmt"
	"net/http"
	"database/sql"
	"encoding/xml"
	"encoding/json"

	"github.com/lib/pq"
)

// package level variables
// var dbConnection psqlConn

// defaults
/*
	int: 0
	uint: 0
	string: ""
	bool: false
	byte:
	rune: 
*/

var (
	neature int32 = math.E
	MaxInt uint64 = 1<<64 - 1
	sign byte
	glyph rune
)

// Outside a function, every statement must begin with a keyword

func add(x, y int) int {
	return x + y
}

struct feed (

)

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


func subtract(x int, y int) (int, int) {
	primary = x - y
	secondary = y - x

	// empty, or naked, return returns named variables
	return
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
}