package partyrobot

import "fmt"

func Welcome(name string) string {
	return fmt.Sprint("Welcome to my party %s", name)
}
