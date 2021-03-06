package controller

import (
	"app/config/env"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

func GetVideos(c echo.Context) error {
	title := c.Param("title")
	developerKey := env.Env.API.Youtube.Key

	client := &http.Client{
		Transport: &transport.APIKey{Key: developerKey},
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	arr := []string{"id", "snippet"}

	// Make the API call to YouTube.
	call := service.Search.List(arr).
		Q(title).
		MaxResults(25)
	response, err := call.Do()
	// handleError(err, "")

	// Group video, channel, and playlist results in separate lists.
	videos := make(map[string]string)

	// Iterate through each item and add it to the correct list.
	for _, item := range response.Items {
		switch item.Id.Kind {
		case "youtube#video":
			videos[item.Id.VideoId] = item.Snippet.Title
		}
	}

	printIDs("Videos", videos)

	return c.JSON(http.StatusOK, response)
}

func GetVideo(c echo.Context) error {
	videoId := c.Param("videoId")
	developerKey := env.Env.API.Youtube.Key

	client := &http.Client{
		Transport: &transport.APIKey{Key: developerKey},
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	arr := []string{"id", "snippet"}

	// Make the API call to YouTube.
	call := service.Search.List(arr).
		Q(videoId).
		MaxResults(1)
	response, err := call.Do()
	// handleError(err, "")

	// Group video, channel, and playlist results in separate lists.
	videos := make(map[string]string)

	// Iterate through each item and add it to the correct list.
	for _, item := range response.Items {
		switch item.Id.Kind {
		case "youtube#video":
			videos[item.Id.VideoId] = item.Snippet.Title
		}
	}

	printIDs("Videos", videos)

	return c.JSON(http.StatusOK, response)
}

// Print the ID and title of each result in a list as well as a name that
// identifies the list. For example, print the word section name "Videos"
// above a list of video search results, followed by the video ID and title
// of each matching video.
func printIDs(sectionName string, matches map[string]string) {
	fmt.Printf("%v:\n", sectionName)
	for id, title := range matches {
		fmt.Printf("[%v] %v\n", id, title)
	}
	fmt.Printf("\n\n")
}
