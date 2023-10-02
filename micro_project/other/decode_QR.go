package other

import (
	"encoding/json"
	"strconv"
)

const defaultLength = 2
const defaultStartTag = "00"

type EmvcoTlvBean struct {
	Length int                     `json:"length"`
	Value  string                  `json:"value"`
	SubTag map[string]EmvcoTlvBean `json:"subTag,omitempty"`
}

func isInteger(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

func parseSubTagValue(value string) map[string]EmvcoTlvBean {
	result := make(map[string]EmvcoTlvBean)
	if len(value) > 4 {
		tag := value[:defaultLength]
		lengthStr := value[defaultLength : defaultLength*2]
		if tag == defaultStartTag && isInteger(tag) && isInteger(lengthStr) {
			length, _ := strconv.Atoi(lengthStr)
			if length > 0 {
				parse := parseQRCodeValue(value, true)
				result = parse
			}
		}
	}
	return result
}

func parseQRCodeValue(qrCodeValue string, isFromSubtag bool) map[string]EmvcoTlvBean {
	result := make(map[string]EmvcoTlvBean)
	i := 0
	for i < len(qrCodeValue) {
		tagInd := i + defaultLength
		tag, _ := strconv.Atoi(qrCodeValue[i:tagInd])
		lengthInd := tagInd + defaultLength
		lengthValue, _ := strconv.Atoi(qrCodeValue[tagInd:lengthInd])
		emvcoTlvBean := EmvcoTlvBean{
			Length: lengthValue,
			Value:  qrCodeValue[lengthInd : lengthInd+lengthValue],
		}

		if !isFromSubtag {
			subTagJSON := parseSubTagValue(emvcoTlvBean.Value)
			if len(subTagJSON) > 0 {
				emvcoTlvBean.SubTag = subTagJSON
			}
		}

		result["tag"+strconv.Itoa(tag)] = emvcoTlvBean

		i = lengthInd + lengthValue
	}
	return result
}

func ParseQRCodeValue(qrCodeValue string) map[string]EmvcoTlvBean {
	return parseQRCodeValue(qrCodeValue, false)
}

func ToJSON(qrCodeValue string) (string, error) {
	parsedResult := ParseQRCodeValue(qrCodeValue)
	jsonResult, err := json.Marshal(parsedResult)
	if err != nil {
		return "", err
	}
	return string(jsonResult), nil
}
