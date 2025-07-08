package partyrobot

import "fmt"

func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party %s", name)
}

func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}
