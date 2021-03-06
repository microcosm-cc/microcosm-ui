package models

import "time"

// Site describes a site
type Site struct {
	ID                 int64          `json:"siteId"`
	SiteURL            string         `json:"siteURL"`
	Title              string         `json:"title"`
	Description        string         `json:"description"`
	SubdomainKey       string         `json:"subdomainKey"`
	Domain             string         `json:"domain"`
	ForceSSL           bool           `json:"forceSSL"`
	OwnedBy            ProfileSummary `json:"ownedBy"`
	ThemeID            int64          `json:"themeId"`
	LogoURL            string         `json:"logoUrl"`
	FaviconURL         string         `json:"faviconUrl,omitempty"`
	BackgroundColor    string         `json:"backgroundColor"`
	BackgroundURL      string         `json:"backgroundUrl,omitempty"`
	BackgroundPosition string         `json:"backgroundPosition,omitempty"`
	LinkColor          string         `json:"linkColor"`
	GaWebPropertyID    string         `json:"gaWebPropertyId,omitempty"`
	Menu               []Link         `json:"menu"`
	Auth0Domain        string         `json:"auth0Domain,omitempty"`
	Auth0ClientID      string         `json:"auth0ClientId,omitempty"`

	Meta struct {
		Created   time.Time      `json:"created"`
		CreatedBy ProfileSummary `json:"createdBy"`

		Flags struct {
			Deleted bool `json:"deleted"`
		} `json:"flags,omitempty"`

		Stats       []Stat     `json:"stats,omitempty"`
		Links       []Link     `json:"links,omitempty"`
		Permissions Permission `json:"permissions,omitempty"`
	} `json:"meta"`
}

// SiteResponse describes the API response that contains a site
type SiteResponse struct {
	BoilerPlate
	Data Site `json:"data"`
}
