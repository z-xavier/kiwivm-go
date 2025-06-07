package kiwivm

import "context"

type BasicShellCDReq struct {
	CurrentDir string `json:"currentDir"`
	NewDir     string `json:"newDir"`
}

type BasicShellCDRsp struct {
	// Result of the "pwd" command after the change.
	PWD   string `json:"pwd"`
	Error int    `json:"error"`
}

// BasicShellCD Simulate change of directory inside of the VPS.
// Can be used to build a shell like Basic shell.
func (c *Client) BasicShellCD(ctx context.Context, req *BasicShellCDReq) (*BasicShellCDRsp, error) {
	return GetWithQueryParams[BasicShellCDRsp](ctx, c, "/v1/basicShell/cd", req)
}

type BasicShellExecReq struct {
	Command string `json:"command"`
}

type BasicShellExecRsp struct {
	// Exit status code of the executed command
	Error int `json:"error"`
	// Console output of the executed command
	Message string `json:"message"`
}

// BasicShellExec Execute a shell command on the VPS (synchronously).
func (c *Client) BasicShellExec(ctx context.Context, req *BasicShellExecReq) (*BasicShellExecRsp, error) {
	return GetWithQueryParams[BasicShellExecRsp](ctx, c, "/v1/basicShell/exec", req)
}
