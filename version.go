package main

import (
	"fmt"
	"runtime"
)

// Version information - this will be updated automatically by GitHub Actions
const (
	VERSION = "0.1.1"
)

// VersionInfo holds extended version information
type VersionInfo struct {
	Version   string
	GoVersion string
	OS        string
	Arch      string
}

// GetVersionInfo returns detailed version information
func GetVersionInfo() VersionInfo {
	return VersionInfo{
		Version:   VERSION,
		GoVersion: runtime.Version(),
		OS:        runtime.GOOS,
		Arch:      runtime.GOARCH,
	}
}

// PrintVersion prints formatted version information
func PrintVersion() {
	info := GetVersionInfo()
	fmt.Printf("Contextor %s\n", info.Version)
	fmt.Printf("  Go version: %s\n", info.GoVersion)
	fmt.Printf("  OS/Arch: %s/%s\n", info.OS, info.Arch)
}
