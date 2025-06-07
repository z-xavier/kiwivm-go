package kiwivm

import "context"

type GetAvailableOSRsp struct {
	// Currently installed Operating System
	Installed string `json:"installed"`
	// Array of available OS
	Templates []string `json:"templates"`
	Error     int      `json:"error"`
}

func (c *Client) GetAvailableOS(ctx context.Context) (*GetAvailableOSRsp, error) {
	return Get[GetAvailableOSRsp](ctx, c, "/v1/getAvailableOS")
}

type ReinstallOSReq struct {
	OS string `json:"os"`
}

type ReinstallOSRsp struct {
	Error int `json:"error"`
}

// ReinstallOS Reinstall the Operating System.
// OS must be specified via "os" variable.
// Use getAvailableOS call to get list of available systems.
// TODO: Need Test
func (c *Client) ReinstallOS(ctx context.Context, req *ReinstallOSReq) (*ReinstallOSRsp, error) {
	return GetWithQueryParams[ReinstallOSRsp](ctx, c, "/v1/reinstallOS", req)
}
