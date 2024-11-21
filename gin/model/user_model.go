package model

type Login struct { // Binding from JSON
	// binding:约束，required必填项 , min字符段最短为3
	User     string `form:"user" json:"user" binding:"required,min=3,max=8"`
	Password string `form:"password" json:"password" binding:"required"`
}

type SignUp struct {
	Age        uint   `json:"age" binding:"gte=1,lte=130"`                     // 年龄，必须大于等于 1，小于等于 130
	Name       string `json:"name" binding:"required"`                         // 用户名，必填字段
	Email      string `json:"email" binding:"required, email"`                 // 邮箱地址，必填字段且必须符合电子邮件格式
	Password   string `json:"password" binding:"required"`                     // 密码，必填字段
	RePassword string `json:"re_password" binding:"required,eqfield=Password"` // 确认密码，必填字段且必须与密码字段的值相同
}
