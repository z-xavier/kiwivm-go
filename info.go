package kiwivm

import (
	"context"
	"encoding/json"
)

type GetServiceInfoRsp struct {
	// Hypervizor type (ovz or kvm)
	VmType string `json:"vm_type"`
	// Hostname of the VPS
	Hostname string `json:"hostname"`
	// IP address of the physical node
	NodeIp string `json:"node_ip"`
	// Internal nickname of the physical node
	NodeAlias string `json:"node_alias"`
	// Physical location (country, state)
	NodeLocation string `json:"node_location"`
	// Whether IPv6 is supported at the current location
	LocationIpv6Ready bool `json:"location_ipv6_ready"`
	// Name of plan
	Plan string `json:"plan"`
	// Disk quota (bytes)
	PlanDisk int64 `json:"plan_disk"`
	// RAM (bytes)
	PlanRam int `json:"plan_ram"`
	// SWAP (bytes)
	PlanSwap int `json:"plan_swap"`
	// Operating system
	OS string `json:"os"`
	// Primary e-mail address of the account
	Email string `json:"email"`
	// Allowed monthly data transfer (bytes).
	// Needs to be multiplied by monthly_data_multiplier - see below.
	PlanMonthlyData int64 `json:"plan_monthly_data"`
	// Data transfer used in the current billing month.
	// Needs to be multiplied by monthly_data_multiplier - see below.
	DataCounter int64 `json:"data_counter"`
	// Some locations offer more expensive bandwidth;
	// this variable contains the bandwidth accounting coefficient.
	MonthlyDataMultiplier int `json:"monthly_data_multiplier"`
	// Date and time of transfer counter reset (UNIX timestamp)
	DataNextReset int64 `json:"data_next_reset"`
	// IPv4 and IPv6 addresses assigned to VPS (Array)
	IpAddresses []string `json:"ip_addresses"`
	// Private IPv4 addresses assigned to VPS (Array)
	PrivateIpAddresses []string `json:"private_ip_addresses"`
	// Information on IP address nullrouting during (D)DoS attacks (Array).
	// Sample output when IP is under attack.
	IpNullRoutes []string `json:"ip_nullroutes"`
	// Mounted image #1
	Iso1 string `json:"iso1"`
	// Mounted image #2 (currently unsupported)
	Iso2 string `json:"iso2"`
	// Array of ISO images available for use
	AvailableISOs []string `json:"available_isos"`
	// Maximum number of IPv6 addresses allowed by plan
	PlanMaxIPV6s int `json:"plan_max_ipv6s"`
	// Whether or not rDNS records can be set via API
	RDNSAPIAvailable bool `json:"rdns_api_available"`
	// Whether or not Private Network features are available on this plan
	PlanPrivateNetworkAvailable bool `json:"plan_private_network_available"`
	// Whether or not Private Network features are available at this location
	LocationPrivateNetworkAvailable bool `json:"location_private_network_available"`
	// rDNS records (Array of two-dimensional arrays: ip=>value)
	Ptr map[string]string `json:"ptr"`
	// Whether VPS is suspended
	Suspended bool `json:"suspended"`
	// Whether there is an active policy violation that needs attention (see getPolicyViolations)
	PolicyViolation bool `json:"policy_violation"`
	// Number of times service was suspended in current calendar year
	SuspensionCount int `json:"suspension_count"`
	// Total abuse points accumulated in current calendar year
	TotalAbusePoints int `json:"total_abuse_points"`
	// Maximum abuse points allowed by plan in a calendar year
	MaxAbusePoints int    `json:"max_abuse_points"`
	NodeLocationId string `json:"node_location_id"`
	NodeDatacenter string `json:"node_datacenter"`
	Error          int    `json:"error"`
}

func (c *Client) GetServiceInfo(ctx context.Context) (*GetServiceInfoRsp, error) {
	return Get[GetServiceInfoRsp](ctx, c, "/v1/getServiceInfo")
}

type GetLiveServiceInfoRsp struct {
	GetServiceInfoRsp

	// Depending on hypervisor this call will return the following information:
	//
	// [OVZ hypervisor]
	// array containing OpenVZ beancounters, system load average,
	// number of processes, open files, sockets, memory usage etc
	VzStatus json.RawMessage `json:"vz_status"`
	// array containing OpenVZ disk size, inodes and usage info
	VzQuota json.RawMessage `json:"vz_quota"`

	// [KVM hypervisor]
	// Starting, Running or Stopped
	VeStatus string `json:"ve_status"`
	// MAC address of primary network interface
	VeMac1 string `json:"ve_mac1"`
	// Occupied (mapped) disk space in bytes
	VeUsedDiskSpaceB int64 `json:"ve_used_disk_space_b"`
	// Actual size of disk image in GB
	VeDiskQuotaGb string `json:"ve_disk_quota_gb"`
	// 0 = Disk I/O is not throttled, 1 = Disk I/O is throttled due to high usage.
	// Throttling resets automatically every 15-180 minutes depending on sustained
	// storage I/O utilization.
	IsDiskThrottled string `json:"is_disk_throttled"`
	// Result of "hostname" command executed inside VPS
	LiveHostname string `json:"live_hostname"`
	// Raw load average string
	LoadAverage string `json:"load_average"`
	// Amount of available RAM in KB
	MemAvailableKB int `json:"mem_available_kb"`
	// Total amount of Swap in KB
	SwapTotalKB int `json:"swap_total_kb"`
	// Amount of available Swap in KB
	SwapAvailableKB int `json:"swap_available_kb"`
	// base64 encoded png screenshot of the VGA console
	ScreenDumpPngBase64 string `json:"screendump_png_base64"`
	VeID                int    `json:"veid"`

	// both
	// 0 = CPU is not throttled, 1 = CPU is throttled due to high usage.
	// Throttling resets automatically every 2 hours.
	IsCpuThrottled string `json:"is_cpu_throttled"`
	// SSH port of the VPS (returned only if VPS is running)
	SSHPort int `json:"ssh_port"`
	Error   int `json:"error"`
}

// GetLiveServiceInfo This function returns all data provided by getServiceInfo.
// In addition, it provides detailed status of the VPS.
// Please note that this call may take up to 15 seconds to complete.
func (c *Client) GetLiveServiceInfo(ctx context.Context) (*GetLiveServiceInfoRsp, error) {
	return Get[GetLiveServiceInfoRsp](ctx, c, "/v1/getLiveServiceInfo")
}
