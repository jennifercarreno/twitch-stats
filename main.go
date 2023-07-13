package main

import (
    "fmt"
	"github.com/gocolly/colly"
	"encoding/csv"
	"os"
	"strings"

)

type Channel struct {
	Name string
	Language string
	Followers string
}

func main() {
	c := colly.NewCollector()
	channels := []Channel{}

	c.OnHTML(".py-4", func(e *colly.HTMLElement) {
		var name string
		var language string
		var followers string

		e.ForEach(".mb-0", func(_ int, el *colly.HTMLElement) {
			name = el.Text
		})

		foundLanguage := false // Flag to track if the desired language is found

    e.ForEach(".mr-3", func(_ int, el *colly.HTMLElement) {
        if foundLanguage {
            return // Skip further processing if the language is already found
        }

        language = el.Text

        if strings.Contains(language, "EN") {
            language = "EN"
            foundLanguage = true 
        } else if strings.Contains(language, "ES") {
            language = "ES"
            foundLanguage = true
        } else if strings.Contains(language, "PT") {
            language = "PT"
            foundLanguage = true
        } else {
			language = "other"
			foundLanguage = true

		}
    })

		e.ForEach("samp", func(_ int, el *colly.HTMLElement) {
			followers = el.Text
		})

		channel := Channel{Name: name, Language: language, Followers: followers}
		channels = append(channels, channel)
	})

	

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://www.twitchmetrics.net/channels/follower")
	for _, channel := range channels {
		fmt.Println(channel.Name,channel.Followers, channel.Language)
	}

    file, err := os.Create("output.csv")
    if err != nil {
		fmt.Println("Error creating CSV file:", err)
    }
    defer file.Close()

    // Create a CSV writer
    writer := csv.NewWriter(file)

    // Write the header row to the CSV file
    header := []string{"Channel", "Language", "Followers"}
    writer.Write(header)

    // Write the data rows to the CSV file
    for _, d := range channels {
        row := []string{d.Name, d.Language, d.Followers}
        writer.Write(row)
    }

    // Flush the CSV writer to ensure all data is written
    writer.Flush()

    // Check for any errors during the writing process
    if err := writer.Error(); err != nil {
        fmt.Println("Error writing CSV:", err)
    }

    fmt.Println("Scraping complete. Data saved to output.csv")
	
}