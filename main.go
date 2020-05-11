package main

import (
    "fmt"
	"poll"
    "os"
)

func printHelp() {
    fmt.Println("USAGE")
    fmt.Println("\t./209poll pSize sSize p")
    fmt.Println("")
    fmt.Println("DESCRIPTION")
    fmt.Println("\tpSize\tsize of the population")
    fmt.Println("\tsSize\tsize of the sample (supposed to be representative)")
    fmt.Println("\tp\tpercentage of voting intentions for a specific candidate")
}

func main() {
    if poll.CheckHelp() {
        printHelp()
        os.Exit(0)
    }
    if !poll.CheckArgs() {
        printHelp()
        os.Exit(84)
    }
    poll.Poll()
    os.Exit(0)
}