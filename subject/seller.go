package subject

import "fmt"

func getSellerValue(val string) string {
	return fmt.Sprintf("%s.%s", prefixes.Seller, val)
}

// Seller ...
var Seller = struct {
	GetSellerInfoByID      string
	GetListSellerInfoByIDs string
}{
	GetSellerInfoByID:      getSellerValue("get_seller_info_by_id"),
	GetListSellerInfoByIDs: getSellerValue("get_list_seller_info_by_ids"),
}

// SupportChat ...
var SupportChat = struct {
	GetListSellerInfoSupportChatByIDs string
}{
	GetListSellerInfoSupportChatByIDs: "SELLY_CHAT.REQUEST.SELLER_INFO",
}
