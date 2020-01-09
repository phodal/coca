package support

import (
	"reflect"
	"testing"
)

func TestBuildMethodPackageInfo(t *testing.T) {
	tests := []struct {
		name string
		want *PackageClassInfo
	}{
		{
			"com.phodal.Coca.hello",
			&PackageClassInfo{"com.phodal", "Coca", "hello"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildMethodPackageInfo(tt.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildMethodPackageInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}