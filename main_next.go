package main

// import (
// 	"github.com/jinzhu/gorm"
// 	_ "github.com/jinzhu/gorm/dialects/postgres"
// )

// type Product struct {
// 	gorm.Model
// 	Code  string
// 	Price uint
// }

// func main() {
// 	db, err := gorm.Open("postgres", "host=myhost port=myport user=gorm dbname=gorm password=mypassword")
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	defer db.Close()

// 	// Migrate the schema
// 	db.AutoMigrate(&Product{})

// 	// Create
// 	db.Create(&Product{Code: "L1212", Price: 1000})

// 	// Read
// 	var product Product
// 	db.First(&product, 1)                   // find product with id 1
// 	db.First(&product, "code = ?", "L1212") // find product with code l1212

// 	// Update - update product's price to 2000
// 	db.Model(&product).Update("Price", 2000)

// 	// Delete - delete product
// 	db.Delete(&product)
// }

// Book struct (Model)
// type Book struct {
// 	ID     string  `json:"id"`
// 	Isbn   string  `json:"isbn"`
// 	Title  string  `json:"title"`
// 	Author *Author `json:"author"`
// }

// // Author struct
// type Author struct {
// 	Firstname string `json:"firstname"`
// 	Lastname  string `json:"lastname"`
// }

// // Init book var as a slice book struct
// var books []Book

// // Get all books
// func getBooks(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("content-Type", "application/json")
// 	json.NewEncoder(w).Encode(books)
// }

// // Get single book
// func getBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("content-Type", "application/json")
// 	params := mux.Vars(r) // Get params

// 	for _, item := range books {
// 		if item.ID == params["id"] {
// 			json.NewEncoder(w).Encode(item)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode(&Book{})
// }

// // Create a new book
// func createBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("content-Type", "application/json")
// 	var book Book
// 	_ = json.NewDecoder(r.Body).Decode(&book)
// 	book.ID = strconv.Itoa(rand.Intn(10000000)) // mock Id - not safe
// 	books = append(books, book)
// 	json.NewEncoder(w).Encode(book)
// }

// // Update a book
// func updateBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("content-Type", "application/json")
// 	params := mux.Vars(r)
// 	for index, item := range books {
// 		if item.ID == params["id"] {
// 			books = append(books[:index], books[index+1:]...)
// 			var book Book
// 			_ = json.NewDecoder(r.Body).Decode(&book)
// 			book.ID = params["id"]
// 			books = append(books, book)
// 			json.NewEncoder(w).Encode(book)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode(books)
// }

// // Delete book
// func deleteBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("content-Type", "application/json")
// 	params := mux.Vars(r)
// 	for index, item := range books {
// 		if item.ID == params["id"] {
// 			books = append(books[:index], books[index+1:]...)
// 			break
// 		}
// 	}
// 	json.NewEncoder(w).Encode(books)
// }

// func main() {
// 	// Init Router
// 	r := mux.NewRouter()

// 	//Mock data
// 	books = append(books, Book{ID: "1", Isbn: "1232456", Title: "Book One", Author: &Author{
// 		Firstname: "John", Lastname: "doe"}})
// 	books = append(books, Book{ID: "2", Isbn: "6732746", Title: "Book Two", Author: &Author{
// 		Firstname: "Gabriel", Lastname: "Micah"}})
// 	books = append(books, Book{ID: "3", Isbn: "3524465", Title: "Book Three", Author: &Author{
// 		Firstname: "sara", Lastname: "matthews"}})

// 	// Route Handlers / endpoints
// 	r.HandleFunc("/api/books", getBooks).Methods("GET")
// 	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
// 	r.HandleFunc("/api/books", createBook).Methods("POST")
// 	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
// 	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

// 	// start the server
// 	log.Fatal(http.ListenAndServe(":1234", r))
// }

// func main() {
// 	c := make(chan string)

// 	go sayHello(c)

// 	fmt.Printf("I'm saying %s", <-c)
// 	time.Sleep(6 * time.Second)
// }

// func sayHello(c chan string) {
// 	fmt.Println("Waking Up...")
// 	time.Sleep(3 * time.Second)
// 	c <- "Hello KiwiPycon"
// }
