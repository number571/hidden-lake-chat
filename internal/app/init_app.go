package app

import (
	"errors"
	"strings"

	"github.com/number571/go-peer/pkg/types"
	"github.com/number571/hidden-lake/pkg/utils/build"
)

// initApp work with the raw data = read files, read args
func InitApp(pPath, pNetwork string) (types.IRunner, error) {
	inputPath := strings.TrimSuffix(pPath, "/")
	_, err := build.SetBuildByPath(inputPath)
	if err != nil {
		return nil, errors.Join(ErrSetBuild, err)
	}
	return NewApp(pNetwork, inputPath), nil
}
