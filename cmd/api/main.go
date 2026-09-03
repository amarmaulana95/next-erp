package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/amarmaulana95/next-erp/internal/database"
	"github.com/amarmaulana95/next-erp/internal/employee"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"ok"}`))
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := database.NewPostgresPool(ctx)
	if err != nil {
		log.Fatalf("database connection failed: %v", err)
	}
	defer pool.Close()

	log.Println("PostgreSQL connection established")

	employeeRepository := employee.NewRepository(pool)
	employeeHandler := employee.NewHandler(employeeRepository)

	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/api/employees", employeeHandler.GetAll)
	http.HandleFunc("/api/employees/{id}", employeeHandler.GetByID)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	log.Printf("Next-ERP API running on :%s", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
