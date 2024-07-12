package models

import "time"

type Albums struct {
	ID                int       `json:"id"`
	Title             string    `json:"title"`
	Description       string    `json:"description"`
	Duration          string    `json:"duration"`
	Artist            string    `json:"artist"`
	Label             string    `json:"label"`
	CreationTimestamp time.Time `json:"creationTimestamp"`
	UpdateTimestamp   time.Time `json:"updateTimestamp"`
}
