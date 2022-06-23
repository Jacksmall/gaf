package conf

// 定义各种配置项entry

// Server 服务器配置项
type Server struct {
	Env          string `json:"env"`
	Version      string `json:"version"`
	Port         int    `json:"port"`
	ReadTimeout  int    `json:"readTimeout"`
	WriteTimeout int    `json:"writeTimeout"`
	Cors         bool   `json:"cors"`
}
