package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/westford14/advent2024/internal/day3"
)

type day3Flags struct {
	file string
	part string
}

func init() {
	rootCmd.AddCommand(runDay3Cmd())
}

func configureDay3Flags(flags *day3Flags, cmd *cobra.Command) {
	cmd.Flags().SortFlags = false
	cmd.Flags().StringVarP(
		&flags.file,
		"file",
		"f",
		"",
		"path to the input test file for day 2",
	)
	cmd.Flags().StringVarP(
		&flags.part,
		"part",
		"p",
		"",
		"which part to run for day 2 (part1|part2)",
	)
	if err := cmd.MarkFlagRequired("file"); err != nil {
		panic(err)
	}
}

func runDay3Cmd() *cobra.Command {
	flags := day3Flags{}
	runDay3Cmd := &cobra.Command{
		Use:   "day3",
		Short: "Run the advent of code 2024 day 3 command",
		Long:  "Run the advent of code 2024 day 3 command",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDay3(cmd, args, flags)
		},
	}

	configureDay3Flags(&flags, runDay3Cmd)
	return runDay3Cmd
}

func runDay3(cmd *cobra.Command, args []string, flags day3Flags) error {
	answer, err := day3.Execute(flags.file, flags.part)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Answer Day 3: %d\n", answer)
	return nil
}
