package scraper

import (
	"fmt"
    "log"
    "strconv"
    "scraper/database"
    "scraper/models"

    "github.com/gocolly/colly/v2"
	

)

func ScrapeTopGainers() {
    c := colly.NewCollector()

    // Use CSS selectors to target the table rows
    c.OnHTML("table tbody tr", func(e *colly.HTMLElement) {
        gainer := models.Gainer{}

        // Extract data using CSS selectors and convert as needed
        gainer.Symbol = e.ChildText("td:nth-child(1) span:nth-child(1) div a span:nth-child(1)")
        gainer.Name = e.ChildText("td:nth-child(1) span:nth-child(1) div a span:nth-child(2)")

        priceStr := e.ChildText("td:nth-child(2)")
        gainer.Price, _ = strconv.ParseFloat(priceStr, 64)

        changeStr := e.ChildText("td:nth-child(3)")
        gainer.Change, _ = strconv.ParseFloat(changeStr, 64)

        percentStr := e.ChildText("td:nth-child(4)")
        percentStr = percentStr[:len(percentStr)-1] // Remove % symbol
        gainer.Percent, _ = strconv.ParseFloat(percentStr, 64)
		fmt.Println(gainer.Price)
        // Save to the database
        result := database.DB.Create(&gainer)
        if result.Error != nil {
            log.Printf("Error saving gainer: %v", result.Error)
        }
    })

    log.Println("Starting scraping...")
    c.Visit("https://finance.yahoo.com/markets/stocks/gainers/")
    log.Println("Scraping finished.")
}