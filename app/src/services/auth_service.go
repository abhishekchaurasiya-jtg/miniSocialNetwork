package services

import (
	fmt "fmt"

	bcrypt "golang.org/x/crypto/bcrypt"

	dto "app/src/dto"
	models "app/src/models"
	repositories "app/src/repositories"
)

type AuthService interface {
	Register(request dto.CreateUserRequest) (*dto.TokensCollection, error)
	Login(request dto.LoginUserRequest) (*dto.TokensCollection, error)
}

type authService struct {
	userRepo repositories.UserRepository
	jwtSvc   JWTService
}

func NewAuthService(
	userRepo repositories.UserRepository,
	jwtSvc JWTService,
) AuthService {
	return &authService{
		userRepo: userRepo,
		jwtSvc:   jwtSvc,
	}
}

// Function to
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// AuthService for registering newuser
func (auth_service *authService) Register(request dto.CreateUserRequest) (*dto.TokensCollection, error) {

	// Email Verification
	user_count, err := auth_service.userRepo.GetUsersCountByEmail(request.Email)
	if err != nil {
		return nil, err
	}
	if user_count > 0 {
		return nil, fmt.Errorf("User Already Exists")
	}

	// Gender Verification
	if request.UserDetails.IsGenderValid(){
		return nil, fmt.Errorf("Invalid Gender Choice.")
	}
	// MaritalStatus Verification
	if request.UserDetails.IsMaritalStatusValid() {
		return nil, fmt.Errorf("Invalid MaritalStatus Choice.")
	}

	// Generating Hashed Password
	hashed_password, err := HashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	office_address := models.OfficeAddress{
		EmployeeId: request.UserDetails.OfficeDetails.EmployeeCode,
		ContactNo:  request.UserDetails.OfficeDetails.ContactNo,
		Address:    request.UserDetails.OfficeDetails.Address,
		City:       request.UserDetails.OfficeDetails.City,
		Country:    request.UserDetails.OfficeDetails.Country,
		State:      request.UserDetails.OfficeDetails.State,
		Name:       request.UserDetails.OfficeDetails.Name,
	}
	residential_address := models.ResidentialAdrress{
		ContactNo1: request.UserDetails.ResidentialDetails.ContactNo1,
		ContactNo2: request.UserDetails.ResidentialDetails.ContactNo2,
		Address:    request.UserDetails.ResidentialDetails.Address,
		City:       request.UserDetails.ResidentialDetails.City,
		Country:    request.UserDetails.ResidentialDetails.Country,
		State:      request.UserDetails.ResidentialDetails.State,
	}

	// Creating user Model
	user := models.User{
		FirstName:            request.UserDetails.FirstName,
		LastName:             request.UserDetails.LastName,
		Email:                request.Email,
		PasswordHash:         hashed_password,
		DateOfBirth:          &request.UserDetails.DateOfBirth,
		Gender:               request.UserDetails.GetEncodeGender(),
		MaritalStatus:        request.UserDetails.GetEncodeMaritalStatus(),
		OfficeAddresses:      []models.OfficeAddress{office_address},
		ResidentialAddresses: []models.ResidentialAdrress{residential_address},
	}
	auth_service.userRepo.AppendSingleNewUser(&user)

	// Generating Tokens
	tokens, err := auth_service.jwtSvc.GenerateNewTokens(int(user.ID), user.Email)
	if err != nil {
		return nil, nil
	}

	// updating refreshToken
	user.RefreshToken = tokens.RefreshToken
	err = auth_service.userRepo.UpdateUser(&user)
	if err != nil {
		return nil, fmt.Errorf("Failed to Generate the Error.")
	}

	return tokens, nil
}

func (auth_service *authService) Login(request dto.LoginUserRequest) (*dto.TokensCollection, error) {
	user, err := auth_service.userRepo.GetUserByEmail(request.Email)
	if err != nil {
		return nil, err
	}
	if !CheckPasswordHash(request.Password, user.PasswordHash) {
		return nil, fmt.Errorf("Failed to Match Password Hash.")
	}
	
	tokens, err := auth_service.jwtSvc.GenerateNewTokens(int(user.ID), user.Email)
	if err != nil {
		return nil, err
	}
	
	return tokens, nil
}
