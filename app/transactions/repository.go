package transactions

type TransactionRepository interface {
	FindAll() ([]Transaction, error)
	FindByID(id int) (Transaction, error)
	FindByUserID(userID int) ([]Transaction, error)
	Create(transaction Transaction) error
	Update(transaction Transaction) error
	Delete(transaction Transaction) error
}
