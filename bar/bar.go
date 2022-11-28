package bar

import (
	"runtime"
)

func PackagePath() string {
	pc, _, _, _ := runtime.Caller(1)
	serviceName := runtime.FuncForPC(pc).Name()

	return serviceName
}
