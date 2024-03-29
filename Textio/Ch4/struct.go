package main

import (
	"fmt"
)

type messageToSend struct {
	phoneNumber int
	message     string
}

func test(m messageToSend) {
	fmt.Printf("Sending message: '%s' to: %v\n",
		m.message, m.phoneNumber)
}

type messageToSend2 struct {
	message string
	sender  user

	recipent user
}

type user struct {
	name   string
	number int
}

func canMessageSend(mToSend messageToSend2) bool {
	if mToSend.sender.name == "" {
		return false
	}

	if mToSend.recipent.name == "" {
		return false
	}

	if mToSend.sender.number == 0 {
		return false
	}

	if mToSend.recipent.number == 0 {
		return false
	}

	return true

}

type senderX struct {
	userX     // embedded struct takes all properties of sender
	ratelimit int
}

type userX struct {
	name   string
	number int
}

func newTest(s senderX) {
	fmt.Println("Sender name is: ", s.name)
	fmt.Println("Sender number is: ", s.number)
	fmt.Println("Rate limit is: ", s.ratelimit)
	fmt.Println("===============")
}
func main() {
	fmt.Println("Hello")

	test(messageToSend{
		phoneNumber: 10,
		message:     "Thanks for signing up",
	})

	test(messageToSend{
		phoneNumber: 192301231,
		message:     "Welcome",
	})

	newTest(senderX{
		ratelimit: 10,
		userX: userX{
			name:   "Saksham",
			number: 1209381230981,
		},
	})

}
