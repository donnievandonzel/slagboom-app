package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func getGroot() string {
	now := time.Now()
	hour := now.Hour()

	if 7 <= hour && hour < 12 {
		return "Goedemorgen"
	} else if 12 <= hour && hour < 18 {
		return "Goedemiddag"
	} else if 18 <= hour && hour < 23 {
		return "Goedenavond"
	} else {
		return ""
	}
}

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		groet := getGroot()
		if groet == "" {
			fmt.Println("Sorry, de parkeerplaats is 's nachts gesloten.")
		} else {
			fmt.Printf("%s! Welkom bij Fonteyn Vakantieparken\n", groet)
		}
	} else {
		fmt.Println("Onbekende parameters")
		os.Exit(1)
	}
}
