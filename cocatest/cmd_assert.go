package cocatest

// based on https://github.com/helm/helm
// license: apache 2.0
// just copy

import (
	"bytes"
	"flag"
	"github.com/pkg/errors"
	"io/ioutil"
	"path/filepath"
	"runtime"
)

// UpdateGolden writes out the Golden files with the latest values, rather than failing the test.
var updateGolden = flag.Bool("update", false, "update Golden files")

// TestingT describes a testing object compatible with the critical functions from the testing.T type
type TestingT interface {
	Fatal(...interface{})
	Fatalf(string, ...interface{})
	HelperT
}

// HelperT describes a test with a helper function
type HelperT interface {
	Helper()
}

// AssertGoldenBytes asserts that the give actual content matches the contents of the given filename
func AssertGoldenBytes(t TestingT, actual []byte, filename string) {
	t.Helper()

	if err := compare(actual, path(filename)); err != nil {
		t.Fatalf("%v", err)
	}
}

// AssertGoldenString asserts that the given string matches the contents of the given file.
func AssertGoldenString(t TestingT, actual, filename string) {
	t.Helper()

	if err := compare([]byte(actual), path(filename)); err != nil {
		t.Fatalf("%v", err)
	}
}

// AssertGoldenFile assers that the content of the actual file matches the contents of the expected file
func AssertGoldenFile(t TestingT, actualFileName string, expectedFilename string) {
	t.Helper()

	actual, err := ioutil.ReadFile(actualFileName)
	if err != nil {
		t.Fatalf("%v", err)
	}
	AssertGoldenBytes(t, actual, expectedFilename)
}

func path(filename string) string {
	if filepath.IsAbs(filename) {
		return filename
	}
	return filepath.Join("", filename)
}

func compare(actual []byte, filename string) error {
	if err := update(filename, actual); err != nil {
		return err
	}

	expected, err := ioutil.ReadFile(filename)
	if err != nil {
		return errors.Wrapf(err, "unable to read testdata %s", filename)
	}

	// TODO: fix path test for Windows...
	if runtime.GOOS == "windows" {
		return nil
	}

	if !bytes.Equal(expected, actual) {
		return errors.Errorf("does not match Golden file %s\n\nWANT:\n'%s'\n\nGOT:\n'%s'\n", filename, expected, actual)
	}
	return nil
}

func update(filename string, in []byte) error {
	if !*updateGolden {
		return nil
	}
	return ioutil.WriteFile(filename, normalize(in), 0666)
}

func normalize(in []byte) []byte {
	return bytes.Replace(in, []byte("\r\n"), []byte("\n"), -1)
}
