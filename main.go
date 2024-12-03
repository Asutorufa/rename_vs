package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Entry struct {
	Video  string
	Ass    string
	NewAss string
}

func main() {
	if len(os.Args) < 3 {
		panic("usage: rename <old(video name)> <new(rss name)> [max]")
	}

	max := 12
	if len(os.Args) > 3 {
		var err error
		max, err = strconv.Atoi(os.Args[3])
		if err != nil {
			panic(err)
		}
	}

	fmt.Println(os.Args[1], os.Args[2], max)

	reg1 := Compile(max, os.Args[1])
	reg2 := Compile(max, os.Args[2])

	files, err := os.ReadDir(".")
	if err != nil {
		panic(err)
	}

	entries := make([]*Entry, max)
	for i := range max {
		entries[i] = &Entry{}
	}

	for _, v := range files {
		if v.IsDir() {
			continue
		}

		i := reg1.match(v.Name())
		if i != -1 {
			entries[i].Video = v.Name()
			continue
		}

		i = reg2.match(v.Name())
		if i != -1 {
			entries[i].Ass = v.Name()
		}
	}

	for _, v := range entries {
		if v.Video == "" || v.Ass == "" {
			continue
		}

		i := strings.IndexByte(v.Video, '.')
		if i == -1 {
			continue
		}

		prefix := v.Video[:i]

		i2 := strings.IndexByte(v.Ass, '.')
		if i2 == -1 {
			continue
		}

		fmt.Println("pair", v.Video, v.Ass)

		suffix := v.Ass[i2:]

		v.NewAss = prefix + suffix
		fmt.Println("rename", v.Ass, prefix+suffix)
	}

	fmt.Print("yes/no: ")

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	if scan.Text() != "yes" {
		return
	}

	for _, v := range entries {
		if v.Video == "" || v.Ass == "" || v.NewAss == "" {
			continue
		}

		fmt.Printf("mv \"%s\" \"%s\"\n", v.Ass, v.NewAss)
		err = os.Rename(v.Ass, v.NewAss)
		if err != nil {
			log.Println(err)
		}
	}
}

type regs []*regexp.Regexp

func (r regs) match(s string) int {
	for i, v := range r {
		if v.MatchString(s) {
			return i
		}
	}
	return -1
}

func Compile(max int, str string) regs {
	regs := make(regs, max)
	for i := range max {
		var err error
		str := strings.ReplaceAll(fmt.Sprintf(str, i+1), "]", "\\]")
		str = strings.ReplaceAll(str, "[", "\\[")
		regs[i], err = regexp.Compile(str)
		if err != nil {
			panic(err)
		}
	}

	return regs
}
