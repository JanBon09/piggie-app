package data

import (
	"database/sql"
	"piggieBackend/content"
	"piggieBackend/utility"
)

// Function to retrive user wallet information for further manipulation or visualization
func GetMainPanelWalletData(username string) (content.Wallet, error) {
	stringStatement := "SELECT balance, currency FROM wallet WHERE username LIKE $1"
	statement, err := DB.Prepare(stringStatement)
	if err != nil {
		return content.Wallet{}, err
	}

	var wallet content.Wallet
	err = statement.QueryRow(username).Scan(&wallet)
	if err != nil {
		if err == sql.ErrNoRows {
			return content.Wallet{}, utility.ErrNoRows
		} else {
			return content.Wallet{}, err
		}
	}

	return content.Wallet{}, nil
}
