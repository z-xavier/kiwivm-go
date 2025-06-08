package kiwivm

import (
	"context"
)

type SnapshotCreateReq struct {
	Description string `json:"description"`
}

type SnapshotCreateRsp struct {
	// E-mail address on file where notification will be sent once task is completed.
	NotificationEmail string `json:"notificationEmail"`
	Error             int    `json:"error"`
}

// SnapshotCreate Create snapshot
func (c *Client) SnapshotCreate(ctx context.Context, req *SnapshotCreateReq) (*SnapshotCreateRsp, error) {
	return GetWithQueryParams[SnapshotCreateRsp](ctx, c, "/v1/snapshot/create", req)
}

type Snapshot struct {
	FileName        string `json:"fileName"`
	Os              string `json:"os"`
	Description     string `json:"description"`
	Size            string `json:"size"`
	Md5             string `json:"md5"`
	Sticky          bool   `json:"sticky"`
	Uncompressed    int64  `json:"uncompressed"`
	PurgesIn        int64  `json:"purgesIn"`
	DownloadLink    string `json:"downloadLink"`
	DownloadLinkSSL string `json:"downloadLinkSSL"`
}

type SnapshotListRsp struct {
	// Array of snapshots (fileName, os, description, size, md5, sticky, purgesIn, downloadLink).
	Snapshots []*Snapshot `json:"snapshots"`
	Error     int         `json:"error"`
}

// SnapshotList Get list of snapshots.
func (c *Client) SnapshotList(ctx context.Context) (*SnapshotListRsp, error) {
	return Get[SnapshotListRsp](ctx, c, "/v1/snapshot/list")
}

type SnapshotDeleteReq struct {
	Snapshot string `json:"snapshot"`
}

type SnapshotDeleteRsp struct {
	NotificationEmail string `json:"notificationEmail"`
	Error             int    `json:"error"`
}

// SnapshotDelete Delete snapshot by fileName (can be retrieved with snapshot/list call).
func (c *Client) SnapshotDelete(ctx context.Context, req *SnapshotDeleteReq) (*SnapshotDeleteRsp, error) {
	return GetWithQueryParams[SnapshotDeleteRsp](ctx, c, "/v1/snapshot/delete", req)
}

type SnapshotRestoreReq struct {
	Snapshot string `json:"snapshot"`
}

type SnapshotRestoreRsp struct {
	Error int `json:"error"`
}

// SnapshotRestore Restores snapshot by fileName (can be retrieved with snapshot/list call).
// This will overwrite all data on the VPS.
// TODO: Need Tests
func (c *Client) SnapshotRestore(ctx context.Context, req *SnapshotRestoreReq) (*SnapshotRestoreRsp, error) {
	return GetWithQueryParams[SnapshotRestoreRsp](ctx, c, "/v1/snapshot/restore", req)
}

type Sticky int8

const (
	StickyRemove Sticky = iota
	StickySet
)

type SnapshotToggleStickyReq struct {
	Snapshot string `json:"snapshot"`
	Sticky   Sticky `json:"sticky"`
}

type SnapshotToggleStickyRsp struct {
	Error int `json:"error"`
}

// SnapshotToggleSticky Set or remove sticky attribute ("sticky" snapshots are never purged).
// Name of snapshot can be retrieved with snapshot/list call – look for fileName variable.
func (c *Client) SnapshotToggleSticky(ctx context.Context, req *SnapshotToggleStickyReq) (*SnapshotToggleStickyRsp, error) {
	return GetWithQueryParams[SnapshotToggleStickyRsp](ctx, c, "/v1/snapshot/toggleSticky", req)
}

type SnapshotExportReq struct {
	Snapshot string `json:"snapshot"`
}

type SnapshotExportRsp struct {
	Token string `json:"token"`
	Error int    `json:"error"`
}

// SnapshotExport Generates a token with which the snapshot can be transferred to another instance.
func (c *Client) SnapshotExport(ctx context.Context, req *SnapshotExportReq) (*SnapshotExportRsp, error) {
	return GetWithQueryParams[SnapshotExportRsp](ctx, c, "/v1/snapshot/export", req)
}

type SnapshotImportReq struct {
	SourceVeID  string `json:"sourceVeid"`
	SourceToken string `json:"sourceToken"`
}

type SnapshotImportRsp struct {
	Error int `json:"error"`
}

// SnapshotImport Imports a snapshot from another instance identified by VEID and Token.
// Both VEID and Token must be obtained from another instance beforehand with a snapshot/export call.
// TODO: Need Tests
func (c *Client) SnapshotImport(ctx context.Context, req *SnapshotImportReq) (*SnapshotImportRsp, error) {
	return GetWithQueryParams[SnapshotImportRsp](ctx, c, "/v1/snapshot/import", req)
}
