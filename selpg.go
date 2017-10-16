package main

import (
	"bufio"
	"errors"
	"flag"
	"io"
	"os"
	"strconv"
)

func errorHandle(s, e, l *int, f *bool, d *string) (err error) {
	if *s <= 0 {
		err = errors.New("the index of the first page should be larger than 0")
	}
	if *s > *e {
		err = errors.New("the index of the first page should not be larger than the last one")
	}
	if *f == true && *l != 72 {
		err = errors.New("you should not specify -f and -l at the same time")
	}
	return
}

func parseFlag() (s, e, l *int, f *bool, d *string) {
	cmdl := flag.CommandLine
	s = cmdl.Int("s", 1, "the first page selected")
	e = cmdl.Int("e", 1, "the last page selected")
	l = cmdl.Int("l", 72, "the capacity of a page [default: 72 lines]")
	d = cmdl.String("d", "", "the target destiny")
	f = cmdl.Bool("f", false, "whether the page is separeted by \\f?")

	cmdl.Parse(os.Args[1:])
	if w := errorHandle(s, e, l, f, d); w != nil {
		os.Stderr.WriteString(w.Error() + "\n")
		// os.Stderr.WriteString("\n")
		cmdl.PrintDefaults()
		// os.Stderr.WriteString(w.Error())
		os.Exit(1)
	}
	return

}

func ouputByDelimiter(rd io.Reader, delimiter byte, e, s, l *int) {
	bufRd := bufio.NewReader(rd)
	linesNum := 1
	pageNum := 1
	for {
		str, err := bufRd.ReadString(delimiter)
		if err != nil {
			if err != io.EOF {
				os.Stderr.WriteString(err.Error())
				os.Exit(1)
			}
			break
		}
		if pageNum > *e {
			break
		}
		if pageNum >= *s && pageNum <= *e {
			os.Stdout.WriteString(str)
			// os.Stdout.Write([]byte{'\n'})
		}
		pageNum = linesNum / *l + 1
		linesNum++
	}
}

func selpg() {
	s, e, l, f, _ := parseFlag()
	var delimiter byte = '\n'
	if *f {
		delimiter = '\f'
	}
	files := flag.Args()
	if len(files) != 0 {
		for index, filename := range files {
			file, err := os.OpenFile(filename, os.O_RDONLY, os.ModeType)
			defer file.Close()
			if err != nil {
				os.Stderr.WriteString(err.Error())
				os.Exit(1)
			}
			// bufRd := bufio.NewReader(file)
			var tip string = "\nthis is file " + strconv.Itoa(index+1) + " :\n"
			if index == 0 {
				tip = "this is file " + strconv.Itoa(index+1) + " :\n"
			}
			os.Stdout.WriteString(tip)
			ouputByDelimiter(file, delimiter, e, s, l)

		}

	} else {
		ouputByDelimiter(os.Stdin, delimiter, e, s, l)

	}
}

func main() {
	selpg()
}
