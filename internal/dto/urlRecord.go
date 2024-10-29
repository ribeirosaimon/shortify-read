package dto

import "time"

type UrlRecord struct {
	Expiration time.Time `json:"expiration"`
	Url        string    `json:"url"`
}
