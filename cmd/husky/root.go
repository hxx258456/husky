package main

import (
	"context"

	"github.com/hxx258456/husky/internal/version"
	"github.com/hxx258456/husky/pkg/husky"
	"github.com/hxx258456/husky/pkg/log"
	"github.com/spf13/cobra"
)

var (
	logger      = log.Logger
	loggerV     = 0
	projectRoot = husky.ResolveGitRoot()
	theHusky    *husky.Husky
)

var CmdRoot = &cobra.Command{
	Use:     "husky",
	Short:   "husky " + version.Version,
	Version: version.Version,
}

func init() {
	CmdRoot.PersistentFlags().IntVarP(&loggerV, "verbose", "v", 0, "log level")

	cobra.OnInitialize(func() {
		log.SetVerbosity(loggerV)
		theHusky = husky.HuskyFrom(log.WithLogger(logger)(context.Background()), projectRoot)
	})
}
