tracker
=======

Status: very much a work in progress.

Tiny time tracker written in Go that will report what you spent you time on 
based on the directories you cd into.

Installation
============

Save your `cd`s somewhere
-------------------------

Because `tracker` uses what directories you traverse during your workday, it 
needs to record all of your `cd`s as and when they happen. At the moment the 
easiest way for it to do that is for you to alias your `cd` command by adding
this to your `~/.bashrc`:

```bash
function cd() {                                                                   
    builtin cd "$1"                                                               
    NOW=$(date +"%Y-%m-%dT%T")                                                    
    echo "$NOW `pwd`" >> ~/.cd_history                                            
}
```

Add your projects to tracker
----------------------------

Currently they're hard-coded in `utils.go`. Add an item to `projects` in 
`GetNewProject`:

```go
    var projects = []Project{
        {Dir: filepath.Join(workspace, "some-project"), Name: "some-project"},
    }
```

Make the binary and give it a whirl
-----------------------------------

Check the project out in your Go src folder and you should be able to run 
`go build` in your `tracker` dir and then run `tracker` in your shell.