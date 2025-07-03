package userModels

import (
	"testApp/config"
)

type Users struct {
	IdUser uint   `json:"id_user"`
	Nama   string `json:"nama"`
	Email  string `json:"email"`
	Alamat string `json:"alamat"`
}

func (Users) TableName() string {
	return "tbl_user"
}

type PaginateUser struct {
	Users       []Users `json:"user"`
	TotalData   int     `json:"total_data"`
	TotalPage   int     `json:"total_page"`
	CurrentPage int     `json:"current_page"`
	Limit       int     `json:"limit"`
}

func GetDataUser(page, limit int) (PaginateUser, error) {
	var result PaginateUser
	var total int64
	offset := (page - 1) * limit

	// err := config.DB.QueryRow("SELECT COUNT(*) FROM tbl_user").Scan(&total)
	// if err != nil {
	// 	return result, err
	// }

	if err := config.DB.Model(&Users{}).Count(&total).Error; err != nil {
		return result, err
	}

	// rows, err := config.DB.Query("SELECT id_user, nama, email, alamat FROM tbl_user LIMIT ? OFFSET ?", limit, offset)
	// if err != nil {
	// 	return result, err
	// }
	// defer rows.Close()

	var users []Users
	if err := config.DB.Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return result, err
	}

	// for rows.Next() {
	// 	var user Users
	// 	if err := rows.Scan(&user.IdUser, &user.Nama, &user.Email, &user.Alamat); err != nil {
	// 		return result, err
	// 	}
	// 	users = append(users, user)
	// }
	result.Users = users
	result.TotalData = int(total)
	result.TotalPage = (int(total) + limit - 1) / limit
	result.CurrentPage = page
	result.Limit = limit

	return result, nil
}
