package kiwivm

import "context"

type Backup struct {
	Size      int    `json:"size"`
	OS        string `json:"os"`
	MD5       string `json:"md5"`
	Timestamp int64  `json:"timestamp"`
}

type BackupListRsp struct {
	// Array of backups (backup_token, size, os, md5, timestamp).
	Backups map[string]Backup `json:"backups"`
	Error   int               `json:"error"`
}

// BackupList Get list of automatic backups.
func (c *Client) BackupList(ctx context.Context) (*BackupListRsp, error) {
	return Get[BackupListRsp](ctx, c, "/v1/backup/list")
}

type BackupCopyToSnapshotReq struct {
	BackupToken string `json:"backupToken"`
}

type BackupCopyToSnapshotRsp struct {
	NotificationEmail string `json:"notificationEmail"`
	Error             int    `json:"error"`
}

// BackupCopyToSnapshot Copies a backup identified by backup_token
// (returned by backup/list) into a restorable Snapshot.
func (c *Client) BackupCopyToSnapshot(ctx context.Context, req *BackupCopyToSnapshotReq) (*BackupCopyToSnapshotRsp, error) {
	return GetWithQueryParams[BackupCopyToSnapshotRsp](ctx, c, "/v1/backup/copyToSnapshot", req)
}
