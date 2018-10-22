package config

import "os"

// 全局配置
const (
	// DEFAULT_ETCD_ADDR 默认etcd地址
	DEFAULT_ETCD_ADDR string = "127.0.0.1:2379"
	// DEFAULT_HTTP_ADDR 默认http监听地址
	DEFAULT_HTTP_ADDR string = "127.0.0.1:18080"
	// DEFAULT_GRPC_ADDR 默认grpc监听地址
	DEFAULT_GRPC_ADDR string = "127.0.0.1:28080"
	// DEFAULT_TCP_ADDR 默认tcp监听地址
	DEFAULT_TCP_ADDR string = "127.0.0.1:38080"
	// DEFAULT_SVC_NAME 默认服务名
	DEFAULT_SVC_NAME string = "default"
	// 当前运行环境，dev or pro
	DEFAULT_MODE string = "dev"
)

// GetETCDAddr 读取etcd服务地址
func GetETCDAddr() string {
	etcdAddr := os.Getenv("ETCD_ADDR")
	if etcdAddr == "" {
		return DEFAULT_ETCD_ADDR
	}
	return etcdAddr
}

// GetHTTPAddr 获取配置值
func GetHTTPAddr() string {
	httpAddr := os.Getenv("HTTP_ADDR")
	if httpAddr == "" {
		return DEFAULT_HTTP_ADDR
	}
	return httpAddr
}

// GetGRPCAddr 读取grpc地址
func GetGRPCAddr() string {
	grpcAddr := os.Getenv("GRPC_ADDR")
	if grpcAddr == "" {
		return DEFAULT_GRPC_ADDR
	}
	return grpcAddr
}

// GetTCPAddr 读取tcp地址
func GetTCPAddr() string {
	grpcAddr := os.Getenv("TCP_ADDR")
	if grpcAddr == "" {
		return DEFAULT_TCP_ADDR
	}
	return grpcAddr
}

// GetSvcName 获取服务名 - [redis 使用不传type，一个服务部分类型使用key]
func GetSvcName(svcType string) string {
	svcName := os.Getenv("SVC_NAME")
	if svcName == "" {
		svcName = DEFAULT_SVC_NAME
	}
	return svcName + "/" + svcType
}

// GetMode 当前运行环境
func GetMode() string {
	mode := os.Getenv("MODE")
	if mode == "" {
		return DEFAULT_MODE
	}
	return mode
}
