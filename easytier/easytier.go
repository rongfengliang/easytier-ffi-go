package easytier

import (
	"os"
	"os/exec"

	"github.com/ebitengine/purego"
)

type KeyValuePair struct {
	Key   uintptr
	Value uintptr
}

var (
	GetErrorMsg           func(uintptr)
	FreeString            func(uintptr)
	ParseConfig           func(uintptr) int32
	RunNetworkInstance    func(uintptr) int32
	RetainNetworkInstance func(uintptr, uintptr) int32
	CollectNetworkInfos   func(*KeyValuePair, uintptr) int32
)

func findLibrary() string {
	// Env var
	if envPath := os.Getenv("EASYTIER_FFI_LIB_PATH"); envPath != "" {
		return envPath
	}

	// ldconfig with Linux
	if path, err := exec.LookPath("libeasytier_ffi.so"); err == nil {
		return path
	}

	// default path
	commonPaths := []string{
		"/usr/local/lib/libeasytier_ffi.so",
		"./lib/libeasytier_ffi.dylib",
		"./lib/libeasytier_ffi.so",
		"./lib/libeasytier_ffi.dll",
		"/usr/lib/libeasytier_ffi.so",
		"/opt/homebrew/lib/libeasytier_ffi.dylib",
	}

	for _, p := range commonPaths {
		if _, err := os.Stat(p); err == nil {
			return p
		}
	}

	return "libeasytier_ffi.so"
}
func init() {
	path := findLibrary()
	libeasytier, err := purego.Dlopen(path, purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	purego.RegisterLibFunc(&GetErrorMsg, libeasytier, "get_error_msg")
	purego.RegisterLibFunc(&FreeString, libeasytier, "free_string")
	purego.RegisterLibFunc(&ParseConfig, libeasytier, "parse_config")
	purego.RegisterLibFunc(&RunNetworkInstance, libeasytier, "run_network_instance")
	purego.RegisterLibFunc(&RetainNetworkInstance, libeasytier, "retain_network_instance")
	purego.RegisterLibFunc(&CollectNetworkInfos, libeasytier, "collect_network_infos")
}
