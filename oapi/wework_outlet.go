package open

import (
	"github.com/MurphyL/lego-kits/oapi/internal/wework"
)

func NewWeworkChatGroupPushService(key string) {
	wework.NewChatGroupPushService(key)
}
