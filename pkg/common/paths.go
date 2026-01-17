package common

import (
	"gurusaranm0025/hyprone/pkg/utils"
	"path/filepath"
)

var SCRIPTS_DIR_PATH = filepath.Join(utils.GetHomeDir(), ".local/share/bin")
var CONFIG_DIR_PATH = filepath.Join(utils.GetHomeDir(), ".config/hyprone")
var ALL_WALLS_DIR_PATH = filepath.Join(CONFIG_DIR_PATH, "walls")
var CURRENT_WALL_DIR_PATH = filepath.Join(CONFIG_DIR_PATH, "current_wall")
var GIT_CLONE_DIR_PATH = filepath.Join(utils.GetHomeDir(), ".config/hyprone/.temp/git_clone")

var HYPR01_CONFIG_PATH = filepath.Join(CONFIG_DIR_PATH, "config.json")

var PlaceholderValues = map[string]string{
	"${CURRENT_WALL_DIR_PATH}": CURRENT_WALL_DIR_PATH,
	"${CONFIG_DIR_PATH}":       CONFIG_DIR_PATH,
	"${ALL_WALLS_DIR_PATH}":    ALL_WALLS_DIR_PATH,
	"${SCRIPTS_DIR_PATH}":      SCRIPTS_DIR_PATH,
	"${GIT_CLONE_DIR_PATH}":    GIT_CLONE_DIR_PATH,
}
