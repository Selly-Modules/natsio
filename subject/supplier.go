package subject

import "fmt"

func getSupplierValue(val string) string {
	return fmt.Sprintf("%s.%s", prefixes.Supplier, val)
}

var Supplier = struct {
	GetListSupplierInfo             string
	GetSupplierContractBySupplierID string
	FindAll                         string
}{
	GetListSupplierInfo:             getSupplierValue("get_list_supplier_info"),
	GetSupplierContractBySupplierID: getSupplierValue("get_supplier_contract_by_supplier_id"),
	FindAll:                         getSupplierValue("find_all"),
}
