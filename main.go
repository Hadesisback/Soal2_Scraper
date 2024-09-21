package main

import (
    "log"
    "scraper/database"
    "scraper/scraper"

    "github.com/robfig/cron/v3"
)

func main() {
    // Initialize the database
    database.InitDatabase()

    // Create a new cron scheduler
    c := cron.New()
	scraper.ScrapeTopGainers()
    // Schedule the scraper to run every day at midnight
    _, err := c.AddFunc("@daily", scraper.ScrapeTopGainers)
    if err != nil {
        log.Fatalf("Error scheduling job: %v", err)
    }

    // Start the cron scheduler
    c.Start()

    log.Println("Scheduler started. Scraping daily...")

    // Block the main thread to keep the application running
    select {}
}
