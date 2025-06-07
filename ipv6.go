package kiwivm

import (
	"context"
	"encoding/json"
)

type IPV6AddReq struct {
	IP string `json:"ip"`
}

type IPV6AddRsp struct {
	// TODO
	AssignedSubnet json.RawMessage `json:"assigned_subnet"`
	Error          int             `json:"error"`
}

// IPV6Add Assigns a new IPv6 address.
// For initial IPv6 assignment an empty IP is required (call without parameters),
// and a new IP from the available pool is assigned automatically.
// All subsequent requested IPv6 addresses must be within the /64 subnet of the
// first IPv6 address.
// TODO: Need Test
func (c *Client) IPV6Add(ctx context.Context, req *IPV6AddReq) (*IPV6AddRsp, error) {
	return GetWithQueryParams[IPV6AddRsp](ctx, c, "/v1/ipv6/add", req)
}

type IPV6DeleteReq struct {
	IP string `json:"ip"`
}

type IPV6DeleteRsp struct {
	Error int `json:"error"`
}

// IPV6Delete Releases specified IPv6 address.
// TODO: Need Test
func (c *Client) IPV6Delete(ctx context.Context, req *IPV6DeleteReq) (*IPV6DeleteRsp, error) {
	return GetWithQueryParams[IPV6DeleteRsp](ctx, c, "/v1/ipv6/delete", req)
}
