package config

type TypeCocaConfig struct {
	ReporterPath string
	ClocDir      string
}

var CocaConfig = &TypeCocaConfig{
	ReporterPath: "coca_reporter",
	ClocDir:      "/cloc",
}

const VERSION = "2.4.1"
