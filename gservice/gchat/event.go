package gchat

import "google.golang.org/api/chat/v1"

// based chat.DeprecatedEvent
type ChatEvent struct {
	Action    *chat.FormAction `json:"action,omitempty"`
	EventTime string           `json:"eventTime,omitempty"`
	Message   *chat.Message    `json:"message,omitempty"`
	Space     *chat.Space      `json:"space,omitempty"`
	Type      string           `json:"type,omitempty"`
	User      *User            `json:"user,omitempty"`
}

// based chat.User
type User struct {
	DisplayName string `json:"displayName,omitempty"`
	DomainId    string `json:"domainId,omitempty"`
	Name        string `json:"name,omitempty"`
	Type        string `json:"type,omitempty"`
	Email       string `json:"email,omitempty"`
}
