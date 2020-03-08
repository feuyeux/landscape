package common

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

/**
redis client
*/
type RedisClient struct {
	client *redis.Client
}

func (c *RedisClient) Open(host, port, pw string) {
	c.client = redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: pw, // no password set
		DB:       1,  // use default DB
	})
}

//PingPong with redis
func (c *RedisClient) PingPong() {
	pong, _ := c.client.Ping().Result()
	fmt.Println(pong)
}

//kv
func (c *RedisClient) SaveString(key, value string) (string, error) {
	return c.client.Set(key, value, 0).Result()
}

func (c *RedisClient) ReadString(key string) string {
	val, err := c.client.Get(key).Result()
	if err == redis.Nil {
		fmt.Println(key, " does not exist")
	} else if err != nil {
		panic(err)
	}
	return val
}

//queue
func (c *RedisClient) PushToQueue(key, value string) (int64, error) {
	return c.client.RPush(key, value).Result()
}

func (c *RedisClient) PushToQueue2(key, value string, snd time.Duration) (int64, error) {
	result, err := c.client.RPush(key, value).Result()
	c.client.Expire(key, snd)
	return result, err
}

func (c *RedisClient) PopFromQueue(key string) (string, error) {
	return c.client.LPop(key).Result()
}

func (c *RedisClient) GetAllFromQueue(key string) ([]string, error) {
	return c.GetFromQueue(key, 0, -1)
}

func (c *RedisClient) GetFromQueue(key string, start, stop int64) ([]string, error) {
	return c.client.LRange(key, start, stop).Result()
}

func (c *RedisClient) GetQueueByIndex(key string, index int64) (string, error) {
	return c.client.LIndex(key, index).Result()
}
func (c *RedisClient) GetQueueLength(key string) (int64, error) {
	return c.client.LLen(key).Result()
}

func (c *RedisClient) GetFirstOne(key string) (string, error) {
	return c.GetQueueByIndex(key, 0)
}

func (c *RedisClient) GetLastOne(key string) (string, error) {
	length, _ := c.GetQueueLength(key)
	return c.GetQueueByIndex(key, length-1)
}

//
func (c *RedisClient) SaveMapValue(key, field string, value interface{}) (bool, error) {
	return c.client.HSet(key, field, value).Result()
}

func (c *RedisClient) SaveMap(key string, fields map[string]interface{}) (string, error) {
	return c.client.HMSet(key, fields).Result()
}

func (c *RedisClient) GetMapValue(key string, field string) string {
	val, err := c.client.HGet(key, field).Result()
	if err == redis.Nil {
		fmt.Println(key, " does not exist")
	} else if err != nil {
		panic(err)
	}
	return val
}
func (c *RedisClient) GetMap(key string) (map[string]string, error) {
	return c.client.HGetAll(key).Result()
}

func (c *RedisClient) DeleteFromMap(key, field string) (int64, error) {
	return c.client.HDel(key, field).Result()
}

//Close redis connection
func (c *RedisClient) Close() {
	c.client.Close()
}
