package main

import ("fmt"
		"os"
		"strings"
		"bufio"
)

type Word struct{
	Noun string
	Verb string
	Adjective string
}

func stringInput(reader *bufio.Reader, prompt string)(string, error){
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil{
		return "", err
	}
	return strings.TrimSpace(input), nil
}

func stringRequiredInput(reader *bufio.Reader, label string)(string){
	for{
		input, err := stringInput(reader, label+": ")
		if err != nil{
			fmt.Println("Error reading", label+":", err )
			return ""
		}
		if input != ""{
			return input
		}
		fmt.Println(label, "can not be empty. Try again.")
	}
}
func main(){
	reader := bufio.NewReader(os.Stdin)
	w :=Word{
		Noun: stringRequiredInput(reader, "Noun"),
		Verb: stringRequiredInput(reader, "Verb"),
		Adjective: stringRequiredInput(reader, "Adjective"),
	}

	fmt.Printf(
		"The %s %s will %s despite unexpected circumstances and external pressures.\n",
		w.Adjective, w.Noun, w.Verb,
	)
}