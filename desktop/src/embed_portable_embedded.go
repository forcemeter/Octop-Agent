//go:build production && linux

package main

import _ "embed"

// Linux production builds embed the matching portable runtime so the release
// tar contains only the desktop binary.
//
//go:embed bundled/portable.zip
var embeddedPortable []byte
