package model

type User struct{
	UserName string  `bson:"_id,omitempty" json:"userName"`
	Password string  `bson:"password"`		
	EmailId string   `bson:"emailId"`
	ContactNo string `bson:"contactNo"`
}

//login DTO
type LoginDto struct{
	EmailId string   `bson:"emailId"`
	Password string  `bson:"password"`		
}

// LoginRespDto
type LoginRespDto struct {
	Token string `json:"token"`
	Name  string `json:"name"`
	Email string `json:"email"`
}