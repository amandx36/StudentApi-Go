package types






// making the  struct to take the response from the client 

type Student struct{
	Id int64
	Name string   `validate:"required"`
	Email string	`validate:"required"`
	Age int 		`validate:"required"`
}