package core

import (
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sys/internal/config"
	"time"
)

func InitDB(c *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(c.DSN()), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,               //不生成外键约束
		SkipDefaultTransaction:                   true,               // 关闭自动事务包装，需手动管理，提升性能
		PrepareStmt:                              true,               //启用 SQL 预编译，复用执行计划，提高查询效率
		Logger:                                   settingLogConfig(), //使用配置的日志记录器输出 SQL 日志
	})
	if err != nil {
		logx.Errorf("数据库连接失败：%s", err.Error())
	}
	sqlDB, err := db.DB()
	if err != nil {
		logx.Errorf("获取数据库实例失败：%s", err.Error())
		return db, err
	}
	sqlDB.SetMaxIdleConns(10)  //最大空闲连接数
	sqlDB.SetMaxOpenConns(100) //最大连接数
	sqlDB.SetConnMaxIdleTime(time.Hour)
	logx.Infof("数据库连接成功")
	return db, nil
}

type Writer struct {
}

func (w Writer) Printf(format string, args ...interface{}) {
	logx.Infof(format, args...)
}

// init log config
func settingLogConfig() logger.Interface {
	newLogger := logger.New(
		Writer{},
		logger.Config{
			SlowThreshold:             200 * time.Millisecond, // Slow SQL threshold
			LogLevel:                  logger.Info,            // Log level
			IgnoreRecordNotFoundError: true,                   // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,                   // Disable color
		},
	)
	return newLogger
}
