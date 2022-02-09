package consumer

import "fmt"

func UserCreatedConsumer(message string) error {

	fmt.Println(message)
	return nil
}

func UserUpdatedConsumer(message string) error {
	fmt.Println(message)
	return nil

}

func UserDeletedConsumer(message string) error {
	fmt.Println(message)
	return nil
}
