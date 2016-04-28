package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.898, -74.23432,
	},
}

func main() {
	fmt.Println(m)
}
