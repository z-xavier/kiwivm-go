package kiwivm

import "context"

type UpdateSSHKeysReq struct {
	SSHKeys string `json:"ssh_keys"`
}

type UpdateSSHKeysRsp struct {
	Error int `json:"error"`
}

// UpdateSSHKeys Update per-VM SSH keys in Hypervisor Vault.
// Keys will be written to /root/.ssh/authorized_keys during a reinstallOS call.
// These keys will override any keys set in Billing Portal.
// TODO: Need Test
func (c *Client) UpdateSSHKeys(ctx context.Context, req *UpdateSSHKeysReq) (*UpdateSSHKeysRsp, error) {
	return GetWithQueryParams[UpdateSSHKeysRsp](ctx, c, "/v1/updateSshKeys", req)
}

type GetSSHKeysRsp struct {
	// Per-VM SSH Keys stored in Hypervisor Vault
	SSHKeysVeID string `json:"ssh_keys_veid"`
	// Per-Account SSH keys stored in the Billing Portal
	SshKeysUser string `json:"ssh_keys_user"`
	// SSH Keys, which will be actually used during a reinstallOS call
	// (Per-VM Keys will always override Per-Account keys)
	SSHKeysPreferred string `json:"ssh_keys_preferred"`
	// Visually shortened keys
	ShortenedSshKeysVeID string `json:"shortened_ssh_keys_veid"`
	// Visually shortened keys
	ShortenedSshKeysUser string `json:"shortened_ssh_keys_user"`
	// Visually shortened keys
	ShortenedSshKeysPreferred string `json:"shortened_ssh_keys_preferred"`
	Error                     int    `json:"error"`
}

// GetSSHKeys Get SSH keys stored in Hypervisor Vault,
// as well as the ones stored in the Billing Portal.
func (c *Client) GetSSHKeys(ctx context.Context) (*GetSSHKeysRsp, error) {
	return Get[GetSSHKeysRsp](ctx, c, "/v1/getSshKeys")
}
