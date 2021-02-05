package readjson

import (
	"encoding/json"
	types "huTao/types"
	"io/ioutil"
)

//ReadJSON JSONファイルの読み込み
func ReadJSON(path string) (chara types.Data, err error) {
	var content []byte
	content, err = ioutil.ReadFile(path)
	if err != nil {
		return
	}

	json.Unmarshal(content, &chara)
	return
}

//GenerateJSON JSONを生成する
func GenerateJSON(s interface{}) ([]byte, error) {
	return json.Marshal(s)
}
