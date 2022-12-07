package day07

import (
	"fmt"

	"github.com/domszyn/adventofcode/2022/mappers"
	"github.com/domszyn/adventofcode/2022/utils"
)

type File struct {
	Name string
	Size int
}

type Directory struct {
	Name        string
	Directories map[string]*Directory
	Files       []File
	Parent      *Directory
}

func (d *Directory) TotalSize() (size int) {
	for _, f := range d.Files {
		size += f.Size
	}

	for _, sd := range d.Directories {
		size += sd.TotalSize()
	}

	return
}

func (d *Directory) CountPart1() (size int) {
	if totalSize := d.TotalSize(); totalSize <= 100000 {
		size += totalSize
	}

	for _, sd := range d.Directories {
		size += sd.CountPart1()
	}

	return
}

func (d *Directory) FindLargeEnough(minSize int) (dirs []*Directory) {
	if d.TotalSize() > minSize {
		dirs = append(dirs, d)
	}

	for _, sd := range d.Directories {
		dirs = append(dirs, sd.FindLargeEnough(minSize)...)
	}

	return
}

func Solve() (part1 int, part2 int) {
	lines := utils.ReadInput("./solutions/day07/input.txt", mappers.ToString)

	root := Directory{
		"/",
		make(map[string]*Directory),
		make([]File, 0),
		nil,
	}

	var currentDir *Directory
	lsMode := false

	for _, l := range lines {
		if l[:4] == "$ cd" {
			lsMode = false
			dirName := l[5:]
			if dirName == ".." {
				currentDir = currentDir.Parent
			} else if dirName == "/" {
				currentDir = &root
			} else {
				currentDir = currentDir.Directories[dirName]
			}
		} else if l == "$ ls" {
			lsMode = true
		} else if lsMode {
			if l[:3] == "dir" {
				dirName := l[4:]
				dir := Directory{
					dirName,
					make(map[string]*Directory, 0),
					make([]File, 0),
					currentDir,
				}

				currentDir.Directories[dirName] = &dir
			} else {
				var size int
				var fileName string
				fmt.Sscanf(l, "%d %s", &size, &fileName)
				currentDir.Files = append(currentDir.Files, File{
					fileName,
					size,
				})
			}
		}
	}

	part1 = root.CountPart1()

	part2 = root.TotalSize()
	dirs := root.FindLargeEnough(part2 - 40000000)

	for _, d := range dirs {
		if ts := d.TotalSize(); ts < part2 {
			part2 = ts
		}
	}

	return
}
