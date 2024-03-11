package model

type RoleList struct {
	RoleId    int       `json:"roleId" xorm:"id"`
	Name      string    `json:"name" xorm:"name"`
	Status    int       `json:"status" xorm:"status"`
}

func (rl RoleList) TableName() string { return "gf_role" }

type RolePassword struct {
	RoleId      int      `json:"roleId" xorm:"role_id"`
	Password    string   `json:"password" xorm:"password"`
}