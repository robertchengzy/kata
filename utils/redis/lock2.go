package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/hashicorp/go-uuid"
	"time"
)

// redis 获取锁和释放锁
// @param acquireTimeout 单位：纳秒
// @param lockTimeout 单位：秒
func AcquireLockWithTimeout2(conn *redis.Client, lockName string, acquireTimeout, lockTimeout int64) string {
	identifier, _ := uuid.GenerateUUID()
	lockKey := "lock:" + lockName
	lockExpire := uint64(lockTimeout)

	end := time.Now().UnixNano() + acquireTimeout
	for {
		if time.Now().UnixNano() < end {
			if boolLock := conn.SetNX(lockKey, identifier, time.Duration(lockExpire)).Val(); boolLock {
				return identifier
			}

			if conn.TTL(lockKey).Val() == -1 {
				conn.Expire(lockKey, time.Duration(lockExpire))
			}

			time.Sleep(5 * time.Millisecond)
		} else {
			break
		}
	}

	return ""
}

func ReleaseLock2(conn *redis.Client, lockName, identifier string) bool {
	lockKey := "lock:" + lockName

	txf := func(tx *redis.Tx) error {
		if v, err := tx.Get(lockKey).Result(); err != nil && err != redis.Nil {
			return err
		} else if v == identifier {
			_, err := tx.Pipelined(func(pipe redis.Pipeliner) error {
				pipe.Del(lockName)
				return nil
			})
			return err
		}
		return nil
	}

	for {
		if err := conn.Watch(txf, lockKey); err == nil {
			return true
		} else if err == redis.TxFailedErr {
			fmt.Println("watch key is modified, retry to release lock. err:", err.Error())
		} else {
			fmt.Println("err:", err.Error())
			return false
		}
	}
}
