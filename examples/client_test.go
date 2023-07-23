package examples

import claude "github.com/all-in-aigc/claude-webapi"

var (
	baseUri    string
	orgid      string
	sessionKey string
	userAgent  string
	debug      bool
)

func getClient() *claude.Client {
	cli := claude.NewClient(
		claude.WithBaseUri(baseUri),
		claude.WithSessionKey(sessionKey),
		claude.WithOrgid(orgid),
		claude.WithUserAgent(userAgent),
		claude.WithDebug(debug),
	)

	return cli
}
