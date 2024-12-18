package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/westford14/advent_of_code_2024/internal/day1"
	"github.com/westford14/advent_of_code_2024/internal/day2"
	"github.com/westford14/advent_of_code_2024/internal/day3"
	"github.com/westford14/advent_of_code_2024/internal/day4"
	"github.com/westford14/advent_of_code_2024/internal/day5"
	"github.com/westford14/advent_of_code_2024/internal/day6"
	"github.com/westford14/advent_of_code_2024/internal/day7"
	"github.com/westford14/advent_of_code_2024/internal/day8"
)

type dayFlags struct {
	file string
	part int
	day  int
}

func init() {
	rootCmd.AddCommand(runDayCmd())
}

func configureDayFlags(flags *dayFlags, cmd *cobra.Command) {
	cmd.Flags().SortFlags = false
	cmd.Flags().IntVarP(
		&flags.day,
		"day",
		"d",
		0,
		"which day to run",
	)
	cmd.Flags().StringVarP(
		&flags.file,
		"file",
		"f",
		"",
		"path to the input test file for day 2",
	)
	cmd.Flags().IntVarP(
		&flags.part,
		"part",
		"p",
		0,
		"which part to run for day 2 (part1|part2)",
	)
	if err := cmd.MarkFlagRequired("file"); err != nil {
		panic(err)
	}
}

func runDayCmd() *cobra.Command {
	flags := dayFlags{}
	runDayCmd := &cobra.Command{
		Use:   "day",
		Short: "Run the advent of code 2024 day command",
		Long:  "Run the advent of code 2024 day command",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDay(cmd, args, flags)
		},
	}

	configureDayFlags(&flags, runDayCmd)
	return runDayCmd
}

func runDay(cmd *cobra.Command, args []string, flags dayFlags) error {
	var ans int
	var err error
	switch flags.day {
	case 1:
		ans, err = day1.Execute(flags.file, flags.part)
	case 2:
		ans, err = day2.Execute(flags.file, flags.part)
	case 3:
		ans, err = day3.Execute(flags.file, flags.part)
	case 4:
		ans, err = day4.Execute(flags.file, flags.part)
	case 5:
		ans, err = day5.Execute(flags.file, flags.part)
	case 6:
		ans, err = day6.Execute(flags.file, flags.part)
	case 7:
		ans, err = day7.Execute(flags.file, flags.part)
	case 8:
		ans, err = day8.Execute(flags.file, flags.part)
	default:
		ans = 0
		err = fmt.Errorf("did not understand %d", flags.day)
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("Answer Day 3: %d\n", ans)
	return nil
}
