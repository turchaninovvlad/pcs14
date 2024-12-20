package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

// Product представляет продукт
type Product struct {
	ID          int
	ImageURL    string
	Name        string
	Description string
	Price       float64
}

// User представляет пользователя
type User struct {
	ID        int
	Password  string
	Name      string
	Surname   string
	Mail      string
	Phone     string
	Favorites []int
	InCart    []int
}
type FavoriteRequest struct {
	UserID    int `json:"user_id"`
	ProductID int `json:"product_id"`
}

// Пример списка продуктов
var products = []Product{
	{ID: 1, ImageURL: "https://ir.ozone.ru/s3/multimedia-1-r/wc1000/7147386879.jpg", Name: "Acer Nitro V ANV15-41", Description: "Acer Nitro V ANV15-41, AMD Ryzen 7 7735HS Игровой ноутбук 15.6\", AMD Ryzen 7 7735HS, RAM 16 ГБ, SSD 512 ГБ, NVIDIA GeForce RTX 3050 (6 Гб), Без системы, (NH.QSHER.002), черный, Русская раскладка", Price: 81690},
	{ID: 2, ImageURL: "https://ir.ozone.ru/s3/multimedia-1-g/wc1000/7050006328.jpg", Name: "Lenovo LOQ 15IAX9", Description: "Lenovo LOQ 15IAX9 Игровой ноутбук 15.6\", Intel Core i5-12450HX, RAM 16 ГБ, SSD 512 ГБ, NVIDIA GeForce RTX 4050 для ноутбуков (6 Гб), Без системы, (83GS00EPRK), серебристый, Русская раскладка.", Price: 84540},
	{ID: 3, ImageURL: "https://ir.ozone.ru/s3/multimedia-1-q/wc1000/7126229138.jpg", Name: "Lenovo LOQ 15IRX9", Description: "Lenovo LOQ 15IRX9 Игровой ноутбук 15.6\", Intel Core i7-13650HX, RAM 16 ГБ, SSD 1024 ГБ, NVIDIA GeForce RTX 4060 (8 Гб), Без системы, (83DV00NJRK), серый, Русская раскладка", Price: 104490},
	{ID: 4, ImageURL: "https://ir.ozone.ru/s3/multimedia-1-a/wc1000/7126230382.jpg", Name: "Lenovo Legion 5 16IRX9", Description: "Lenovo Legion 5 16IRX9 Игровой ноутбук 16\", Intel Core i7-14650HX, RAM 32 ГБ, SSD 1024 ГБ, NVIDIA GeForce RTX 4070 для ноутбуков (8 Гб), Без системы, (83DG00E0RK), серебристый, Русская раскладка.", Price: 151990},
	{ID: 5, ImageURL: "https://ir.ozone.ru/s3/multimedia-1-7/wc1000/7076666203.jpg", Name: "Ninkear Super G16 Pro", Description: "Ninkear Super G16 Pro Игровой ноутбук 16\", Intel Core i9-10885H, RAM 32 ГБ, SSD 1024 ГБ, NVIDIA GeForce GTX 1650 Ti (4 Гб), Windows Pro, серый металлик, Русская раскладка", Price: 77732},
	{ID: 6, ImageURL: "https://ir.ozone.ru/s3/multimedia-n/wc1000/6834200027.jpg", Name: "VETAS 2024 ", Description: "VETAS 2024 Новое Последний выпуск Windows была активирована Игровой ноутбук 15.6\", Intel Celeron N5095, RAM 16 ГБ, SSD 512 ГБ, Intel UHD Graphics 750, Windows Pro, (N5905), серебристый, Русская раскладка.", Price: 21473},
	{ID: 7, ImageURL: "https://ir.ozone.ru/s3/multimedia-1-c/wc1000/7152362724.jpg", Name: "N4000", Description: "N4000 Игровой ноутбук 15\", Intel Celeron N4000C, RAM 16 ГБ, SSD, Windows Pro, (M66-1), черно-серый, прозрачный, Русская раскладка", Price: 16504},
	{ID: 8, ImageURL: "https://ir.ozone.ru/s3/multimedia-v/wc1000/6776590459.jpg", Name: "UZZAI Lenovo Por x50", Description: "UZZAI Lenovo Por x50 Игровой ноутбук 15.6\", Intel Celeron N5095, RAM 24 ГБ, SSD, Intel HD Graphics 610, Windows Pro, (SC-976), черный, оливковый, Русская раскладка", Price: 23260},
	{ID: 9, ImageURL: "https://ir.ozone.ru/s3/multimedia-1-7/wc1000/7034232355.jpg", Name: "TANSHI X15F RTX3050", Description: "TANSHI X15F RTX3050, RAM и SSD с возможностью расширения, новинка 2024 года Игровой ноутбук 15.6\", AMD Ryzen 5 6600H, RAM 16 ГБ, SSD 512 ГБ, NVIDIA GeForce RTX 3050 для ноутбуков (4 Гб), Linux, черный, Русская раскладка", Price: 71780},
	{ID: 10, ImageURL: "https://ir.ozone.ru/s3/multimedia-1-1/wc1000/7152993169.jpg", Name: "Lenovo Legion Pro 5 16IRX9", Description: "Lenovo Legion Pro 5 16IRX9 Игровой ноутбук 16\", Intel Core i7-14650HX, RAM 32 ГБ, SSD 1024 ГБ, NVIDIA GeForce RTX 4060 (8 Гб), Без системы, (83DF00E3RK), серый, Русская раскладка", Price: 182900},
	{ID: 11, ImageURL: "https://ir.ozone.ru/s3/multimedia-1-y/wc1000/7142706394.jpg", Name: "VANWIN N156", Description: "VANWIN N156 Игровой ноутбук 15.6\", Intel N95, RAM 16 ГБ, SSD 512 ГБ, Intel UHD Graphics 770, Windows Pro, (ноутбук для работы и учебы), черный, Русская раскладка", Price: 32500},
	{ID: 12, ImageURL: "https://ir.ozone.ru/s3/multimedia-1-4/wc1000/7152993172.jpg", Name: "Lenovo Legion 7 16IRX9", Description: "Lenovo Legion 7 16IRX9 Игровой ноутбук 16\", Intel Core i7-14700HX, RAM 32 ГБ, SSD 1024 ГБ, NVIDIA GeForce RTX 4060 (8 Гб), Без системы, (83FD007DRK), черный, Русская раскладка", Price: 210990},
	{ID: 13, ImageURL: "https://ir.ozone.ru/s3/multimedia-1-a/wc1000/7057184662.jpg", Name: "ASUS TUF Gaming A15 FA506NC-HN065", Description: "ASUS TUF Gaming A15 FA506NC-HN065 Игровой ноутбук, RAM 16 ГБ, черный", Price: 73566},
	{ID: 14, ImageURL: "https://ir.ozone.ru/s3/multimedia-r/wc1000/6834200067.jpg", Name: "VETAS 2024", Description: "VETAS 2024 Новое Последний выпуск Windows активирована Игровой ноутбук 15.6\", Intel Celeron N5095, RAM 32 ГБ, SSD 1024 ГБ, Intel UHD Graphics 750, Windows Pro, ( N5095), серебристый, Русская раскладка", Price: 31790},
	{ID: 15, ImageURL: "https://ir.ozone.ru/s3/multimedia-1-5/wc1000/7134536489.jpg", Name: "Lenovo LOQ 3 Series 15IAX9", Description: "Lenovo LOQ 3 Series 15IAX9 Игровой ноутбук 15.6\", Intel Core i5-12450HX, RAM 16 ГБ, SSD, NVIDIA GeForce RTX 4050 для ноутбуков (6 Гб), Без системы, (LOQ 3 Series 15IAX9), серый, Английская раскладка", Price: 112900},
}

