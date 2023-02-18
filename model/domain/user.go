package domain

type User struct {
	ID         uint64       `gorm:"primary_key:auto_increment" json:"id"`
	Name       string       `gorm:"type:varchar(255)" json:"name"`
	Email      string       `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	UserGroups *[]UserGroup `gorm:"foreignKey:UserID;references:ID" json:"user_groups,omitempty"`
}
