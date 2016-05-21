package main

import (
	"fmt"
	"http"
)

// var dbConnection psqlConn

func add(x, y int) int {
	return x + y
}

struct feed (

)

func fetchFeed(uri string) (*feed) {

}


func subtract(x int, y int) (int, int) {
	primary = x - y
	secondary = y - x

	// empty, or naked, return returns named variables
	return
}

func main() {
	fmt.Println(add(42, 13))

	pos, neg := subtract(42, 13)
	fmt.Println(pos, neg)
}