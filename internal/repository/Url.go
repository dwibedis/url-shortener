package repository

import "time"

type Url struct {
	URL string `json:"url"`
}

type UrlDb struct {
	ID string
	ParentUrl string
	addedOn time.Time
}