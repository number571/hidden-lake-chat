package settings

import (
	"github.com/number571/hidden-lake/pkg/utils/appname"
)

var (
	gAppName = appname.LoadAppName(CProjectFullName)
)

func GetAppName() appname.IFmtAppName {
	return gAppName
}

const (
	CProjectName        = "chat"
	CProjectShortPrefix = "hlp"
	CProjectFullPrefix  = "hidden-lake-project"
)

const (
	CProjectFullName    = CProjectFullPrefix + "=" + CProjectName
	CProjectDescription = "console group chat"
)

const (
	CPathDB = CProjectShortPrefix + "-" + CProjectName + ".db"
)
