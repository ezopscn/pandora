package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"pandora/common"
	"time"
)

// MySQL 连接初始化
func MySQL() {
	// 数据库连接串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&collation=%s&timeout=%dms&%s",
		common.Config.Common.MySQL.Username,
		common.Config.Common.MySQL.Password,
		common.Config.Common.MySQL.Host,
		common.Config.Common.MySQL.Port,
		common.Config.Common.MySQL.Database,
		common.Config.Common.MySQL.Charset,
		common.Config.Common.MySQL.Collation,
		common.Config.Common.MySQL.Timeout,
		common.Config.Common.MySQL.ExtraParam)

	// 连接数据库
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn, // 数据库连接字符串
		DefaultStringSize: 170, // varchar 默认长度，太长影响查询
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 单数表名
			// TablePrefix:   "tb_", // 表名前缀
		},
		DisableForeignKeyConstraintWhenMigrating: true,  // 禁用外键
		IgnoreRelationshipsWhenMigrating:         false, // 开启会导致 many2many 的表创建失败
		QueryFields:                              true,  // 解决查询索引失效问题
	})

	// 错误处理
	if err != nil {
		common.SystemLog.Error(err.Error())
		panic("MySQL 连接初始化失败")
	}

	// 设置数据库连接池
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(common.Config.Common.MySQL.MaxOpenConns)
	sqlDB.SetMaxIdleConns(common.Config.Common.MySQL.MaxIdleConns)
	sqlDB.SetConnMaxIdleTime(time.Duration(common.Config.Common.MySQL.MaxIdleTime) * time.Minute)

	// 设置全局数据库连接，方便后续使用
	common.MySQLDB = db
	common.SystemLog.Info("MySQL 连接初始完化成：", fmt.Sprintf("%s:%d/%s",
		common.Config.Common.MySQL.Host,
		common.Config.Common.MySQL.Port,
		common.Config.Common.MySQL.Database))
}
