package service

import (
	"fmt"
	"github.com/shiningrush/demo-service/internal/domain/entity"
	"github.com/shiningrush/demo-service/internal/domain/repo/account"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestAccountService_NewAccount(t *testing.T) {
	tests := []struct {
		caseDesc     string          // 用例简述
		givePhone    string          // 执行用例的输入电话
		giveAcc      *entity.Account // mock的仓储中返回的账户
		giveErr      error           // mock的仓储函数返回的错误
		wantErr      error           // 执行后期望的错误
		wantAcc      *entity.Account // 执行后 期望获得的新增账户
		wantRepoCall bool            // 用于判断仓储的指定函数是否得到调用
	}{
		{
			caseDesc:  "正常用例",
			givePhone: "right",
			wantAcc: &entity.Account{
				Phone:  "right",
				Status: entity.AccountStatusNormal,
			},
			wantRepoCall: true,
		},
		{
			caseDesc:     "手机号重复",
			givePhone:    "repeat",
			giveAcc:      &entity.Account{},
			wantErr:      fmt.Errorf("phone alreay be used"),
			wantRepoCall: true,
		},
		{
			caseDesc:     "调用 repo 错误",
			givePhone:    "right",
			giveErr:      fmt.Errorf("repo error"),
			wantErr:      fmt.Errorf("repo error"),
			wantRepoCall: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.caseDesc, func(t *testing.T) {
			isCalled := false
			mockRepo := &account.MockRepo{}
			// 对 GetWithPhone 进行打桩
			mockRepo.On("GetWithPhone", mock.Anything).Run(func(args mock.Arguments) {
				// 当桩函数被调用时，我们做一个标记
				isCalled = true
			}).Return(tc.giveAcc, tc.giveErr)

			svc := accountService{
				accRepo: mockRepo,
			}
			acc, err := svc.NewAccount(tc.givePhone)

			// 检查期望的结果
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantAcc, acc)
			assert.Equal(t, tc.wantRepoCall, isCalled)
		})
	}
}

func TestAccountService_ChangePhone(t *testing.T) {
	tests := []struct {
		caseDesc             string          // 用例简述
		giveAccId            int             // 要修改的账户ID
		givePhone            string          // 新的电话号码
		giveTargetAcc        *entity.Account // mock的仓储中 Get 返回的账户
		giveAcc              *entity.Account // mock的仓储中 GetWithPhone 返回的账户
		giveGetErr           error           // mock仓储的 Get 函数返回错误
		giveGetWithPhoneErr  error           // mock仓储的 GetWithPhone 函数返回错误
		giveUpdateErr        error           // mock仓储的 Update 函数返回错误
		wantErr              error           // 执行后期望的错误
		wantGetCall          bool            // 用于判断仓储的指定函数是否得到调用
		wantGetWithPhoneCall bool            // 用于判断仓储的指定函数是否得到调用
		wantUpdateCall       bool            // 用于判断仓储的指定函数是否得到调用
	}{
		{
			caseDesc:             "正常用例",
			giveAccId:            123,
			givePhone:            "right",
			giveTargetAcc:        &entity.Account{},
			wantGetCall:          true,
			wantGetWithPhoneCall: true,
			wantUpdateCall:       true,
		},
		{
			caseDesc:             "手机号重复",
			givePhone:            "repeat",
			giveAcc:              &entity.Account{},
			wantErr:              fmt.Errorf("phone alreay be used"),
			wantGetWithPhoneCall: true,
			wantUpdateCall:       false,
		},
		{
			caseDesc:             "要修改的账户不存在",
			giveAccId:            123,
			givePhone:            "right",
			wantErr:              fmt.Errorf("account id: 123 not found"),
			wantGetWithPhoneCall: true,
			wantUpdateCall:       false,
		},
		{
			caseDesc:             "调用仓储 getWithPhone 错误",
			givePhone:            "right",
			giveGetWithPhoneErr:  fmt.Errorf("giveGetWithPhoneErr error"),
			wantErr:              fmt.Errorf("giveGetWithPhoneErr error"),
			wantGetWithPhoneCall: true,
			wantUpdateCall:       false,
		},
		{
			caseDesc:             "调用仓储 get 错误",
			givePhone:            "right",
			giveGetErr:           fmt.Errorf("get error"),
			wantErr:              fmt.Errorf("get error"),
			wantGetWithPhoneCall: true,
			wantGetCall:          true,
			wantUpdateCall:       false,
		},
		{
			caseDesc:             "调用仓储 update 错误",
			givePhone:            "right",
			giveTargetAcc:        &entity.Account{},
			giveGetErr:           fmt.Errorf("update error"),
			wantErr:              fmt.Errorf("update error"),
			wantGetWithPhoneCall: true,
			wantGetCall:          true,
			wantUpdateCall:       false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.caseDesc, func(t *testing.T) {
			isUpdateCalled, isGetCalled, isGetWithPhoneCalled := false, false, false
			mockRepo := &account.MockRepo{}
			// 对 Get 进行打桩
			mockRepo.On("Get", mock.Anything).Run(func(args mock.Arguments) {
				// 当桩函数被调用时，我们做一个标记
				isGetCalled = true
				// 确保输入的账户为用例中的账户ID
				assert.Equal(t, tc.giveAccId, args.Int(0))
			}).Return(tc.giveTargetAcc, tc.giveGetErr)

			// 对 GetWithPhone 进行打桩
			mockRepo.On("GetWithPhone", mock.Anything).Run(func(args mock.Arguments) {
				// 当桩函数被调用时，我们做一个标记
				isGetWithPhoneCalled = true
			}).Return(tc.giveAcc, tc.giveGetWithPhoneErr)

			// 对 Update 进行打桩
			mockRepo.On("Update", mock.Anything).Run(func(args mock.Arguments) {
				// 当桩函数被调用时，我们做一个标记
				isUpdateCalled = true
				// 确保输入的账户为用例中的账户ID
				acc := args.Get(0).(entity.Account)
				assert.Equal(t, tc.giveTargetAcc.ID, acc.ID)
				assert.Equal(t, tc.givePhone, acc.Phone)
			}).Return(tc.giveAcc, tc.giveGetErr)

			svc := accountService{
				accRepo: mockRepo,
			}
			err := svc.ChangePhone(tc.giveAccId, tc.givePhone)

			// 检查期望的结果
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantGetCall, isGetCalled)
			assert.Equal(t, tc.wantGetWithPhoneCall, isGetWithPhoneCalled)
			assert.Equal(t, tc.wantUpdateCall, isUpdateCalled)
		})
	}
}

type stubAccountRepo struct {
}

func (r *stubAccountRepo) Get(id int) (*entity.Account, error) {
	return nil, nil
}

func (r *stubAccountRepo) GetWithPhone(phone string) (*entity.Account, error) {
	if phone == "testphone" {
		return &entity.Account{}, nil
	}
	return nil, nil
}

func (r *stubAccountRepo) Create(account entity.Account) (*entity.Account, error) {
	return nil, nil
}

func (r *stubAccountRepo) Update(account entity.Account) (*entity.Account, error) {
	return nil, nil
}

func (r *stubAccountRepo) Delete(id int) error {
	return nil
}
