package main

import (
	"fmt"
	"strconv"
	"strings"
)

const shopName = "AltPhone"

type UserData struct {
	firstName string
	lastName  string
	email     string
	phoneId   int
	quantity  int
}

var orderList = make([]UserData, 0)

var shopItems = map[int]map[string]string{
	1: {"name": "iphone6", "qty": "10"},
	2: {"name": "iphone7", "qty": "10"},
	3: {"name": "iphone7", "qty": "10"},
	4: {"name": "iphone7", "qty": "10"},
}

func main() {
	welcomePage()
	orderCreation()
}

func welcomePage() {
	fmt.Printf("Welcome to %v stores \n", shopName)
	fmt.Println("Here are items in our shop", shopItems)
	fmt.Println("Please enter your detals and a valid phone id to purchase an item")
}

func getUserInputToFromUser() (string, string, string, uint, uint) {
	var firstName string
	var lastName string
	var email string
	var phoneId uint
	var quantity uint

	fmt.Println("Please enter your first name")
	fmt.Scan(&firstName)
	fmt.Println("Please enter your last name name")
	fmt.Scan(&lastName)
	fmt.Println("Please enter your email ")
	fmt.Scan(&email)
	fmt.Println("Please enter the phone id to purchase")
	fmt.Scan(&quantity)
	fmt.Println("Please enter the quantity of phone to purchase")
	fmt.Scan(&quantity)

	return firstName, lastName, email, phoneId, quantity

}

func validateInputsFromUser(firstName string, lastName string, email string, phoneId int, qty int) (bool, bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isPhoneIdValid := phoneId > 0 && phoneId <= 4
	individualPhone := shopItems[phoneId]
	fmt.Println(individualPhone)
	balance, _ := strconv.Atoi(individualPhone["qty"])
	isQuantityRequestValid := balance >= qty

	return isValidName, isValidEmail, isPhoneIdValid, isQuantityRequestValid
}

func inventory(shopItems map[int]map[string]string, phoneId int, qty int) map[int]map[string]string {
	innerValueOfMap, ok := shopItems[phoneId]
	if !ok {
		fmt.Println("PLEASE INPUT A CORRECT ID")
		return nil
	}
	balance, _ := strconv.Atoi(innerValueOfMap["qty"])
	balance = balance - qty

	innerValueOfMap["qty"] = strconv.Itoa(balance)
	return shopItems
}

func orderCreation(firstName string, lastName string, email string, phoneId int, qty int) {

	var userInfo = UserData{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		phoneId:   phoneId,
		quantity:  qty,
	}
	orderList = append(orderList, userInfo)

	fmt.Println("This is the list of our orders", orderList)

}
