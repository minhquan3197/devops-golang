package structs

// UpdateUser struct
type UpdateUser struct {
	Username string `json:"username" valid:"required~Tên không được để trống,length(8|50)~Độ dài tối thiểu 8 kí tự"`
}

// CreateUser struct
type CreateUser struct {
	Username string `json:"username" valid:"required~Tên không được để trống,length(8|50)~Độ dài tối thiểu 8 kí tự"`
	Password string `json:"password" valid:"required~Mật khẩu không được để trống,length(8|50)~Độ dài tối thiểu 8 kí tự"`
}

// PaginateUser struct
type PaginateUser struct {
	Data    []UserSchema `json:"data"`
	PerPage int64        `json:"perPage"`
	Total   int64        `json:"total"`
}
