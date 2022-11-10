package config

type config struct {
	Environment     environment `yaml:"environment" validate:"gte=1,lte=2"`
	LogLevel        uint32      `yaml:"log_level" validate:"required"`
	Address         string      `yaml:"address" validate:"required"`
	Cors            []string    `yaml:"cors" validate:"required"`
	ProjectorSecret string      `yaml:"projector_secret" validate:"required"`
	AdminSecret     string      `yaml:"admin_secret" validate:"required"`
	StudentSecret   string      `yaml:"student_secret" validate:"required"`
}
