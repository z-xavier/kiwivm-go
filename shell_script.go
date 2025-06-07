package kiwivm

import "context"

type ShellScriptExecReq struct {
	Script string `json:"script"`
}

type ShellScriptExecRsp struct {
	NodeIP string `json:"node_ip"`
	// Name of the output log file.
	Log            string `json:"log"`
	OutputStreamID string `json:"output_stream_id"`
	Error          int    `json:"error"`
}

// ShellScriptExec Execute a shell script on the VPS (asynchronously).
func (c *Client) ShellScriptExec(ctx context.Context, req *ShellScriptExecReq) (*ShellScriptExecRsp, error) {
	return GetWithQueryParams[ShellScriptExecRsp](ctx, c, "/v1/shellScript/exec", req)
}
