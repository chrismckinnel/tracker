package main

import (
	"bufio"
	"fmt"
	"github.com/chrismckinnel/tracker/utils"
	"os"
	"time"
)

func main() {
	printSummary()
}

type ProjectTime struct {
	name      string
	timeSpent time.Duration
}

var projectMap = map[string]time.Duration{
	"nailsinc":   time.Since(time.Now()),
	"tracker":    time.Since(time.Now()),
	"no-project": time.Since(time.Now()),
}

func printSummary() {
	file, err := os.Open(utils.DirHistoryFile())
	utils.Check(err)
	scanner := bufio.NewScanner(file)
	firstLine := true
	var startTime time.Time
	var lastTime time.Time
	var lastProject string
	for scanner.Scan() {
		dirChangeTime := utils.GetTime(scanner.Text())
		if firstLine {
			startTime = dirChangeTime
			lastTime = dirChangeTime
			fmt.Println(fmt.Sprintf("First time entry: %s", startTime.String()))
			lastProject = "no-project"
			firstLine = false
			continue
		}
		timeElapsed := dirChangeTime.Sub(lastTime)
		dir := utils.GetDir(scanner.Text())
		project := utils.GetNewProject(dir)
		projectMap[lastProject] = projectMap[lastProject] + timeElapsed
		lastTime = dirChangeTime
		lastProject = project
	}

	for project, duration := range projectMap {
		fmt.Println(fmt.Sprintf("%s: %s", project, duration.String()))
	}
	err = scanner.Err()
	utils.Check(err)
}