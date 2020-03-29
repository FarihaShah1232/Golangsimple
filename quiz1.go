package main

import(
	"fmt"
	"os"
)

func main(){

var i int
var msg string = "Leave the entire universe and never come back."
var option string

Pakistan := 1528
SouthKorea := 2548
France := 1023

Select := "Please select an option:\n\n1-Print Covid19 cases in Pakistan\n2-Print Covid19 cases in SouthKorea.\n3-Print Covid19 cases in France\n4-Print my personalized message to Coronavirus\n0-Exit\n"

for i < 3{

fmt.Println(Select)
fmt.Scanf("%s\n",&option)

switch option{

case "1":
	fmt.Println("Pakistan = ", Pakistan , "\n")
case "2":
	fmt.Println("SouthKorea = ",SouthKorea, "\n")
case "3":
	fmt.Println("France = ", France, "\n")
case "4":
	fmt.Println(msg)
case "0":
	os.Exit(1)	
default:
	fmt.Println("Select an option to get results.")
}

if (option == "1" || option == "2" || option == "3"){
	i++
}
}

}