package dto

import "testing"

func TestChannelDto_Validate(t *testing.T) {
	type fields struct {
		UserId       int
		TgChatIds    string
		FormatString string
		ChatType     string
		ChannelLink  ChannelLinkDto
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
				TgChatIds:    "134134143234",
				FormatString: "test",
				ChatType:     "MyTeam",
				ChannelLink:  0,
			},
			wantErr: true,
		},
		{
			name: "FailedEmptyUserTest",
			fields: fields{
				UserId:       0,
				TgChatIds:    "2234234234",
				FormatString: "test",
				ChatType:     "telegram",
				ChannelLink:  0,
			},
			wantErr: true,
		},
		{
			name: "FailedEmptyChannelTypeTest",
			fields: fields{
				UserId:       2,
				TgChatIds:    "234234234",
				FormatString: "test",
				ChatType:     "",
				ChannelLink:  0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ChannelDto{
				UserId:       tt.fields.UserId,
				TgChatIds:    tt.fields.TgChatIds,
				FormatString: tt.fields.FormatString,
				ChatType:     tt.fields.ChatType,
				ChannelLink:  tt.fields.ChannelLink,
			}
			if err := c.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
