package config

import (
	"context"
	"errors"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/naoina/toml"
)

/* 读取各种配置 */

/* 读取mysql配置 */

// DbConfig 数据库配置
type DbConfig struct {
	Debug        bool   `toml:"debug"`          // 是否调试模式
	Address      string `toml:"address"`        // 数据库连接地址
	Port         int    `toml:"port"`           // 数据库端口
	MaxIdleConns int    `toml:"max_idle_conns"` // 连接池最大连接数
	MaxOpenConns int    `toml:"max_open_conns"` // 默认打开连接数
	User         string `toml:"user"`           // 数据库用户名
	Passwd       string `toml:"passwd"`         // 数据库密码
	DbName       string `toml:"db_name"`        // 数据库名
	Prefix       string `toml:"prefix"`         // 数据库表前缀
}

// GetDBConfig 获取数据库配置
func GetDBConfig(cli *clientv3.Client, node string, updateConfig func(*DbConfig)) error {
	if node != "master" && node != "slave" {
		return errors.New("Node can only be master or slave.")
	}
	key := "home/config/mysql/account/" + node + ".toml"
	// 初始化master连接
	etcdResp, err := cli.Get(context.Background(), key, clientv3.WithPrefix())
	if err != nil {
		return err
	}
	if len(etcdResp.Kvs) == 0 {
		return errors.New("The database configuration MySQL master node is configured to be empty.")
	}
	dbConfig := new(DbConfig)
	err = toml.Unmarshal(etcdResp.Kvs[0].Value, dbConfig)
	if err != nil {
		return err
	}
	// 监视key变化
	go func() {
		rch := cli.Watch(context.Background(), key)
		for wresp := range rch {
			for _, ev := range wresp.Events {
				if string(ev.Kv.Key) != key || ev.Type != mvccpb.PUT {
					continue
				}
				err = toml.Unmarshal(ev.Kv.Value, dbConfig)
				if err != nil {
					continue
				}
				updateConfig(dbConfig)
			}
		}
	}()

	// 调用更新配置回调函数
	updateConfig(dbConfig)

	return nil
}

/* redis配置 */

// RedisConfg redis配置
type RedisConfg struct {
	Address   []string `toml:"address"`    // redis 服务器地址,包括地址和端口 127.0.0.1:6379
	Password  string   `toml:"password"`   // redis 密码
	Db        int      `toml:"db"`         // 连接的数据库
	PoolSize  int      `toml:"pool_size"`  // 连接池大小
	IsCluster bool     `toml:"is_cluster"` // 是否集群模式
}

// GetRedisConfg 获取redis配置
func GetRedisConfg(cli *clientv3.Client, updateConfig func(*RedisConfg)) error {
	key := "home/config/redis/cfg.toml"
	// 初始化master连接
	etcdResp, err := cli.Get(context.Background(), key, clientv3.WithPrefix())
	if err != nil {
		return err
	}
	if len(etcdResp.Kvs) == 0 {
		return errors.New("The redis configuration configured to be empty.")
	}
	redisConfg := new(RedisConfg)
	err = toml.Unmarshal(etcdResp.Kvs[0].Value, redisConfg)
	if err != nil {
		return err
	}

	// 监视key变化
	go func() {
		rch := cli.Watch(context.Background(), key)
		for wresp := range rch {
			for _, ev := range wresp.Events {
				if string(ev.Kv.Key) != key || ev.Type != mvccpb.PUT {
					continue
				}
				err = toml.Unmarshal(ev.Kv.Value, redisConfg)
				if err != nil {
					continue
				}
				updateConfig(redisConfg)
			}
		}
	}()

	// 调用更新配置回调函数
	updateConfig(redisConfg)
	return nil
}
