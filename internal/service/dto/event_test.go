package dto

import "testing"

func TestEventDto_Validate(t *testing.T) {
	type fields struct {
		Key         string
		UserId      int
		ChannelLink ChannelLinkDto
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "FailedEmptyEvent",
			fields: fields{
				Key: "",
			},
			wantErr: true,
		},
		{
			name: "PassedEvent",
			fields: fields{
				Key: "error",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &EventDto{
				Key:         tt.fields.Key,
				UserId:      tt.fields.UserId,
				ChannelLink: tt.fields.ChannelLink,
			}
			if err := e.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
