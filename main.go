package main

import (
	"generaterHead/generater"
	"log"
	"os"
	"strings"
)

func main() {
	fname := generater.GetVal("head", "fname")
	title := generater.GetVal("head", "title")
	subject := generater.GetVal("subtitle", "subject")
	event := generater.GetVal("subtitle", "event")
	otherword := generater.GetVal("subtitle", "otherword")
	subtitle := generater.MarketingGenerater(subject, event, otherword)
	author := generater.GetVal("head", "author")
	catalog := generater.GetVal("head", "catalog")
	tags := generater.GetVal("head", "tags")

	fname, ret := generater.Head(fname, title, subtitle, author, catalog, tags)
	writeLines(fname, ret)
}
func writeLines(fname string, s []string) {
	f, err := os.OpenFile(fname, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0776)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	for _, v := range s {
		_, err := f.WriteString(strings.Join([]string{v, "\n"}, ""))
		if err != nil {
			return
		}
	}
}
