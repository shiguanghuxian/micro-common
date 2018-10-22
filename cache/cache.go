package cache

import (
	"errors"
	"time"

	"github.com/shiguanghuxian/micro-common/config"
	"github.com/shiguanghuxian/micro-common/etcdcli"
	"github.com/shiguanghuxian/micro-common/log"
	goredis "gopkg.in/redis.v5"
)

/* 缓存连接 redis */

var (
	// Client redis 连接对象
	Client goredis.Cmdable
)

// GetClient 获取redis连接对象
func GetClient() goredis.Cmdable {
	if Client == nil {
		initRedis()
	}
	return Client
}

func init() {
	initRedis()
}

// 初始化redis
func initRedis() {
	var err error
	err = config.GetRedisConfg(etcdcli.EtcdCli, func(cfg *config.RedisConfg) {
		Client, err = NewClient(cfg)
		if err != nil {
			log.Logger.Panicw("Creating redis connection errors", "err", err)
		}
	})
	if err != nil {
		log.Logger.Panicw("Get redis configuration error", "err", err)
	}
}

// NewClient 创建客户端连接
func NewClient(cfg *config.RedisConfg) (client goredis.Cmdable, err error) {
	if cfg == nil {
		err = errors.New("The redis configuration file can not be empty.")
		return
	}
	log.Logger.Infow("Start connecting to redis database")
	if cfg.IsCluster == true {
		// redis集群
		client = goredis.NewClusterClient(&goredis.ClusterOptions{
			Addrs:    cfg.Address,
			Password: cfg.Password,
			PoolSize: cfg.PoolSize,
		})
	} else {
		// redis单机
		client = goredis.NewClient(&goredis.Options{
			Addr:     cfg.Address[0],
			Password: cfg.Password,
			DB:       cfg.Db,
			PoolSize: cfg.PoolSize,
		})
	}
	// ping 防止断开
	go func() {
		for {
			err := client.Ping().Err()
			if err != nil {
				log.Logger.Errorw("redis ping error", "err", err)
			}

			time.Sleep(time.Second * 30)
		}
	}()
	log.Logger.Infow("Connect to redis database successfully")
	return
}
