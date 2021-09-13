package services

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/suraboy/go-fiber-api/config"
	"github.com/suraboy/go-fiber-api/models"
)

func GetAllUser(c *fiber.Ctx) error {
	db := config.PostgresConnection()

	var user []models.Users
	result := db.First(&user)
	return c.JSON(fiber.Map{"data": result})

}

func CreateNewProduct(c *fiber.Ctx) error {
	//	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	//		"password=%s dbname=%s sslmode=disable",
	//		os.Getenv("POSTGRES_HOST"),
	//		5432,
	//		os.Getenv("POSTGRES_USER"),
	//		os.Getenv("POSTGRES_PASSWORD"),
	//		os.Getenv("POSTGRES_DATABASE"))
	//
	//	conn, err := sql.Open("postgres", psqlInfo)
	//	if err != nil {
	//		log.Fatalf("cannot open postgres connection:%s", err)
	//	}
	//
	//	defer conn.Close()
	//	// insert
	//	insertStmt := `insert into "users"("Name", "Roll") values('John', 1)`
	//	_, errInsert := conn.Exec(insertStmt)
	//
	//	if errInsert != nil {
	//		log.Fatalf("cannot open postgres create:%s", errInsert)
	//	}
	//
	return c.JSON(fiber.Map{"data": "done"})
}
