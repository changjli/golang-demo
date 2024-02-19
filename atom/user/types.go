package atom_user

import (
	"time"
)

type UserDatasModel struct {
	UserDatas []UserDataModel
}

// get
type UserDataModel struct {
	Id_user          int       `json:"id_user"`
	Employee_name    string    `json:"employee_name" binding:"required"`
	Employee_address string    `json:"employee_address" binding:"required"`
	Employee_postal  int       `json:"employee_postal" binding:"required"`
	Created_dt       time.Time `json:"created_dt"`
	Created_by       string    `json:"created_by" binding:"required"`
	Updated_dt       time.Time `json:"updated_dt"`
	Updated_by       string    `json:"updated_by"`
	Deleted_dt       time.Time `json:"deleted_dt"`
	Deleted_by       string    `json:"deleted_by"`
	Is_active        bool      `json:"is_active" binding:"required"`
}

// create
type InputUserModel struct {
	Employee_name    string `json:"employee_name" binding:"required"`
	Employee_address string `json:"employee_address" binding:"required"`
	Employee_postal  int    `json:"employee_postal" binding:"required"`
	Created_by       string `json:"created_by" binding:"required"`
}

// update
type UpdateUserModel struct {
	Id_User          int       `json:"id_user"`
	Employee_name    string    `json:"employee_name" binding:"required"`
	Employee_address string    `json:"employee_address" binding:"required"`
	Employee_postal  int       `json:"employee_postal" binding:"required"`
	Updated_dt       time.Time `json:"updated_dt"`
	Updated_by       string    `json:"updated_by" binding:"required"`
}
