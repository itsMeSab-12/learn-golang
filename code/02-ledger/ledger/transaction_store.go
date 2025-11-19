package ledger

var transactionDB = make(map[string]*Transaction)

func AddTransaction(t *Transaction) bool {
	_, exists := transactionDB[t.ID]
	if !exists {
		transactionDB[t.ID] = t
		return true
	}
	return false
}

func GetTransactionByID(id string) *Transaction {
	t, ok := transactionDB[id]
	if ok {
		return t
	}

	return nil

}

func GetTransactionByUsers(userID string) []*Transaction {
	result := make([]*Transaction, 0)
	for _, t := range transactionDB {
		if t.From.ID == userID || t.To.ID == userID {
			result = append(result, t)
		}
	}
	return result

}

func ListTransactions() []*Transaction {
	list := make([]*Transaction, 0, len(transactionDB))
	for _, t := range transactionDB {
		list = append(list, t)
	}

	return list
}

func DeleteTransaction(id string) {
	delete(transactionDB, id)
}
