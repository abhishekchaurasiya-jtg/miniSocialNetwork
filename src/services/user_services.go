package services

import (
	dto "app/src/dto"
	"app/src/models"
	repositories "app/src/repositories"
	responses "app/src/responses"
)


type UserService interface {
	UpdateUser(request dto.UpdateUserRequest, userID int) (*dto.UpdateUserResponse,*responses.APIError)
	DeleteUser(userID int) (*dto.DeleteUserResponse, *responses.APIError)
	GetUserDetails(userID int) (*dto.GetUserResponse, *responses.APIError)
	GetActiveUsers(pageNo, pageSize int) (*[]dto.GetActiveUsersResponseItem, *responses.APIError)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService{
	return &userService{
		userRepo: userRepo,
	}
}

func (us *userService)UpdateUser(request dto.UpdateUserRequest, userID int) (*dto.UpdateUserResponse,*responses.APIError) {
	user, err := us.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, &responses.ErrInternalServerError
	}
	if request.DateOfBirth != nil {
		user.DateOfBirth = *request.DateOfBirth
	}
	if request.FirstName != nil {
		user.FirstName = *request.FirstName
	}
	if request.LastName != nil {
		user.LastName = *request.LastName
	}
	if request.Gender != nil {
		user.Gender = models.GenderChoices[*request.Gender]
	}
	if request.MaritalStatus != nil {
		user.MaritalStatus = models.MaritalStatusChoices[*request.MaritalStatus]
	}
	if err = us.userRepo.Save(user); err != nil {
		return nil, &responses.ErrInternalServerError 
	}

	responsePayload := dto.UpdateUserResponse{
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

	}
	return &responsePayload, nil
}

func (us *userService)DeleteUser(userID int) (*dto.DeleteUserResponse, *responses.APIError) {
	user, err := us.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, &responses.ErrInternalServerError
	}
	if err = us.userRepo.DeleteUser(user); err != nil {
		return nil, &responses.ErrInternalServerError
	}
	return &dto.DeleteUserResponse{
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
	}, nil
}

func (us *userService)GetUserDetails(userID int) (*dto.GetUserResponse, *responses.APIError) {
	user, err := us.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, &responses.ErrInternalServerError
	}
	return &dto.GetUserResponse{
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
	}, nil
}

func (us *userService)GetActiveUsers(pageNo, pageSize int) (*[]dto.GetActiveUsersResponseItem, *responses.APIError){
	users, err := us.userRepo.GetActiveUsers(pageNo, pageSize)

	if err != nil {
		return nil, &responses.ErrInternalServerError
	}

	return users, nil
}