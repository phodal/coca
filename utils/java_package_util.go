package utils

import (
	. "./models"
	"strings"
)

func BuildMethodPackageInfo(name string) *PackageClassInfo {
	pkgInfo := &PackageClassInfo{"", "", ""}
	split := strings.Split(name, ".")

	pkgInfo.Method = split[len(split)-1]
	pkgInfo.Class = split[len(split)-2]
	pkgInfo.Package = strings.Join(split[:len(split)-3], ".")
	return pkgInfo
}
