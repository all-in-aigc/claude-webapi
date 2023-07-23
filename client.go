package claude

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"github.com/tidwall/gjson"
)

const (
	BASE_URI = "https://claude.ai"
)

// MixMap is a type alias for map[string]interface{}
type MixMap = map[string]interface{}

// Client is a ChatGPT request client
type Client struct {
	opts    Options // custom options
	httpCli *http.Client
}

// NewClient will return a ChatGPT request client
func NewClient(options ...Option) *Client {
	cli := &Client{
		opts: Options{
			BaseUri:   BASE_URI,
			UserAgent: "xxx",
			Timeout:   120 * time.Second, // set default timeout
			Model:     "claude-2",        // set default chat model
			Debug:     false,
		},
	}

	// load custom options
	for _, option := range options {
		option(cli)
	}

	cli.initHttpClient()

	return cli
}

func (c *Client) initHttpClient() {
	transport := &http.Transport{}

	if c.opts.Proxy != "" {
		proxy, err := url.Parse(c.opts.Proxy)
		if err == nil {
			transport.Proxy = http.ProxyURL(proxy)
		}
	}

	c.httpCli = &http.Client{
		Timeout:   c.opts.Timeout,
		Transport: transport,
	}
}

// Get will request api with Get method
func (c *Client) Get(uri string) (*gjson.Result, error) {
	if !strings.HasPrefix(uri, "http") {
		uri = c.opts.BaseUri + uri
	}

	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("new request failed: %v", err)
	}

	resp, err := c.doRequest(req)

	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}

	return c.parseBody(resp)
}

// Post will request api with Post method
func (c *Client) Post(uri string, params MixMap) (*gjson.Result, error) {
	resp, err := c.post(uri, params)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}

	return c.parseBody(resp)
}

func (c *Client) post(uri string, params MixMap) (*http.Response, error) {
	data, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("marshal request body failed: %v", err)
	}

	if !strings.HasPrefix(uri, "http") {
		uri = c.opts.BaseUri + uri
	}

	req, err := http.NewRequest(http.MethodPost, uri, bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("new request failed: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	return c.doRequest(req)
}

func (c *Client) doRequest(req *http.Request) (*http.Response, error) {
	req.Header.Set("Cookie", fmt.Sprintf("sessionKey=%s", c.opts.SessionKey))
	req.Header.Set("User-Agent", c.opts.UserAgent)

	if c.opts.Debug {
		reqInfo, _ := httputil.DumpRequest(req, true)
		log.Printf("http request info: \n%s\n", reqInfo)
	}

	resp, err := c.httpCli.Do(req)

	if c.opts.Debug {
		respInfo, _ := httputil.DumpResponse(resp, false)
		log.Printf("http response info: \n%s\n", respInfo)
	}

	return resp, err
}

func (c *Client) parseBody(resp *http.Response) (*gjson.Result, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	res := gjson.ParseBytes(body)

	return &res, nil
}

// GetOrgid will get orgid set in option
func (c *Client) GetOrgid() string {
	return c.opts.Orgid
}

// GetModel will get model set in option
func (c *Client) GetModel() string {
	return c.opts.Model
}
