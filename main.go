package main

import (
	"database/sql"
	autoriz "diplom/src/autorization"
	models "diplom/src/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func main() {
	models.CreateDatabase()

	// Инициализация подключения к базе данных
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=diplom_rob sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	//r.StaticFS("/static", http.Dir("/home/lutik/Desktop/diplom/frontend"))
	r.LoadHTMLFiles(
		"/home/lutik/Desktop/diplom/frontend/login/index.html",
		"/home/lutik/Desktop/diplom/frontend/login/cards.html",

		// Другие шаблоны, которые нужно загрузить
	)
	r.StaticFile("/vendor/bootstrap/css/bootstrap.min.css", "./frontend/login/vendor/bootstrap/css/bootstrap.min.css")
	r.StaticFile("/fonts/font-awesome-4.7.0/css/font-awesome.min.css", "./frontend/login/fonts/font-awesome-4.7.0/css/font-awesome.min.css")
	r.StaticFile("/vendor/select2/select2.min.css", "./frontend/login/vendor/select2/select2.min.css")
	r.StaticFile("/vendor/css-hamburgers/hamburgers.min.css", "./frontend/login/vendor/css-hamburgers/hamburgers.min.css")
	r.StaticFile("/vendor/animate/animate.css", "./frontend/login/vendor/animate/animate.css")
	r.StaticFile("/css/util.css", "./frontend/login/css/util.css")
	r.StaticFile("/css/main.css", "./frontend/login/css/main.css")
	r.StaticFile("/images/img-01.png", "./frontend/login/images/img-01.png")
	r.StaticFile("/vendor/jquery/jquery-3.2.1.min.js", "./frontend/login/vendor/jquery/jquery-3.2.1.min.js")
	r.StaticFile("/vendor/bootstrap/js/popper.js", "./frontend/login/vendor/bootstrap/js/popper.js")
	r.StaticFile("/vendor/bootstrap/js/bootstrap.min.js", "./frontend/login/vendor/bootstrap/js/bootstrap.min.js")
	r.StaticFile("/vendor/select2/select2.min.js", "./frontend/login/vendor/select2/select2.min.js")
	r.StaticFile("/vendor/tilt/tilt.jquery.min.js", "./frontend/login/vendor/tilt/tilt.jquery.min.js")
	r.StaticFile("/js/main.js", "./frontend/login/js/main.js")
	r.StaticFile("/fonts/poppins/Poppins-Bold.ttf", "./frontend/login/fonts/poppins/Poppins-Bold.ttf")
	r.StaticFile("/poppins/Poppins-Medium.ttf", "./frontend/login/fonts/poppins/Poppins-Medium.ttf")
	r.StaticFile("/fonts/poppins/Poppins-Regular.ttf", "./frontend/login/fonts/poppins/Poppins-Regular.ttf")
	r.StaticFile("/fonts/montserrat/Montserrat-Bold.ttf", "./frontend/login/fonts/montserrat/Montserrat-Bold.ttf")
	r.StaticFile("/fonts/font-awesome-4.7.0/fonts/fontawesome-webfont.woff2", "./frontend/login/fonts/font-awesome-4.7.0/fonts/fontawesome-webfont.woff2?v=4.7.0")
	r.StaticFile("/fonts/font-awesome-4.7.0/fonts/fontawesome-webfont.woff", "./frontend/login/fonts/font-awesome-4.7.0/fonts/fontawesome-webfont.woff?v=4.7.0")
	r.StaticFile("/fonts/font-awesome-4.7.0/fonts/fontawesome-webfont.ttf", "./frontend/login/fonts/font-awesome-4.7.0/fonts/fontawesome-webfont.ttf?v=4.7.0")
	//loadTemplates(r, "frontend/login")
	//r.LoadHTMLGlob("frontend/****/***/**/*")
	//r.Static("/static", "./frontend")

	//r.Static("/images", "./frontend/login/images")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})

	})
	r.GET("/cards.html", func(c *gin.Context) {

		c.HTML(http.StatusOK, "cards.html", nil)
	})
	r.GET("/get_user_data", func(c *gin.Context) {
		autoriz.GetUserData(c, db)
	})

	r.POST("/signup", func(c *gin.Context) {
		autoriz.SignupHandler(c, db)
	})

	r.POST("/login", func(c *gin.Context) {
		autoriz.LoginHandler(c, db)
	})

	// r.GET("/mod", func(c *gin.Context) {
	// 	models.PrintUsers(c)
	// })
	r.POST("forgotpassword", func(c *gin.Context) {
		mail := c.PostForm("mail") // Здесь предполагается, что форма содержит поле "email" для ввода адреса электронной почты
		err := autoriz.SendPasswordResetEmail(mail)
		if err != nil {
			// Обработка ошибки отправки письма
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// Успешно отправлено
		c.JSON(http.StatusOK, gin.H{"message": "Письмо для сброса пароля отправлено на указанный адрес."})
	})

	// r.GET("/cards", GetCardsHandler)

	// // Нужно быть авторизованным, чтобы работать с эндпоинтами для последовательностей карточек
	// authRouter := r.Group("/sequences")
	// authRouter.Use(AuthenticationMiddleware)
	// {
	// 	authRouter.POST("", CreateSequenceHandler)
	// 	authRouter.GET("/:user_id", autoriz.GetSequencesHandler)
	// }

	log.Println("Starting server...")
	r.Run(":8080")
}

// func GetCardsHandler(c *gin.Context) {
// 	rows, err := db.Query("SELECT * FROM cards")
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return

// 		defer rows.Close()

// 		cards := []models.Card{}
// 		for rows.Next() {
// 			var card models.Card
// 			err := rows.Scan(&card.ID, &card.Name)
// 			if err != nil {
// 				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 				return
// 			}
// 			cards = append(cards, card)
// 		}

// 		c.JSON(http.StatusOK, cards)
// 	}

// }

// func CreateSequenceHandler(c *gin.Context) {
// 	claims := c.MustGet("claims").(*models.Claims)
// 	var sequence models.Sequence
// 	err := c.BindJSON(&sequence)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Проверяем, что пользователь может создавать последовательности только для себя
// 	userID := strconv.Itoa(sequence.UserID) // Преобразуем UserID в строку
// 	if claims.Username != "admin" && claims.Username != userID {
// 		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
// 		return
// 	}

// 	// Проверяем, что переданные id карт существуют в БД
// 	var count int
// 	err = db.QueryRow("SELECT COUNT(*) FROM cards WHERE id IN ($1, $2)", sequence.Card1ID, sequence.Card2ID).Scan(&count)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	if count != 2 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid card ids"})
// 		return
// 	}

// 	// Сохраняем новую последовательность в БД
// 	createdAt := time.Now().Unix()
// 	err = db.QueryRow("INSERT INTO sequences (user_id, card_1_id, card_2_id, created_at) VALUES ($1, $2, $3, $4) RETURNING id", sequence.UserID, sequence.Card1ID, sequence.Card2ID, createdAt).Scan(&sequence.ID)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, sequence)
// }

// func AuthenticationMiddleware(c *gin.Context) {
// 	authHeader := c.GetHeader("Authorization")
// 	if authHeader == "" {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
// 		return
// 	}

// 	tokenString := authHeader[len("Bearer "):]

// 	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
// 		return []byte("secret"), nil
// 	})
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
// 		return
// 	}

// 	claims, ok := token.Claims.(*models.Claims)
// 	if !ok || !token.Valid {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
// 		return
// 	}

// 	c.Set("claims", claims)

// 	c.Next()
// }
