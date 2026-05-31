package db_data_controller

import (
	"errors"
	"goRedisAdmin/global/global_redis"
	"goRedisAdmin/utils/log_utils"
	"time"

	"github.com/go-redis/redis"
)

// DbDataHelpModel is a generic payload for Redis add/get operations.
type DbDataHelpModel struct {
	DbNum       int    `json:"db_num"`
	VType       string `json:"type"`
	Key         string `json:"key"`
	Val         string `json:"val"`
	Expire      int    `json:"expire"`
	HashKey     string `json:"hash_key,omitempty"`
	StreamField string `json:"stream_field,omitempty"`
	Score       int    `json:"score,omitempty"`
}

// DbDataHelpCont wraps a redis client and operation payload.
type DbDataHelpCont struct {
	redis *redis.Client
	DbDataHelpModel
}

// NewDbDataHelpController validates payload and initializes redis client.
func NewDbDataHelpController(v *DbDataHelpModel) (*DbDataHelpCont, error) {
	if v.DbNum < 0 || v.DbNum > 15 {
		return nil, errors.New("dbNum must be between 0 and 15")
	}
	cont := new(DbDataHelpCont)
	cont.DbDataHelpModel = *v
	rd, err := cont.getRedisClient()
	if err != nil {
		log_utils.WriteLog("err", err, nil)
		return nil, err
	}
	cont.redis = rd
	return cont, nil
}

// getRedisClient creates client bound to the payload db index.
func (d *DbDataHelpCont) getRedisClient() (*redis.Client, error) {
	return global_redis.GetRedisClient(d.DbNum)
}

// CloseClient closes held redis client.
func (d *DbDataHelpCont) CloseClient() {
	_ = d.redis.Close()
}

// AddString sets a string value with optional expiration.
func (d *DbDataHelpCont) AddString() error {
	_, err := d.redis.Set(d.Key, d.Val, time.Duration(int64(d.Expire))*time.Second).Result()
	return err
}

// GetString gets string value by key.
func (d *DbDataHelpCont) GetString() (string, error) {
	return d.redis.Get(d.Val).Result()
}

// AddList pushes value into list and applies optional expiration.
func (d *DbDataHelpCont) AddList() error {
	_, err := d.redis.LPush(d.Key, d.Val).Result()
	if err != nil {
		return err
	}
	return setExpire(d)
}

// AddSet inserts value into set and applies optional expiration.
func (d *DbDataHelpCont) AddSet() error {
	_, err := d.redis.SAdd(d.Key, d.Val).Result()
	if err != nil {
		return err
	}
	return setExpire(d)
}

// AddZSet inserts value into sorted-set and applies optional expiration.
func (d *DbDataHelpCont) AddZSet() error {
	zVal := redis.Z{
		Score:  1,
		Member: d.Val,
	}
	_, err := d.redis.ZAdd(d.Key, zVal).Result()
	if err != nil {
		return err
	}
	return setExpire(d)
}

// AddHash sets hash field/value and applies optional expiration.
func (d *DbDataHelpCont) AddHash() error {
	_, err := d.redis.HSet(d.Key, d.HashKey, d.Val).Result()
	if err != nil {
		return err
	}
	return setExpire(d)
}

// AddStream appends an entry into stream and applies optional expiration.
func (d *DbDataHelpCont) AddStream() error {
	field := d.StreamField
	if field == "" {
		field = "value"
	}
	_, err := d.redis.XAdd(
		&redis.XAddArgs{
			Stream: d.Key,
			ID:     "*",
			Values: map[string]interface{}{
				field: d.Val,
			},
		},
	).Result()
	if err != nil {
		return err
	}
	return setExpire(d)
}

// setExpire applies expiration only when Expire > 0.
func setExpire(d *DbDataHelpCont) error {
	if d.Expire > 0 {
		_, err := d.redis.Expire(d.Key, time.Duration(int64(d.Expire))*time.Second).Result()
		return err
	} else {
		return nil
	}
}
