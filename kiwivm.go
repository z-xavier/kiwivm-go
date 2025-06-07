package kiwivm

import (
	"context"
	"encoding/json"
)

type BandwidthAlert struct {
	FriendlyDescription string `json:"friendly_description"`
	IsEnabled           int    `json:"is_enabled"`
	ChangedTimestamp    int64  `json:"changed_timestamp"`
	SValue              string `json:"s_value"`
}

type EmailPreferences struct {
	BandwidthAlerts       map[string]BandwidthAlert `json:"Bandwidth Alerts"`
	SecurityNotifications map[string]BandwidthAlert `json:"Security Notifications"`
	SystemNotifications   map[string]BandwidthAlert `json:"System Notifications"`
	Miscellaneous         map[string]BandwidthAlert `json:"Miscellaneous"`
}

type KiwiVMGetNotificationPreferencesRsp struct {
	// Array of available notifications and their state
	EmailPreferences EmailPreferences `json:"email_preferences"`
	// Currently configured e-mail address where notifications are sent
	NotificationEmail string `json:"notificationEmail"`
	Error             int    `json:"error"`
}

// KiwiVMGetNotificationPreferences Returns all available notification settings, as well as their state
func (c *Client) KiwiVMGetNotificationPreferences(ctx context.Context) (*KiwiVMGetNotificationPreferencesRsp, error) {
	return Get[KiwiVMGetNotificationPreferencesRsp](ctx, c, "/v1/kiwivm/getNotificationPreferences")
}

type KiwiVMSetNotificationPreferencesReq struct {
	// JsonNotificationPreferences (json formatted array, preference_id:0/1)
	// Changes notification preferences
	JsonNotificationPreferences string `json:"json_notification_preferences"`
}

type KiwiVMSetNotificationPreferencesRsp struct {
	// TODO
	// Array of submitted changes
	SubmittedEmailPreferences json.RawMessage `json:"submitted_email_preferences"`
	// Array of actually changed preferences
	UpdatedEmailPreferences json.RawMessage `json:"updated_email_preferences"`
	// Friendly descriptions of all preferences
	FriendlyDescriptions string `json:"friendly_descriptions"`
	Error                int    `json:"error"`
}

// KiwiVMSetNotificationPreferences Changes notification preferences
// TODO: Need Test
func (c *Client) KiwiVMSetNotificationPreferences(ctx context.Context, req *KiwiVMSetNotificationPreferencesReq) (*KiwiVMSetNotificationPreferencesRsp, error) {
	return GetWithQueryParams[KiwiVMSetNotificationPreferencesRsp](ctx, c, "/v1/kiwivm/setNotificationPreferences", req)
}
