package subject

import "fmt"

func getNewsValue(val string) string {
	return fmt.Sprintf("%s.%s", prefixes.News, val)
}

var News = struct {
	GetProductNoticesByInventory string
}{
	GetProductNoticesByInventory: getNewsValue("get_product_notices_by_inventory"),
}
