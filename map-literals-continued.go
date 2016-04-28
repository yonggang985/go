package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.134, -74.243},
	"Google":    {37.2343, -122.23432},
}

func main() {
	fmt.Println(m)
}
