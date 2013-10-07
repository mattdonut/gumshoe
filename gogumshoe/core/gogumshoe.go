package gogumshoe

import (
    "strings"
    "net/http"
    "pathways"
    "events"
)

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    basepath := &pathways.Fragment{Pattern: "api", Forks: []*pathways.Fragment{events.Fragment()}}
    steps := strings.Split(r.URL.Path,"/")[2:]
    destination := pathways.Follow(steps, r.Method, basepath, r)
    destination(w, r)
}