package core

import (
	"github.com/forktopot/ksubdomain/pkg/core/conf"
	"github.com/forktopot/ksubdomain/pkg/core/gologger"
)

const banner = `
 _              _         _                       _       
| | _____ _   _| |__   __| | ___  _ __ ___   __ _(_)_ __  
| |/ / __| | | | '_ \ / _' |/ _ \| '_ ' _ \ / _| | | '_ \
|   <\__ \ |_| | |_) | (_| | (_) | | | | | | (_| | | | | |
|_|\_\___/\__,_|_.__/ \__,_|\___/|_| |_| |_|\__,_|_|_| |_|

`

func ShowBanner() {
	gologger.Printf(banner)
	gologger.Infof("Current Version: %s\n", conf.Version)
}
