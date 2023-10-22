package events

import (
	"github.com/golang/mock/gomock"
	"go_alert_bot/internal/db_actions"
	"go_alert_bot/internal/service/dto"
	"testing"
)

func TestEventService_AddEventInChannel(t *testing.T) {
	channelLink := db_actions.ChannelLink(14243423414134134)
	mockStorage := NewMockEventRepo(gomock.NewController(t))

	type fields struct {
		storage         EventRepo
		eventMap        *StorageMap
		eventCounterMap *CounterMap
		SendEventRepos  map[string]SendEventRepo
	}
	type args struct {
		event          dto.EventDto
		channelLinkDto dto.ChannelLinkDto
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "add_event_test",
			fields: fields{
				storage:         mockStorage,
				eventMap:        NewStorageMap(),
				eventCounterMap: NewCounterMap(),
				SendEventRepos: map[string]SendEventRepo{
					"telegram": NewMockSendEventRepo(gomock.NewController(t)),
					"stdout":   NewMockSendEventRepo(gomock.NewController(t)),
				},
			},
			args: args{
				event: dto.EventDto{
					Key:         "key",
					UserId:      1,
					ChannelLink: 14243423414134134,
				},
				channelLinkDto: 14243423414134134,
			},
			want:    "Event added",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockStorage.EXPECT().IsExistChannelByChannelLink(channelLink).Return(true)
			es := &EventService{
				storage:         tt.fields.storage,
				eventMap:        tt.fields.eventMap,
				eventCounterMap: tt.fields.eventCounterMap,
				SendEventRepos:  tt.fields.SendEventRepos,
			}

			got, err := es.AddEventInChannel(tt.args.event, tt.args.channelLinkDto)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddEventInChannel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AddEventInChannel() got = %v, want %v", got, tt.want)
			}
		})
	}
}
