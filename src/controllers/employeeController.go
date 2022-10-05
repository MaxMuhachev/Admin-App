package controllers

import (
	"content/src/app"
	"content/src/models"
	"content/src/storage"
	"content/src/utils"
	"net/http"
)

func HandlerEmployees(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieAdminWriteHeader(w, r) {
		res, err := GetEmployees()
		app.RenderTemplate(w, "employees/content-employees", &app.Page{Title: utils.EMPLOYEES, EmployeeList: res}, &err)
	}
}

func GetEmployees() ([]*models.Employee, error) {
	connect := app.NewConnect()

	var res []*models.Employee

	rows, err := connect.Mysql.Queryx(storage.GetEmployees)
	if err != nil {
		return nil, err
	} else {
		for rows.Next() {
			var employee = models.Employee{}
			err := rows.StructScan(&employee)
			if err != nil {
				return nil, err
			}
			res = append(res, &employee)
		}
	}

	return res, err
}

func HandlerEditEmployee(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieAdminWriteHeader(w, r) {
		r.ParseForm()
		email := r.Form.Get("email")

		employee, err := GetEmployeeById(email)
		app.RenderTemplate(w, "employees/edit/content-edit-employee", &app.Page{Title: utils.EMPLOYEES, Employee: employee}, &err)
	}
}

func HandlerCreateEmployee(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieAdminWriteHeader(w, r) {
		app.RenderTemplate(w, "employees/edit/content-edit-employee", &app.Page{}, nil)
	}
}

func HandlerCreatePostEmployee(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieAdminWriteHeader(w, r) {
		employee := ParseEmployeeForm(r)

		connect := app.NewConnect()

		_, err := connect.Mysql.Queryx(
			storage.CreateEmployee,
			employee.Email,
			employee.FIO,
			employee.Login,
			employee.Password,
			employee.Floor,
			employee.Address,
			employee.Birthday,
			employee.CompanyID,
		)
		if !utils.ThrowError(err, w) {
			r.Form.Set("oldEmail", employee.Email)
			http.Redirect(w, r, "/employees/edit", http.StatusTemporaryRedirect)
		}

	}
}

func GetEmployeeById(email string) (*models.Employee, error) {
	connect := app.NewConnect()

	var res *models.Employee

	rows, err := connect.Mysql.Queryx(storage.GetEmployeeByEmail, email)
	if err != nil {
		return nil, err
	} else {
		for rows.Next() {
			var employee = models.Employee{}
			err = rows.StructScan(&employee)
			if err != nil {
				return nil, err
			}
			res = &employee
		}
	}

	return res, nil
}

func GetEmployeeByEmailPassword(email string, password string) (*models.Employee, error) {
	connect := app.NewConnect()

	var res *models.Employee

	rows, err := connect.Mysql.Queryx(storage.GetEmployeeByEmailPassword, email, password)
	if err != nil {
		return nil, err
	} else {
		for rows.Next() {
			var employee = models.Employee{}
			err = rows.StructScan(&employee)
			if err != nil {
				return nil, err
			}
			res = &employee
		}
	}

	return res, nil
}

func HandlerEditPostEmployee(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieAdminWriteHeader(w, r) {
		employee := ParseEmployeeForm(r)
		oldEmail := r.Form.Get("oldEmail")
		connect := app.NewConnect()

		_, err := connect.Mysql.Queryx(
			storage.UpdateEmployee,
			employee.Email,
			employee.FIO,
			employee.Login,
			employee.Password,
			employee.Floor,
			employee.Address,
			employee.Birthday,
			employee.CompanyID,
			oldEmail,
		)

		var updatedEmployee *models.Employee
		if err == nil {
			updatedEmployee, err = GetEmployeeById(employee.Email)
		}

		app.RenderTemplate(w, "employees/edit/content-edit-employee", &app.Page{Title: utils.EMPLOYEES, Employee: updatedEmployee, Success: utils.EMPLOYEE_SAVED}, &err)
	}
}

func HandlerDeletePostEmployee(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieAdminWriteHeader(w, r) {
		r.ParseForm()
		employeeId := r.Form.Get("id")

		connect := app.NewConnect()

		_, err := connect.Mysql.Queryx(
			storage.DeleteEmployee,
			employeeId,
		)
		if !utils.ThrowError(err, w) {
			w.Write([]byte("1"))
			w.WriteHeader(http.StatusOK)
		}

	}
}

func ParseEmployeeForm(r *http.Request) *models.Employee {
	r.ParseForm()
	fio := r.Form.Get("FIO")
	email := r.Form.Get("email")
	birthday := r.Form.Get("birthday")
	address := r.Form.Get("address")
	floor := r.Form.Get("floor")
	login := r.Form.Get("login")
	password := r.Form.Get("password")
	company := r.Form.Get("company")

	var employee = models.Employee{}
	employee.FIO = fio
	employee.Email = email
	employee.Login = login
	employee.Password = password
	employee.Floor = floor
	employee.Address = address
	employee.Birthday = birthday
	employee.CompanyID = uint8(utils.ConvertUint(company))
	return &employee
}
