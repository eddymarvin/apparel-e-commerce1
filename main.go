package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Product struct to hold product information
type Product struct {
	ID    int
	Name  string
	Price float64
	Image string
}

// Sample product data
var products = []Product{
	{ID: 1, Name: "T-Shirt", Price: 10.99, Image: "/static/images/product1.jpg"},
	{ID: 2, Name: "Jeans", Price: 15.99, Image: "/static/images/product2.jpg"},
	// Add more products as needed
}

// Cart to hold selected products
var cart []Product

// Render the homepage
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Render the product page
func productHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/products.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Add product to cart
func addToCartHandler(w http.ResponseWriter, r *http.Request) {
	productIDStr := r.URL.Query().Get("id")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	for _, product := range products {
		if product.ID == productID {
			cart = append(cart, product)
			break
		}
	}
	http.Redirect(w, r, "/cart", http.StatusFound)
}

// Render the cart page
func cartHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/cart.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, cart)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Render the delivery location page
func deliveryHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/delivery.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Render the login handler
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Handle login logic here
		http.Redirect(w, r, "/products", http.StatusFound)
		return
	}
	tmpl, err := template.ParseFiles("templates/login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	// Serve static assets (HTML, CSS, JS, images) from the "static" folder
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Handle routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/products", productHandler)
	http.HandleFunc("/add-to-cart", addToCartHandler)
	http.HandleFunc("/cart", cartHandler)
	http.HandleFunc("/delivery", deliveryHandler)
	http.HandleFunc("/login", loginHandler) // Add the login handler

	// Start the server on port 8080
	log.Println("Server started at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
