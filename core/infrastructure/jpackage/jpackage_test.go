package jpackage

import "testing"

func TestGetMethodName(t *testing.T) {
	tests := []struct {
		name   string
		origin string
		want   string
	}{
		{"get package", "com.phodal.coca.JPackage.getMethodName", "getMethodName"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMethodName(tt.origin); got != tt.want {
				t.Errorf("GetMethodName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetClassName(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{"get package", "JPackage.getMethodName", "JPackage"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetClassName(tt.path); got != tt.want {
				t.Errorf("GetClassName() = %v, want %v", got, tt.want)
			}
		})
	}
}