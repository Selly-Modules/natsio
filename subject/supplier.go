package subject

import "fmt"

func getSupplierValue(val string) string {
	return fmt.Sprintf("%s.%s", prefixes.Supplier, val)
}

var Supplier = struct {
	GetSupplierInfo string
}{
	GetSupplierInfo: getSupplierValue("get_supplier_info"),
}
