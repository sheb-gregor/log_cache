package config

import "github.com/lancer-kit/uwe/v2"

const ServiceName = "log_cache"

// The variables are set when compiling the binary and used to make sure which version of the backend is running.
// Example: go build -ldflags "-X log_cache/config.version=$VERSION\
// -X log_cache/config.build=$COMMIT \
// -X log_cache/config.tag=$TAG" .

// nolint:gochecknoglobals
var (
	version = "n/a"
	build   = "n/a"
	tag     = "n/a"
)

func AppInfo() uwe.AppInfo {
	return uwe.AppInfo{
		Name:    ServiceName,
		Version: version,
		Build:   build,
		Tag:     tag,
	}
}
