package config

type TypeCocaConfig struct {
	ReporterPath string
}

var CocaConfig = &TypeCocaConfig{
	ReporterPath: "coca_reporter",
}

const VERSION = "1.0.0"
