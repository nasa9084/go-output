package output

import (
	"errors"
	"flag"
	"io"
	"os"
	"sync"
)

// once for flag initialization
var once sync.Once

var (
	mu           sync.RWMutex
	openFileFlag             = os.O_RDWR | os.O_CREATE | os.O_TRUNC
	openFilePerm os.FileMode = 0644
)

var (
	long  string
	short string
)

func init() {
	once.Do(func() {
		flag.StringVar(&long, "output", "", "where output to")
		flag.StringVar(&short, "o", "", "short hand for -output")
	})
}

func UseFlagSet(flagset *flag.FlagSet) {
	flagset.StringVar(&long, "output", "", "where output to")
	flagset.StringVar(&short, "o", "", "short hand for -output")
}

func SetFileFlag(flg int) {
	mu.Lock()
	openFileFlag = flg
	mu.Unlock()
}

func SetFilePerm(perm os.FileMode) {
	mu.Lock()
	openFilePerm = perm
	mu.Unlock()
}

func Writer() (io.WriteCloser, error) {
	if long != "" && short != "" {
		return nil, errors.New("-o and -output are mutualy exclusive")
	}

	var output string

	if long == "" {
		if short != "" && short != "-" {
			output = short
		}
	} else {
		if long != "" && long != "-" {
			output = long
		}
	}

	if output == "" {
		return nopCloser{Writer: os.Stdout}, nil
	}

	mu.RLock()
	f, err := os.OpenFile(output, openFileFlag, openFilePerm)
	mu.RUnlock()
	return f, err
}

// nopCloser just wraps given io.Writer and provides Close method which does nothing.
// This wrapper is for providing transparent os.Stdout but not allow user to close it.
type nopCloser struct {
	io.Writer
}

func (nopCloser) Close() error { return nil }
