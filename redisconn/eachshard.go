package redisconn

import "github.com/xsyr/redispipe/redis"

// EachShard implements redis.Sender.EachShard.
// It just calls callback once with Connection itself.
func (c *Connection) EachShard(cb func(redis.Sender, error) bool) {
	cb(c, nil)
}
