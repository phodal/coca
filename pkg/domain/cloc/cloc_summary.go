package cloc

// LanguageSummary to generate output like cloc
type LanguageSummaryCloc struct {
	Name    string `yaml:"name"`
	Code    int64  `yaml:"code"`
	Comment int64  `yaml:"comment"`
	Blank   int64  `yaml:"blank"`
	Count   int64  `yaml:"nFiles"`
}

type SummaryStruct struct {
	Code    int64 `yaml:"code"`
	Comment int64 `yaml:"comment"`
	Blank   int64 `yaml:"blank"`
	Count   int64 `yaml:"nFiles"`
}

type HeaderStruct struct {
	Url            string  `yaml:"url"`
	Version        string  `yaml:"version"`
	ElapsedSeconds float64 `yaml:"elapsed_seconds"`
	NFiles         int64   `yaml:"n_files"`
	NLines         int64   `yaml:"n_lines"`
	FilesPerSecond float64 `yaml:"files_per_second"`
	LinesPerSecond float64 `yaml:"lines_per_second"`
}

type LanguageReportStart struct {
	Header HeaderStruct
}

type LanguageReportEnd struct {
	Sum SummaryStruct `yaml:"SUM"`
}

type ClocSummary struct {
	Header HeaderStruct
	Sum    SummaryStruct `yaml:"SUM,omitempty"`
	Java   *LanguageSummaryCloc `yaml:"Java,omitempty"`
	Kotlin *LanguageSummaryCloc `yaml:"Kotlin,omitempty"`
	Groovy *LanguageSummaryCloc `yaml:"Groovy,omitempty"`
}
