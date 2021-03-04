package cmd_test

import (
	"errors"
	"testing"

	"github.com/JosiahWitt/ensure"
	"github.com/JosiahWitt/ensure-cli/internal/cmd"
	"github.com/JosiahWitt/ensure-cli/internal/ensurefile"
	"github.com/JosiahWitt/ensure-cli/internal/mocks/mock_ensurefile"
	"github.com/JosiahWitt/ensure-cli/internal/mocks/mock_mockgen"
	"github.com/JosiahWitt/ensure/ensurepkg"
	"github.com/golang/mock/gomock"
)

func TestGenerateMocks(t *testing.T) {
	ensure := ensure.New(t)

	type Mocks struct {
		EnsureFileLoader *mock_ensurefile.MockLoaderIface
		MockGen          *mock_mockgen.MockGeneratorIface
	}

	exampleError := errors.New("something went wrong")
	defaultWd := func() (string, error) {
		return "/test", nil
	}

	table := []struct {
		Name          string
		ExpectedError error
		Flags         []string

		Getwd      func() (string, error)
		Mocks      *Mocks
		SetupMocks func(*Mocks)
		Subject    *cmd.App
	}{
		{
			Name:  "with valid execution",
			Getwd: defaultWd,
			SetupMocks: func(m *Mocks) {
				m.EnsureFileLoader.EXPECT().
					LoadConfig("/test").
					Return(&ensurefile.Config{
						RootPath: "/some/root/path",
					}, nil)

				m.MockGen.EXPECT().
					GenerateMocks(&ensurefile.Config{
						RootPath: "/some/root/path",
					}, gomock.Any()).
					Return(nil)
			},
		},

		{
			Name:  "with valid execution: disabled parallel generation",
			Flags: []string{"--disable-parallel"},
			Getwd: defaultWd,
			SetupMocks: func(m *Mocks) {
				m.EnsureFileLoader.EXPECT().
					LoadConfig("/test").
					Return(&ensurefile.Config{
						RootPath: "/some/root/path",
					}, nil)

				m.MockGen.EXPECT().
					GenerateMocks(&ensurefile.Config{
						RootPath:                  "/some/root/path",
						DisableParallelGeneration: true,
					}, gomock.Any()).
					Return(nil)
			},
		},

		{
			Name:          "when error loading working directory",
			Getwd:         func() (string, error) { return "", exampleError },
			ExpectedError: exampleError,
		},

		{
			Name:          "when cannot load config",
			Getwd:         defaultWd,
			ExpectedError: exampleError,
			SetupMocks: func(m *Mocks) {
				m.EnsureFileLoader.EXPECT().LoadConfig("/test").Return(nil, exampleError)
			},
		},

		{
			Name:          "when cannot generate mocks",
			Getwd:         defaultWd,
			ExpectedError: exampleError,
			SetupMocks: func(m *Mocks) {
				m.EnsureFileLoader.EXPECT().
					LoadConfig("/test").
					Return(&ensurefile.Config{
						RootPath: "/some/root/path",
					}, nil)

				m.MockGen.EXPECT().
					GenerateMocks(&ensurefile.Config{
						RootPath: "/some/root/path",
					}, gomock.Any()).
					Return(exampleError)
			},
		},
	}

	ensure.RunTableByIndex(table, func(ensure ensurepkg.Ensure, i int) {
		entry := table[i]
		entry.Subject.Getwd = entry.Getwd

		err := entry.Subject.Run(append([]string{"ensure", "generate", "mocks"}, entry.Flags...))
		ensure(err).IsError(entry.ExpectedError)
	})
}
