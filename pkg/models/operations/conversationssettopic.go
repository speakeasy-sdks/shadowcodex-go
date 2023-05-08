// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type ConversationsSetTopicSecurity struct {
	SlackAuth string `security:"scheme,type=oauth2,name=Authorization"`
}

type ConversationsSetTopicApplicationJSON struct {
	// Conversation to set the topic of
	Channel *string `json:"channel,omitempty"`
	// The new topic string. Does not support formatting or linkification.
	Topic *string `json:"topic,omitempty"`
}

type ConversationsSetTopicRequest struct {
	RequestBody *ConversationsSetTopicApplicationJSON `request:"mediaType=application/json"`
	// Authentication token. Requires scope: `conversations:write`
	Token *string `header:"style=simple,explode=false,name=token"`
}

type ConversationsSetTopicResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Typical error response
	ConversationsSetTopicErrorSchema map[string]map[string]interface{}
	// Typical success response
	ConversationsSetTopicSuccessSchema map[string]map[string]interface{}
}
