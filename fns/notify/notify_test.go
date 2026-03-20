package notify

import (
	"testing"
)

type MockNotificationMessage struct {
}

func (m *MockNotificationMessage) Content() string {
	return "Mock Notification Message"
}

func TestNotificationManager(t *testing.T) {
}
