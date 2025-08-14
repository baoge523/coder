// server 服务启动
package server

import (
	"code_design/data_stream/component"
	"context"
	"sync"
)

type ComponentManager struct {
	components map[component.ComponentType]map[component.ComponentName]component.ComponentLife
	extenders  map[component.ComponentName]component.ComponentLife
}

type ConfigSetting struct {
	cm ComponentManager
}

type Server struct {
	cs     ConfigSetting
	cancel context.CancelFunc
	wait   sync.WaitGroup
}

func New(cs ConfigSetting) *Server {

	server := &Server{cs: cs}

	// some setup ....

	return server
}

func (s *Server) Run(ctx context.Context) error {

	return nil
}

func (s *Server) Stop() error {
	if s.cancel != nil {
		s.cancel() // 通知内部协程，表示当前服务关闭了
	}
	s.wait.Wait() // 等待内部服务执行完成
	return nil
}
