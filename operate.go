package kiwivm

import "context"

type StartRsp struct {
	IsMounted int `json:"is_mounted"`
	Error     int `json:"error"`
}

// Start the VPS
func (c *Client) Start(ctx context.Context) (*StartRsp, error) {
	return Get[StartRsp](ctx, c, "/v1/start")
}

type StopRsp struct {
	Error int `json:"error"`
}

// Stop the VPS
func (c *Client) Stop(ctx context.Context) (*StopRsp, error) {
	return Get[StopRsp](ctx, c, "/v1/stop")
}

type RestartRsp struct {
	Error int `json:"error"`
}

// Restart Reboots the VPS
func (c *Client) Restart(ctx context.Context) (*RestartRsp, error) {
	return Get[RestartRsp](ctx, c, "/v1/restart")
}

type KillRsp struct {
	Error int `json:"error"`
}

// Kill Allows to forcibly stop a VPS that is stuck and cannot be stopped by normal means.
// Please use this feature with great care as any unsaved data will be lost.
// TODO: Need Test
func (c *Client) Kill(ctx context.Context) (*KillRsp, error) {
	return Get[KillRsp](ctx, c, "/v1/kill")
}
