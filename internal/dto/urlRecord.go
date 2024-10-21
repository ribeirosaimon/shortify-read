package dto

import "time"

type UrlRecord struct {
	Url        string    `json:"url"`
	Expiration time.Time `json:"expiration"`
}
