package util

//func LoadJson(fileName string) base.Operation {
//	file, err := os.Open(fileName)
//	if err != nil {
//		panic("failed to open json file: " + err.Error())
//	}
//	defer file.Close()
//
//	bytes, err := io.ReadAll(file)
//	if err != nil {
//		panic("failed to read json file: " + err.Error())
//	}
//
//	var v json.RawMessage
//	var op base.Operation
//	var ok bool
//	if err = json.Unmarshal(bytes, &v); err != nil {
//		panic("failed to unmarshal json: " + err.Error())
//	} else if hinter, err := enc.Decode(bytes); err != nil {
//		panic("failed to decode json: " + err.Error())
//	} else if op, ok = hinter.(base.Operation); !ok {
//		panic("decoded object is not Operation")
//	}
//
//	return op
//}
