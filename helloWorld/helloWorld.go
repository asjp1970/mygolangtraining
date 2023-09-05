// The first exercise one always do when learning a new programming language
package main

import (
	"flag"
	"fmt"
	"log"
)

type Lang int32

// language abreviations
const (
	Es Lang = iota
	En
	Fr
	It
	Unkown
)

// program arguments names and default values
const langArgName = "lang"
const defaultLang = "es"

// global vars for program input flags, defined at init time
var lang string

func init() {
	flag.StringVar(&lang, langArgName, defaultLang, "What language you want me to say hello")
	flag.Parse()
}

// Converts the more or less user frienly language abreviation provided in the command
// line into the corresponding value of type Lang.
// Returns error if the language name is not recognized.
func convertLangNameToLang(languageName string) (Lang, error) {
	switch languageName {
	case "es", "spanish", "español":
		return Es, nil
	case "en", "english":
		return En, nil
	case "fr", "french", "français":
		return Fr, nil
	case "it", "italian", "italiano":
		return It, nil
	default:
		return Unkown, fmt.Errorf("not a known language (%v)", languageName)
	}

}

// Taking Lang as input, produces the salutation message in the language it knows;
// if the language is unkown to it, then it will just let you know in plain English.
func babelSaluter(l Lang) string {
	log.Printf("I'll say hello in %v", l.toString())
	switch l {
	case Es:
		return "Hola mundo ¿Qué tal?"
	case En:
		return "Hello workd. What's up?"
	case Fr:
		return "Bonjour tour le monde ¿Comment ça va?"
	case It:
		return "Ciao mondo"
	default:
		return fmt.Sprintf("Sorry, I don't understand this language")
	}
}

// Method with receiver of type Lang to return the corresponding language name,
// mainly for logging
func (l Lang) toString() string {
	switch l {
	case Es:
		return "Spanish"
	case En:
		return "English"
	case Fr:
		return "French"
	case It:
		return "Italian"
	}
	return "unknown"
}

func main() {
	l, err := convertLangNameToLang(lang)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(babelSaluter(l))
}
