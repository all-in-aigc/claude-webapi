package claude

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/launchdarkly/eventsource"
	"github.com/tidwall/gjson"
)

// ChatStream: chat strem reply
type ChatStream struct {
	Stream chan *gjson.Result // chat message stream
	Err    error              // error message
}

// GetChatStream is used to get chat stream
func (c *Client) GetChatStream(params MixMap) (*ChatStream, error) {
	uri := "/api/append_message"

	resp, err := c.post(uri, params)
	if err != nil {
		return nil, err
	}

	contentType := resp.Header.Get("Content-Type")
	// not event-strem response
	if !strings.HasPrefix(contentType, "text/event-stream") {
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)

		if strings.HasPrefix(contentType, "application/json") {
			res := gjson.ParseBytes(body)
			errmsg := res.Get("error.message").String()
			errcode := res.Get("error.type").String()
			if errmsg != "" {
				return nil, fmt.Errorf("response failed: [%s] %s", errcode, errmsg)
			}
		}

		return nil, fmt.Errorf("response failed: [%s] %s", resp.Status, body)
	}

	chatStream := &ChatStream{
		Stream: make(chan *gjson.Result),
		Err:    nil,
	}

	decoder := eventsource.NewDecoderWithOptions(resp.Body,
		eventsource.DecoderOptionReadTimeout(c.opts.Timeout))

	go func() {
		defer resp.Body.Close()
		defer close(chatStream.Stream)

		for {
			event, err := decoder.Decode()
			if err != nil {
				chatStream.Err = fmt.Errorf("decode data failed: %v", err)
				return
			}

			text := event.Data()

			jres := gjson.Parse(text)

			if jres.Get("model").String() == "" {
				chatStream.Err = fmt.Errorf("invalid stream data: %s", text)
				return
			}

			chatStream.Stream <- &jres
		}
	}()

	return chatStream, nil
}
