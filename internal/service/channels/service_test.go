package channels

import (
	"testing"

	"github.com/golang/mock/gomock"

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

	expectedChannelLinkId := internal.ChannelLinkDto(7969375211542538373)

	agrs := args{channel: internal.ChannelDto{
		UserId:       1,
		TgChatIds:    "1111",
		FormatString: "test format",
		ChatType:     "telegram",
		ChannelLink:  expectedChannelLinkId,
	}}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    internal.ChannelLinkDto
		wantErr bool
	}{
		// TODO: Add test cases.
		// TODO переписать тесты с вовзратом канала
		{
			name:    "first test",
			fields:  newStorage,
			args:    agrs,
			want:    expectedChannelLinkId,
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
