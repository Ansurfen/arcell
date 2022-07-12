package build

import (
	"arcell/utils"
	"bufio"
	"os"
	"strings"
)

type Lexer struct {
	CMD  string
	ARG  map[string]string
	ENV  map[string]string
	MAP  map[string]string
	EXEC map[string]string
}

var (
	lines []string
	index int
)

func NewLexer() *Lexer {
	return &Lexer{
		CMD:  "",
		ARG:  make(map[string]string),
		ENV:  make(map[string]string),
		MAP:  make(map[string]string),
		EXEC: make(map[string]string),
	}
}

func (l *Lexer) Read(filename string) {
	fp, err := os.Open(filename)
	utils.Panic(err)
	defer fp.Close()
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := fmtLine(scanner.Text())
		if line != "" {
			lines = append(lines, line)
		}
	}
	l.Analyzer()
}

func (l *Lexer) ToString() {

}

func fmtLine(line string) string {
	line = strings.Replace(line, " ", "", -1)
	line = strings.Replace(line, "\n", "", -1)
	line = strings.Replace(line, "\t", "", -1)
	if i := strings.IndexByte(line, ';'); i >= 0 {
		line = line[:i]
	}
	return line
}

func (l *Lexer) Analyzer() {
	index = 0
	for index < len(lines) {
		switch lines[index] {
		case "[ENV]":
			l.ReadEnv()
		case "[CMD]":
			l.ReadCMD()
		case "[ARG]":
			l.ReadARG()
		case "[MAP]":
			l.ReadMAP()
		case "[EXEC]":
			l.ReadEXEC()
		default:
			panic("Fail to lexer conf")
		}
	}
}

func (l *Lexer) ReadEnv() {
	next()
	for index != len(lines) && lines[index][0] != '[' {
		env := strings.Split(lines[index], "=")
		l.ENV[env[0]] = env[1]
		next()
	}
}

func (l *Lexer) ReadCMD() {
	next()
	for index != len(lines) && lines[index][0] != '[' {
		l.CMD += lines[index]
		next()
	}
}

func (l *Lexer) ReadARG() {
	next()
	for index != len(lines) && lines[index][0] != '[' {
		arg := strings.Split(lines[index], "=")
		l.ARG[arg[0]] = arg[1]
		next()
	}
}

func (l *Lexer) ReadMAP() {
	next()
	for index != len(lines) && lines[index][0] != '[' {
		_map := strings.Split(lines[index], "=")
		l.MAP[_map[0]] = _map[1]
		next()
	}
}

func (l *Lexer) ReadEXEC() {
	next()
	for index != len(lines) && lines[index][0] != '[' {
		exec := strings.Split(lines[index], "=")
		l.MAP[exec[0]] = exec[1]
		next()
	}
}

func next() {
	if index < len(lines) {
		index++
	}
}
