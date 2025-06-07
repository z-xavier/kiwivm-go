package kiwivm

import (
	"context"
	"testing"
)

func TestClient_GetAuditLog(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GetAuditLog",
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDefaultTestClient().GetAuditLog(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAuditLog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(MarshalString(got))
			for _, entry := range got.LogEntries {
				t.Logf("IPV4: %s", IntToNetIP4Str(entry.RequestorIpv4))
			}
		})
	}
}

func TestClient_GetSuspensionDetails(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GetSuspensionDetails",
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDefaultTestClient().GetSuspensionDetails(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSuspensionDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(MarshalString(got))
		})
	}
}

func TestClient_GetPolicyViolations(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GetPolicyViolations",
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDefaultTestClient().GetPolicyViolations(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPolicyViolations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(MarshalString(got))
		})
	}
}

func TestClient_GetRateLimitStatus(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GetRateLimitStatus",
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDefaultTestClient().GetRateLimitStatus(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRateLimitStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(MarshalString(got))
		})
	}
}
