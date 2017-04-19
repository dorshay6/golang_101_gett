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

func getNamesFromKey(classNameKey string) (classNameTitle, classLocale string) {
	classNameTitle = strings.Title(strings.Replace(classNameKey, "_", " ", -1))
	classNameParsed := strings.Split(classNameKey, "_")
	classLocale = classNameParsed[len(classNameParsed)-1]
	return classNameTitle, classLocale
}
