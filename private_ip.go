package kiwivm

import "context"

type PrivateIPGetAvailableIPsRsp struct {
	// TODO
	//assigned_ips: Array of successfully assigned private IP addresses
	Error int `json:"error"`
}

// PrivateIPGetAvailableIPs Returns all available (free) IPv4 addresses which you can activate on VM
// TODO: Need Test
func (c *Client) PrivateIPGetAvailableIPs(ctx context.Context) (*PrivateIPGetAvailableIPsRsp, error) {
	return Get[PrivateIPGetAvailableIPsRsp](ctx, c, "/v1/privateIp/getAvailableIps")
}

type PrivateIpAssignReq struct {
	IP string `json:"ip"`
}

type PrivateIpAssignRsp struct {
	// TODO
	//assigned_ips: Array of successfully assigned private IP addresses
	Error int `json:"error"`
}

// PrivateIpAssign Assign private IP address.
// If IP address not specified, a random address will be assigned.
// TODO: Need Test
func (c *Client) PrivateIpAssign(ctx context.Context, req *PrivateIpAssignReq) (*PrivateIpAssignRsp, error) {
	return GetWithQueryParams[PrivateIpAssignRsp](ctx, c, "/v1/privateIp/assign", req)
}

type PrivateIpDeleteReq struct {
	IP string `json:"ip"`
}

type PrivateIpDeleteRsp struct {
	Error int `json:"error"`
}

// PrivateIpDelete Delete private IP address.
// TODO: Need Test
func (c *Client) PrivateIpDelete(ctx context.Context, req *PrivateIpDeleteReq) (*PrivateIpDeleteRsp, error) {
	return GetWithQueryParams[PrivateIpDeleteRsp](ctx, c, "/v1/privateIp/delete", req)
}
