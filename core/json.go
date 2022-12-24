package core

import (
	"czfinger/structs"
	"encoding/json"
	"io/ioutil"
)

//var FofaFingerPrints []structs.FofaFingerPrint

//func init() {
//	err := json.Unmarshal(asset.FofaFingerPrintString, &FofaFingerPrints)
//	if err != nil {
//		utils.OptionsError("Json parse error", 1)
//	}
//}

//返回切片数据
func Parse(filename string) ([]structs.FofaFinger, error) {

	Json, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var dataArray []structs.FofaFinger
	err = json.Unmarshal(Json, &dataArray)
	if err != nil {
		return nil, err
	}

	return dataArray, nil
}
