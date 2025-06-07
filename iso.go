package kiwivm

import "context"

type ISOMountReq struct {
	ISO string `json:"iso"`
}

type ISOMountRsp struct {
	Error int `json:"error"`
}

// ISOMount Sets ISO image to boot from.
// VM must be completely shut down and restarted after this API call.
// TODO: Need Tests
func (c *Client) ISOMount(ctx context.Context, req *ISOMountReq) (*ISOMountRsp, error) {
	return GetWithQueryParams[ISOMountRsp](ctx, c, "/v1/iso/mount", req)
}

type ISOUnmountRsp struct {
	Error int `json:"error"`
}

// ISOUnmount Removes ISO image and configures VM to boot from primary storage.
// VM must be completely shut down and restarted after this API call.
// TODO: Need Tests
func (c *Client) ISOUnmount(ctx context.Context) (*ISOUnmountRsp, error) {
	return Get[ISOUnmountRsp](ctx, c, "/v1/iso/unmount")
}
