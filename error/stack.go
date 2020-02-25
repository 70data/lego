package rerror

import (
	"errors"
	"fmt"
	"io"
	"log"
	"path"
	"runtime"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

type RError struct {
	cause error
	msg   string
	*stack
}

type stackTracer interface {
	StackTrace() StackTrace
}

// Frame represents a program counter inside a stack frame.
type Frame uintptr

// StackTrace is stack of Frames from innermost (newest) to outermost (oldest).
type StackTrace []Frame

// stack represents a stack of program counters.
type stack []uintptr

type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var DEBUG bool

func (r *RError) Cause() error { return r.cause }

func (r *RError) Msg() string { return r.msg }

func (r RError) Error() string {
	return r.msg + ":" + r.cause.Error()
}

func (r *RError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			_, _ = fmt.Fprintf(s, "%+v", r.Cause())
			r.stack.Format(s, verb)
			return
		}
		fallthrough
	case 's':
		_, _ = fmt.Fprintf(s, "%+v\n", r.Cause())
		_, _ = io.WriteString(s, r.Error())
	case 'q':
		_, _ = fmt.Fprintf(s, "%q", r.Error())
		_, _ = io.WriteString(s, r.Error())
	}
}

func callers() *stack {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(3, pcs[:])
	var st stack = pcs[0:n]
	return &st
}

func NewRErr(message string) error {
	err := errors.New(message)
	return RError{
		err,
		message,
		callers(),
	}
}

func NewRErrf(format string, args ...interface{}) error {
	err := errors.New(fmt.Sprintf(format, args...))
	return RError{
		err,
		fmt.Sprintf(format, args...),
		callers(),
	}
}

func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}
	return RError{
		err,
		message,
		callers(),
	}
}

func Wrapf(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return RError{
		err,
		fmt.Sprintf(format, args...),
		callers(),
	}
}

func HandleError(err error) string {
	key := bson.NewObjectId().Hex()
	log.SetPrefix(fmt.Sprintf("[logID: %v]: ", key))
	log.Printf("%+v", err)
	if err, ok := err.(stackTracer); DEBUG && ok {
		for _, f := range err.StackTrace() {
			log.Printf("%+s:%d,func %n ", f, f, f)
		}
	}
	return key
}

func NewCodeError(code int, msg string) *CodeError {
	return &CodeError{code, msg}
}

func NewCodeErrorf(code int, format string, args ...interface{}) *CodeError {
	return &CodeError{code, fmt.Sprintf(format, args...)}
}

func (e CodeError) Error() string {
	return fmt.Sprintf("%d:%s", e.Code, e.Msg)
}

// pc returns the program counter for this frame;
// multiple frames may have the same PC value.
func (f Frame) pc() uintptr { return uintptr(f) - 1 }

// file returns the full path to the file that contains the
// function for this Frame's pc.
func (f Frame) file() string {
	fn := runtime.FuncForPC(f.pc())
	if fn == nil {
		return "unknown"
	}
	file, _ := fn.FileLine(f.pc())
	return file
}

// line returns the line number of source code of the
// function for this Frame's pc.
func (f Frame) line() int {
	fn := runtime.FuncForPC(f.pc())
	if fn == nil {
		return 0
	}
	_, line := fn.FileLine(f.pc())
	return line
}

// funcname removes the path prefix component of a function's name reported by func.Name().
func funcname(name string) string {
	i := strings.LastIndex(name, "/")
	name = name[i+1:]
	i = strings.Index(name, ".")
	return name[i+1:]
}

// Format formats the frame according to the fmt.Formatter interface.
//
//    %s    source file
//    %d    source line
//    %n    function name
//    %v    equivalent to %s:%d
//
// Format accepts flags that alter the printing of some verbs, as follows:
//
//    %+s   function name and path of source file relative to the compile time
//          GOPATH separated by \n\t (<funcname>\n\t<path>)
//    %+v   equivalent to %+s:%d
func (f Frame) Format(s fmt.State, verb rune) {
	switch verb {
	case 's':
		switch {
		case s.Flag('+'):
			pc := f.pc()
			fn := runtime.FuncForPC(pc)
			if fn == nil {
				_, _ = io.WriteString(s, "unknown")
			} else {
				file, _ := fn.FileLine(pc)
				_, _ = fmt.Fprintf(s, "%s\n\t%s", fn.Name(), file)
			}
		default:
			_, _ = io.WriteString(s, path.Base(f.file()))
		}
	case 'd':
		_, _ = fmt.Fprintf(s, "%d", f.line())
	case 'n':
		name := runtime.FuncForPC(f.pc()).Name()
		_, _ = io.WriteString(s, funcname(name))
	case 'v':
		f.Format(s, 's')
		_, _ = io.WriteString(s, ":")
		f.Format(s, 'd')
	}
}

// Format formats the stack of Frames according to the fmt.Formatter interface.
//
//    %s	lists source files for each Frame in the stack
//    %v	lists the source file and line number for each Frame in the stack
//
// Format accepts flags that alter the printing of some verbs, as follows:
//
//    %+v   Prints filename, function, and line number for each Frame in the stack.
func (st StackTrace) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		switch {
		case s.Flag('+'):
			for _, f := range st {
				_, _ = fmt.Fprintf(s, "\n%+v", f)
			}
		case s.Flag('#'):
			_, _ = fmt.Fprintf(s, "%#v", []Frame(st))
		default:
			_, _ = fmt.Fprintf(s, "%v", []Frame(st))
		}
	case 's':
		_, _ = fmt.Fprintf(s, "%s", []Frame(st))
	}
}

func (s *stack) Format(st fmt.State, verb rune) {
	switch verb {
	case 'v':
		switch {
		case st.Flag('+'):
			for _, pc := range *s {
				f := Frame(pc)
				_, _ = fmt.Fprintf(st, "\n%+v", f)
			}
		}
	}
}

func (s *stack) StackTrace() StackTrace {
	f := make([]Frame, len(*s))
	for i := 0; i < len(f); i++ {
		f[i] = Frame((*s)[i])
	}
	return f
}
