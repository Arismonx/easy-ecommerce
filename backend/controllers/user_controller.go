package controllers

// func CreateUser(c *fiber.Ctx) error {
// 	newUser := new(models.Users)

// 	if err := c.BodyParser(newUser); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": "Cannot parse JSON",
// 		})
// 	}

// 	if err := config.DB.Create(newUser); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": "Cannot parse User",
// 		})
// 	}
// 	return c.JSON(newUser)
// }
