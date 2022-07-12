package build

import (
	"arcell/utils"
	"errors"
	"strings"
)

func ParserCmd(s *Stream, cmd *Command) {
	if s.CurByte() != '{' {
		utils.Panic(errors.New("Unknown syntax"))
	}
	s.Next()
	parserHelper(s, cmd)
}

func parserHelper(s *Stream, cmd *Command) {
	for s.index < s.len {
		s.SkipComma()
		if s.CurByte() == '}' {
			return
		}
		switch s.CurByte() {
		case '-':
			utils.Log("Parser flag")
			ParserFlag(s, cmd)
		case '{':
			utils.Log("Parser block")
			ParserBlock(s, cmd)
		case '@':
			utils.Log("Parser flagfn")
			ParserFlagFn(s, cmd)
		default:
			if s.isString() {
				utils.Log("Parser cmd")
			} else {
				utils.Panic(errors.New("A error occur at Parser()"))
			}
		}
	}
}

func ParserFlag(s *Stream, cmd *Command) {
	if s.CurByte() != '-' {
		utils.Panic(errors.New("Fail to verify flag"))
	}
	s.Next()
	fname := ""
	ftype := ""
	hasCommda := false
	for s.CurByte() != ',' {
		if s.CurByte() == ':' {
			s.Next()
			hasCommda = true
		}
		if !hasCommda {
			fname += string(s.CurByte())
		} else {
			ftype += string(s.CurByte())
		}
		s.Next()
	}
	switch strings.ToLower(ftype) {
	case "bool":
		cmd.flags[fname] = BOOL
	case "string":
		cmd.flags[fname] = STRING
	case "int":
		cmd.flags[fname] = INT
	default:
		utils.Panic(errors.New("Unknown flagType"))
	}
	s.Next()
}

func ParserFlagFn(s *Stream, cmd *Command) {

}

func ParserBlock(s *Stream, cmd *Command) {
	if s.CurByte() != '{' {
		utils.Panic(errors.New("Fail to match block"))
	}
	s.Next()
	str := ""
	for s.CurByte() != '}' {
		if s.CurByte() == ',' {
			cmd.script = append(cmd.script, str)
			str = ""
			s.Next()
			continue
		}
		str += string(s.CurByte())
		s.Next()
	}
	if len(str) > 0 {
		cmd.script = append(cmd.script, str)
	}
	s.Next()
}

func ParserSub(s *Stream, cmd *Command) {
	if !s.isString() {
		utils.Panic(errors.New("Fail to match string"))
	}
	cname := ""
	hasDomain := false
	for {
		if s.CurByte() == '{' {
			hasDomain = true
		}
		if s.CurByte() == '}' {
			break
		}
		if s.isString() && !hasDomain {
			cname += string(s.CurByte())
		} else if hasDomain {
			cmd.sub[cname] = NewCommand(cname)
		}
		parserHelper(s, cmd.sub[cname])
		s.Next()
	}
	s.Next()
}
