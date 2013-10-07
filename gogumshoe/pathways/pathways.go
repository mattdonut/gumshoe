package pathways

import (
	"net/http"
)

type Action func (w http.ResponseWriter, r *http.Request)

type Fragment struct {
	Pattern string
	Get Action
	Put Action
	Post Action
	Delete Action
	Forks []*Fragment
}

func (p *Fragment) Match(s string, r *http.Request) bool {
	if p.Pattern[0] == '<' {
		r.Form.Add(p.Pattern[1:len(p.Pattern) - 1], s)
		return true
	} else {
		if p.Pattern == s {
			return true
		}
	}
	return false
}

func DeadEnd(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "No handler for that Fragment", http.StatusNotFound)
}

func Follow(steps []string, method string, basepath *Fragment, r *http.Request) Action {
	endpoint := basepath
	for _, segment := range steps {
		matched := false
		for _, fork := range endpoint.Forks {
			if fork.Match(segment, r) {
				endpoint = fork
				matched = true
				break
			}
		}
		if matched == false {
			return DeadEnd
		}
	}
	var destination Action
	switch method {
		case "GET":
			destination = endpoint.Get
		case "POST":
			destination = endpoint.Post
		case "PUT":
			destination = endpoint.Put
		case "DELETE":
			destination = endpoint.Delete
	}
	if destination != nil {
		return destination
	} else {
		return DeadEnd
	}
}