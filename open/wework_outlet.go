package open

import (
	"github.com/MurphyL/lego-kits/open/internal/wework"
)

func NewWeworkChatGroupPushService(key string) {
	wework.NewChatGroupPushService(key)
}
