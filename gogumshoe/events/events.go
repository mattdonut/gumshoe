package events

import (
	"fmt"
	"net/http"
	"pathways"
)

func Fragment() *pathways.Fragment {
	return &pathways.Fragment{Pattern: "events", Get: getEvents }
}

func getEvents(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from the events module!")
}

type Event struct {
	Time string
	place string
}

