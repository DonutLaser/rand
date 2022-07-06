package main

import (
	"fmt"
	"os"
	"strconv"
)

type Args struct {
	Command     string
	CommandArgs []string
}

func printUsage() {
	fmt.Println("Usage: rand <command> <args>")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  range <min> <max>	Get random number in the range [min, max]")
	fmt.Println("  list <path/to/file>	Get a random line in the specified file")
	fmt.Println("  coin 	Flip a coin")
	fmt.Println("  d6 	Roll a six-sided die. The same as calling `rand range 1 6`")
}

func parseArgs() (result Args, success bool) {
	args := os.Args[1:]

	if len(args) < 1 {
		return Args{}, false
	}

	result = Args{
		Command:     args[0],
		CommandArgs: args[1:],
	}

	return result, true
}

func runCommandRange(args []string) {
	if len(args) != 2 {
		fmt.Println("Incorrect `range` arguments")
		printUsage()
		return
	}

	min, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println(err)
		printUsage()
		return
	}

	max, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println(err)
		printUsage()
		return
	}

	fmt.Println(RandomRange(min, max))
}

func runCommandList(args []string) {
	if len(args) != 1 {
		fmt.Println("Incorrect `list` arguments")
		printUsage()
		return
	}

	lines, success := ReadFile(args[0])
	if !success {
		return
	}

	fmt.Println(RandomEntry(lines))
}

func runCommandCoin() {
	fmt.Println(FlipACoin())
}

func runCommandD6() {
	fmt.Println(RandomRange(1, 6))
}

func main() {
	args, success := parseArgs()
	if !success {
		printUsage()
		return
	}

	RandomInit()

	if args.Command == "range" {
		runCommandRange(args.CommandArgs)
	} else if args.Command == "list" {
		runCommandList(args.CommandArgs)
	} else if args.Command == "coin" {
		runCommandCoin()
	} else if args.Command == "d6" {
		runCommandD6()
	} else {
		fmt.Printf("Unknown command %q", args.Command)
	}
}
