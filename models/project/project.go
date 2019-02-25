package project

import "time"

type Project struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	URL            string    `json:"url"`
	State          string    `json:"state"`
	Revision       int       `json:"revision"`
	Visibility     string    `json:"visibility"`
	LastUpdateTime time.Time `json:"lastUpdateTime"`
	Description    string    `json:"description,omitempty"`
}

type Projects struct {
	Count   int       `json:"count"`
	Project []Project `json:"value"`
}
