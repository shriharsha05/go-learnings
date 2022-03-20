package main

import "fmt"

func main() {

	// var colors map[int]string

	// colors := make(map[int]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"white": "#ffffff",
	}

	// colors[10] = "#ffffff"
	// colors[20] = "#eeeeee"

	// delete(colors, 20)

	// fmt.Println(colors)

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for color", color, "is", hex)
	}

}
