package main

// Defining Customer structure
type Customer struct {
	CustomerId  int     `json:"customerId"`
	FirstName   string  `json:"firstName"`
	LastName    string  `json:"lastName"`
	Email       string  `json:"email"`
	AddressInfo Address `json:"addressInfo"`
	Orders      []Order `json:"orders"`
}

// Defining Address structure
type Address struct {
	StreetNumber string `json:"streetNumber"`
	StreetName   string `json:"streetName"`
	City         string `json:"city"`
	Country      string `json:"country"`
}

// Defining Book structure
type Book struct {
	BookId       int    `json:"bookId"`
	Title        string `json:"title"`
	Language     string `json:"language"`
	NumberOfPage int    `json:"numberOfPage"`
	Price        int    `json:"price"`
	AuthorId     int    `json:"authorId"`
}

// Defining Order structure
type Order struct {
	OrderId    int    `json:"orderId"`
	BookId     int    `json:"bookId"`
	CustomerId int    `json:"customerId"`
	Status     string `json:"status"`
}

// Defining Author structure
type Author struct {
	AuthorId   int          `json:"authorId"`
	AuthorName string       `json:"authorName"`
	TotalBook  int          `json:"totalBook"`
	BookSale   int          `json:"bookSale"`
	MyBooks    map[int]bool `json:"myBooks"`
}

// Defining Databases
type BookDB map[int]*Book
type CustomerDB map[int]*Customer
type OrderDB map[int]*Order
type AuthorDB map[int]*Author

// Declaring Databases
var BookList BookDB
var CustomerList CustomerDB
var OrderList OrderDB
var AuthorList AuthorDB

// Initializing Databases
func Init() {
	BookList = make(BookDB)
	CustomerList = make(CustomerDB)
	OrderList = make(OrderDB)
	AuthorList = make(AuthorDB)
}

// Generating a demo Author
func GenerateAuthor() {
	AuthorList[1] = &Author{
		AuthorId:   1,
		AuthorName: "CALEB DOXSEY",
		TotalBook:  1,
		BookSale:   1,
		MyBooks:    map[int]bool{1: true},
	}
}

// Generating a demo Book
func GenerateBook() {
	BookList[1] = &Book{
		BookId:       1,
		Title:        "AN INTRODUCTION TO PROGRAMMING IN GO",
		Language:     "English",
		NumberOfPage: 161,
		Price:        200,
		AuthorId:     1,
	}
}

// Generating a demo Order
func GenerateOrder() {
	OrderList[1] = &Order{
		OrderId:    1,
		BookId:     1,
		CustomerId: 1,
		Status:     "Waiting For Delivery",
	}
}

// Generating a demo Customer
func GenerateCustomer() {
	CustomerList[1] = &Customer{
		CustomerId: 1,
		FirstName:  "Mobarak",
		LastName:   "Hossain",
		Email:      "mobarak@appscode.com",
		AddressInfo: Address{
			StreetNumber: "10/A",
			StreetName:   "Sector 10",
			City:         "Dhaka",
			Country:      "Bangladesh",
		},
		Orders: []Order{*OrderList[1]},
	}
}
