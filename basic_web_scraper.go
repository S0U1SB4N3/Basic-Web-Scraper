/*
 * Copyright (c) 2022. All Rights Reserved Next Generation Cybernetics, LLC
 */
// Resources used:
// GoQuery Docs: https://pkg.go.dev/github.com/PuerkitoBio/goquery#section-documentation
// CSV Writer: https://golang.org/pkg/encoding/csv/#example_Writer
//Further reading
// GoQuery Example: https://www.golangprograms.com/web-scraping-with-go.html
// GoQuery Example: https://www.golangprograms.com/scrape-data-from-web-page.html
package main

import (
	"encoding/csv"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	myUrl := "https://shop.76fireworks.com/category/Artillery-Shells"
	response, err1 := http.Get(myUrl)
	if err1 != nil {
		log.Fatal(err1)
	}
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("output.csv")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	//Setup CSV File
	writer := csv.NewWriter(file)
	content := []string{"Name", "Link"}
	errWriter := writer.Write(content)
	if errWriter != nil {
		return
	}
	writer.Flush()

	// Use goquery to select specific elements from the page
	doc.Find("table.product_list").Find("tbody").Find("tr").Find("td").Find("a").Each(func(i int, s *goquery.Selection) {
		// For each item found, link text and href are extracted
		linkTag := s.Find("span").First()
		link, _ := s.Attr("href")
		//prepend link with domain
		link = "https://shop.76fireworks.com" + link
		text := strings.TrimSpace(linkTag.Text())
		if text != "" {
			fmt.Printf("Link: %s Text: %s \n", link, text)

			content := []string{text, link}
			errWriterContent := writer.Write(content)
			if errWriterContent != nil {
				return
			}
		}
	})

	//Flush the writer and close the file
	writer.Flush()
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
}
