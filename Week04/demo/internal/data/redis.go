package data

import (
	"demo/pkg/cache/redis"
	"demo/pkg/config"
	c "demo/pkg/config/cache/redis"
	"github.com/go-kratos/kratos/pkg/log"
)

func NewRedis() (r *redis.Redis, cf func(), err error) {
	var (
		yml string
		cfg *c.Config
	)
	yml, err = config.ReadYaml("../configs/redis.yml")
	if err != nil {
		return
	}

	pbCfg := new(c.PbConfig)
	err = config.ApplyYAML(yml, pbCfg)
	if err != nil {
		return
	}

	cfg, err = c.NewConfig(pbCfg, pbCfg.Options()...)
	if err != nil {
		return
	}

	r = redis.New(cfg)
	cf = func(){r.Close()}

	log.Info("redis config: %+v\n", r.Config)
	return
}