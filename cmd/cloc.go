package cmd

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/boyter/scc/processor"
	"github.com/phodal/coca/cmd/config"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type CocaClocConfig struct {
	ByDirectory bool
}

var (
	clocConfig CocaClocConfig
)

var clocCmd = &cobra.Command{
	Use:     "cloc",
	Short:   "count lines of code with complexity estimation",
	Long:    fmt.Sprintf("Sloc, Cloc and Code. Count lines of code in a directory with complexity estimation.\nVersion %s\nBen Boyter <ben@boyter.org> + Contributors", processor.Version),
	Version: processor.Version,
	Run: func(cmd *cobra.Command, args []string) {
		if clocConfig.ByDirectory {
			var dirs []string
			firstFile, err := ioutil.ReadDir(args[0])
			if err != nil {
				log.Fatal(err)
			}

			for _, f := range firstFile {
				if f.IsDir() {
					dirs = append(dirs, filepath.FromSlash(args[0] + "/" + f.Name()))
				}
			}

			processor.Format = "json"

			_ = createClocDir()
			baseCloc := config.CocaConfig.ReporterPath + "/base_cloc.json"
			processBaseCloc(filepath.FromSlash(args[0]), baseCloc)
			keys := buildBaseKey(baseCloc)

			outputFiles := process_dirs(dirs)
			convertToCsv(outputFiles, keys)

			return
		} else {
			processor.DirFilePaths = args
		}

		if processor.ConfigureLimits != nil {
			processor.ConfigureLimits()
		}
		processor.ConfigureGc()
		processor.ConfigureLazy(true)
		processor.Process()
	},
}

func buildBaseKey(baseDir string) []string {
	contents, _ := ioutil.ReadFile(baseDir)
	var languages []processor.LanguageSummary
	err := json.Unmarshal(contents, &languages)
	if err != nil {
		fmt.Println("Error parsing JSON: ", err)
	}

	var keys []string
	for _, data := range languages {
		keys = append(keys, data.Name)
	}

	return keys
}

func processBaseCloc(input string, output string) {
	processor.DirFilePaths= []string{input}
	processor.FileOutput = filepath.FromSlash(output)
	processor.ConfigureGc()
	processor.ConfigureLazy(true)
	processor.Process()
}

func createClocDir() error {
	return os.Mkdir(config.CocaConfig.ReporterPath+"/cloc/", os.ModePerm)
}

func process_dirs(dirs []string) []string {
	var outputFiles []string

	for _, dir := range dirs {
		baseName := filepath.Base(dir)
		if baseName == ".git" || baseName == ".svn" || baseName == ".hg" || baseName == ".idea" {
			continue
		}
		processor.DirFilePaths = []string{dir}
		outputFile := filepath.FromSlash(config.CocaConfig.ReporterPath + "/cloc/" + baseName + ".json")
		outputFiles = append(outputFiles, outputFile)
		processor.FileOutput = outputFile
		processor.ConfigureGc()
		processor.ConfigureLazy(true)
		processor.Process()
	}

	return outputFiles
}

