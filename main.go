package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

const apiKey = "****************"

func main() {
	ctx := context.Background()

	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Error creating YouTube service: %v", err)
	}

	part := []string{"id", "snippet", "contentDetails", "statistics"}
	channelStats, err := service.Channels.List(part).Id("UCpDJl2EmP7Oh90Vylx0dZtA").Do()
	if err != nil {
		log.Fatalf("Error fetching channel statistics: %v", err)
	}

	if len(channelStats.Items) > 0 {
		channel := channelStats.Items[0]
		fmt.Printf("Channel ID: %s\n", channel.Id)
		fmt.Printf("Channel Title: %s\n", channel.Snippet.Title)
		fmt.Printf("Subscriber Count: %d\n", channel.Statistics.SubscriberCount)
		fmt.Printf("Video Count: %d\n", channel.Statistics.VideoCount)

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
