package config

import (
    "github.com/gin-gonic/gin"
)

var appConfig map[string]map[string]string = map[string]map[string]string{
    "debug" : map[string]string{
        "port" : "3002",
        "datamoreRedisUrl" : "rc7.dc.DCMETA.db:50007",
        "datamoreRedisPassword" : "redis@dc",
        "datamoreMysql" : "idata:idata11@tcp(10.224.133.183:20000)/db_datamore?charset=utf8",
        "allowOrigin" : "*",
        "needAuth" : "false",
    },

    "test" : map[string]string{
        "port" : "80",
        "datamoreRedisUrl" : "rc7.dc.DCMETA.db:50007",
        "datamoreRedisPassword" : "redis@dc",
        "datamoreMysql" : "idata:idata11@tcp(10.224.133.183:20000)/db_datamore?charset=utf8",
        "allowOrigin" : "http://pvp.qq.com",
        "needAuth" : "true",
    },

    "release" : map[string]string{
        "port" : "80",
        "datamoreRedisUrl" : "rc7.dc.DCMETA.db:50007",
        "datamoreRedisPassword" : "redis@dc",
        "datamoreMysql" : "idata:idata11@tcp(10.224.133.183:20000)/db_datamore?charset=utf8",
        "allowOrigin" : "http://pvp.qq.com",
        "needAuth" : "true",
    },
}

func Get(key string) string {
    mode := gin.Mode()

    return appConfig[mode][key]
}