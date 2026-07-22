package services

import (
	fmt "fmt"
	time "time"

	bcrypt "golang.org/x/crypto/bcrypt"

	dto "app/src/dto"
	models "app/src/models"
	repositories "app/src/repositories"
)

type AuthService interface {
	Register(request dto.CreateUserRequest) (*dto.UserCreateResponse, *dto.TokensCollection, error)
	Login(request dto.LoginUserRequest) (*dto.UserLoginResponse, *dto.TokensCollection, error)
	RefreshAccessToken(refreshToken string) (*string, error)
	UpdatePassword(request dto.UpdatePasswordRequest, email string) error
	LogOut(email string) error
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
func (auth_service *authService) Register(request dto.CreateUserRequest) (
	*dto.UserCreateResponse, *dto.TokensCollection, error) {

	// Email Verification
	user_count, err := auth_service.userRepo.GetUsersCountByEmail(request.Email)
	if err != nil {
		return nil, nil, err
	}
	if user_count > 0 {
		return nil,nil, fmt.Errorf("User Already Exists")   // Explore
	}

	// Gender Verification
	if !request.UserDetails.IsGenderValid(){
		return nil,nil, fmt.Errorf("Invalid Gender Choice.")  // make const error messages.
	}
	// MaritalStatus Verification
	if !request.UserDetails.IsMaritalStatusValid() {
		return nil, nil, fmt.Errorf("Invalid MaritalStatus Choice.")
	}

	// Generating Hashed Password
	hashed_password, err := HashPassword(request.Password)
	if err != nil {
		return nil,nil,err
	}

	office_address := models.OfficeDetails{
		EmployeeId: request.UserDetails.OfficeDetails.EmployeeCode,
		ContactNo:  request.UserDetails.OfficeDetails.ContactNo,
		Address:    request.UserDetails.OfficeDetails.Address,
		City:       request.UserDetails.OfficeDetails.City,
		Country:    request.UserDetails.OfficeDetails.Country,
		State:      request.UserDetails.OfficeDetails.State,
		Email:      request.UserDetails.OfficeDetails.Email,
		Name:       request.UserDetails.OfficeDetails.Name,
	}
	residential_address := models.ResidentialDetails{
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
		DateOfBirth:          request.UserDetails.DateOfBirth,
		Gender:               request.UserDetails.GetEncodeGender(),
		MaritalStatus:        request.UserDetails.GetEncodeMaritalStatus(),
		OfficeDetails:      []models.OfficeDetails{office_address},
		ResidentialDetails: []models.ResidentialDetails{residential_address},
	}
	auth_service.userRepo.AppendSingleNewUser(&user)

	err = auth_service.userRepo.PreloadAddresses(&user)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to reload user associations: %w", err)
	}

	// Generating Tokens
	tokens, err := auth_service.jwtSvc.GenerateNewTokens(int(user.ID), user.Email)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate auth tokens: %w", err)
	}

	// updating refreshToken
	user.RefreshToken = &tokens.RefreshToken
	err = auth_service.userRepo.UpdateUser(&user)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to update user session records: %w", err)
	}


	// Response Payload
	responsePayload := &dto.UserCreateResponse{
		ID: int(user.ID),
		UserId: int(user.ID),
		Email: user.Email,
		LastModified: user.UpdatedAt,
		UserDetail: dto.UserDetailsResponse{
			ID: int(user.ID),
			UserId: int(user.ID),
			FirstName: user.FirstName,
			LastName: user.LastName,
			DateOfBirth: user.DateOfBirth,
			Gender: user.Gender.String(),
			MaritalStatus: user.MaritalStatus.String(),
			ResidentialDetails: dto.ResidentialDetailsResponse{
				Id: int(user.ResidentialDetails[0].ID),
				UserId: int(user.ID),
				Address: user.ResidentialDetails[0].Address,
				City: user.ResidentialDetails[0].City,
				State: user.ResidentialDetails[0].State,
				Country: user.ResidentialDetails[0].Country,
				ContactNo1: user.ResidentialDetails[0].ContactNo1,
				ContactNo2: user.ResidentialDetails[0].ContactNo2,
			},
			OfficeDetails: dto.OfficeDetailsResponse{
				Id: int(user.OfficeDetails[0].ID),
				UserId: int(user.ID),
				EmployeeCode: user.OfficeDetails[0].EmployeeId,
				Address: user.OfficeDetails[0].Address,
				City: user.OfficeDetails[0].City,
				State: user.OfficeDetails[0].State,
				Country: user.OfficeDetails[0].Country,
				ContactNo: user.OfficeDetails[0].ContactNo,
				Email: user.OfficeDetails[0].Email,
				Name: user.OfficeDetails[0].Name,
			},
		},
		Token: dto.TokenResponse{
			Token: tokens.AccessToken,
			ExpiryTime: time.Now().Add(15 * time.Minute),
		},
	}

	return responsePayload, tokens, nil
}

func (auth_service *authService) Login(request dto.LoginUserRequest) (
	*dto.UserLoginResponse, *dto.TokensCollection, error) {
	user, err := auth_service.userRepo.GetUserByEmail(request.Email)
	if err != nil {
		return nil, nil, err
	}
	if !CheckPasswordHash(request.Password, user.PasswordHash) {
		return nil, nil, fmt.Errorf("Failed to Match Password Hash.")
	}
	
	tokens, err := auth_service.jwtSvc.GenerateNewTokens(int(user.ID), user.Email)
	if err != nil {
		return nil, nil, err
	}

	loginResponse := dto.UserLoginResponse{
		Token: dto.TokenResponse{
			Token: tokens.AccessToken,
			ExpiryTime: time.Now().Add(15*time.Minute),
		},
		User: dto.UserDetailsLoginResponse{
			ID: int(user.ID),
			Email: user.Email,
			FirstName: user.FirstName,
			LastName: user.LastName,
			LastModified: user.UpdatedAt,
		},
	}
	
	return &loginResponse, tokens, nil
}


func (auth_service *authService) RefreshAccessToken(refresh_token string) (*string, error) {
	accessToken, err := auth_service.jwtSvc.GenerateAccessTokenFromRefresh(refresh_token)
	if err != nil {
		return nil, err
	}
	return &accessToken, nil
}

func (auth_service *authService) UpdatePassword(request dto.UpdatePasswordRequest, email string) error {
	if request.NewPassword == request.OldPassword {
		return fmt.Errorf("Same Password.")
	}

	user, err := auth_service.userRepo.GetUserByEmail(email)
	if err != nil {
		return err
	}
	// Validation
	if !CheckPasswordHash(request.OldPassword, user.PasswordHash) {
		return fmt.Errorf("unauthorized: wrong old Password ")
	}
	newPasswordHash, err := HashPassword(request.NewPassword)
	if err != nil {
		return err
	}

	user.PasswordHash = newPasswordHash
	err = auth_service.userRepo.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
}


func (auth_service *authService)LogOut(email string) error {
	user, err := auth_service.userRepo.GetUserByEmail(email)
	if err != nil {
		return err
	}

	user.RefreshToken = nil
	err = auth_service.userRepo.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
}