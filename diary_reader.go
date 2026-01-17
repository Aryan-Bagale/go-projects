package main

import ("fmt"
		"bufio"
		"os"
)

func main(){
	file, err := os.Open("diary.txt")
	if err != nil{
		panic(err)
	}
	defer file.Close()


	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		fmt.Println("Readi line:", scanner.Text())
	}

}