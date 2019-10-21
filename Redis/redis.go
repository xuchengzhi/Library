package Redis

import (
	"github.com/xuchengzhi/redigo/redis"
)

func Redis_(ip, port, pwd string, db int) (redis.Conn, error) {
	address := fmt.Sprintf("%v:%v", ip, port)
	RedisClient, err := redis.Dial("tcp", address, redis.DialDatabase(db), redis.DialPassword(pwd))
	if err != nil {
		fmt.Println("connect redis error :", err)
		return nil, err
	}

	return RedisClient, nil
}
