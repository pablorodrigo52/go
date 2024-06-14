package main

import "fmt"

type Payment struct {
	OriginalAmount float32
	PaymentAmount float32
	Barcode string
	DueDate string
	UserID int
}

type setters func(*Payment)
type PaymentBuilder struct {
	actions []setters
}

func NewPaymentBuilder() *PaymentBuilder {
	return &PaymentBuilder{}
}

func (p *PaymentBuilder) WithOriginalAmount(amount float32) *PaymentBuilder{
	p.actions = append(p.actions, func(p *Payment) {
		p.OriginalAmount = amount
	})

	return p
}

func (p *PaymentBuilder) WithPaymentAmount(amount float32) *PaymentBuilder{
	p.actions = append(p.actions, func(p *Payment) {
		p.PaymentAmount = amount
	})

	return p
}

func (p *PaymentBuilder) WithBarcode(barcode string) *PaymentBuilder {
	p.actions = append(p.actions, func(p *Payment) {
		p.Barcode = barcode
	})

	return p
}

func (p *PaymentBuilder) WithDueDate(duedate string) *PaymentBuilder {
	p.actions = append(p.actions, func(p *Payment) {
		p.DueDate = duedate
	})

	return p
}

func (p *PaymentBuilder) WithUserID(userID int) *PaymentBuilder {
	p.actions = append(p.actions, func(p *Payment) {
		p.UserID = userID
	})

	return p
}

func (p *PaymentBuilder) Build() Payment {
	payment := Payment{}
	for _, setter := range p.actions {
		setter(&payment)
	}

	return payment
}

func main () {
	builder := NewPaymentBuilder()
	p := builder.
		WithOriginalAmount(10).
		WithPaymentAmount(12.1).		
		WithBarcode("1234mock_barcode").
		WithDueDate("10/08/2024").
		WithUserID(123).
		Build()

	p2 := Payment{
		10, 12.1, "1234mock_barcode", "10/08/2024", 123,
	}

	p3 := Payment{}
	p3.OriginalAmount = 10
	p3.PaymentAmount = 12.1
	p3.Barcode = "1234mock_barcode"
	p3.DueDate = "10/08/2024"
	p3.UserID = 123


	fmt.Println("p : ", p)
	fmt.Println("p2: ", p2)
	fmt.Println("p3: ", p3)
}