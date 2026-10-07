package oauth2dbmodel

import (
	"time"

	oauth2domain "github.com/twin-te/twin-te/back/module/oauth2/domain"
)

type Client struct {
	ID           string
	Name         string
	ClientURI    string
	RedirectURIs []string `gorm:"column:redirect_uris;serializer:json"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Client) TableName() string { return "oauth2_clients" }

func FromDBClient(dbClient *Client) (*oauth2domain.Client, error) {
	return oauth2domain.ConstructClient(func(c *oauth2domain.Client) (err error) {
		c.ID = dbClient.ID
		c.Name = dbClient.Name
		c.ClientURI = dbClient.ClientURI
		c.RedirectURIs = dbClient.RedirectURIs
		return
	})
}

func ToDBClient(client *oauth2domain.Client) *Client {
	return &Client{
		ID:           client.ID,
		Name:         client.Name,
		ClientURI:    client.ClientURI,
		RedirectURIs: client.RedirectURIs,
	}
}
