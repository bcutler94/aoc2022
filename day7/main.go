package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	TotalAvailableSpace := 70000000
	UnusedMinSpace := 30000000
	fs := BuildFileSystem()
	TotalAvailableSpace -= fs["/"].CalculateSize()
	fmt.Println(fs["/"].FindSpace(TotalAvailableSpace, UnusedMinSpace))
}

type FS map[string]*Dir

func (dir *Dir) FindSpace(totalAvailableSpace, unusedMinSpace int) int {
	candidate := totalAvailableSpace
	for fskey := range dir.Fs {
		if !dir.Fs[fskey].File {
			if totalAvailableSpace + dir.Fs[fskey].Size >= unusedMinSpace && dir.Fs[fskey].Size < candidate {
				candidate = dir.Fs[fskey].Size
			}
			possibleCandidate := dir.Fs[fskey].FindSpace(totalAvailableSpace, unusedMinSpace)
			if possibleCandidate < candidate {
				candidate = possibleCandidate
			}
		} 
	}	
	return candidate
}

func (dir *Dir) CalculateSize() int {
	size := 0
	for fskey := range dir.Fs {
		if dir.Fs[fskey].File {
			size += dir.Fs[fskey].Size
		} else {
			size += dir.Fs[fskey].CalculateSize()
		}
	}	
	dir.Size = size
	return size
}

func (dir *Dir) TotalSize(maxSize int) int {
	size := 0
	for fskey := range dir.Fs {
		if !dir.Fs[fskey].File {
			if maxSize >= dir.Fs[fskey].Size {
				size += dir.Fs[fskey].Size
			}
			size += dir.Fs[fskey].TotalSize(maxSize)
		}
	}	
	return size
}

type Dir struct {
	Parent *Dir
	Size int;
	Name string;
	File bool;
	Fs FS;
}

type Tracker struct {
	CurrentDir *Dir
}



func BuildFileSystem() FS {
	fd, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fd)
	scanner.Split(bufio.ScanLines)

	fs := make(FS)
	homeDir := Dir{
		Name: "/",
		Size: 0,
		File: false,
		Fs: make(FS),
	}
	fs["/"] = &homeDir
	tracker := new(Tracker)
	tracker.CurrentDir = &homeDir


	totalSize := 0
	for scanner.Scan() {
		text := scanner.Text()
		args := strings.Split(text, " ")

		switch (args[0]) {
		case "$":
			// Handle cmd
			switch (args[1]) {
			case "cd":
				// Handle CD
				switch (args[2]) {
				case "/":
					tracker.CurrentDir = fs["/"]
				case "..":
					tracker.CurrentDir = tracker.CurrentDir.Parent
				default:
					tracker.CurrentDir = tracker.CurrentDir.Fs[args[2]]
				}
			case "ls":
				// Handle LS - I dont think we do shit
			}
		case "dir":
			// Add dir
			newDir := Dir{
				Name: args[1],
				Size: 0,
				Parent: tracker.CurrentDir,
				File: false,
				Fs: FS{},
			}
			tracker.CurrentDir.Fs[args[1]] = &newDir
		default:
			// Add file
			size, err := strconv.Atoi(args[0])
			if err != nil {
				panic("Cannot get size of file")
			}
			totalSize += size
			newFile := Dir{
				Name: args[1],
				Size: size,
				Parent: tracker.CurrentDir,
				File: true,
			}
			tracker.CurrentDir.Fs[args[1]] = &newFile
		}

	}
	return fs
}
