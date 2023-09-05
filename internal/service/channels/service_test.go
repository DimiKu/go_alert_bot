package channels

import (
	"github.com/golang/mock/gomock"
	"testing"

	"go_alert_bot/internal"
)

func TestChannelService_CreateChannel(t *testing.T) {
	type fields struct {
		storage ChannelRepo
	}
	type args struct {
		channel internal.ChannelDto
	}

	newStorage := fields{storage: NewMockChannelRepo(gomock.NewController(t))}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    internal.ChannelLinkDto
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "first test",
			fields:  newStorage,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chs := &ChannelService{
				storage: tt.fields.storage,
			}
			got, err := chs.CreateChannel(tt.args.channel)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateChannel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateChannel() got = %v, want %v", got, tt.want)
			}
		})
	}
}
