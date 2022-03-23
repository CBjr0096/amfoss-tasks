package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {

	fName := "rich.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("File not created, err :%q", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector(
		colly.AllowedDomains("https://www.forbes.com/real-time-billionaires/#131241903d78"),
	)

	c.OnHTML(".base ng-scope", func(e *colly.HTMLElement) {
		writer.Write([]string{
			e.ChildText(".ng-scope~a"),
			e.ChildText("span"),
			e.ChildText("ng-binging"),
			e.ChildText("ng-scope~span"),
		})
	})

	log.Printf("Scraping Done\n")
	log.Println(c)

}
