package global
//存放一些redis表名，同一管理,使用map还是struct
type RedisCon struct {
	JwtBlacklist string
}

var (
	RedisC = &RedisCon{
		JwtBlacklist: "JwtBlacklist",
	}
)
