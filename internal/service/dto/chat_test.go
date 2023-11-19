package dto

import "testing"

func TestChatDto_Validate(t *testing.T) {
	type fields struct {
		UserId       int
		TgChatId     string
		ChatType     string
		FormatString string
		ChannelLink  int64
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "FailedChatTypeTest",
			fields: fields{
				UserId:       1,
				TgChatId:     "1134134234134",
				ChatType:     "MyTeam",
				FormatString: "test",
				ChannelLink:  0,
			},
			wantErr: true,
		},
		{
			name: "FailedEmptyUserTest",
			fields: fields{
				UserId:       0,
				TgChatId:     "1134134234134",
				ChatType:     "telegram",
				FormatString: "test",
				ChannelLink:  0,
			},
			wantErr: true,
		},
		{
			name: "FailedEmptyTelegramChat",
			fields: fields{
				UserId:       3,
				TgChatId:     "",
				ChatType:     "telegram",
				FormatString: "test",
				ChannelLink:  0,
			},
			wantErr: true,
		},
		{
			name: "PassedStdoutEmptyTgChatIdTest",
			fields: fields{
				UserId:       2,
				TgChatId:     "",
				ChatType:     "stdout",
				FormatString: "test",
				ChannelLink:  0,
			},
			wantErr: false,
		},
		{
			name: "PassedChatTest",
			fields: fields{
				UserId:       1,
				TgChatId:     "-1341342341",
				ChatType:     "telegram",
				FormatString: "test",
				ChannelLink:  0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ChatDto{
				UserId:       tt.fields.UserId,
				TgChatId:     tt.fields.TgChatId,
				ChatType:     tt.fields.ChatType,
				FormatString: tt.fields.FormatString,
				ChannelLink:  tt.fields.ChannelLink,
			}
			if err := c.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
