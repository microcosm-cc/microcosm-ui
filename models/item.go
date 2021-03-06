package models

import (
	"encoding/json"
	"time"
)

// Item describes a generic item that can be stored within a microcosm
type Item struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	MicrocosmID int64  `json:"microcosmId"`
	Breadcrumb  []struct {
		Rel      string `json:"rel,omitempty"` // REST
		Href     string `json:"href"`
		Title    string `json:"title,omitempty"`
		Text     string `json:"text,omitempty"` // HTML
		LogoURL  string `json:"logoUrl,omitempty"`
		ID       int64  `json:"id"`
		Level    int64  `json:"level,omitempty"`
		ParentID int64  `json:"parentId,omitempty"`
	} `json:"breadcrumb,omitempty"`

	ItemType string `json:"itemType"`

	CommentCount         int64          `json:"totalComments"`
	ViewCount            int64          `json:"totalViews"`
	LastCommentID        int64          `json:"lastCommentId,omitempty"`
	LastCommentCreatedBy ProfileSummary `json:"lastCommentCreatedBy,omitempty"`
	LastCommentCreated   string         `json:"lastCommentCreated,omitempty"`

	Meta DefaultMeta `json:"meta"`

	// Search only fields
	Item           json.RawMessage `json:"item"`
	ParentItemType string          `json:"parentItemType,omitempty"`
	ParentItem     json.RawMessage `json:"parentItem,omitempty"`
	Unread         bool            `json:"unread"`
	Rank           float64         `json:"rank"`
	LastModified   time.Time       `json:"lastModified"`
	Highlight      string          `json:"highlight"`
}

// ItemDetail describes an item within a microcosm
type ItemDetail struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	MicrocosmID int64  `json:"microcosmId"`
	Breadcrumb  []struct {
		Rel      string `json:"rel,omitempty"` // REST
		Href     string `json:"href"`
		Title    string `json:"title,omitempty"`
		Text     string `json:"text,omitempty"` // HTML
		LogoURL  string `json:"logoUrl,omitempty"`
		ID       int64  `json:"id"`
		Level    int64  `json:"level,omitempty"`
		ParentID int64  `json:"parentId,omitempty"`
	} `json:"breadcrumb,omitempty"`
}

// ItemDetailCommentsAndMeta describes the comments associated to an item
type ItemDetailCommentsAndMeta struct {
	Comments Array       `json:"comments"`
	Meta     DefaultMeta `json:"meta"`
}

// ItemSummary describes an item within a microcosm for the purpose of listing
// microcosm contents
type ItemSummary struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`

	CommentCount         int64          `json:"totalComments"`
	ViewCount            int64          `json:"totalViews"`
	LastCommentID        int64          `json:"lastCommentId,omitempty"`
	LastCommentCreatedBy ProfileSummary `json:"lastCommentCreatedBy,omitempty"`
	LastCommentCreated   string         `json:"lastCommentCreated,omitempty"`

	MicrocosmID int64 `json:"microcosmId"`
	Breadcrumb  []struct {
		Rel      string `json:"rel,omitempty"` // REST
		Href     string `json:"href"`
		Title    string `json:"title,omitempty"`
		Text     string `json:"text,omitempty"` // HTML
		LogoURL  string `json:"logoUrl,omitempty"`
		ID       int64  `json:"id"`
		Level    int64  `json:"level,omitempty"`
		ParentID int64  `json:"parentId,omitempty"`
	} `json:"breadcrumb,omitempty"`
}

// ItemSummaryMeta describes the metadata of an ItemSummary
type ItemSummaryMeta struct {
	CommentCount int64       `json:"totalComments"`
	ViewCount    int64       `json:"totalViews"`
	LastComment  LastComment `json:"lastComment,omitempty"`
	Meta         SummaryMeta `json:"meta"`
}

// LastComment describes the last visible comment made against an item
type LastComment struct {
	ID        int64          `json:"id"`
	Created   time.Time      `json:"created"`
	CreatedBy ProfileSummary `json:"createdBy"`
}
