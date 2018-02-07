package redis

import (
	"time"
	"github.com/satori/go.uuid"
	"menteslibres.net/gosexy/redis"
)

// redis 获取锁和释放锁

func AcquireLockWithTimeout(conn *redis.Client, lockName string, acquireTimeout, lockTimeout int64) string {
	identifier := uuid.NewV4().String()
	lockKey := "lock:" + lockName
	lockExpire := uint64(lockTimeout)

	end := time.Now().Unix() + acquireTimeout
	for {
		if time.Now().Unix() < end {
			if boolLock, _ := conn.SetNX(lockKey, identifier); boolLock == true {
				conn.Expire(lockKey, lockExpire)
				return identifier
			}

			if ret, _ := conn.TTL(lockKey); ret == -1 {
				conn.Expire(lockKey, lockExpire)
			}

			time.Sleep(time.Duration(5) * time.Millisecond)
		} else {
			break
		}
	}

	return ""
}

func ReleaseLock(conn *redis.Client, lockName, identifier string) bool {
	lockKey := "lock:" + lockName

	for {
		conn.Watch(lockKey)
		identifierRedis, _ := conn.Get(lockKey)
		if identifier == identifierRedis {
			conn.Multi()
			conn.Del(lockKey)

			res, _ := conn.Exec()
			if res == nil {
				continue
			}
			return true
		}

		conn.Unwatch()
		break
	}

	return false
}
