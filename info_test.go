package kiwivm

import (
	"context"
	"testing"
)

func TestClient_GetServiceInfo(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GetServiceInfo",
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDefaultTestClient().GetServiceInfo(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetServiceInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(MarshalString(got))
		})
	}
}

func TestClient_GetLiveServiceInfo(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GetLiveServiceInfo",
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDefaultTestClient().GetLiveServiceInfo(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLiveServiceInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(MarshalString(got))
		})
	}
}
