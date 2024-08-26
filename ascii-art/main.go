package main

import (
	"ascii/methods"
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	errGeneral    = errors.New("usage: go run . [OPTION] [STRING]\nEX: go run . --color=<color> <letters to be colored> something")
	errOutputFlag = errors.New("usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
	errAlignFlag  = errors.New("usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --align=<type> something standard")
	//errFlag       = errors.New("usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --align=<type> something standard\nEX: go run . --output=<fileName.txt> something standard")
	//errTemplates  = errors.New("iNVALID TEMPLATE, use those:\n-->  shadow\n-->  standard\n-->  thinkertoy")
	errTypeFlag = errors.New("iNVALID ALIGN TYPE, use those:\n-->  left\n-->  right\n-->  center\n-->  justify")
)

func main() {
	args := os.Args
	var input string
	if len(args) == 2 {
		// go run . [STRING]
		input = args[1]
		if !methods.CheckFlag(args[1]) {
			fmt.Print(methods.ProccessOutput("", input, methods.ProcessBanner("./templates/standard.txt", input)))
		} else {
			if strings.HasPrefix(args[1], "--output") {
				fmt.Println(errOutputFlag)
			} else if strings.HasPrefix(args[1], "--align") {
				fmt.Println(errAlignFlag)
			} else {
				fmt.Println(errGeneral)
			}
		}
	} else if len(args) == 3 {
		if methods.CheckFlag(args[1]) {
			// go run . [OPTION = --output] [STRING]
			if strings.HasPrefix(args[1], "--output") {
				if outputFile := methods.OutputFile(args[1]); outputFile != "" {
					input = args[2]
					methods.WriteOutput(outputFile, methods.ProccessOutput("", input, methods.ProcessBanner("./templates/standard.txt", input)))
				} else {
					fmt.Println(errOutputFlag)
				}
			} else if strings.HasPrefix(args[1], "--align") {
				if alignColorOutput := methods.AlignColorOutput(args[1]); alignColorOutput != "" {
					input = args[2]
					if methods.IsValidType(alignColorOutput) {
						fmt.Print(methods.ProccessOutput(alignColorOutput, input, methods.ProcessBanner("./templates/standard.txt", input)))
					} else {
						fmt.Println(errTypeFlag)
					}
				} else {
					fmt.Println(errAlignFlag)
				}
			} else {
				if alignColorOutput := methods.AlignColorOutput(args[1]); alignColorOutput != "" {
					input = args[2]
					res := methods.ProccessOutputColor(alignColorOutput, "", input, methods.ProcessBanner("./templates/standard.txt", input))
					fmt.Print(res)
				} else {
					fmt.Println(errGeneral)
				}
			}
		} else {
			// go run . [STRING] [BANNER]
			input = args[1]
			templatefilename := methods.Template(args[2])
			fmt.Print(methods.ProccessOutput("", input, methods.ProcessBanner("./templates/"+templatefilename, input)))
		}
	} else if len(args) == 4 {
		// go run . [OPTION = --output] [STRING] [BANNER]
		if strings.HasPrefix(args[1], "--output") {
			if outputFile := methods.OutputFile(args[1]); outputFile != "" {
				input = args[2]
				templatefilename := methods.Template(args[3])
				methods.WriteOutput(outputFile, methods.ProccessOutput("", input, methods.ProcessBanner("./templates/"+templatefilename, input)))
			} else {
				fmt.Println(errOutputFlag)
			}
			// go run . [OPTION = --align] [STRING] [BANNER]
		} else if strings.HasPrefix(args[1], "--align") {
			if alignColorOutput := methods.AlignColorOutput(args[1]); alignColorOutput != "" {
				input = args[2]
				templatefilename := methods.Template(args[3])
				if methods.IsValidType(alignColorOutput) {
					fmt.Print(methods.ProccessOutput(alignColorOutput, input, methods.ProcessBanner("./templates/"+templatefilename, input)))
				} else {
					fmt.Println(errTypeFlag)
				}
			} else {
				fmt.Println(errAlignFlag)
			}
			// go run . [OPTION = --color] [SubString] [String]
		} else if strings.HasPrefix(args[1], "--color") {
			if alignColorOutput := methods.AlignColorOutput(args[1]); alignColorOutput != "" {
				if templatefilename := args[3]; methods.IsValidTemplate(templatefilename) {
					templatefilename := methods.Template(args[3])
					input = args[2]
					res := methods.ProccessOutputColor(alignColorOutput, "", input, methods.ProcessBanner("./templates/"+templatefilename, input))
					fmt.Print(res)
				} else {
					subStr := args[2]
					input = args[3]
					res := methods.ProccessOutputColor(alignColorOutput, subStr, input, methods.ProcessBanner("./templates/standard.txt", input))
					fmt.Print(res)
				}
			} else {
				fmt.Println(errGeneral)
			}
		} else {
			fmt.Println(errGeneral)
		}
	} else if len(args) == 5 && strings.HasPrefix(args[1], "--color") {
		if alignColorOutput := methods.AlignColorOutput(args[1]); alignColorOutput != "" {
			subStr := args[2]
			input = args[3]
			templatefilename := methods.Template(args[4])
			res := methods.ProccessOutputColor(alignColorOutput, subStr, input, methods.ProcessBanner("./templates/"+templatefilename, input))
			fmt.Print(res)
		} else {
			fmt.Println(errGeneral)
		}
	} else {
		fmt.Println(errGeneral)
	}
}
