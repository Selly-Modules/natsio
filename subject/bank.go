package subject

import "fmt"

func getBankValue(val string) string {
	return fmt.Sprintf("%s.%s", prefixes.Bank, val)
}

var Bank = struct {
	GetBankById       string
	GetBankBranchById string
}{
	GetBankBranchById: getBankValue("get_bank_by_id"),
}
