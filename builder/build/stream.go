package build

import "strings"

type Stream struct {
	src   string
	index int
	len   int
}

func NewStream(in string) *Stream {
	in = fmtIn(in)
	return &Stream{
		src:   in,
		index: 0,
		len:   len(in),
	}
}

func (s *Stream) CurByte() byte {
	return s.src[s.index]
}

func (s *Stream) Next() {
	if s.index < s.len {
		s.index++
	} else {
		println("[WARNING] Over range!")
	}
}

func (s *Stream) SkipComma() {
	if ch := s.CurByte(); ch == ',' {
		s.index++
	}
}

func (s *Stream) isString() bool {
	ch := s.CurByte()
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z')
}

func (s *Stream) toString() string {
	return s.src[s.index:]
}

func fmtIn(in string) string {
	in = strings.Replace(in, "\n", "", -1)
	in = strings.Replace(in, " ", "", -1)
	in = strings.Replace(in, "\t", "", -1)
	return in
}
