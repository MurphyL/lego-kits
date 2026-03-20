package internal

/* 通知渠道（Notification Channels） */
type NotificationChannelType string

const (
	// 邮件
	ChannelEmail NotificationChannelType = "email"
	// 短信
	ChannelSMS NotificationChannelType = "sms"
	// 企业微信群组会话机器人
	ChannelWecomChatBot NotificationChannelType = "wecom_chat_bot"
	// 飞书机群组会话器人
	ChannelFeishuChatBot NotificationChannelType = "feishu_chat_bot"
)
