package logs

import (
	"errors"
	"fmt"
	"io"
)

type ParseError struct {
	msg string
	err error
}

func (e ParseError) Error() string { return e.msg }

func (e ParseError) Unwrap() error { return e.err }

func IsParseError(err error) bool { var v *ParseError; return errors.As(err, &v) }

type (
	LogLine interface {
		Assign(name string, value string) error
	}

	Parser interface {
		ReadLine(LogLine) error
		Parse(row []byte, line LogLine) error
		Info() string
	}
)

const (
	TypeCSV    = "csv"
	TypeLTSV   = "ltsv"
	TypeRegExp = "regexp"
	TypeJSON   = "json"
)

type ParserConfig struct {
	LogType string       `yaml:"log_type"`
	CSV     CSVConfig    `yaml:"csv_config"`
	LTSV    LTSVConfig   `yaml:"ltsv_config"`
	RegExp  RegExpConfig `yaml:"regexp_config"`
	JSON    JSONConfig   `yaml:"json_config"`
}

func NewParser(config ParserConfig, in io.Reader) (Parser, error) {
	switch config.LogType {
	case TypeCSV:
		return NewCSVParser(config.CSV, in)
	case TypeLTSV:
		return NewLTSVParser(config.LTSV, in)
	case TypeRegExp:
		return NewRegExpParser(config.RegExp, in)
	case TypeJSON:
		return NewJSONParser(config.JSON, in)
	default:
		return nil, fmt.Errorf("invalid type: %q", config.LogType)
	}
}
