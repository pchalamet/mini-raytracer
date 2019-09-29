package main

//import "fmt"
import "github.com/pchalamet/mini-raytracer/foundation"

func main() {
	c := foundation.NewCanvas(320, 200)
	err := c.Save("d:\\toto.png")
	if err != nil {
		panic("Failed to save canvas")
	}
}