package db

import (
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/shiguanghuxian/micro-common/config"
	"github.com/shiguanghuxian/micro-common/etcdcli"
	"github.com/shiguanghuxian/micro-common/log"
)

var (
	// MasterDB db主节点
	MasterDB *gorm.DB
	// SlaveDB db从节点
	SlaveDB *gorm.DB
)

// GetMasterDB 获取db主节点连接对象
func GetMasterDB() *gorm.DB {
	if MasterDB == nil {
		initDB()
	}
	return MasterDB
}

// GetSlaveDB 获取db从节点连接对象
func GetSlaveDB() *gorm.DB {
	if SlaveDB == nil {
		initDB()
	}
	return SlaveDB
}

func init() {
	initDB()
}

// 初始化主节点
func initDB() {
	var err error
	// 初始化MasterDB
	err = config.GetDBConfig(etcdcli.EtcdCli, "master", func(cfg *config.DbConfig) {
		log.Logger.Infow("Start connecting to database master node")
		// 创建数据库连接
		MasterDB, err = NewDbClient(cfg)
		if err != nil {
			log.Logger.Panicw("Connect database master node error", "err", err)
		}
		log.Logger.Infow("Connect database master node successfully")
	})
	if err != nil {
		log.Logger.Panicw("Get db master configuration error", "err", err)
	}

	// 初始化SlaveDB
	err = config.GetDBConfig(etcdcli.EtcdCli, "slave", func(cfg *config.DbConfig) {
		log.Logger.Infow("Start connecting to database slave node")
		// 创建数据库连接
		SlaveDB, err = NewDbClient(cfg)
		if err != nil {
			log.Logger.Panicw("Connect database slave node error", "err", err)
		}
		log.Logger.Infow("Connect database slave node successfully")
	})
	if err != nil {
		log.Logger.Panicw("Get db slave configuration error", "err", err)
	}

}

// NewDbClient 创建数据库连接
func NewDbClient(cfg *config.DbConfig) (*gorm.DB, error) {
	if cfg == nil {
		return nil, errors.New("The database configuration file can not be empty.")
	}
	// 拼接连接数据库字符串
	connStr := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
		cfg.User,
		cfg.Passwd,
		cfg.Address,
		cfg.Port,
		cfg.DbName)
	// 连接数据库
	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}
	// 设置表名前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return cfg.Prefix + defaultTableName
	}
	// 禁用表名多元化
	db.SingularTable(true)
	// 是否开启debug模式
	if cfg.Debug {
		// debug 模式
		db = db.Debug()
	}
	// 连接池最大连接数
	db.DB().SetMaxIdleConns(cfg.MaxIdleConns)
	// 默认打开连接数
	db.DB().SetMaxOpenConns(cfg.MaxOpenConns)
	// 开启协程ping MySQL数据库查看连接状态
	go func() {
		for {
			// ping
			err = db.DB().Ping()
			if err != nil {
				log.Logger.Errorw("mysql ping error", "err", err)
			}
			// 间隔30s ping一次
			time.Sleep(time.Second * 30)
		}
	}()

	return db, err
}
