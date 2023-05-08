// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type ChatScheduleMessageSecurity struct {
	SlackAuth string `security:"scheme,type=oauth2,name=Authorization"`
}

type ChatScheduleMessageApplicationJSON struct {
	// Pass true to post the message as the authed user, instead of as a bot. Defaults to false. See [chat.postMessage](chat.postMessage#authorship).
	AsUser *bool `json:"as_user,omitempty"`
	// A JSON-based array of structured attachments, presented as a URL-encoded string.
	Attachments *string `json:"attachments,omitempty"`
	// A JSON-based array of structured blocks, presented as a URL-encoded string.
	Blocks *string `json:"blocks,omitempty"`
	// Channel, private group, or DM channel to send message to. Can be an encoded ID, or a name. See [below](#channels) for more details.
	Channel *string `json:"channel,omitempty"`
	// Find and link channel names and usernames.
	LinkNames *bool `json:"link_names,omitempty"`
	// Change how messages are treated. Defaults to `none`. See [chat.postMessage](chat.postMessage#formatting).
	Parse *string `json:"parse,omitempty"`
	// Unix EPOCH timestamp of time in future to send the message.
	PostAt *string `json:"post_at,omitempty"`
	// Used in conjunction with `thread_ts` and indicates whether reply should be made visible to everyone in the channel or conversation. Defaults to `false`.
	ReplyBroadcast *bool `json:"reply_broadcast,omitempty"`
	// How this field works and whether it is required depends on other fields you use in your API call. [See below](#text_usage) for more detail.
	Text *string `json:"text,omitempty"`
	// Provide another message's `ts` value to make this message a reply. Avoid using a reply's `ts` value; use its parent instead.
	ThreadTs *float64 `json:"thread_ts,omitempty"`
	// Pass true to enable unfurling of primarily text-based content.
	UnfurlLinks *bool `json:"unfurl_links,omitempty"`
	// Pass false to disable unfurling of media content.
	UnfurlMedia *bool `json:"unfurl_media,omitempty"`
}

type ChatScheduleMessageRequest struct {
	RequestBody *ChatScheduleMessageApplicationJSON `request:"mediaType=application/json"`
	// Authentication token. Requires scope: `chat:write`
	Token *string `header:"style=simple,explode=false,name=token"`
}

type ChatScheduleMessageResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Typical error response if the `post_at` is invalid (ex. in the past or too far into the future)
	ChatScheduleMessageErrorSchema map[string]map[string]interface{}
	// Typical success response
	ChatScheduleMessageSuccessSchema map[string]map[string]interface{}
}
