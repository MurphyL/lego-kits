package notify

import "murphyl.com/lego/fns/notify/internal"

/* 通知管理器（Notification Manager） */
type NotificationManager interface {
	Send(internal.NotificationMessage) error
}
