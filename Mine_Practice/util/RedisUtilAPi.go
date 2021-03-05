package util

/*import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/ansel1/merry"
	"github.com/go-redis/redis"
	"github.com/astaxie/beego/adapter/logs"
)

var RedisClient *redis.ClusterClient

func NewClient() *redis.ClusterClient {
	if RedisClient != nil {
		return RedisClient
	}

	RedisClient = redis.NewClusterClient(&redis.ClusterOptions{
		//-------------------------------------------------------------------------------------------
		//集群相关的参数

		//集群节点地址，理论上只要填一个可用的节点客户端就可以自动获取到集群的所有节点信息。但是最好多填一些节点以增加容灾能力，因为只填一个节点的话，如果这个节点出现了异常情况，则Go应用程序在启动过程中无法获取到集群信息。
		Addrs: []string{"127.0.0.1:7000", "127.0.0.1:7001", "127.0.0.1:7002", "127.0.0.1:7003", "127.0.0.1:7004", "127.0.0.1:7005"},

		MaxRedirects: 8, // 当遇到网络错误或者MOVED/ASK重定向命令时，最多重试几次，默认8
	})

	return RedisClient
}

func RedisSet(key string, value interface{}, expire int) error {
	if expire > 0 {
		err := RedisClient.Do("SET", key, value, "PX", expire).Err()
		if err != nil {
			logs.Error("RedisSet Error! key:", key, "Details:", err.Error())
			return err
		}
	} else {
		err := RedisClient.Do("SET", key, value).Err()
		if err != nil {
			logs.Error("RedisSet Error! key:", key, "Details:", err.Error())
			return err
		}
	}

	return nil
}

func RedisKeyExists(key string) (bool, error) {
	ok, err := RedisClient.Do("EXISTS", key).Bool()
	return ok, err
}

func RedisGet(key string) (string, error) {
	value, err := RedisClient.Do("GET", key).String()
	if err != "nil" {
		return "", nil
	}

	return value, nil
}

func RedisGetResult(key string) (interface{}, error) {
	v, err := RedisClient.Do("GET", key).Result()
	if err == redis.Nil {
		return v, nil
	}
	return v, err
}

func RedisGetInt(key string) (int, error) {
	v, err := RedisClient.Do("GET", key).Int()
	if err == redis.Nil {
		return 0, nil
	}
	return v, err
}

func RedisGetInt64(key string) (int64, error) {
	v, err := RedisClient.Do("GET", key).Int64()
	if err == redis.Nil {
		return 0, nil
	}
	return v, err
}

func RedisGetUint64(key string) (uint64, error) {
	v, err := RedisClient.Do("GET", key).Uint64()
	if err == redis.Nil {
		return 0, nil
	}
	return v, err
}

func RedisGetFloat64(key string) (float64, error) {
	v, err := RedisClient.Do("GET", key).Float64()
	if err == redis.Nil {
		return 0.0, nil
	}
	return v, err
}

func RedisExpire(key string, expire int) error {
	err := RedisClient.Do("EXPIRE", key, expire).Err()
	if err != nil {
		logs.Error("RedisExpire Error!", key, "Details:", err.Error())
		return err
	}

	return nil
}

func RedisPTTL(key string) (int, error) {
	ttl, err := RedisClient.Do("PTTL", key).Int()
	if err != nil {
		return -1, err
	}

	return ttl, nil
}

func RedisTTL(key string) (int, error) {
	ttl, err := RedisClient.Do("TTL", key).Int()
	if err != nil {
		return -1, err
	}

	return ttl, nil
}

func RedisSetJson(key string, value interface{}, expire int) error {
	jsonData, _ := json.Marshal(value)
	if expire > 0 {
		err := RedisClient.Do("SET", key, jsonData, "PX", expire).Err()
		if err != nil {
			logs.Error("RedisSetJson Error! key:", key, "Details:", err.Error())
			return err
		}
	} else {
		err := RedisClient.Do("SET", key, jsonData).Err()
		if err != nil {
			logs.Error("RedisSetJson Error! key:", key, "Details:", err.Error())
			return err
		}
	}

	return nil
}

func RedisGetJson(key string) ([]byte, error) {
	value, err := RedisClient.Do("GET", key).String()
	if err != nil {
		return nil, nil
	}

	return []byte(value), nil
}

func RedisDel(key string) error {
	err := RedisClient.Do("DEL", key).Err()
	if err != nil {
		logs.Error("RedisDel Error! key:", key, "Details:", err.Error())
	}
	return err
}

func RedisHGet(key, field string) (string, error) {
	value, err := RedisClient.Do("HGET", key, field).String()
	if err != nil {
		return "", nil
	}

	return value, nil
}

func RedisHSet(key, field, value string) error {
	err := RedisClient.Do("HSET", key, field, value).Err()
	if err != nil {
		logs.Error("RedisHSet Error!", key, "field:", field, "Details:", err.Error())
	}
	return err
}

func RedisHDel(key, field string) error {
	err := RedisClient.Do("HDEL", key, field).Err()
	if err != nil {
		logs.Error("RedisHDel Error!", key, "field:", field, "Details:", err.Error())
	}
	return err
}

func RedisZAdd(key, member, score string) error {
	err := RedisClient.Do("ZADD", key, score, member).Err()
	if err != nil {
		logs.Error("RedisZAdd Error!", key, "member:", member, "score:", score, "Details:", err.Error())
	}
	return err
}

func RedisZRank(key, member string) (int, error) {
	rank, err := RedisClient.Do("ZRANK", key, member).Int()
	if err == redis.Nil {
		return -1, nil
	}

	if err != nil {
		logs.Error("RedisZRank Error!", key, "member:", member, "Details:", err.Error())
		return -1, nil
	}

	return rank, err
}

func RedisZRange(key string, start, stop int) (values []string, err error) {
	values, err = RedisClient.ZRange(key, int64(start), int64(stop)).Result()
	if err != nil {
		logs.Error("RedisZRange Error!", key, "start:", start, "stop:", stop, "Details:", err.Error())
		return
	}

	return
}

func RedisZRangeWithScores(key string, start, stop int) (values []redis.Z, err error) {
	values, err = RedisClient.ZRangeWithScores(key, int64(start), int64(stop)).Result()
	if err != nil {
		logs.Error("RedisZRange Error!", key, "start:", start, "stop:", stop, "Details:", err.Error())
		return
	}

	return
}

func RedisZRem(key, member string) error {
	err := RedisClient.Do("ZREM", key, member).Err()
	if err != nil {
		logs.Error("RedisZRem Error!", key, "member:", member, "Details:", err.Error())
	}
	return err
}

func RedisRPUSH(key string, member string) (err error) {
	err = RedisClient.Do("RPUSH", key, member).Err()
	if err != nil {
		logs.Error("RedisRPUSH Error!", key, member, "Details:", err.Error())
		return
	}

	return
}

func RedisBLPOP(timeout time.Duration, keys ...string) (value []string, err error) {
	value, err = RedisClient.BLPop(timeout, keys...).Result()
	if err == redis.Nil {
		err = nil
		return
	}

	if err != nil {
		logs.Error("BLPop Error!", keys, timeout, "Details:", err.Error())
		return
	}
	return
}

func RedisLLEN(key string) (value int64, err error) {
	value, err = RedisClient.LLen(key).Result()
	if err != nil {
		logs.Error("RedisLLEN Error!", key, "Details:", err.Error())
		return
	}

	return
}

func RedisLRange(key string, start, stop int) (values []string, err error) {
	values, err = RedisClient.LRange(key, int64(start), int64(stop)).Result()
	if err != nil {
		logs.Error("RedisLRange Error!", key, "start:", start, "stop:", stop, "Details:", err.Error())
		return
	}

	return
}

func RedisKeys(pattern string) (keys []string, err error) {
	keys, err = RedisClient.Keys(pattern).Result()
	if err != nil {
		logs.Error("RedisKeys Error!", pattern, "Details:", err.Error())
		return
	}

	return
}

// RedisListAllValuesWithPrefix will take in a key prefix and return the value of all the keys that contain that prefix
func RedisListAllValuesWithPrefix(prefix string) (map[string]string, error) {
	// Grab all the keys with the prefix
	keys, err := getKeys(fmt.Sprintf("%s*", prefix))
	if err != nil {
		return nil, err
	}

	// We will now iterate through all of the values to
	values, err := getKeyAndValuesMap(keys, prefix)

	return values, nil
}

// getKeys will take a certain prefix that the keys share and return a list of all the keys
func getKeys(prefix string) ([]string, error) {
	var allKeys []string
	var cursor uint64
	count := int64(10) // count specifies how many keys should be returned in every Scan call

	for {
		var keys []string
		var err error
		keys, cursor, err = RedisClient.Scan(cursor, prefix, count).Result()
		if err != nil {
			return nil, merry.Appendf(err, "error retrieving '%s' keys", prefix)
		}

		allKeys = append(allKeys, keys...)

		if cursor == 0 {
			break
		}

	}

	return allKeys, nil
}

// getKeyAndValuesMap generates a [string]string map structure that will associate an ID with the token value stored in Redis
func getKeyAndValuesMap(keys []string, prefix string) (map[string]string, error) {
	values := make(map[string]string)
	for _, key := range keys {
		value, err := RedisClient.Do("GET", key).String()
		if err != nil {
			return nil, merry.Appendf(err, "error retrieving value for key %s", key)
		}

		// Strip off the prefix from the key so that we save the key to the user ID
		strippedKey := strings.Split(key, prefix)
		values[strippedKey[1]] = value
	}

	return values, nil
}

func RedisBatchDel(key ...string) error {
	err := RedisClient.Del(key...).Err()
	if err != nil {
		logs.Error("RedisBatchDel Error! key:", key, "Details:", err.Error())
	}
	return err
}

func RedisMset(pairs ...interface{}) error {
	err := RedisClient.MSet(pairs...).Err()
	if err != nil {
		logs.Error("RedisMset Error! pairs:", pairs, "Details:", err.Error())
	}
	return err
}
*/
