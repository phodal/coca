package config

type TypeCocaConfig struct {
	ReporterPath string
}

var CocaConfig = &TypeCocaConfig{
	ReporterPath: "coca_reporter",
}

const VERSION = "1.2.0"

// TBS Config List

var (
	DuplicatedAssertionLimitLength = 5
	ASSERTION_LIST = []string{
		"assert",
		"should",
		"check",    // ArchUnit,
		"maynotbe", // ArchUnit,
		"is",       // RestAssured,
		"spec",     // RestAssured,
		"verify",   // Mockito,
	}
)
