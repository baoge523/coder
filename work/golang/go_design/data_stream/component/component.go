package component

import "context"

type Component interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

type ComponentType string
type ComponentName string

// Factory 组件的通用工厂，提供类型和名称
type Factory interface {
	ID() ID // 组件的唯一标识
}

type BuildInfo struct {
	Version string
}

// ID 组件标识
type ID struct {
	Type ComponentType
	Name ComponentName
}
