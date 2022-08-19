/*
Поиск файла в заданном формате и его обработка
Данная задача поможет вам разобраться в пакете encoding/csv и path/filepath,
хотя для решения может быть использован также пакет archive/zip (поскольку файл с заданием предоставляется именно в этом формате).

В тестовом архиве, который вы можете скачать из нашего репозитория на github.com,
содержится набор папок и файлов. Один из этих файлов является файлом с данными в формате CSV, прочие же файлы структурированных данных не содержат.

Требуется найти и прочитать этот единственный файл со структурированными данными (это таблица 10х10, разделителем является запятая),
а в качестве ответа необходимо указать число, находящееся на 5 строке и 3 позиции (индексы 4 и 2 соответственно).
*/

package main

import (
	"archive/zip"
	"encoding/csv"
	"fmt"
)

func main() {
	zipFile, _ := zip.OpenReader("task.zip")
	defer zipFile.Close()

	for _, file := range zipFile.File {
		f, _ := file.Open()

		data, _ := csv.NewReader(f).ReadAll()
		if len(data) > 1 {
			fmt.Printf("Искомое число: %s. Файл в котором лежат данные: %s", data[4][2], file.Name)
		}
		f.Close()
	}
}

/*func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err // Если по какой-то причине мы получили ошибку, проигнорируем эту итерацию
	}

	if info.IsDir() {
		return nil // Проигнорируем директории
	}

	f, _ := os.Open(path)
	r := csv.NewReader(f)
	record, _ := r.ReadAll()
	if len(record) > 1 {
		fmt.Println(record[4][2])
		fmt.Printf("Name: %s\tSize: %d byte\tPath: %s\n", info.Name(), info.Size(), path)
	}
	f.Close()
	return nil
}

func main() {
	const root = "./task"

	if err := filepath.Walk(root, walkFunc); err != nil {
		fmt.Printf("Какая-то ошибка: %v\n", err)
	}
}*/
