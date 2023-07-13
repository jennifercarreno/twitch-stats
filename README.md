# twitch-stats
![Screen Shot 2023-07-13 at 3 20 58 PM](https://github.com/jennifercarreno/twitch-stats/assets/60307928/354ecbb2-6081-488b-a427-7b8114ce14f0)

## Description
This project is a webscraper that writes a csv file containing the channel name, language, and follower count for the top 50 most followed creators on Twitch. It uses the Colly web scraping framework to extract data from twitchmetrics.net.
## Installation
1. Clone this repo
```
git clone https://github.com/jennifercarreno/twitch-stats.git
```
2. Install project dependencies
```
go mod download
```
## How to use
To run the Twitch Stats Web Scraper build and run the project using the Go command
```
go run main.go
```
To view the data visualization, open ```index.html``` with Live Server
