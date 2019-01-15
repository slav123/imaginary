package main

import "github.com/slav123/bimg"

// Version stores the current package semantic version
const Version = "1.0.15"

// Version represents the supported version
type Versions struct {
	ImaginaryVersion string `json:"imaginary"`
	BimgVersion      string `json:"bimg"`
	VipsVersion      string `json:"libvips"`
}

// CurrentVersions stores the current runtime system version metadata
var CurrentVersions = Versions{Version, bimg.Version, bimg.VipsVersion}
