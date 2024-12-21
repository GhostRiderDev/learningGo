package main

import (
	"errors"
	"fmt"
	"sync"
)

type occupation struct {
	Name    string
	Level   uint8
	Payment float64
}

type user struct {
	Name       string
	Age        uint8
	Email      string
	Occupation occupation
}

func main() {
	user1 := user{
		Name:  "Olvadis",
		Age:   28,
		Email: "olvadis2004@gmail.com",
		Occupation: occupation{
			Name:    "Student",
			Level:   3,
			Payment: 700000.50,
		},
	}

	user2 := user{
		Name:  "Golang",
		Age:   13,
		Email: "golang@gmail.com",
		Occupation: occupation{
			Name:    "Phd",
			Level:   6,
			Payment: 87700000.50,
		},
	}

	var users = make([]user, 0)

	users = append(users, user1, user2)

	for i, v := range users {
		fmt.Println(i)
		fmt.Println(v.Email)
	}

	stringUser, err := user1.generateInfo()

	if err != nil {
		fmt.Println("Something was wrong")
	}

	fmt.Println(stringUser)

	totalPayments, err := calculatePayments(&users)

	if err != nil {
		fmt.Println("Error....")
	}
	fmt.Printf("Total payments: %.2f\n", totalPayments)

	higher18 := func(u user) bool { return u.Age >= 18 }

	highers18 := filter(users, higher18)

	for _, s := range highers18 {
		fmt.Println(s.Email)
	}

	
	extractEmail := func (u user) any { return u.Email}

	emails := mapFunc(users, extractEmail)
	stringEmails := make([]string, len(emails))
	for i, email := range emails {
		stringEmails[i] = email.(string)
	}

	var wg sync.WaitGroup

	wg.Add(1)

	go sendEmails(&stringEmails, &wg)

	wg.Wait()


	
}

func (u *user) generateInfo() (stringUser string, err error) {

	if u.Name == "" || u.Email == "" || u.Age <= 0 {
		return "", errors.New("invalid user")
	}

	stringUser = fmt.Sprintf("{\n  name: %s,\n  age: %d,\n  email: %s\n}", u.Name, u.Age, u.Email)

	return stringUser, nil
}

func calculatePayments(users *[]user) (float64, error) {
	var total float64 = 0.0

	for _, u := range *users {
		total += u.Occupation.Payment
	}
	return total, nil
}

func filter[T any](ss []T, apply func(T) bool) (ret []T) {
	for _, s := range ss {
		if apply(s) {
			ret = append(ret, s)
		}
	}
	return
}

func mapFunc[T any](data []T, apply func(T) any) (result []any) {
	for _, s := range data {
		result = append(result, apply(s))
	}
	return
}

func sendEmails(emails *[]string, wg *sync.WaitGroup) {
	for _, email := range *emails {
		fmt.Println("Sending email to: ", email)
	}
	wg.Done()
}
