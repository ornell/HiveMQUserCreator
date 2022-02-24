package main

import (
	"encoding/base64"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func main() {
	sslmode, present := os.LookupEnv("DB_SSLMODE")
	if present == false {
		sslmode = "disable"
	}
	dsn := "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " port=" + os.Getenv("DB_PORT") + " sslmode=" + sslmode
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	switch cmd := os.Args[1]; cmd {
	case "init":
		initDB(*db)
		fmt.Println("Database has been initialized")
	case "create":
		switch usertype := os.Args[2]; usertype {
		case "ccuser":
			createCommandCenterUser(db)
		case "mqttuser":
			// TODO create MQTT user
		default:
			// TODO create function with explination of how to use tooling.
		}
	default:
		// TODO create function with explination of how to use tooling.
	}
}

func createCommandCenterUser(db *gorm.DB) {

	user := CcUser{
		Model:              gorm.Model{},
		Username:           "ornell",
		Password:           base64.StdEncoding.EncodeToString([]byte("password123")),
		PasswordIterations: 100,
		PasswordSalt:       "IThEd9r1+ALx/d++tdcOmw==",
		Algorithm:          "SHA512",
	}
	result := db.Create(&user)
	fmt.Println(result.Statement)
}

func seedDB() {
	//insert into public.users
	//(username, password, password_iterations, password_salt, algorithm)
	//values
	//('backendservice', 'wtUo2dri+ttHGHRpngg9uG21piWLiKSX7IaNSnU/BfN9pt+ZOLQByG/3JlPPQ7t/pl8S3tjR2+Um/DPBdAQULg==', 100, 'Nv6NU9XY7tvHdSGaKmNTOw==', 'SHA512'),
	//('frontendclient', 'ZHg/rNJel1BHOYMEvc40ekCRUE5vVLcsPF6mk9GPDcdEmX3stm50MplaqjGb8Lxhy6rNFQZSQRSbOxmFZ8ps1Q==', 100, 'JhpW27QU9WfIaG6FJT5MkQ==', 'SHA512'),
	//('superuser', 'nOgr9xVnkt51Lr68KS/rAKm/LqxAt8oEki7vCerRod3qDbyMFfDBGT8obnkw+AGygxCQDWdaA2sQnXXoAbVK6Q==', 100, 'wxw+3diCV4bWXQHb6LLniA==', 'SHA512');
	//
	//insert into public.permissions
	//(id, topic, publish_allowed, subscribe_allowed, qos_0_allowed, qos_1_allowed, qos_2_allowed, retained_msgs_allowed, shared_sub_allowed, shared_group)
	//values
	//(1, 'topic/+/status', false, true, true, true, true, false, false, ''),
	//(2, 'topic/${mqtt-clientid}/status', true, false, true, true, true, true, false, ''),
	//(3, '#', true, true, true, true, true, true, true, '');
	//
	//insert into public.roles
	//(id, name, description)
	//values
	//(1, 'backendservice', 'only allowed to subscribe to topics'),
	//(2, 'frontendclients', 'only allowed to publish to topics'),
	//(3, 'superuser', 'is allowed to do everything');
	//
	//insert into public.user_roles
	//(user_id, role_id)
	//values
	//(1, 1),
	//(2, 2),
	//(3, 3);
	//
	//insert into public.role_permissions
	//(role, permission)
	//values
	//(1, 1),
	//(2, 2),
	//(3, 3);
}

// insert into users (username, password, password_iterations, password_salt, algorithm) values ('my_user', 'bSdf47hY52qPBShvJE+mgcGtuQyveNdGWtO11BSm6l6Bp6cicW3ulb9GYaVTxCLyuIbzOgb5VM6KysxXhNgGrA==', 100, 'IThEd9r1+ALx/d++tdcOmw==', 'SHA512');
func initDB(db gorm.DB) {
	// Migrate the schema
	err := db.AutoMigrate(&CcUser{})
	if err != nil {
		fmt.Println(err)
	}
	err = db.AutoMigrate(&CcRole{})
	if err != nil {
		fmt.Println(err)
	}
	err = db.AutoMigrate(&CcPermission{})
	if err != nil {
		fmt.Println(err)
	}
	err = db.AutoMigrate(&CcRolePermission{})
	if err != nil {
		fmt.Println(err)
	}
	err = db.AutoMigrate(&CcUserRole{})
	if err != nil {
		fmt.Println(err)
	}
	err = db.AutoMigrate(&CcUserPermission{})
	if err != nil {
		fmt.Println(err)
	}
}
