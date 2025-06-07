package kiwivm

import "context"

type MigrateGetLocationsRsp struct {
	// ID of current location
	CurrentLocation string `json:"currentLocation"`
	// IDs of locations available for migration into
	Locations []string `json:"locations"`
	// Friendly descriptions of available locations
	Descriptions map[string]string `json:"descriptions"`
	// Some locations may offer more expensive bandwidth
	// where monthly allowance will be lower.
	// This array contains monthly data transfer allowance
	// multipliers for each location.
	DataTransferMultipliers map[string]int `json:"dataTransferMultipliers"`
	Error                   int            `json:"error"`
}

// MigrateGetLocations Return all possible migration locations.
func (c *Client) MigrateGetLocations(ctx context.Context) (*MigrateGetLocationsRsp, error) {
	return Get[MigrateGetLocationsRsp](ctx, c, "/v1/migrate/getLocations")
}

type MigrateStartReq struct {
	Location string `json:"location"`
}

type MigrateStartRsp struct {
	Error int `json:"error"`
}

// MigrateStart Start VPS migration to new location.
// Takes new location ID as input.
// Note that this will result in all IPv4 addresses
// to be replaced with new ones, and all IPv6 addresses
// will be released.
// TODO: Need Test
func (c *Client) MigrateStart(ctx context.Context, req *MigrateStartReq) (*MigrateStartRsp, error) {
	return GetWithQueryParams[MigrateStartRsp](ctx, c, "/v1/migrate/start", req)
}
