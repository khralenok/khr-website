package utilities

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type youtubeResponse struct {
	Items []struct {
		Snippet struct {
			Title       string `json:"title"`
			Description string `json:"description"`
		} `json:"snippet"`
	} `json:"items"`
}

func FetchYoutubeVideo(videoId string) (title, description string, err error) {
	apiKey := os.Getenv("YOUTUBE_API_KEY")

	if apiKey == "" {
		return "", "", errors.New("there is no YOUTUBE_API_KEY in .env")
	}

	url := fmt.Sprintf("https://www.googleapis.com/youtube/v3/videos?part=snippet&id=%s&key=%s", videoId, apiKey)

	data, err := getDataFromYoutubeApi(url)

	if err != nil {
		return "", "", err
	}

	title = data.Items[0].Snippet.Title
	description = data.Items[0].Snippet.Description

	return title, description, nil
}

func getDataFromYoutubeApi(url string) (youtubeResponse, error) {
	resp, err := http.Get(url)

	if err != nil {
		return youtubeResponse{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return youtubeResponse{}, fmt.Errorf("youtube: %s: %s", resp.Status, string(body))
	}

	var data youtubeResponse

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return youtubeResponse{}, fmt.Errorf("can't parse response JSON: %w", err)
	}

	if len(data.Items) == 0 {
		return youtubeResponse{}, fmt.Errorf("no items in response: %v", data.Items) //errors.New("no items in response")
	}

	return data, nil
}
