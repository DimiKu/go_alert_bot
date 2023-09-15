package users

import (
	"github.com/golang/mock/gomock"
	"go_alert_bot/internal/db_actions"
	"go_alert_bot/internal/service/dto"
	"reflect"
	"testing"
)

func TestUserService_CreateUser(t *testing.T) {
	testUser := dto.UserDto{
		UserId: 32,
		ChatId: 123123,
	}

	testUserDB := db_actions.UserDb{
		UserID: testUser.UserId,
		ChatId: testUser.ChatId,
	}

	var wantErr bool = false

	mockCtrl := gomock.NewController(t)
	mockUserRepo := NewMockUserRepo(mockCtrl)
	defer mockCtrl.Finish()

	mockUserRepo.EXPECT().CheckIfExistUser(testUserDB).Return(false).AnyTimes()
	mockUserRepo.EXPECT().CreateUser(testUserDB).Return(nil).AnyTimes()

	testUserService := NewUserService(mockUserRepo)

	sutUserID, err := testUserService.CreateUser(testUser)
	if (err != nil) != wantErr {
		t.Errorf("CreateUser() error = %v, wantErr %v", err, wantErr)
		return
	}
	if !reflect.DeepEqual(sutUserID, testUser.UserId) {
		t.Errorf("CreateUser() got = %v, want %v", sutUserID, testUser.UserId)
	}
}
