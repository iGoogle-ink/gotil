package orm

import (
	"testing"
	"time"

	"github.com/iGoogle-ink/gotil/xlog"
	"github.com/iGoogle-ink/gotil/xtime"
)

var (
	dsn = "root:root@tcp(mysql:3306)/school?parseTime=true&loc=Local&charset=utf8mb4"
)

type FmUser struct {
	Id     int    `gorm:"column:id;primaryKey" xorm:"'id' pk"`
	UName  string `gorm:"column:uname" xorm:"'uname'"`
	Passwd string `gorm:"column:passwd"`
	Openid string `gorm:"column:openid"`
}

func (m *FmUser) TableName() string {
	return "fm_user"
}

func TestInitGormV2(t *testing.T) {
	// 初始化 Gorm
	gc1 := &MySQLConfig{
		DSN:            dsn,
		MaxOpenConn:    10,
		MaxIdleConn:    10,
		MaxConnTimeout: xtime.Duration(10 * time.Second),
	}

	g := InitGormV2(gc1)
	u := &FmUser{
		UName: "jerry",
	}
	// create
	err := g.Create(u).Error
	if err != nil {
		xlog.Error(err)
		return
	}
	var uQs []*FmUser
	// query
	err = g.Table(u.TableName()).Where("uname = ?", "jerry").Find(&uQs).Error
	if err != nil {
		xlog.Error(err)
		return
	}
	for _, v := range uQs {
		xlog.Debugf("%+v", v)
	}
}