type Message struct {
	SenderID   int    `json:"sender_id"`
	ReceiverID int    `json:"receiver_id"`
	Timestamp  string `json:"timestamp"`
	Text       string `json:"text"`
}

// Список сообщений
var messages []Message

var users = []User{}

func getProductsHandler(w http.ResponseWriter, r *http.Request) {
	// Устанавливаем заголовки для правильного формата JSON
	w.Header().Set("Content-Type", "application/json")
	// Преобразуем список заметок в JSON
	json.NewEncoder(w).Encode(products)
}

// обработчик для POST-запроса, добавляет продукт
func createProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	fmt.Println("*")
	var newProduct Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		fmt.Println("Error decoding request body:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Received new Product: %+v\n", newProduct)
	var lastID int = len(products)

	for _, productItem := range products {
		if productItem.ID > lastID {
			lastID = productItem.ID
		}
	}
	newProduct.ID = lastID + 1
	products = append(products, newProduct)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newProduct)
}

// Регистрация нового пользователя
func registerUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Генерация нового ID
	newUser.ID = len(users) + 1
	users = append(users, newUser)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

// Проверка пользователя (вход)
func loginUserHandler(w http.ResponseWriter, r *http.Request) {
	print()
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var credentials struct {
		Mail     string
		Password string
	}
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(credentials)
	for _, user := range users {
		if user.Mail == credentials.Mail && user.Password == credentials.Password {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	fmt.Println(credentials)
	http.Error(w, "Invalid credentials", http.StatusUnauthorized)
}

// Добавление товара в избранное
func addToFavoritesHandler(w http.ResponseWriter, r *http.Request) {
	userID, productID, err := parseIDsFromRequest(r)
	print(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for i, user := range users {
		if user.ID == userID {
			users[i].Favorites = append(users[i].Favorites, productID)
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)
}

// Удаление товара из избранного
func removeFromFavoritesHandler(w http.ResponseWriter, r *http.Request) {

	userID, productID, err := parseIDsFromRequest(r)
	fmt.Println(userID, productID, err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, user := range users {
		if user.ID == userID {
			for j, id := range user.Favorites {
				if id == productID {
					users[i].Favorites = append(user.Favorites[:j], user.Favorites[j+1:]...)
					w.WriteHeader(http.StatusOK)
					return
				}
			}
		}
	}

	http.Error(w, "User or product not found", http.StatusNotFound)
}

// Добавление товара в корзину
func addToCartHandler(w http.ResponseWriter, r *http.Request) {
	userID, productID, err := parseIDsFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, user := range users {
		if user.ID == userID {
			users[i].InCart = append(users[i].InCart, productID)
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)
}

// Удаление товара из корзины
func removeFromCartHandler(w http.ResponseWriter, r *http.Request) {
	userID, productID, err := parseIDsFromRequest(r)
	print(userID, productID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, user := range users {
		if user.ID == userID {
			for j, id := range user.InCart {
				if id == productID {
					users[i].InCart = append(user.InCart[:j], user.InCart[j+1:]...)
					w.WriteHeader(http.StatusOK)
					return
				}
			}
		}
	}

	http.Error(w, "User or product not found", http.StatusNotFound)
}
func getProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем ID из URL
	idStr := r.URL.Path[len("/Products/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	// Ищем продукт с данным ID
	for _, Product := range products {
		if Product.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(Product)
			return
		}
	}

	// Если продукт не найден
	http.Error(w, "Product not found", http.StatusNotFound)
}

// удаление продукта по id
func deleteProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Получаем ID из URL
	idStr := r.URL.Path[len("/Products/delete/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	// Ищем и удаляем продукт с данным ID
	for i, Product := range products {
		if Product.ID == id {
			// Удаляем продукт из среза
			products = append(products[:i], products[i+1:]...)
			w.WriteHeader(http.StatusNoContent) // Успешное удаление, нет содержимого
			return
		}
	}

	// Если продукт не найден
	http.Error(w, "Product not found", http.StatusNotFound)
}

// Обновление продукта по id
func updateProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Получаем ID из URL
	idStr := r.URL.Path[len("/Products/update/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	// Декодируем обновлённые данные продукта
	var updatedProduct Product
	err = json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Ищем продукт для обновления
	for i, Product := range products {
		if Product.ID == id {

			products[i].ImageURL = updatedProduct.ImageURL
			products[i].Name = updatedProduct.Name
			products[i].Description = updatedProduct.Description
			products[i].Price = updatedProduct.Price

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(products[i])
			return
		}
	}

	// Если продукт не найден
	http.Error(w, "Product not found", http.StatusNotFound)
}

// Парсинг ID пользователя и товара из запроса

func parseIDsFromRequest(r *http.Request) (int, int, error) {
	var req FavoriteRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return 0, 0, errors.New("invalid JSON format")
	}

	if req.UserID <= 0 || req.ProductID <= 0 {
		return 0, 0, errors.New("invalid user ID or product ID")
	}

	return req.UserID, req.ProductID, nil
}
func sendMessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var msg Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Добавляем сообщение в список
	messages = append(messages, msg)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(msg)
}

func getMessagesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		ID int `json:"id"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var userMessages []Message
	for _, msg := range messages {
		if msg.SenderID == req.ID || msg.ReceiverID == req.ID {
			userMessages = append(userMessages, msg)
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8") // Установка кодировки
	json.NewEncoder(w).Encode(userMessages)
}

func main() {
	http.HandleFunc("/users/register", registerUserHandler)                // Регистрация
	http.HandleFunc("/users/login", loginUserHandler)                      // Вход
	http.HandleFunc("/users/favorites/add", addToFavoritesHandler)         // Добавление в избранное
	http.HandleFunc("/users/favorites/remove", removeFromFavoritesHandler) // Удаление из избранного
	http.HandleFunc("/users/cart/add", addToCartHandler)                   // Добавление в корзину
	http.HandleFunc("/users/cart/remove", removeFromCartHandler)           // Удаление из корзины

	http.HandleFunc("/products", getProductsHandler)           // Получить все продукты
	http.HandleFunc("/products/create", createProductHandler)  // Создать продукт
	http.HandleFunc("/products/", getProductByIDHandler)       // Получить продукт по ID
	http.HandleFunc("/products/update/", updateProductHandler) // Обновить продукт
	http.HandleFunc("/products/delete/", deleteProductHandler) // Удалить продукт

	http.HandleFunc("/messages/send", sendMessageHandler) // Отправить сообщение
	http.HandleFunc("/messages", getMessagesHandler)      // Получить сообщения

	fmt.Println("Server is running on port 8080!")
	http.ListenAndServe(":8080", nil)
}
