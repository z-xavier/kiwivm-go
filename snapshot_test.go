package kiwivm

import (
	"context"
	"testing"
)

func TestClient_SnapshotList(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "SnapshotList",
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDefaultTestClient().SnapshotList(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("SnapshotList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(MarshalString(got))
		})
	}
}
