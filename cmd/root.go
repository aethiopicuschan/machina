package cmd

import (
	"fmt"

	"github.com/aethiopicuschan/machina/internal/service"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "machina [source]",
	Long:    "Machina is automatic generator for kataribe.toml",
	Version: "0.1.0",
	Args:    cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(1)),
	RunE:    run,
}

func run(cmd *cobra.Command, args []string) (err error) {
	fp := args[0]

	s, err := service.New(fp)
	if err != nil {
		return
	}
	defer s.Close()

	bundles, err := s.Generate()
	if err != nil {
		return
	}

	for _, bundle := range bundles {
		fmt.Printf("%s\n", bundle.String())
	}

	return
}

func Execute() {
	rootCmd.Execute()
}
