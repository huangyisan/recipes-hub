package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	MysqlConf MysqlConf
}

type MysqlConf struct {
	DataSource string
	Timeout    int64
}
