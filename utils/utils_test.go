package utils

import (
	"testing"
	"time"
)

type getNewProjectTestPair struct {
	dir             string
	expectedProject string
}

type getTimeTestPair struct {
	line string
	time time.Time
}

type getDirTestPair struct {
	line string
	dir  string
}

func TestGetNewProject(t *testing.T) {
	var getNewProjectTests = []getNewProjectTestPair{
		{"/home/cmckinnel/workspace/nailsinc-us", "nailsinc"},
		{"/home/cmckinnel/workspace", "no-project"},
		{"/home/cmckinnel/workspace/go/src/github.com/chrismckinnel/tracker", "tracker"},
	}

	for _, pair := range getNewProjectTests {
		project := GetNewProject(pair.dir)
		if project != pair.expectedProject {
			t.Error(
				"For", pair.dir,
				"expected", pair.expectedProject,
				"got", project,
			)
		}
	}
}

func TestGetTime(t *testing.T) {
	var getTimeTests = []getTimeTestPair{
		{"2014-01-01T00:00:00 /home/cmckinnel/workspace/go/src/github.com/chrismckinnel/tracker",
			time.Date(2014, time.January, 1, 0, 0, 0, 0, time.UTC)},
		{"2014-02-01T07:12:07 /home/cmckinnel/workspace",
			time.Date(2014, time.February, 1, 7, 12, 7, 0, time.UTC)},
		{"2014-02-01T06:04:04 /home/cmckinnel/workspace/nailsinc-us/",
			time.Date(2014, time.February, 1, 6, 4, 4, 0, time.UTC)},
	}

	for _, pair := range getTimeTests {
		time := GetTime(pair.line)
		if time.Format(timeFormat) != pair.time.Format(timeFormat) {
			t.Error(
				"For", pair.line,
				"expected", pair.time.Format(timeFormat),
				"got", time.Format(timeFormat),
			)
		}
	}
}

func TestGetDir(t *testing.T) {
	var getDirTests = []getDirTestPair{
		{"01-01-2014T00:00:00 /home/cmckinnel/workspace/go/src/github.com/chrismckinnel/tracker",
			"/home/cmckinnel/workspace/go/src/github.com/chrismckinnel/tracker"},
		{"01-02-2014T07:12:07 /home/cmckinnel/workspace",
			"/home/cmckinnel/workspace"},
		{"01-02-2014T06:04:04 /home/cmckinnel/workspace/nailsinc-us/",
			"/home/cmckinnel/workspace/nailsinc-us/"},
	}

	for _, pair := range getDirTests {
		dir := GetDir(pair.line)
		if dir != pair.dir {
			t.Error(
				"For", pair.line,
				"expected", pair.dir,
				"got", dir,
			)
		}
	}
}
