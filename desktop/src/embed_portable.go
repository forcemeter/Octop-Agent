//go:build !production || darwin || windows

package main

// Development builds look for Octop-<plat>.zip beside the executable.
// macOS production keeps the zip in the signed .app Resources directory.
// Windows production ships the zip next to Octop.exe via the NSIS installer.
var embeddedPortable []byte
