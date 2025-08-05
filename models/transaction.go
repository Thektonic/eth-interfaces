package models

// TransactionReturn contains the result of a transaction execution
type TransactionReturn struct {
	Hash string
	Err  error
}
