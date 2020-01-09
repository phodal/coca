package testcase

// CmdTestCase describes a test case that works with releases.
type CmdTestCase struct {
	Name      string
	Cmd       string
	Golden    string
	WantError bool
}

