package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"
)

const timeFormat = "2006-01-02 15:04:05"

type Project struct {
	Dir  string
	Name string
}

func GetNewProject(dir string) string {
	var workspace = filepath.Join(currentUser().HomeDir, "workspace")
	var projects = []Project{
		{Dir: filepath.Join(workspace, "go/src/github.com/chrismckinnel/tracker"), Name: "tracker"},
		{Dir: filepath.Join(workspace, "nailsinc-us"), Name: "nailsinc"},
	}

	for _, project := range projects {
		if strings.Contains(dir, project.Dir) {
			return project.Name
		}
	}
	return "no-project"
}

func GetTime(line string) time.Time {
	timeString := strings.Split(line, " ")[0]
	timeString = strings.Replace(timeString, "T", " ", -1)
	t, err := time.Parse(timeFormat, timeString)
	Check(err)
	return t
}

func GetDir(line string) string {
	return strings.Split(line, " ")[1]
}

func PrintDirHistory() {
	file, err := os.Open(DirHistoryFile())
	Check(err)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	err = scanner.Err()
	Check(err)
}

func DirHistoryFile() string {
	return filepath.Join(currentUser().HomeDir, ".cd_history")
}

func currentUser() *user.User {
	usr, err := user.Current()
	Check(err)
	return usr
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
