package conf

import "os"

var HomeDirPath, _ = os.UserHomeDir()

var AutoCPUFreqConfPath = "/etc/auto-cpufreq.conf"
