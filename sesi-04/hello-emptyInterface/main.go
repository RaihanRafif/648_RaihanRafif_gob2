package main

func main() {
	// var randomValues interface{}
	// _ = randomValues

	// randomValues = "jasdkijkasd"
	// randomValues = 23
	// randomValues = true
	// randomValues = []string{"airell", "nanda"}

	// ----------------------------------------------
	// var v interface{}
	// v = 20
	// if value, ok := v.(int); ok == true {
	// 	v = value * 9
	// }

	// ----------------------------------------------
	rs := []interface{}{1, "Airell", true, 2, "Ananda", true}

	rm := map[string]interface{}{
		"Name":   "Airell",
		"Status": true,
		"Age":    23,
	}
	_, _ = rs, rm
}
