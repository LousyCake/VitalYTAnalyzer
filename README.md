
# VitalYTAnalyzer - YouTube Channel Statistics Analyzer

## Overview
VitalYTAnalyzer is a Go application that utilizes the YouTube API to fetch statistics for a specific channel. It provides insights into subscriber count, video count, and channel start date.

__Important Note:__ The YouTube API provides real-time data, and the numbers retrieved represent the most up-to-date information as of the time of the API request. Differences between API statistics and those displayed on the YouTube website may arise due to real-time updates, processing delays, or data caching on the website.

## Prerequisites
1. __Google Cloud Project:__

+ Create a Google Cloud Project:  [Google Cloud Console](https://console.cloud.google.com/).
+ Enable the YouTube Data API v3 for your project.
2. __API Key:__

+ Generate an API key in the Google Cloud Console.
3. __Go Language:__

+ Install Go: [Go Installation Guide](https://go.dev/doc/install).

## Installation
1. __Clone the repository:__
```
git clone https://github.com/yourusername/VitalYTAnalyzer.git
cd VitalYTAnalyzer
```
2. __Set up your API key:__

+ Open __main.go.__
+ Replace __"YOUR_API_KEY"__ with your actual API key.
3. __Run the application:__

```
go run main.go
```
## Usage

### Fetch Channel Statistics
To fetch statistics for a specific channel, modify the __'Id'__ parameter in __'service.Channels.List().Id("YOUR_CHANNEL_ID").Do()'__.

### Interpretation
+ Subscriber Count: Represents the number of subscribers to the channel.
+ Video Count: Indicates the total number of videos on the channel.
+ Channel Start Date: Displays the date when the channel was created.

### Example
```
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

const apiKey = "YOUR_API_KEY"

func main() {
	ctx := context.Background()

	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Error creating YouTube service: %v", err)
	}

	// Example: Fetching channel statistics for a specific channel
	part := []string{"id", "snippet", "contentDetails", "statistics"}
	channelStats, err := service.Channels.List(part).Id("YOUR_CHANNEL_ID").Do()
	if err != nil {
		log.Fatalf("Error fetching channel statistics: %v", err)
	}

	// Accessing and printing channel details
	if len(channelStats.Items) > 0 {
		channel := channelStats.Items[0]
		fmt.Printf("Channel ID: %s\n", channel.Id)
		fmt.Printf("Channel Title: %s\n", channel.Snippet.Title)
		fmt.Printf("Subscriber Count: %d\n", channel.Statistics.SubscriberCount)
		fmt.Printf("Video Count: %d\n", channel.Statistics.VideoCount)

		// Channel start date (publishedAt in snippet)
		startDate, err := time.Parse(time.RFC3339, channel.Snippet.PublishedAt)
		if err == nil {
			fmt.Printf("Channel Start Date: %s\n", startDate.Format("2006-01-02"))
		} else {
			fmt.Println("Error parsing channel start date:", err)
		}
	} else {
		fmt.Println("No channel statistics found.")
	}
}
```
