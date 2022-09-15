package subject

import "fmt"

func getSellerValue(val string) string {
	return fmt.Sprintf("%s.%s", prefixes.Seller, val)
}

// Seller ...
var Seller = struct {
	GetSellerInfoByID string
}{
	GetSellerInfoByID: getSellerValue("get_seller_info_by_id"),
}
