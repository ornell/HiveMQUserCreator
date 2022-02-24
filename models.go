package main

import "gorm.io/gorm"

type CcUser struct {
	gorm.Model
	ID                 uint   `gorm:"primaryKey,uniqueIndex"`
	Username           string `gorm:"unique,uniqueIndex"`
	Password           string
	PasswordIterations uint
	PasswordSalt       string
	Algorithm          string
}

type CcRole struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey,uniqueIndex"`
	Name        string `gorm:"unique,uniqueIndex"`
	Description string
}

type CcUserRole struct {
	gorm.Model
	UserId uint `gorm:"foreignKey:users_pkey"`
	RoleId uint `gorm:"foreignKey:roles_pkey"`
}

type CcPermission struct {
	gorm.Model
	ID                  uint   `gorm:"primaryKey,uniqueIndex"`
	Topic               string `gorm:"index"`
	PublishAllowed      bool
	SubscribeAllowed    bool
	Qos0Allowed         bool
	Qos1Allowed         bool
	Qos2Allowed         bool
	RetainedMsgsAllowed bool
	SharedSubAllowed    bool
	SharedGroup         string
}

type CcRolePermission struct {
	gorm.Model
	Role       uint `gorm:"not null,foreignKey:role_pkey"`
	Permission uint `gorm:"not null,foreignKey:permissions_pkey"`
}

type CcUserPermission struct {
	gorm.Model
	userId     uint `gorm:"not null,foreignKey:users_pkey"`
	permission uint `gorm:"foreignKey:permissions_pkey"`
}

//
//comment on column users.password is 'Base64 encoded raw byte array';
//comment on column users.password_salt is 'Base64 encoded raw byte array';
//comment on table permissions is 'All permissions are whitelist permissions';
