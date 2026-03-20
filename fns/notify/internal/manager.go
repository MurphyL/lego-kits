package internal

/* 通知消息（Notification Message） */
type NotificationMessage interface {
	Content() string
}

type NotificationManager struct {
}

func (n *NotificationManager) Send(msg NotificationMessage, channels ...NotificationChannelType) {
	// 实现发送通知的逻辑
}
