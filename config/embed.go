package config

import "embed"

//go:embed *.toml
var Embed embed.FS
