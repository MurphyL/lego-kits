package notify

import "murphyl.com/lego/fns/notify/internal"

/* 创建通知管理器 */
func NewNotificationManager() NotificationManager {
	return &internal.NotificationManager{}
}

/* 通知管理器（Notification Manager） */
type NotificationManager interface {
	Send(internal.NotificationMessage) error
}
