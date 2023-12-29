package version

import (
	"encoding/json"
	"fmt"
	"runtime"

	"github.com/gosuri/uitable"
)

var (
	GitBranch = "master"
	BuildDate = "1970-01-01T00:00:00Z"
	GitCommit = "$Format:%H$"
)

type Info struct {
	GitBranch string `json:"gitBranch"`
	GitCommit string `json:"gitCommit"`
	BuildDate string `json:"buildDate"`
	GoVersion string `json:"goVersion"`
	Compiler  string `json:"compiler"`
	Platform  string `json:"platform"`
}

func (info Info) String() string {
	if s, err := info.Text(); err == nil {
		return string(s)
	}

	return info.GitBranch
}

func (info Info) Text() ([]byte, error) {
	table := uitable.New()

	table.RightAlign(0)
	table.MaxColWidth = 80
	table.Separator = " "

	table.AddRow("GitBranch:", info.GitBranch)
	table.AddRow("GitCommit:", info.GitCommit)
	table.AddRow("BuildDate:", info.BuildDate)
	table.AddRow("GoVersion:", info.GoVersion)
	table.AddRow("Compiler:", info.Compiler)
	table.AddRow("Platform:", info.Platform)

	return table.Bytes(), nil
}

func (info Info) ToJSON() string {
	s, _ := json.Marshal(info)

	return string(s)
}

func Get() Info {
	return Info{
		GitBranch: GitBranch,
		GitCommit: GitCommit,
		BuildDate: BuildDate,
		GoVersion: runtime.Version(),
		Compiler:  runtime.Compiler,
		Platform:  fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
