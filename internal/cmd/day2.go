package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/westford14/advent2024/internal/day2"
)

type day2Flags struct {
	file string
	part string
}

func init() {
	rootCmd.AddCommand(runDay2Cmd())
}

func configureDay2Flags(flags *day2Flags, cmd *cobra.Command) {
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

func runDay2Cmd() *cobra.Command {
	flags := day2Flags{}
	runDay2Cmd := &cobra.Command{
		Use:   "day2",
		Short: "Run the advent of code 2024 day 2 command",
		Long:  "Run the advent of code 2024 day 2 command",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDay2(cmd, args, flags)
		},
	}

	configureDay2Flags(&flags, runDay2Cmd)
	return runDay2Cmd
}

func runDay2(cmd *cobra.Command, args []string, flags day2Flags) error {
	answer, err := day2.Execute(flags.file, flags.part)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Answer Day 2: %d\n", answer)
	return nil
}
