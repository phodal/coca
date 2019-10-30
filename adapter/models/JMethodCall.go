package models

type JMethodCall struct {
	Pkg               string
	Dlz               string
	MethodName        string
	StartLine         int
	StartLinePosition int
	StopLine          int
	StopLinePosition  int
}

