package wework

import (
	"testing"
)

func TestPushMessageToChatGroup(t *testing.T) {
	pushService := NewChatGroupPushService("8174eee9-77a3-452e-b015-17053003750a")
	t.Run("text message", func(t *testing.T) {
		pushService.SendTextMessage("hello")
	})
}
