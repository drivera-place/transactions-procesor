package balance

type Summary struct{
	Total float64
	MonthlyTransactions []MonthlyTransactions
	AverageDebit float64
	AverageCredit float64
}

type MonthlyTransactions struct {
	Month string
	Number int
}
