package main

import (
	"github.com/feuyeux/landscape/src/cli"
	"github.com/feuyeux/landscape/src/common"
)

func main() {
	redisClient := common.RedisClient{}
	cli.Run(redisClient)
}
