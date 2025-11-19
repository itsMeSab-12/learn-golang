package ledger

var graph = make(map[string][]*Transaction)

func AddEdge(senderID string, tx *Transaction) {
	list := graph[senderID]
	list = append(list, tx)
	graph[senderID] = list
}

func GetOutTransactions(userID string) []*Transaction {
	list := graph[userID]
	return list
}

func GetInTransactions(userID string) []*Transaction {
	//traverse the map
	//add all transactions where Transaction.ID = userID
	incTx := make([]*Transaction, 0)
	for _, list := range graph {
		for _, tx := range list {
			if tx.ID == userID {
				//note tx is a pointer
				incTx = append(incTx, tx)
			}
		}
	}
	return incTx
}

func GetNeighbours(userID string) []*User {
	outTx, ok := graph[userID]
	neighbours := make(map[string]*User, 0)
	if ok {
		for _, tx := range outTx {
			_, exists := neighbours[tx.From.ID]
			if !exists {
				neighbours[tx.From.ID] = tx.From
			}
		}
	}

	outcome := make([]*User, 0, len(neighbours))
	for _, neighbour := range neighbours {
		outcome = append(outcome, neighbour)
	}

	return outcome

}

func DetectCycle() {

}
