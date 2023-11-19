package dto

import "testing"

func TestUserDto_Validate(t *testing.T) {
	type fields struct {
		UserId int
		ChatId int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "FailedEmptyUserID",
			fields: fields{
				UserId: 0,
				ChatId: 234234234234,
			},
			wantErr: true,
		},
		{
			name: "FailedEmptyChatID",
			fields: fields{
				UserId: 3,
				ChatId: 0,
			},
			wantErr: true,
		},
		{
			name: "PassedUserTest",
			fields: fields{
				UserId: 3,
				ChatId: 234234234234,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserDto{
				UserId: tt.fields.UserId,
				ChatId: tt.fields.ChatId,
			}
			if err := u.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
