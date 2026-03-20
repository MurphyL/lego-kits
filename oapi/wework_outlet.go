package open

import (
	"murphyl.com/lego/oapi/internal/wework"
)

func NewWeworkChatGroupPushService(key string) {
	wework.NewChatGroupPushService(key)
}
