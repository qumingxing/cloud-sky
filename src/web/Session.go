package web

import (
	"logs"
	"time"
)

var sessionServiceMap map[string]map[string]interface{} = make(map[string]map[string]interface{})
var sessionMap map[string]*Session = make(map[string]*Session)
var sessionTime map[string]int64 = make(map[string]int64)
var defaultTimeout int64 = 1000 * 60 * 10

type Session struct {
	Id string
}

func init() {
	logs.Debug("init!")
	var session Session
	session.loopTime()
}
func (session Session) loopTime() {
	logs.Debug("loopTime!")
	go func() {
		logs.Debug("go loopTime!")
		for {
			currentTime := time.Now().Unix() * 1000
			for key, value := range sessionTime {
				if (value + defaultTimeout) < currentTime {
					logs.Debug(key, value, "session timeout!")
					expire(key)
				}
			}
			time.Sleep(3000 * time.Millisecond)
		}
	}()
}

func (session *Session) SetTimeout(timeout int64) {
	defaultTimeout = timeout
}
func (session *Session) GetSessionId() string {
	return session.Id
}
func (session *Session) GetTimeout() int64 {
	return defaultTimeout
}
func (session *Session) SetAttribute(key string, value interface{}) {
	serviceMap := sessionServiceMap[session.Id]
	if serviceMap != nil {
		serviceMap[key] = value
	} else {
		userMap := make(map[string]interface{})
		userMap[key] = value
		sessionServiceMap[session.Id] = userMap
	}
	sessionTime[session.Id] = time.Now().Unix() * 1000
}
func (session *Session) GetAttribute(key string) interface{} {
	serviceMap := sessionServiceMap[session.Id]
	if serviceMap != nil {
		sessionTime[session.Id] = time.Now().Unix() * 1000
		return serviceMap[key]
	}
	return nil
}
func (session *Session) RemoveAttribute(key string) {
	serviceMap := sessionServiceMap[session.Id]
	if serviceMap != nil {
		delete(serviceMap, key)
	}
}
func expire(sessionId string) {
	//delete(sessionServiceMap, sessionId)
	delete(sessionMap, sessionId)
	delete(sessionTime, sessionId)
}
