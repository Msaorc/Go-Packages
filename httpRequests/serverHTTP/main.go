package main

import "net/http"

func main() {
	http.HandleFunc("/add", addExpense)
	http.HandleFunc("/listExpense", listExpense)
	http.ListenAndServe(":8080", nil)
}

func addExpense(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Expense add with success"))
}

func listExpense(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("List the Expenses: 1 - Gasoline"))
}
