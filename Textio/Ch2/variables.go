package main

import "fmt"

func main() {

	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var userName string
	fmt.Printf(
		"%v %f %v %q\n",
		smsSendingLimit,
		costPerSMS,
		hasPermission,
		userName,
	)

	// short assignment operator
	congrats := "Happy Birthday"
	fmt.Println(congrats)

	penniesPerText := 2.0

	fmt.Printf("The type of penniesPerText: %T\n", penniesPerText)

	averageOpenRate, displayMessage := 0.3, "is the average open rate of your message"

	fmt.Println(averageOpenRate, displayMessage)

	accountAge := 2.6
	accountAgeInt := int(accountAge)
	fmt.Println("Your account has existed for", accountAgeInt, "Years")

	// stick unless reason

	/*
		int
		uint
		float64
		uint32
		bool
		byte
		rune
		complex128
	*/

	const premiumPlanName = "Premium Plan"
	// premiumPlanName = "Basic Plan" can't mutate constants in go
	const basicPlan = "Basic Plan"
	fmt.Println(premiumPlanName, basicPlan)

	const secondsInMinute = 60
	const minuteInHours = 60
	const secondsInHour = secondsInMinute * minuteInHours

	fmt.Println("Number of seconds in Hours: ", secondsInHour)

	const name = "Saul GoodMan"
	const openRate = 0.5

	msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent", name, openRate)
	fmt.Println(msg)

	messageLen := 10
	maxMessageLen := 20
	fmt.Printf("Trying to send a message of length %v\n", messageLen)

	if messageLen > maxMessageLen {
		fmt.Println("Message Sent")
	} else {
		fmt.Println("Message not Sent")
	}
}
