package entities

type Customer struct {
	ID        string `bson:"uuid,omitempty"`
	Name      string
	Email     string
	Addresses []Address
}

type CreateCustomer struct {
	Name    string  `form:"name" binding:"required"`
	Email   string  `form:"email" binding:"required"`
	Address Address `form:"address" binding:"omitempty"`
}
