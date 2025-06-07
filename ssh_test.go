package kiwivm

import (
	"context"
	"testing"
)

func TestClient_GetSSHKeys(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GetSSHKeys",
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDefaultTestClient().GetSSHKeys(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSSHKeys() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(MarshalString(got))
		})
	}
}
