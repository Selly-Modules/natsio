package subject

import "fmt"

func getSupplierValue(val string) string {
	return fmt.Sprintf("%s.%s", prefixes.Supplier, val)
}

var Supplier = struct {
	FindAll string
}{
	FindAll: getSupplierValue("find_all"),
}
