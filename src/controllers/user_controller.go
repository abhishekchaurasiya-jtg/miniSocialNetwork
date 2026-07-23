package controllers

import (
	strconv "strconv"
	http "net/http"
	log "log"

	gin "github.com/gin-gonic/gin"

	dto "app/src/dto"
	responses "app/src/responses"
	services "app/src/services"
)




type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func getUserIDFromContext(context *gin.Context) (int, error) {
    idVal, exists := context.Get("UserID")
    if !exists {
        return 0, responses.CorruptedTokenData
    }
	log.Printf("--- DEBUG: UserID Value: %v, Type: %T ---", idVal, idVal)

    id, ok := idVal.(int)
    if !ok {
        return 0, responses.CorruptedTokenData
    }
    return id, nil
}

func (userCnt *UserController)UpdateUser(context *gin.Context) {
	// Validations
	userID, err := getUserIDFromContext(context)
	if err != nil {
		responses.WriteAPIError(context, responses.ErrCourruptedToken)
		return
	}

	var request *dto.UpdateUserRequest;
	
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusUnprocessableEntity,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	responsePayload, err_ := userCnt.userService.UpdateUser(*request, userID)
	if err_ != nil {
		responses.WriteAPIError(context, *err_)
		return
	}

	context.JSON(
		http.StatusCreated, responsePayload,
	)
}

func (userCnt *UserController)DeleteUser(context *gin.Context) {
	userID, err := getUserIDFromContext(context)
	if err != nil {
		responses.WriteAPIError(context, responses.ErrCourruptedToken)
		return
	}

	responsePlayload, err_ := userCnt.userService.DeleteUser(userID)
	if err_ != nil {
		responses.WriteAPIError(context, *err_)
		return
	}

	context.JSON(
		http.StatusNoContent,
		responsePlayload,
	)
} 

func (userCnt *UserController)GetUser(context *gin.Context) {
	userID, err := getUserIDFromContext(context)
	if err != nil {
		responses.WriteAPIError(context, responses.ErrCourruptedToken)
		return
	}

	responsePayload, err_ := userCnt.userService.GetUserDetails(userID)
	if err_ != nil {
		responses.WriteAPIError(context, *err_)
		return 
	}
	context.JSON(
		http.StatusOK, responsePayload)
}

func (userCnt *UserController)GetActiveUsers(context *gin.Context) {
	pageNoStr := context.Query("page_no")
	pageSizeStr := context.Query("page_size")
	if pageNoStr == "" {
		pageNoStr = "1"
	}
	if pageSizeStr == "" {
		pageSizeStr = "50"
	}

	pageNo, err := strconv.Atoi(pageNoStr)
	if err != nil {
		responses.WriteAPIError(context, responses.ErrInvalidQueryParams)
		return
	}
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		responses.WriteAPIError(context, responses.ErrInvalidQueryParams)
	}


	responsePayload, apiErr := userCnt.userService.GetActiveUsers(pageNo, pageSize) 
	if apiErr != nil {
		responses.WriteAPIError(context, *apiErr)
	}

	context.JSON(
		http.StatusOK,
		responsePayload,
	)
}