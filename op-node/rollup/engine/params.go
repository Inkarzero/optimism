package engine

import "time"

const (
	buildSealTimeout      = time.Second * 100
	buildStartTimeout     = time.Second * 100
	buildCancelTimeout    = time.Second * 100
	payloadProcessTimeout = time.Second * 100
)
