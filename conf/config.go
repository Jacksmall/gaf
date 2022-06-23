package conf

var (
	Config SystemConfig
)

type SystemConfig struct {
	Server Server `json:"server"`
}
