// +build integration

package repo

import (
	"github.com/shiningrush/demo-service/internal/domain/entity"
	"github.com/shiningrush/demo-service/internal/utils/dao"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAccountRepo(t *testing.T) {
	err := dao.InitMySQL("root:root@tcp(127.0.0.1:3306)/demo_service?charset=utf8&parseTime=True&loc=Local")
	assert.NoError(t, err)

	repo := NewAccountRepo(dao.DefaultDB)

	// 测试创建
	newAccount := entity.NewAccount("123456")
	reAcc, err := repo.Create(newAccount)
	assert.NoError(t, err)
	assert.NotEqual(t, 0, reAcc.ID)

	getAcc, err := repo.Get(reAcc.ID)
	assert.NoError(t, err)

	// 由于数据库的精度不同，所以需要将时间调整到秒级
	reAcc.CreatedAt = reAcc.CreatedAt.Round(time.Second)
	reAcc.UpdatedAt = reAcc.UpdatedAt.Round(time.Second)
	assert.Equal(t, reAcc, getAcc)

	// 测试更新
	getAcc.Balance = 10
	_, err = repo.Update(*getAcc)
	assert.NoError(t, err)
	updatedAcc, err := repo.Get(reAcc.ID)
	assert.NoError(t, err)
	assert.Equal(t, updatedAcc, getAcc)

	// 测试删除
	err = repo.Delete(getAcc.ID)
	assert.NoError(t, err)
	_, err = repo.Get(reAcc.ID)
	assert.Equal(t, "execute query failed: record not found", err.Error())

	//  关闭连接，结束测试
	err = dao.CloseMySQL()
	assert.NoError(t, err)
}
