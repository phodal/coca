package models

type JMethod struct {
	Name              string
	Type              string
	StartLine         int
	StartLinePosition int
	StopLine          int
	StopLinePosition  int
	Parameters        []JParameter
}

type JParameter struct {
	Name string
	Type string
}

type JMethodInfo struct {
	Name       string
	Type       string
	Parameters []JParameter
	Length  	string
}
