package main

import "strings"

func main() {
	classNameKey := "gett_express_uk"
	classNameTitle := testIf(classNameKey)
}

func testIf(classNameKey string) (classLocale string) {
	classNameTitle = strings.Title(strings.Replace(classNameKey, "_", " ", -1))
	classNameParsed := strings.Split(classNameKey, "_")
	classLocale = classNameParsed[len(classNameParsed)-1]
	return classNameTitle, classLocale
}
