package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/westford14/advent2024/internal/day1"
)

type day1Flags struct {
	file string
	part string
}

func init() {
	rootCmd.AddCommand(runDayCmd())
}

func configureDay1Flags(flags *day1Flags, cmd *cobra.Command) {
	cmd.Flags().SortFlags = false
	cmd.Flags().StringVarP(
		&flags.file,
		"file",
		"f",
		"",
		"path to the input test file for day 1",
	)
	cmd.Flags().StringVarP(
		&flags.part,
		"part",
		"p",
		"",
		"which part to run for day 1 (part1|part2)",
	)
	if err := cmd.MarkFlagRequired("file"); err != nil {
		panic(err)
	}
}

func runDayCmd() *cobra.Command {
	flags := day1Flags{}
	runDayCmd := &cobra.Command{
		Use:   "day1",
		Short: "Run the advent of code 2023 day 1 command",
		Long:  "Run the advent of code 2023 day 1 command",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDay1(cmd, args, flags)
		},
	}

	configureDay1Flags(&flags, runDayCmd)
	return runDayCmd
}

func runDay1(cmd *cobra.Command, args []string, flags day1Flags) error {
	answer, err := day1.Execute(flags.file, flags.part)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Answer Day 1: %d\n", answer)
	return nil
}
