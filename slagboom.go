package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var allowedKentekens = []string{"AB-123-CD", "EF-456-GH", "XY-789-ZZ"}

func getGroet() string {
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

func kentekenIsAllowed(kenteken string) bool {
	for _, k := range allowedKentekens {
		if k == kenteken {
			return true
		}
	}
	return false
}

func main() {
	flag.Parse()

	if flag.NArg() == 1 {
		kenteken := flag.Arg(0)
		if kentekenIsAllowed(kenteken) {
			groet := getGroet()
			if groet == "" {
				fmt.Println("Sorry, de parkeerplaats is 's nachts gesloten.")
			} else {
				fmt.Printf("%s! Welkom bij Fonteyn Vakantieparken\n", groet)
			}
		} else {
			fmt.Println("U heeft helaas geen toegang tot het parkeerterrein")
		}
	} else if flag.NArg() > 1 {
		fmt.Println("Onbekende parameters")
		os.Exit(1)
	} else {
		fmt.Println("Geen kenteken opgegeven")
		os.Exit(1)
	}
}
