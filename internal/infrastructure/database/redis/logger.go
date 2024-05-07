package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"net"
)

type loggerHook struct {
}

func (loggerHook) DialHook(hook redis.DialHook) redis.DialHook {
	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		logrus.Infof("dialing %s %s\n", network, addr)
		conn, err := hook(ctx, network, addr)
		logrus.Infof("finished dialing %s %s\n", network, addr)
		return conn, err
	}
}

func (loggerHook) ProcessHook(hook redis.ProcessHook) redis.ProcessHook {
	return func(ctx context.Context, cmd redis.Cmder) error {
		logrus.Infof("processing cmd %s\n", cmd.String())
		err := hook(ctx, cmd)
		logrus.Infof("finished processing: <%s>\n", cmd)
		return err
	}
}

func (loggerHook) ProcessPipelineHook(hook redis.ProcessPipelineHook) redis.ProcessPipelineHook {
	return func(ctx context.Context, cmds []redis.Cmder) error {
		logrus.Infof("pipeline starting processing: %v\n", cmds)
		err := hook(ctx, cmds)
		logrus.Infof("pipeline finished processing: %v\n", cmds)
		return err
	}
}
