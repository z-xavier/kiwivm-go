package kiwivm

import (
	"context"
)

type ResetRootPasswordRsp struct {
	Error int `json:"error"`
}

// ResetRootPassword Generates and sets a new root password.
// TODO: Need Test
func (c *Client) ResetRootPassword(ctx context.Context) (*ResetRootPasswordRsp, error) {
	return Get[ResetRootPasswordRsp](ctx, c, "/v1/resetRootPassword")
}

type AuditLog struct {
	Timestamp     int64  `json:"timestamp"`
	RequestorIpv4 int64  `json:"requestor_ipv4"`
	Type          int    `json:"type"`
	Summary       string `json:"summary"`
}

type GetAuditLogRsp struct {
	LogEntries []AuditLog `json:"log_entries"`
	Error      int        `json:"error"`
}

// GetAuditLog Returns an array with the detailed audit
// log shown under Audit Log in KiwiVM.
func (c *Client) GetAuditLog(ctx context.Context) (*GetAuditLogRsp, error) {
	return Get[GetAuditLogRsp](ctx, c, "/v1/getAuditLog")
}

type SetHostnameReq struct {
	NewHostname string `json:"newHostname"`
}

type SetHostnameRsp struct {
	Error int `json:"error"`
}

// SetHostname Sets new hostname.
func (c *Client) SetHostname(ctx context.Context, req *SetHostnameReq) (*SetHostnameRsp, error) {
	return GetWithQueryParams[SetHostnameRsp](ctx, c, "/v1/setHostname", req)
}

type SetPTRReq struct {
	IP  string `json:"ip"`
	PTR string `json:"ptr"`
}

type SetPTRRsp struct {
	Error int `json:"error"`
}

// SetPTR Sets new PTR (rDNS) record for IP.
func (c *Client) SetPTR(ctx context.Context, req *SetPTRReq) (*SetPTRRsp, error) {
	return GetWithQueryParams[SetPTRRsp](ctx, c, "/v1/setPTR", req)
}

type CloneFromExternalServerReq struct {
	ExternalServerIP           string `json:"externalServerIP"`
	ExternalServerSSHPort      string `json:"externalServerSSHport"`
	ExternalServerRootPassword string `json:"externalServerRootPassword"`
}

type CloneFromExternalServerRsp struct {
	Error int `json:"error"`
}

// CloneFromExternalServer (OVZ only) Clone a remote server or VPS.
// See Migrate from another server for example on how this works.
// TODO: Need Test
func (c *Client) CloneFromExternalServer(ctx context.Context, req *CloneFromExternalServerReq) (*CloneFromExternalServerRsp, error) {
	return GetWithQueryParams[CloneFromExternalServerRsp](ctx, c, "/v1/cloneFromExternalServer", req)
}

type IsSoft int8

const (
	// IsSoftContact must contact support to unsuspend
	IsSoftContact IsSoft = iota
	// IsSoftAPI can unsuspend via API call
	IsSoftAPI
)

type Suspensions struct {
	// Case ID, needed to unsuspend the service via "unsuspend" API call
	RecordId int `json:"record_id"`
	// Type of abuse
	Flag   string `json:"flag"`
	IsSoft IsSoft `json:"is_soft"`
	// Detailed abuse report ID (see below)
	EvidenceRecordId string `json:"evidence_record_id"`
	// Each abuse incident increases total_abuse_points counter
	AbusePoints int `json:"abuse_points"`
}

type GetSuspensionDetailsRsp struct {
	// Number of times service was suspended in current calendar year
	SuspensionCount int `json:"suspension_count"`
	// Total abuse points accumulated in current calendar year
	TotalAbusePoints int `json:"total_abuse_points"`
	// Maximum abuse points allowed by plan in a calendar year
	MaxAbusePoints int `json:"max_abuse_points"`
	// array of all outstanding issues along with supporing evidence of abuse. See example below.
	Suspensions []*Suspensions `json:"suspensions"`
	// Full text of the complaint or more details about the issue
	Evidence map[int64]string `json:"evidence"`
	Error    int              `json:"error"`
}

// GetSuspensionDetails Retrieve information related to service suspensions.
func (c *Client) GetSuspensionDetails(ctx context.Context) (*GetSuspensionDetailsRsp, error) {
	return Get[GetSuspensionDetailsRsp](ctx, c, "/v1/getSuspensionDetails")
}

type PolicyViolations struct {
	// Case ID, for resolvePolicyViolation
	RecordId int `json:"record_id"`
	// Unix timestamp when record was created
	Timestamp int64 `json:"timestamp"`
	// Service will be suspended if not resolved by this time
	SuspendAt int64 `json:"suspend_at"`
	// Type of abuse
	Flag   string `json:"flag"`
	IsSoft IsSoft `json:"is_soft"`
	// Each abuse incident increases total_abuse_points counter
	AbusePoints int `json:"abuse_points"`
	// Details of violation (text)
	EvidenceData string `json:"evidence_data"`
}

type GetPolicyViolationsRsp struct {
	// Total abuse points accumulated in current calendar year
	TotalAbusePoints int `json:"total_abuse_points"`
	// Maximum abuse points allowed by plan in a calendar year
	MaxAbusePoints int `json:"max_abuse_points"`
	// array of all outstanding issues along with supporing evidence of abuse. See example below.
	PolicyViolations []*PolicyViolations `json:"policy_violations"`
	Error            int                 `json:"error"`
}

// GetPolicyViolations Retrieve information related to active policy violations.
func (c *Client) GetPolicyViolations(ctx context.Context) (*GetPolicyViolationsRsp, error) {
	return Get[GetPolicyViolationsRsp](ctx, c, "/v1/getPolicyViolations")
}

type UnsuspendReq struct {
	RecordID string `url:"record_id"`
}

type UnsuspendRsp struct {
	Error int `json:"error"`
}

// Unsuspend Clear abuse issue identified by record_id and unsuspend the VPS.
// Refer to getSuspensionDetails call for details.
// TODO: Need Test
func (c *Client) Unsuspend(ctx context.Context, req *UnsuspendReq) (*UnsuspendRsp, error) {
	return GetWithQueryParams[UnsuspendRsp](ctx, c, "/v1/unsuspend", req)
}

type ResolvePolicyViolationReq struct {
	RecordID string `url:"record_id"`
}

type ResolvePolicyViolationRsp struct {
	Error int `json:"error"`
}

// ResolvePolicyViolation Mark policy violation as resolved.
// This is required to avoid service suspension.
// Refer to getPolicyViolations call for details.
// TODO: Need Test
func (c *Client) ResolvePolicyViolation(ctx context.Context, req *ResolvePolicyViolationReq) (*ResolvePolicyViolationRsp, error) {
	return GetWithQueryParams[ResolvePolicyViolationRsp](ctx, c, "/v1/resolvePolicyViolation", req)
}

type GetRateLimitStatusRsp struct {
	// Number of "points" available to use in the current 15-minute interval
	RemainingPoints15Min int `json:"remaining_points_15min"`
	// Number of "points" available to use in the current 24-hour interval
	RemainingPoints24H int `json:"remaining_points_24h"`
	Error              int `json:"error"`
}

// GetRateLimitStatus When you perform too many API calls in a short amount of time,
// KiwiVM API may start dropping your requests for a few minutes.
// This call allows monitoring this matter.
func (c *Client) GetRateLimitStatus(ctx context.Context) (*GetRateLimitStatusRsp, error) {
	return Get[GetRateLimitStatusRsp](ctx, c, "/v1/getRateLimitStatus")
}
