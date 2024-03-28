package enum

type RedisSdkKey = string

const (
	LockPrefix                 RedisSdkKey = `business:lock:`
	JsonPrefix                 RedisSdkKey = `business:json:`
	SessionPrefix              RedisSdkKey = `session:`
	LoginIdMapSessionIdsPrefix RedisSdkKey = `business:loginId2sids:`
)
