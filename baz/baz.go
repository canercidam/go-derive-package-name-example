package baz

import "foo/bar"

func PackagePath() string {
	return bar.PackagePath()
}
