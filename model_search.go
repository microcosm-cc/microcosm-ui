package ui

import (
	"encoding/json"
	"time"
)

// SearchResults is a list of SearchResult
type SearchResults struct {
	Query     SearchQuery `json:"query"`
	TimeTaken int64       `json:"timeTakenInMs,omitempty"`
	// TODO: This needs to be an array of SearchResult
	Results interface{} `json:"results,omitempty"`
}

type SearchResult struct {
	ItemType       string          `json:"itemType"`
	Item           json.RawMessage `json:"item"`
	ParentItemType string          `json:"parentItemType,omitempty"`
	ParentItem     json.RawMessage `json:"parentItem,omitempty"`
	Unread         bool            `json:"unread"`
	Rank           float64         `json:"rank"`
	LastModified   time.Time       `json:"lastModified"`
	Highlight      string          `json:"highlight"`
}

type SearchQuery struct {
	Query             string   `json:"q,omitempty"`
	InTitle           bool     `json:"inTitle,omitempty"`
	Hashtags          []string `json:"hashtags,omitempty"`
	MicrocosmIDsQuery []int64  `json:"forumId,omitempty"`
	ItemTypesQuery    []string `json:"type,omitempty"`
	ItemIDsQuery      []int64  `json:"id,omitempty"`
	ProfileID         int64    `json:"authorId,omitempty"`
	Emails            []string `json:"email,omitempty"`
	Following         bool     `json:"following,omitempty"`
	Since             string   `json:"since,omitempty"`
	Until             string   `json:"until,omitempty"`
	EventAfter        string   `json:"eventAfter,omitempty"`
	EventBefore       string   `json:"eventBefore,omitempty"`
	Attendee          bool     `json:"attendee,omitempty"`
	Has               []string `json:"has,omitempty"`
	Sort              string   `json:"sort,omitempty"`

	Ignored  string `json:"ignored,omitempty"`
	Searched string `json:"searched,omitempty"`
}
