package db

import (
	"encoding/json"
	"reflect"
)

type DbClient struct {
	rawData []string
	cur     int
}

func (c *DbClient) HasMore() bool {
	return c.cur < len(c.rawData)
}

func (c *DbClient) Close() {

}

func (c *DbClient) ReadDocument(result interface{}) error {
	if err := json.Unmarshal([]byte(c.rawData[c.cur]), result); err != nil {
		return err
	}
	c.cur++
	return nil
}

var client DbClient

func init() {
	client = DbClient{}
	client.rawData = []string{}

	client.rawData = append(client.rawData, `{ "id": "12537336", "name": "Raghav Last", "email": "raaanand@gmail.com", "phone": "9411395698" }`)
	client.rawData = append(client.rawData, `{ "id": "18649039", "name": "Deepak kumar", "email": "deepr8490@gmail.com", "phone": "9864567207" }`)
	client.rawData = append(client.rawData, `{ "id": "65482043", "name": "Sumit Kumar", "email": "sumita14@gmail.com", "phone": "8050571234" }`)
	client.rawData = append(client.rawData, `{ "id": "65550179", "name": "Nithish Kumar", "email": "ihsh56@gmail.com", "phone": "8197678374" }`)
	client.rawData = append(client.rawData, `{ "id": "75708025", "name": "RIKA ARJLCHON", "email": "karl.arnan@gmail.com", "phone": "9766667689" }`)
}

// 1)
func QueryResults_CallBack_ReflectionFunc(callBackFuncI interface{}) error {
	defer client.Close()

	funcType := reflect.TypeOf(callBackFuncI)
	docType := funcType.In(0).Elem() // skip check for parameter count
	callBackFunc := reflect.ValueOf(callBackFuncI)
	arg := make([]reflect.Value, 1)

	doc := reflect.New(docType)
	for client.HasMore() {
		if err := client.ReadDocument(doc.Interface()); err != nil {
			return err
		}
		arg[0] = doc
		callBackFunc.Call(arg)
	}

	return nil
}

// 2)
func QueryResults_CallBack_ObjNFunc(doc interface{}, callBackFunc func(interface{})) error {
	defer client.Close()

	for client.HasMore() {
		if err := client.ReadDocument(doc); err != nil {
			return err
		}
		callBackFunc(doc)
	}
	return nil
}

// 3)
func QueryResults_CallBack_DBClient(callBackFunc func(DbClient) error) error {
	defer client.Close()

	for client.HasMore() {
		if err := callBackFunc(client); err != nil {
			return err
		}
	}
	return nil
}

// 4)
func QueryResults_ExposeClient() DbClient {
	return client
}
