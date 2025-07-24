package app

const (
	errPrefix = "internal/kernel/pkg/app = "
)

type SAppError struct {
	str string
}

func (err *SAppError) Error() string {
	return errPrefix + err.str
}

var (
	ErrSetBuild = &SAppError{"set build"}
)
