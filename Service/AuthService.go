package service

import (
	config "bookapplication/Config"
	model "bookapplication/Model"
	repository "bookapplication/Repository"
	"errors"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

//Creating the skeletal Interface
type AuthService interface {
	LoginUser(LoginUser *model.LoginDto) (*model.LoginRespDto, error)
	RegisterNewUser(NewUser *model.User) error
}

//Actual service Implementation
type AuthServiceImpl struct {
	UserRepo repository.UserRepository
}

//constructor injection
func NewAuthService(UserRepo repository.UserRepository) AuthService {
	return &AuthServiceImpl{
		UserRepo: UserRepo,
	}
}

// LoginUser implements [AuthService].
func (AuthService *AuthServiceImpl) LoginUser(LoginUser *model.LoginDto) (*model.LoginRespDto, error) {
	//Get the user by EmailId
	ExistingUser, err:= AuthService.UserRepo.FindUserByEmail(LoginUser.EmailId)
	if err!= nil{
		log.Printf("Fetch By emailId failed to retreive a valid User... %v",err)
		return nil, errors.New("user not found")
	}
	
	if ExistingUser == nil {
		return nil, errors.New("user not found")
	}

	//compare the hashed passwords to see if they are same
	if err:= bcrypt.CompareHashAndPassword([]byte(ExistingUser.Password),[]byte(LoginUser.Password)); err!= nil{
		return nil, errors.New("The password or EmailId is wrong")
	}

	//Generate and return an Auth token
	Token,err:= config.GenerateAuthToken(ExistingUser.EmailId)
	if err!= nil{
		log.Printf("The AuthTokenGeneration has failed... %v",err)
		return nil,err
	}
	
	Response := model.LoginRespDto{
		Token: Token,
		Name:  ExistingUser.UserName,
		Email: ExistingUser.EmailId,
	}
	return &Response, nil
}

// RegisterNewUser implements [AuthService].
func (AuthService *AuthServiceImpl) RegisterNewUser(NewUser *model.User) error {
	ExistingUser, err := AuthService.UserRepo.FindUserByEmail(NewUser.EmailId)
	if err != nil && ExistingUser == nil {
		// log the error but continue - repository may return error for other reasons
		log.Printf("error checking existing user: %v", err)
	}

	if ExistingUser != nil {
		return fmt.Errorf("the user is already registered, please try to login")
	}

	//Hash the password before saving it to the collection
	HashedPassword,err := bcrypt.GenerateFromPassword([]byte(NewUser.Password),bcrypt.DefaultCost)
	if err!=nil{
		return fmt.Errorf("The hash method failed to hash Password... %v",err)
	}
	NewUser.Password=string(HashedPassword)

	//create the new User
	if err := AuthService.UserRepo.CreateNewUser(NewUser); err != nil {
		log.Printf("failed to create new user: %v", err)
		return err
	}

	log.Print("The user has been successfully registered...")
	return nil
}

