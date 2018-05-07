package conf_struct

type BaseYaml struct {
	Env string `yaml:"env"`
}

func GetBaseYamlConfig(data interface{}) *BaseYaml {
	var baseConfig *BaseYaml
	switch config := data.(type) {
	case *BaseYaml:
		baseConfig = config
	default:
		return nil
	}
	return baseConfig
}
