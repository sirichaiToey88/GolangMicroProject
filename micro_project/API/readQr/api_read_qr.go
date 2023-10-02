package readQr

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CompareQR(e echo.Context) error {
	QR_test := "00020101021129370016A000000677010111021313099008937905802TH5303764630401F2"
	DEFAULT_LENGTH := 2
	// DEFAULT_START_TAG := "00"

	i := 0

	for i < len(QR_test) {
		tagInd := i + DEFAULT_LENGTH
		tag := parsInt(i, QR_test)
		lengthValue_sub := QR_test[tagInd : tagInd+DEFAULT_LENGTH]
		lengthValue_pars, _ := strconv.Atoi(lengthValue_sub)

		fmt.Println("QR => ", lengthValue_pars)
		fmt.Println("tag => ", tag)
		return e.JSON(http.StatusOK, string(lengthValue_pars))
	}
	return e.JSON(http.StatusOK, "OK")
}

func parsInt(tagInd int, qrCode string) int {
	prefix := 0
	substring := qrCode[prefix:tagInd]
	parsedInt, _ := strconv.Atoi(substring)
	return parsedInt
}

// You can edit this code!
// Click here and start typing.
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	qrCodeValue := "00020101021129370016A000000677010111021313099008937905802TH5303764630401F2"
// 	qrCodeEwall := "00020101021129390016A000000677010111031500499900000895253037645802TH63045388"
// 	qrMobileT29 := "00020101021129370016A000000677010111011300669545359565303764540510.735802TH63041BA7"

// 	tag2 := qrCodeValue[40:53]

// 	tag29 := qrCodeValue[20:36]

// 	tag37 := qrCodeValue[16:53]

// 	tag63 := qrCodeValue[70:74]
// 	tag53 := qrCodeValue[63:66]
// 	tag58 := qrCodeValue[57:59]

// 	fmt.Println(tag2)
// 	fmt.Println(tag29)
// 	fmt.Println(tag37)
// 	fmt.Println(tag63)
// 	fmt.Println(tag53)
// 	fmt.Println(tag58)
// }
