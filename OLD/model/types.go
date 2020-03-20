package model

import (
	"time"
)

type DuoCMSpost struct {
	Title     string `json:"title"`
	Subtitle  string `json:"subtitle"`
	User      string
	Created   time.Time
	Modified  time.Time
	Content   []string
	Command   interface{} `json:"command"`
	Data      interface{} `json:"data"`
	Revisions []DuoCMSpost
}

type DuoCMSpage struct {
	Title     string `json:"title"`
	Subtitle  string `json:"subtitle"`
	User      string
	Created   time.Time
	Modified  time.Time
	Content   []string
	Command   interface{} `json:"command"`
	Data      interface{} `json:"data"`
	Revisions []DuoCMSpage
}

type DuoCMSimg struct {
	Name     string `json:"name"`
	User     string
	Created  time.Time
	Modified time.Time
}

type DuoCMSitem struct {
	Name     string `json:"name"`
	User     string
	Created  time.Time
	Modified time.Time
}
