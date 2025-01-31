package main

import "fmt"

//Interface ->contract work
type paymenter interface{
	pay(amount float32) //It is a method
}


//lets say we want payment integration
type payment struct{
	gateway paymenter
}

func (p payment) makePayment(amount float32){
	// razorpayPaymentGateway:=razorpay{}
	// razorpayPaymentGateway.pay(amount)
	//p.gateway.payStripe(amount)
}

type razorpay struct{}
//method
func (r razorpay) pay(amount float32){ //interface implemented here
	//logic to make payment
	fmt.Println("making payment using razorpay",amount)
}

type stripe struct{}
func (r stripe) payStripe(amount float32){
	//logic to make payment
	fmt.Println("making payment using stripe",amount)
}

type paypal struct{}
func(p paypal) pay(amount float32){ //interface implemented here
	fmt.Println("Payment Done",amount)
}
func main(){
	// stripePayment := stripe{}
	// newPayment := payment{
	// 	gateway: stripePayment,
	// }
	// newPayment.makePayment(600)

	//razorpayPay := razorpay{}

	paypalGateway:=paypal{}	
	newPayment:=payment{
		gateway: paypalGateway,
	}

	newPayment.gateway.pay(2323)


}