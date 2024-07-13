package config

type Jwt struct {
	SignKey string `mapstructure:"signKey" json:"signKey" yaml:"signKey"`
	JwtTtl  int64  `mapstructure:"jwtTtl" json:"jwtTtl" yaml:"jwtTtl"` // token 有效期（秒）
}
