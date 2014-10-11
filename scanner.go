package aautomata

import (
	"bufio"
	"io"
)

type scanner struct {
	scanner *bufio.Scanner
	stack stack
	text *string
}

func newScanner() *scanner {
	return &scanner{
		nil,
		newStack(),
		nil,
	}
}

func (s *scanner) Init(r io.Reader) {
	s.scanner = bufio.NewScanner(r)
	s.scanner.Split(bufio.ScanRunes)
}

func (s *scanner) Scan() bool {
	if !s.stack.empty() {
		tmp := s.stack.pop();
		s.text = &tmp
		return true
	}

	return s.scanner.Scan()
}

func (s *scanner) Text() string {
	if s.text != nil {
		tmp := *s.text
		s.text = nil
		return tmp
	}

	return s.scanner.Text()
}

func (s *scanner) Push(i string) {
	s.stack.push(i)
}
