package domain

type Group struct {
	ID         uint64       `gorm:"primary_key:auto_increment" json:"id"`
	Name       string       `gorm:"type:varchar(255)" json:"name"`
	UsersGroup *[]UserGroup `gorm:"foreignKey:GroupID;references:ID" json:"users_groups,omitempty"`
}
