package main

import (
	"fmt"
	"strings"
)

func main() {
	classNameKey := "gett_express_uk"
	classNameTitle, classLocale := getNamesFromKey(classNameKey)
	fmt.Println("classNameTitle: ", classNameTitle)
	fmt.Println("classLocale: ", classLocale)
}

// This function will get a classNameKey wich it a string contain class name
// and locale.
// The function will return `classNameTitle` and `classLocale` both are strings
// Notice that I declared then both in return paranthises so no need to `:=` them
func getNamesFromKey(classNameKey string) (classNameTitle, classLocale string) {
	classNameTitle = strings.Title(strings.Replace(classNameKey, "_", " ", -1))
	classNameParsed := strings.Split(classNameKey, "_")
	classLocale = classNameParsed[len(classNameParsed)-1]
	return classNameTitle, classLocale
}
