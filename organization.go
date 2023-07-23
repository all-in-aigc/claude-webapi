package claude

import (
	"github.com/tidwall/gjson"
)

// GetOrganizations is used to get account organizations
func (c *Client) GetOrganizations() (*gjson.Result, error) {
	uri := "/api/organizations"

	return c.Get(uri)
}
