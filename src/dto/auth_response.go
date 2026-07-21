package dto 


import (
	time "time"
)


type ResidentialDetailsResponse struct {
	Id int						`json:"id"`
	UserId int					`json:"user_id"`
	Address string				`json:"address"`
	City string					`json:"city"`
	State string				`json:"state"`
	Country string				`json:"country"`
	ContactNo1 string			`json:"contact_no_1"`
	ContactNo2 string			`json:"contact_no_2"`
}

type OfficeDetailsResponse struct {
	Id int					`json:"id"`			
	UserId int				`json:"user_id"`				
	EmployeeCode string		`json:"employee_code"`	
	Address string			`json:"address"`
	City string				`json:"city"`
	State string			`json:"state"`
	Country string			`json:"country"`
	ContactNo string		`json:"contact_no"`	
	Email string			`json:"email"`					
	Name string				`json:"name"`		


}

type TokenResponse struct {
	Token string				`json:"key"`
	ExpiryTime time.Time		`json:"expiry_time"`
}


type UserDetailsResponse struct {
	ID int 											`json:"id"`
	FirstName string								`json:"first_name"`
	LastName string									`json:"last_name"`
	DateOfBirth string								`json:"date_of_birth"`
	Gender string									`json:""`
	MaritalStatus string							`json:""`
	ResidentialDetails ResidentialDetailsResponse	`json:""`
	OfficeDetails OfficeDetailsResponse				`json:""`

}

type UserCreateResponse struct {
	ID int 								`json:"user_id"`
	Email int							`json:"email"`
	LastModified time.Time				`json:"last_modified"`
	UserDetail UserDetailsResponse		`json:"user_details"`
	Token TokenResponse					`json:"token"`
}


       
