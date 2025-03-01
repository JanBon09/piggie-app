package content

var (
	currencies      = make(map[int]string)
	transactionType = make(map[int]string)
)

func InitWallet() {
	currencies[0] = "EUR"
	currencies[1] = "USD"
	currencies[2] = "PLN"

	transactionType[0] = "In"
	transactionType[1] = "Out"
}

// Struct to represent amount of money
// In int64 to avoid decimal point error
type Cash struct {
	Euros     int64 `json:"euros"`
	Eurocents int64 `json:"eurocents"`
}

// Struct to represent user wallet
type Wallet struct {
	Balance  Cash
	Currency int8
}

// Struct representing singular transaction
// that was made by user or to user
type Transaction struct {
	TransactionType int8 `json:"transactionType"`
	Amount          Cash
	TransactionDate string `json:"transactionDate"`
}

// Struct representing history of transactions
// containing all in/out transaction that was made
// in user account
type TransactionsHistory struct {
	NumberOfTransactions int64 `json:"numberOfTransaction"`
	Earnings             Cash
	Spendings            Cash
}

type ICash interface {
	CalculateCents(eurocecnts int64) int64
}

// Calculating full euros
// I think it's easier on server to just do it always after transaction
// than checking and the maybe doing it
func (cash *Cash) CalculateCents() {
	cash.Euros += cash.Eurocents / 100
}
