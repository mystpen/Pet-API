package redis

import (
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Middleware для Basic Authentication
func BasicAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получаем заголовок "Authorization" из запроса
		authHeader := c.GetHeader("Authorization")

		// Проверяем, что заголовок Authorization передан и начинается с "Basic "
		if authHeader == "" || len(authHeader) < 6 || authHeader[:6] != "Basic " {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Необходима авторизация"})
			return
		}

		// Получаем и декодируем значение Basic Auth
		authToken := authHeader[6:]
		decodedToken, err := base64.StdEncoding.DecodeString(authToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Ошибка декодирования токена"})
			return
		}

		// Разделяем имя пользователя и пароль
		credentials := string(decodedToken)
		username, password, ok := c.Request.BasicAuth()

		// Проверяем, совпадают ли полученные учетные данные с теми, что переданы в Basic Auth
		if !ok || credentials != (username+":"+password) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Неверные учетные данные"})
			return
		}

		// Если учетные данные верны, передаем управление следующему обработчику
		c.Next()
	}
}

// Обработчик для запросов на регистрацию
func SignUpHandler(c *gin.Context) {
	// Здесь обрабатываем запрос на регистрацию
	// Например, извлекаем данные из тела запроса и добавляем нового пользователя в базу данных

	// Возвращаем успешный ответ
	c.JSON(http.StatusOK, gin.H{
		"message": "Регистрация прошла успешно",
	})
}
