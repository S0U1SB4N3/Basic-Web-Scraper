/*
 * Copyright (c) 2022. All Rights Reserved Next Generation Cybernetics, LLC
 */
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

	file, err2 := os.Create("output.csv")
	if err2 != nil {
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

//how to build an executable from project
//go build -o build/basic_web_scraper basic_web_scraper.go
//Run: go tool dist list
//For MacOS: GOOS=darwin GOARCH=amd64 go build -o build/basic_web_scraper_macOS basic_web_scraper.go
//To Execute the MacOS executable: ./build/basic_web_scraper
//For Linux: GOOS=linux GOARCH=amd64 go build -o build/basic_web_scraper_linux basic_web_scraper.go
//To Execute the Linux executable: ./build/basic_web_scraper
//For Windows: GOOS=windows GOARCH=amd64 go build -o build/basic_web_scraper_64.exe basic_web_scraper.go
//To Execute the Windows executable: basic_web_scraper.exe
