package kiwivm

import (
	"context"
	"sort"
	"time"
)

type RawUsage struct {
	Timestamp       int64 `json:"timestamp"`
	CpuUsage        int64 `json:"cpu_usage"`
	NetworkInBytes  int64 `json:"network_in_bytes"`
	NetworkOutBytes int64 `json:"network_out_bytes"`
	DiskReadBytes   int64 `json:"disk_read_bytes"`
	DiskWriteBytes  int64 `json:"disk_write_bytes"`
}

type GetRawUsageStatsRsp struct {
	Data   []*RawUsage `json:"data"`
	VmType string      `json:"vm_type"`
	Error  int         `json:"error"`
}

func (s *GetRawUsageStatsRsp) GetData(beginTime, endTime time.Time) []*RawUsage {
	begin, end := 0, len(s.Data)
	if !beginTime.IsZero() {
		begin, _ = sort.Find(len(s.Data), func(i int) int {
			return int(beginTime.Unix() - s.Data[i].Timestamp)
		})
	}

	if !endTime.IsZero() {
		end, _ = sort.Find(len(s.Data), func(i int) int {
			return int(endTime.Unix() - s.Data[i].Timestamp)
		})
	}
	return s.Data[begin:end]
}

func (s *GetRawUsageStatsRsp) NetworkInBytes(beginTime, endTime time.Time) int64 {
	var sum int64
	for _, usage := range s.GetData(beginTime, endTime) {
		sum += usage.NetworkInBytes
	}
	return sum
}

func (s *GetRawUsageStatsRsp) NetworkOutBytes(beginTime, endTime time.Time) int64 {
	var sum int64
	for _, usage := range s.GetData(beginTime, endTime) {
		sum += usage.NetworkOutBytes
	}
	return sum
}

func (s *GetRawUsageStatsRsp) DiskReadBytes(beginTime, endTime time.Time) int64 {
	var sum int64
	for _, usage := range s.GetData(beginTime, endTime) {
		sum += usage.DiskReadBytes
	}
	return sum
}

func (s *GetRawUsageStatsRsp) DiskWriteBytes(beginTime, endTime time.Time) int64 {
	var sum int64
	for _, usage := range s.GetData(beginTime, endTime) {
		sum += usage.DiskWriteBytes
	}
	return sum
}

// GetRawUsageStats Returns a two-dimensional array with the detailed
// usage statistics shown under Detailed Statistics in KiwiVM.
func (c *Client) GetRawUsageStats(ctx context.Context) (*GetRawUsageStatsRsp, error) {
	return Get[GetRawUsageStatsRsp](ctx, c, "/v1/getRawUsageStats")
}
