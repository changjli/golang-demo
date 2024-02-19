package atom_user

import (
	"fmt"
	"log"
	user_db_postgresql "user_service/config/db/postgresql"
)

// get all user
func GetAllDataUserDB() ([]UserDataModel, error) {
	db := user_db_postgresql.OpenConnection()

	// defer baru jalan setelah function selesai
	defer db.Close()

	queryString := `SELECT * FROM user_employees`

	rows, err := db.Query(queryString)
	if err != nil {
		log.Println("[atom][user][resource_db][GetAllDataUserDB] errors in query row", err.Error())
		return []UserDataModel{}, err
	}

	defer rows.Close()

	var datas []UserDataModel
	for rows.Next() {
		var data UserDataModel
		rows.Scan(
			&data.Id_user,
			&data.Employee_name,
			&data.Employee_address,
			&data.Employee_postal,
			&data.Created_dt,
			&data.Created_by,
			&data.Updated_dt,
			&data.Updated_by,
			&data.Deleted_dt,
			&data.Deleted_by,
			&data.Is_active,
		)
		datas = append(datas, data)
	}

	return datas, nil

}

func InsertUserDB(inputData InputUserModel) error {
	db := user_db_postgresql.OpenConnection()
	defer db.Close()

	// template literals dari golang
	// queryString := `
	// 	INSERT INTO "user_employees"
	// 	("employee_name")
	// 	VALUES
	// 	('%s');
	// `
	// queryStringFormat = fmt.Sprintf(queryString, inputData.Employee_name)

	// template literals dari db udah otomatis dicleaning
	queryFormat := `
		INSERT INTO "user_employees"
		("employee_name", "employee_address", "employee_postal", "created_by")
		VALUES 
		($1, $2, $3, $4); 
	`

	_, err := db.Exec(queryFormat, inputData.Employee_name, inputData.Employee_address, inputData.Employee_postal, inputData.Created_by)

	if err != nil {
		log.Println("[atom][user][resource_db][InsertUserDB] errors in insert row", err.Error())
		return err
	}

	return nil
}

func GetUserByUsername(username string) ([]UserDataModel, error) {
	db := user_db_postgresql.OpenConnection()

	defer db.Close()

	queryString := `SELECT * FROM "user_employees" WHERE "employee_name" = '%s'`

	queryStringFormat := fmt.Sprintf(queryString, username)

	rows, err := db.Query(queryStringFormat)
	if err != nil {
		log.Println("[atom][user][resource_db][GetUserByUsername] errors in query row", err.Error())
		return []UserDataModel{}, err
	}
	defer rows.Close()

	var datas []UserDataModel
	for rows.Next() {
		var data UserDataModel
		rows.Scan(
			&data.Id_user,
			&data.Employee_name,
			&data.Employee_address,
			&data.Employee_postal,
			&data.Created_dt,
			&data.Created_by,
			&data.Updated_dt,
			&data.Updated_by,
			&data.Deleted_dt,
			&data.Deleted_by,
			&data.Is_active,
		)
		datas = append(datas, data)
	}

	return datas, nil
}

func GetUserById(id int) ([]UserDataModel, error) {
	db := user_db_postgresql.OpenConnection()

	defer db.Close()

	queryString := `SELECT * FROM "user_employees" WHERE "id_user" = %d`

	queryStringFormat := fmt.Sprintf(queryString, id)

	rows, err := db.Query(queryStringFormat)
	if err != nil {
		log.Println("[atom][user][resource_db][GetUserByUsername] errors in query row", err.Error())
		return []UserDataModel{}, err
	}
	defer rows.Close()

	var datas []UserDataModel
	for rows.Next() {
		var data UserDataModel
		rows.Scan(
			&data.Id_user,
			&data.Employee_name,
			&data.Employee_address,
			&data.Employee_postal,
			&data.Created_dt,
			&data.Created_by,
			&data.Updated_dt,
			&data.Updated_by,
			&data.Deleted_dt,
			&data.Deleted_by,
			&data.Is_active,
		)
		datas = append(datas, data)
	}

	return datas, nil
}

func UpdateUserById(updateData UpdateUserModel) error {
	db := user_db_postgresql.OpenConnection()
	defer db.Close()

	queryString :=
		`
		UPDATE "user_employees" 
		SET "employee_name" = '%s', 
		"employee_address" = '%s', 
		"employee_postal" = '%d', 
		"updated_dt" = CURRENT_TIMESTAMP,
		"updated_by" = '%s'
		WHERE "id_user" = %d;
	`

	queryFormat := fmt.Sprintf(queryString, updateData.Employee_name, updateData.Employee_address, updateData.Employee_postal, updateData.Updated_by, updateData.Id_User)

	log.Println(queryFormat)

	_, err := db.Exec(queryFormat)
	if err != nil {
		log.Println("[atom][user][resource_db][UpdateUserById] errors in insert row", err.Error())
		return err
	}

	return nil
}

func DeactivateUser(id int) error {
	db := user_db_postgresql.OpenConnection()
	defer db.Close()

	queryString :=
		`
		UPDATE "user_employees" 
		SET "is_active" = false
		WHERE "id_user" = %d;
	`

	queryFormat := fmt.Sprintf(queryString, id)

	log.Println(queryFormat)

	_, err := db.Exec(queryFormat)
	if err != nil {
		log.Println("[atom][user][resource_db][UpdateUserById] errors in insert row", err.Error())
		return err
	}

	return nil
}

func ActivateUser(id int) error {
	db := user_db_postgresql.OpenConnection()
	defer db.Close()

	queryString :=
		`
		UPDATE "user_employees" 
		SET "is_active" = true
		WHERE "id_user" = %d;
	`

	queryFormat := fmt.Sprintf(queryString, id)

	log.Println(queryFormat)

	_, err := db.Exec(queryFormat)
	if err != nil {
		log.Println("[atom][user][resource_db][UpdateUserById] errors in insert row", err.Error())
		return err
	}

	return nil
}

func UpdateUserByIdV2(inputData InputUserModel, id int) error {
	db := user_db_postgresql.OpenConnection()
	defer db.Close()

	log.Println(id)

	queryString :=
		`
		UPDATE "user_employees" 
		SET "employee_name" = '%s'
		WHERE "id_user" = %d;
	`

	queryFormat := fmt.Sprintf(queryString, inputData.Employee_name, id)

	log.Println(queryFormat)

	_, err := db.Exec(queryFormat)

	if err != nil {
		log.Println("[atom][user][resource_db][UpdateUserById] errors in insert row", err.Error())
		return err
	}

	return nil
}

func DeleteUserById(id int) error {
	db := user_db_postgresql.OpenConnection()
	defer db.Close()

	queryString :=
		`
		DELETE FROM "user_employees" 
		WHERE "id_user" = %d;
	`

	queryFormat := fmt.Sprintf(queryString, id)

	log.Println(queryFormat)

	_, err := db.Exec(queryFormat)

	if err != nil {
		log.Println("[atom][user][resource_db][UpdateUserById] errors in insert row", err.Error())
		return err
	}

	return nil
}
