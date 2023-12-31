package main

import (
    "fmt"
    "bufio"
    "strings"
    "os"
    "strconv"
)

type bill struct {
    name string
    items map[string]float64
    tip float64
}

func GetInput(prompt string, reader *bufio.Reader) (string, error) {
    fmt.Println(prompt)
    input, err := reader.ReadString('\n')
    return strings.TrimSpace(input), err
}

func BillBreakdown(bill bill) string {
    breakdown := bill.name + "'s bill\n"
    total := 0.0

    for item, price := range bill.items {
        breakdown += fmt.Sprintf("%-25v $%0.2f\n", item, price)
        total += price
    }
    total += bill.tip
    breakdown += "\n"
    breakdown += fmt.Sprintf("%-25v $%0.2f\n", "tip:", bill.tip)
    breakdown += fmt.Sprintf("%-25v $%0.2f\n", "total:", total)
    return breakdown
}

func GetFloatInput(prompt string, reader *bufio.Reader) (float64, error, error) {
    fmt.Println(prompt)
    input, readingErr := reader.ReadString('\n')
    input = strings.TrimSpace(input)
    float, parseErr := strconv.ParseFloat(input, 64)

    return float, readingErr, parseErr
}

func (bill *bill) AddItem(name string, price float64) {
    bill.items[name] = price
}

func (bill *bill) SetTip(tip float64) {
    bill.tip = tip
}

func NewBill(name string) bill {
    bill := bill{
        name: name,
        items: map[string]float64{},
        tip: 0.0,
    }
    return bill
}

func DisplayOptions(bill bill) {
    reader := bufio.NewReader(os.Stdin) 
    input, _ := GetInput("'a' to add an item - 't' to add a tip - 's' to save the bill - 'q' to quit", reader)

    switch input { case "a":
        name, _ := GetInput("name of the item: ", reader)
        price, _, _ := GetFloatInput("price of the item: ", reader)
        bill.AddItem(name, price)

        DisplayOptions(bill)
    case "t":
        tip, _, _ := GetFloatInput("set tip amount: ", reader)
        bill.SetTip(tip)

        DisplayOptions(bill)
    case "s":
        billBreakdown := BillBreakdown(bill)
        os.WriteFile("bills/"+bill.name+"'s bill", []byte(billBreakdown), 0666)
        fmt.Println(billBreakdown)

        DisplayOptions(bill)
    case "q":
        break
    default:
        fmt.Print("not an option")

        DisplayOptions(bill)
    }
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    name, _ := GetInput("Name of the bill: ", reader)
    myBill := NewBill(name)
    DisplayOptions(myBill)
}
