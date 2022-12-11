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
	my_url := "https://shop.76fireworks.com/category/Artillery-Shells"
	response, error := http.Get(my_url)
	if error != nil {
		log.Fatal(error)
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
	writer.Write(content)
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
			writer.Write(content)
		}
	})

	writer.Flush()
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

}
