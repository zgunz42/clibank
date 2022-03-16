package transactions

type TransactionService interface {
	Create(transaction Transaction) (Transaction, error)
	Update(transaction Transaction) (Transaction, error)
	Delete(transaction Transaction) error
	FindByID(id int) (Transaction, error)
	FindByUserID(userID int) ([]Transaction, error)
}
