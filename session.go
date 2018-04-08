package main

import (
	"sync"
	"time"
)

type SessionModel struct {
	SessionId string
	Session   map[string]interface{}
	Expire    int64
	Lock      sync.Mutex
}

var (
	allSession    = make(map[string]*SessionModel)
	lock          = new(sync.Mutex)
	lastCheckTime int64
)

type User struct {
	Id int
	Name string
}

func GetSession(sessionId string) *SessionModel {
	if sessionId == "" {
		return nil
	}
	lock.Lock()
	defer lock.Unlock()
	session := allSession[sessionId]
	if session == nil {
		return nil
	}
	return session
}
func NewOrRefresh(sessionId string) map[string]interface{} {
	if sessionId == "" {
		return nil
	}
	now := time.Now().UnixNano()
	lock.Lock()
	defer lock.Unlock()
	session := allSession[sessionId]
	if session == nil {
		session = new(SessionModel)
		session.SessionId = sessionId
		session.Session = make(map[string]interface{})
		session.Expire = now + 30*time.Minute.Nanoseconds()
		allSession[sessionId] = session
	} else {
		session.Expire = now + 30*time.Minute.Nanoseconds()
		// 至少经过五分钟检查session
		if lastCheckTime+5*time.Minute.Nanoseconds() < now {
			allSessionCheck()
		}
	}
	return session.Session
}

func allSessionCheck() {
	now := time.Now().UnixNano()
	lastCheckTime = now
	arr := make([]string, 0, 10)
	for key, ss := range allSession {
		if ss.Expire < now {
			arr = append(arr, key)
		}
	}
	for _, key := range arr {
		delete(allSession, key)
	}
}
func Get(sessionId string, key string) interface{} {
	if sessionId == "" {
		return nil
	}
	session := GetSession(sessionId)
	if session == nil {
		return nil
	}
	session.Lock.Lock()
	defer session.Lock.Unlock()
	return session.Session[key]
}

func GetUser(ctx string) *User {
	obj := Get(ctx, "user")
	if obj == nil {
		return nil
	}
	return obj.(*User)
}

func Set(sessionId string, key string, value interface{}) {
	if sessionId == "" {
		return
	}
	session := GetSession(sessionId)
	if session == nil {
		return
	}
	session.Lock.Lock()
	defer session.Lock.Unlock()
	session.Session[key] = value
}
func SetUser(ctx string, user *User) {
	Set(ctx, "user", user)
}
func GetString(ctx string, key string) string {
	obj := Get(ctx, key)
	if obj == nil {
		return ""
	}
	return obj.(string)
}
func GetInt(ctx string, key string) int {
	obj := Get(ctx, key)
	if obj == nil {
		return 0
	}
	return obj.(int)
}
func GetInt64(ctx string, key string) int64 {
	obj := Get(ctx, key)
	if obj == nil {
		return 0
	}
	return obj.(int64)
}
func GetBool(ctx string, key string) bool {
	obj := Get(ctx, key)
	if obj == nil {
		return false
	}
	return obj.(bool)
}

func Filter(sessionId string) {
	apsid := sessionId
	if apsid == "" {
		return
	}
	// 创建session
	NewOrRefresh(apsid)
}
