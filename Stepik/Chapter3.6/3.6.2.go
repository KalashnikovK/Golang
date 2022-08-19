/*
Данная задача ориентирована на реальную работу с данными в формате JSON.
Для решения мы будем использовать справочник ОКВЭД (Общероссийский классификатор видов экономической деятельности),
который был размещен на web-портале data.gov.ru.

Необходимая вам информация о структуре данных содержится в файле structure-20190514T0000.json,
а сами данные, которые вам потребуется декодировать, содержатся в файле data-20190514T0100.json.
Файлы размещены в нашем репозитории на github.com.

Для того, чтобы показать, что вы действительно смогли декодировать документ вам необходимо в качестве ответа
записать сумму полей global_id всех элементов, закодированных в файле.
*/

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Gid struct {
	Id int `json:"global_id"`
}

func main() {
	var gid []Gid

	file, err := os.Open("data-20190514T0100.json")
	if err != nil {
		fmt.Printf("File open: %v\n", err)
	}
	defer file.Close()

	if err = json.NewDecoder(file).Decode(&gid); err != nil {
		fmt.Printf("Decode: %v\n", err)
	}

	var res = 0
	for _, value := range gid {
		res += value.Id
	}
	fmt.Println(res)
}
