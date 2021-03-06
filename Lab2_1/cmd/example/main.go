package main
 
import (
    "fmt"
	"bufio"
	"os"
	lab1 "github.com/Iluysha/Lab2_1"
)


func main() {
	fmt.Printf("Input of postfix expression: ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
    
	if res, err := lab1.PostfixToInfix(text); err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println("\nThe infix expression is ", res)
    }
	
	fmt.Println(buildVersion)
}
