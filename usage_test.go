package kiwivm

import (
	"context"
	"testing"
	"time"
)

func TestClient_GetRawUsageStats(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GetRawUsageStats",
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDefaultTestClient().GetRawUsageStats(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRawUsageStats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(MarshalString(got))
		})
	}
}

func TestGetRawUsageStatsRsp_GetData(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GetData",
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDefaultTestClient().GetRawUsageStats(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRawUsageStats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(MarshalString(got.GetData(time.Date(2025, 6, 1, 0, 0, 0, 0, time.Local),
				time.Date(2025, 6, 9, 0, 0, 0, 0, time.Local))))
		})
	}
}
