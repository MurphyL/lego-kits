package notify

import (
	"testing"

	"murphyl.com/lego/fns/notify/internal"
)

type MockNotificationMessage struct {
}

func (m *MockNotificationMessage) Content() string {
	return "Mock Notification Message"
}

func TestNotificationManager(t *testing.T) {
	manager := NewNotificationManager()
	manager.Send(&MockNotificationMessage{}, internal.ChannelEmail)
}
