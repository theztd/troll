package backend

import (
	"crypto/tls"
	"strings"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	db *redis.Client
}

func NewRedis(dsn string) (*Redis, error) {
	opt, err := redis.ParseURL(dsn)
	if err != nil {
		return nil, err
	}

	if strings.HasPrefix(dsn, "rediss://") {
		opt.TLSConfig = &tls.Config{MinVersion: tls.VersionTLS12}
	}

	return &Redis{db: redis.NewClient(opt)}, nil
}

func (red *Redis) RunAndRenderTpl(query string, tmpl string) ([]byte, error) {
	return []byte{}, nil
}

func (r *Redis) Close() error { return r.db.Close() }
