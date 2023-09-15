package channels

import (
	"github.com/golang/mock/gomock"
	"go_alert_bot/internal/db_actions"
	"go_alert_bot/internal/service/dto"
	"go_alert_bot/pkg"
	"reflect"
	"testing"
)

func TestChannelService_CreateChannel(t *testing.T) {
	type fields struct {
		storage ChannelRepo
	}
	type args struct {
		channel dto.ChannelDto
	}

	testChannelDto := dto.ChannelDto{
		UserId:       1,
		TgChatIds:    "-1001654890472",
		FormatString: "its test",
		ChatType:     "telegram",
		ChannelLink:  8706556132975901137,
	}

	mockCtrl := gomock.NewController(t)
	mockChannRepo := NewMockChannelRepo(mockCtrl)

	tests := []struct {
		name    string
		prepare func(f *fields)
		fields  fields
		args    args
		want    *dto.ChannelDto
		wantErr bool
	}{
		{
			name: "fist_test",
			prepare: func(f *fields) {
				mockChannRepo.EXPECT().IsExistChannel(testChannelDto).Return(true).AnyTimes()
			},
			fields:  fields{storage: mockChannRepo},
			args:    args{channel: testChannelDto},
			want:    &testChannelDto,
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateChannel() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChannelService_CreateChannel2(t *testing.T) {

	testChannelDto := dto.ChannelDto{
		UserId:       1,
		TgChatIds:    "-1001654890472",
		FormatString: "its test",
		ChatType:     "telegram",
		ChannelLink:  8706556132975901137,
	}

	tgIds, err := pkg.ConvertStrToInt64Slice(testChannelDto.TgChatIds)
	if err != nil {
		return
	}

	testChannelDb := db_actions.ChannelDb{
		UserId:       testChannelDto.UserId,
		TgChatIds:    tgIds,
		ChannelLink:  db_actions.ChannelLink(testChannelDto.ChannelLink),
		ChannelType:  testChannelDto.ChatType,
		FormatString: testChannelDto.FormatString,
	}

	var wantErr bool = false
	want := &testChannelDto

	mockCtrl := gomock.NewController(t)
	mockChannRepo := NewMockChannelRepo(mockCtrl)
	defer mockCtrl.Finish()

	chs := &ChannelService{
		storage: mockChannRepo,
	}

	mockChannRepo.EXPECT().IsExistChannel(testChannelDb).Return(false).AnyTimes()
	mockChannRepo.EXPECT().CreateTelegramChannel(testChannelDb).Return(nil).AnyTimes()

	sut, err := chs.CreateChannel(testChannelDto)
	if (err != nil) != wantErr {
		t.Errorf("CreateChannel() error = %v, wantErr %v", err, wantErr)
		return
	}
	if !reflect.DeepEqual(sut, want) {
		t.Errorf("CreateChannel() got = %v, want %v", sut, want)
	}

}
