package docker

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Tag is a single Docker Image Tag.
type Tag struct {
	Name        string    `json:"name"`
	LastUpdated time.Time `json:"last_updated"`
}

type listTagsResponse struct {
	Count    int     `json:"count"`
	Previous *string `json:"previous"`
	Next     *string `json:"next"`
	Items    []Tag   `json:"results"`
}

// LatestTags lists the latest tags on a repository.
func LatestTags(repository string) ([]Tag, error) {
	tags := []Tag{}
	url := fmt.Sprintf("https://hub.docker.com/v2/repositories/%s/tags/", repository)
	for {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		if resp.StatusCode != 200 {
			return nil, fmt.Errorf("bad response code: %s", resp.Status)
		}
		payload := &listTagsResponse{}
		err = json.NewDecoder(resp.Body).Decode(payload)
		resp.Body.Close()
		if err != nil {
			return nil, err
		}
		for _, tag := range payload.Items {
			tags = append(tags, tag)
		}
		if payload.Next == nil {
			break
		}
		url = *payload.Next
	}
	return tags, nil
}
