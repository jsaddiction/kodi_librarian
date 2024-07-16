package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type RadarrTestEvent struct {
	ApplicationUrl string `json:"applicationUrl"`
	EventType      string `json:"eventType"`
	InstanceName   string `json:"instanceName"`
	Movie          struct {
		FolderPath  string   `json:"folderPath"`
		ID          int      `json:"id"`
		ReleaseDate string   `json:"releaseDate"`
		Tags        []string `json:"tags"`
		Title       string   `json:"title"`
		TmdbID      int      `json:"tmdbId"`
		Year        int      `json:"year"`
	} `json:"movie"`
	Release struct {
		CustomFormatScore int    `json:"customFormatScore"`
		Indexer           string `json:"indexer"`
		Quality           string `json:"quality"`
		QualityVersion    int    `json:"qualityVersion"`
		ReleaseGroup      string `json:"releaseGroup"`
		ReleaseTitle      string `json:"releaseTitle"`
		Size              int    `json:"size"`
	} `json:"release"`
	RemoteMovie struct {
		ImdbID string `json:"imdbId"`
		Title  string `json:"title"`
		TmdbID int    `json:"tmdbId"`
		Year   int    `json:"year"`
	} `json:"remoteMovie"`
}

func handleRadarr(w http.ResponseWriter, r *http.Request) {
	fmt.Println(strings.Repeat("=", 80)) // ASCII line
	fmt.Println("Radarr Event received")
	fmt.Println(strings.Repeat("=", 80)) // ASCII line

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	var payload map[string]interface{}
	if err := json.Unmarshal(body, &payload); err != nil {
		log.Printf("Error parsing JSON: %v", err)
		http.Error(w, "can't parse JSON", http.StatusBadRequest)
		return
	}

	eventType, ok := payload["eventType"].(string)
	if !ok {
		log.Printf("Error: eventType not found or invalid")
		http.Error(w, "invalid event type", http.StatusBadRequest)
		return
	}

	switch eventType {
	case "Test":
		handleRadarrTest(w, body)
	// Add cases for other event types
	default:
		log.Printf("Unhandled event type: %s", eventType)
		http.Error(w, "unhandled event type", http.StatusNotImplemented)
	}
	fmt.Println(strings.Repeat("-", 80)) // ASCII line
	fmt.Println()
}

func handleRadarrTest(w http.ResponseWriter, body []byte) {
	fmt.Println(strings.Repeat("=", 80)) // ASCII line
	fmt.Println("Radarr Test Event received")
	fmt.Println(strings.Repeat("=", 80)) // ASCII line

	var event RadarrTestEvent
	if err := json.Unmarshal(body, &event); err != nil {
		log.Printf("Error parsing JSON: %v", err)
		http.Error(w, "can't parse JSON", http.StatusBadRequest)
		return
	}

	prettyJSON, err := json.MarshalIndent(event, "", "  ")
	if err != nil {
		log.Printf("Error pretty printing JSON: %v", err)
		return
	}
	log.Println(string(prettyJSON))
	fmt.Println(strings.Repeat("-", 80)) // ASCII line
	fmt.Println()
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(strings.Repeat("=", 80)) // ASCII line
		fmt.Println("Request received")
		fmt.Println(strings.Repeat("=", 80)) // ASCII line

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading body: %v", err)
			http.Error(w, "can't read body", http.StatusBadRequest)
			return
		}

		var payload map[string]interface{}
		if err := json.Unmarshal(body, &payload); err != nil {
			log.Printf("Error parsing JSON: %v", err)
			return
		}

		if eventResource, ok := payload["event_resource"].(string); ok {
			var parsedEventResource interface{}
			if err := json.Unmarshal([]byte(eventResource), &parsedEventResource); err != nil {
				log.Printf("Error parsing event resource: %v", err)
				return
			}
			payload["event_resource"] = parsedEventResource
		}

		prettyJSON, err := json.MarshalIndent(payload, "", "  ")
		if err != nil {
			log.Printf("Error pretty printing JSON: %v", err)
			return
		}
		log.Println(string(prettyJSON))
		fmt.Println(strings.Repeat("-", 80)) // ASCII line
		fmt.Println()
	})
	fmt.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
