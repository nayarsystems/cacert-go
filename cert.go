package cacert

import (
	_ "embed"
)

//go:embed cacert.pem
var CaCert []byte
