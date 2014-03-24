package main

import (
	"flag"
	"fmt"
	"os/exec"
	"os/user"
	"strings"
)

func main() {
	length := flag.Int("length", 50, "max length")
	flag.Parse()

	cmd := exec.Command("pwd")
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	directory := strings.TrimSpace(string(output))
	if strings.HasPrefix(directory, usr.HomeDir) {
		directory = "~" + directory[len(usr.HomeDir):]
	}

	if len(directory) <= *length {
		fmt.Println(directory)
		return
	}

	dirs := strings.Split(directory, "/")

	// keep the first segment
	result := []string{dirs[0]}
	dirs = dirs[1:]

	// figure out how far we can get starting at the other end
	var total int
	var startingIndex int
	for i := len(dirs)-1; i >= 0; i -- {
		total += len(dirs[i])
		if total > *length {
			startingIndex = i+1
			break
		}
	}

	if startingIndex >= len(dirs) {
		// weird. how did this happen?
		fmt.Println(directory)
		return
	}

	result = append(result, "...")
	for i := startingIndex; i < len(dirs); i++ {
		result = append(result, dirs[i])
	}

	fmt.Println(strings.Join(result, "/"))



}