func convertToCsv(outputFiles []string, keys []string) {
	var basemap = make(map[string]processor.LanguageSummary)
	for _, key := range keys {
		basemap[key] = processor.LanguageSummary{}
	}

	var languageMap = make(map[string]map[string]processor.LanguageSummary)
	for _, file := range outputFiles {
		var f []processor.LanguageSummary
		contents, _ := ioutil.ReadFile(file)
		err := json.Unmarshal(contents, &f)
		if err != nil {
			fmt.Println("Error parsing JSON: ", err)
		}

		baseName := strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
		languageMap[baseName] = make(map[string]processor.LanguageSummary)

		for _, key := range keys {
			var hasSet = false
			for _, lang := range f {
				if key == lang.Name {
					hasSet = true
					languageMap[baseName][key] = lang
				}
			}
			if !hasSet {
				languageMap[baseName][key] = processor.LanguageSummary{};
			}
		}
	}

	var data [][]string
	baseKey := []string{"package", "summary"}
	data = append(data, append(baseKey, keys...))

	for baseName, langSummary := range languageMap {
		var column []string
		column = append(column, baseName)

		var codes []string
		var summary int64
		for _, lang := range langSummary {
			summary = summary + lang.Code
			codes = append(codes, strconv.Itoa(int(lang.Code)))
		}

		column = append(column, strconv.Itoa(int(summary)));
		column = append(column, codes...)
		data = append(data, column);
	}

	file, err := os.Create(filepath.FromSlash(config.CocaConfig.ReporterPath + "/" + "cloc.csv"))
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		fmt.Fprintln(output, strings.Join(value, ","))
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func init() {
	rootCmd.AddCommand(clocCmd)

	addClocConfigs()
}

func addClocConfigs() {
	flags := clocCmd.PersistentFlags()

	flags.BoolVar(&clocConfig.ByDirectory, "by-directory", false, "list directory and out csv")

	flags.Int64Var(&processor.AverageWage, "avg-wage", 56286, "average wage value used for basic COCOMO calculation")
	flags.BoolVar(&processor.DisableCheckBinary, "binary", false, "disable binary file detection")
	flags.BoolVar(&processor.Files, "by-file", false, "display output for every file")
	flags.BoolVar(&processor.Ci, "ci", false, "enable CI output settings where stdout is ASCII")
	flags.BoolVar(&processor.Ignore, "no-ignore", false, "disables .ignore file logic")
	flags.BoolVar(&processor.GitIgnore, "no-gitignore", false, "disables .gitignore file logic")
	flags.BoolVar(&processor.Debug, "debug", false, "enable debug output")
	flags.StringSliceVar(&processor.PathDenyList, "exclude-dir", []string{".git", ".hg", ".svn"}, "directories to exclude")
	flags.IntVar(&processor.GcFileCount, "file-gc-count", 10000, "number of files to parse before turning the GC on")
	flags.StringVarP(&processor.Format, "format", "f", "tabular", "set output format [tabular, wide, json, csv, cloc-yaml, html, html-table]")
	flags.StringSliceVarP(&processor.AllowListExtensions, "include-ext", "i", []string{}, "limit to file extensions [comma separated list: e.g. go,java,js]")
	flags.BoolVarP(&processor.Languages, "languages", "l", false, "print supported languages and extensions")
	flags.BoolVar(&processor.Cocomo, "no-cocomo", false, "remove COCOMO calculation output")
	flags.BoolVar(&processor.Size, "no-size", false, "remove size calculation output")
	flags.StringVar(&processor.SizeUnit, "size-unit", "si", "set size unit [si, binary, mixed, xkcd-kb, xkcd-kelly, xkcd-imaginary, xkcd-intel, xkcd-drive, xkcd-bakers]")
	flags.BoolVarP(&processor.Complexity, "no-complexity", "c", false, "skip calculation of code complexity")
	flags.BoolVarP(&processor.Duplicates, "no-duplicates", "d", false, "remove duplicate files from stats and output")
	flags.BoolVarP(&processor.MinifiedGenerated, "min-gen", "z", false, "identify minified or generated files")
	flags.BoolVarP(&processor.Minified, "min", "", false, "identify minified files")
	flags.BoolVarP(&processor.Generated, "gen", "", false, "identify generated files")
	flags.StringSliceVarP(&processor.GeneratedMarkers, "generated-markers", "", []string{"do not edit"}, "string markers in head of generated files")
	flags.BoolVar(&processor.IgnoreMinifiedGenerate, "no-min-gen", false, "ignore minified or generated files in output (implies --min-gen)")
	flags.BoolVar(&processor.IgnoreMinified, "no-min", false, "ignore minified files in output (implies --min)")
	flags.BoolVar(&processor.IgnoreGenerated, "no-gen", false, "ignore generated files in output (implies --gen)")
	flags.IntVar(&processor.MinifiedGeneratedLineByteLength, "min-gen-line-length", 255, "number of bytes per average line for file to be considered minified or generated")
	flags.StringArrayVarP(&processor.Exclude, "not-match", "M", []string{}, "ignore files and directories matching regular expression")
	flags.StringVarP(&processor.FileOutput, "output", "o", "", "output filename (default stdout)")
	flags.StringVarP(&processor.SortBy, "sort", "s", "files", "column to sort by [files, name, lines, blanks, code, comments, complexity]")
	flags.BoolVarP(&processor.Trace, "trace", "t", false, "enable trace output (not recommended when processing multiple files)")
	flags.BoolVarP(&processor.Verbose, "verbose", "v", false, "verbose output")
	flags.BoolVarP(&processor.More, "wide", "w", false, "wider output with additional statistics (implies --complexity)")
	flags.BoolVar(&processor.NoLarge, "no-large", false, "ignore files over certain byte and line size set by max-line-count and max-byte-count")
	flags.BoolVar(&processor.IncludeSymLinks, "include-symlinks", false, "if set will count symlink files")
	flags.Int64Var(&processor.LargeLineCount, "large-line-count", 40000, "number of lines a file can contain before being removed from output")
	flags.Int64Var(&processor.LargeByteCount, "large-byte-count", 1000000, "number of bytes a file can contain before being removed from output")
	flags.StringVar(&processor.CountAs, "count-as", "", "count extension as language [e.g. jsp:htm,chead:\"C Header\" maps extension jsp to html and chead to C Header]")
	flags.StringVar(&processor.FormatMulti, "format-multi", "", "have multiple format output overriding --format [e.g. tabular:stdout,csv:file.csv,json:file.json]")
	flags.StringVar(&processor.RemapUnknown, "remap-unknown", "", "inspect files of unknown type and remap by checking for a string and remapping the language [e.g. \"-*- C++ -*-\":\"C Header\"]")
	flags.StringVar(&processor.RemapAll, "remap-all", "", "inspect every file and remap by checking for a string and remapping the language [e.g. \"-*- C++ -*-\":\"C Header\"]")
}
