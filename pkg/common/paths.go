package common

import (
	"fmt"
	"gurusaranm0025/hyprone/pkg/utils"
)

var SCRIPTS_PATH = fmt.Sprintf("%s/.local/share/bin", utils.GetHomeDir())
var WALLS_PATH = fmt.Sprintf("%s/.HyprOne/Walls", utils.GetHomeDir())
