package atom_user

import (
	"errors"
	"log"
)

func GetAllDataUserUseCase() ([]UserDataModel, error) {
	getDatas, err := GetAllDataUserDB()
	if err != nil {
		return []UserDataModel{}, err
	}

	if len(getDatas) < 1 {
		log.Println("[atom][user][resource][GetAllDataUserUseCase] no rows returned")
		return getDatas, errors.New(`no data returned`)
	}

	return getDatas, nil
}

func InsertUserUseCase(inputData InputUserModel) ([]UserDataModel, error) {
	// validasi
	getDatas, err := GetUserByUsername(inputData.Employee_name)
	if err != nil {
		return []UserDataModel{}, err
	}

	if len(getDatas) >= 1 {
		log.Println("[atom][user][resource][InsertUserUseCase] username already exists")
		return []UserDataModel{}, errors.New(`username already exists`)
	} else {
		err := InsertUserDB(inputData)

		if err != nil {
			return []UserDataModel{}, errors.Join()
		}

		getData, err := GetAllDataUserDB()
		if err != nil {
			return []UserDataModel{}, err
		}

		return getData, nil
	}
}

func GetUserByUsernameCaseV2(username string) (UserDataModel, error) {
	getDatas, err := GetUserByUsername(username)
	if err != nil {
		return UserDataModel{}, err
	}

	if len(getDatas) < 1 {
		log.Println("[atom][user][resource][GetUserByUsernameUseCase] no rows returned")
		return UserDataModel{}, errors.New(`no data returned`)
	}

	return getDatas[0], nil
}

func GetUserByUsernameUseCase(inputData InputUserModel) ([]UserDataModel, error) {
	getDatas, err := GetUserByUsername(inputData.Employee_name)
	if err != nil {
		return []UserDataModel{}, err
	}

	if len(getDatas) < 1 {
		log.Println("[atom][user][resource][GetUserByUsernameUseCase] no rows returned")
		return []UserDataModel{}, errors.New(`no data returned`)
	}

	return getDatas, nil
}

func UpdateUserByIdUseCase(updateData UpdateUserModel) ([]UserDataModel, error) {

	// validasi usernya ada ga
	getId, err := GetUserById(updateData.Id_User)
	if err != nil {
		return []UserDataModel{}, err
	}

	// validasi username unique ga
	getUsername, err := GetUserByUsername(updateData.Employee_name)
	if err != nil {
		return []UserDataModel{}, err
	}

	// validasi user active ga
	if !getId[0].Is_active {
		return []UserDataModel{}, errors.New(`user status deactivated`)
	}

	if len(getId) < 1 {
		log.Println("[atom][user][resource][InsertUserUseCase] user doesnt exist")
		return []UserDataModel{}, errors.New(`user doesnt exist`)
	}

	if len(getUsername) >= 1 {
		log.Println("[atom][user][resource][InsertUserUseCase] username already exists")
		return []UserDataModel{}, errors.New(`username already exists`)
	}

	err = UpdateUserById(updateData)
	if err != nil {
		return []UserDataModel{}, errors.Join()
	}

	getData, err := GetAllDataUserDB()
	if err != nil {
		return []UserDataModel{}, err
	}

	return getData, nil

}

func UpdateUserByIdUseCaseV2(inputData InputUserModel, id int) ([]UserDataModel, error) {

	// validasi
	getDatas, err := GetUserByUsername(inputData.Employee_name)
	if err != nil {
		return []UserDataModel{}, err
	}

	if len(getDatas) >= 1 {
		log.Println("[atom][user][resource][InsertUserUseCase] username already exists")
		return []UserDataModel{}, errors.New(`username already exists`)
	} else {
		err := UpdateUserByIdV2(inputData, id)

		if err != nil {
			return []UserDataModel{}, errors.Join()
		}

		getData, err := GetAllDataUserDB()
		if err != nil {
			return []UserDataModel{}, err
		}

		return getData, nil
	}
}

func DeleteUserByUseCase(id int) ([]UserDataModel, error) {
	err := DeleteUserById(id)

	if err != nil {
		return []UserDataModel{}, errors.Join()
	}

	getData, err := GetAllDataUserDB()
	if err != nil {
		return []UserDataModel{}, err
	}

	return getData, nil
}

func DeactiveUserByUseCase(updateData UpdateUserModel) ([]UserDataModel, error) {
	// validasi usernya ada ga
	getId, err := GetUserById(updateData.Id_User)
	if err != nil {
		return []UserDataModel{}, err
	}

	if len(getId) < 1 {
		log.Println("[atom][user][resource][DeactivateUserUseCase] user doesnt exist")
		return []UserDataModel{}, errors.New(`user doesnt exist`)
	}

	if getId[0].Is_active {
		err = DeactivateUser(updateData.Id_User)
	} else {
		err = ActivateUser(updateData.Id_User)
	}

	if err != nil {
		return []UserDataModel{}, errors.Join()
	}

	getData, err := GetAllDataUserDB()
	if err != nil {
		return []UserDataModel{}, err
	}

	return getData, nil
}
