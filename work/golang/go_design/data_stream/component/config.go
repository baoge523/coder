package component

type ComponentConfig any

// ConfigValidate 验证配置
type ConfigValidate interface {
	ValidateConfig() bool
}

// ValidateConfig 如果config实现了ConfigValidate，则验证
func ValidateConfig(config ComponentConfig) bool {
	if validate, ok := config.(ConfigValidate); ok {
		return validate.ValidateConfig()
	}
	return false
}
