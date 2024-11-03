package cmd

import (
	"fmt"

	"github.com/aethiopicuschan/machina/internal/service"
	"github.com/aethiopicuschan/machina/internal/service/common"
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

	// Default config from https://github.com/matsuu/kataribe/blob/master/toml.go
	general := common.General{
		RankingCount:   20,
		SlowCount:      37,
		ShowStddev:     true,
		ShowStatusCode: true,
		ShowBytes:      true,
		Percentiles:    []float64{50.0, 90.0, 95.0, 99.0},
	}
	scale := common.Scale{
		Scale:          0,
		EffectiveDigit: 3,
	}

	bundles, err := s.GenerateBundles()
	if err != nil {
		return
	}

	// Print the general config
	fmt.Printf("# General\n")
	fmt.Printf("%s\n", general.String())

	// Print the scale config
	fmt.Printf("# Scale\n")
	fmt.Printf("%s\n", scale.String())

	// Print the bundles
	fmt.Printf("# Bundles\n")
	for _, bundle := range bundles {
		fmt.Printf("%s\n", bundle.String())
	}

	return
}

func Execute() {
	rootCmd.Execute()
}
