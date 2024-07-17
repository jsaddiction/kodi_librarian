package radarr

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// Handle all events from radarr
func HandleEvent(w http.ResponseWriter, r *http.Request) {
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
		handleTestEvent(w, body)
	// Add cases for other event types
	default:
		log.Printf("Unhandled event type: %s", eventType)
		log.Printf("Payload: %v", payload)
		http.Error(w, "unhandled event type", http.StatusNotImplemented)
	}

	fmt.Println(strings.Repeat("-", 80)) // ASCII line
	fmt.Println()
}

func handleTestEvent(w http.ResponseWriter, body []byte) {
	var event TestEvent
	if err := json.Unmarshal(body, &event); err != nil {
		log.Printf("Error parsing JSON: %v", err)
		http.Error(w, "can't parse JSON", http.StatusBadRequest)
		return
	}

	log.Printf("Parsed event: %+v", event)
}
