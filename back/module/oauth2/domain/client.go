package oauth2domain

import (
	"fmt"
)

// Client is identified by the following fields.
//   - ID
//
// Only public clients (PKCE required) are supported.
type Client struct {
	ID           string
	Name         string
	ClientURI    string
	RedirectURIs []string
}

func ConstructClient(fn func(c *Client) (err error)) (*Client, error) {
	c := new(Client)
	if err := fn(c); err != nil {
		return nil, err
	}

	if c.ID == "" ||
		c.Name == "" ||
		len(c.RedirectURIs) == 0 {
		return nil, fmt.Errorf("failed to construct %+v", c)
	}

	return c, nil
}
