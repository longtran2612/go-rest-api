package domain

type UserGroup struct {
	UserID  uint64 `gorm:"primary_key;uniqueIndex:idx_userid_groupid_role" json:"user_id"`
	GroupID uint64 `gorm:"primary_key:uniqueIndex:idx_userid_groupid_role" json:"group_id"`
	Role    string `gorm:"type:varchar(255);uniqueIndex:idx_userid_groupid_role" json:"role"`
	User    User   `gorm:"foreignKey:ID;references:UserID" json:"user,omitempty"`
	Group   Group  `gorm:"foreignKey:ID;references:GroupID" json:"group,omitempty"`
}
