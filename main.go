/**
读取excel
*/
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Luxurioust/excelize"
	"github.com/sergeilem/xls"
)

var file = flag.String("file", "", "")

func main() {
	flag.Parse()
	var arr = strings.Split(*file, ".")
	if arr[1] == "xlsx" {

		xlsx, err := excelize.OpenFile(*file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		var sheetName = xlsx.GetSheetName(xlsx.GetActiveSheetIndex())
		rows, _ := xlsx.GetRows(sheetName)
		list := make(map[int]map[int]string)
		i := 0
		for _, row := range rows {
			r := make(map[int]string)
			var j = 0
			for _, colCell := range row {
				r[j] = colCell
				j++
			}

			list[i] = r
			i++
			//os.Exit(1)
		}
		result, err := json.MarshalIndent(list, "", "  ")
		if err != nil {
			fmt.Println("err = ", err)
			return
		}
		fmt.Println("result = ", string(result))
	} else if arr[1] == "xls" {

		if xlFile, err := xls.Open(*file, "utf-8"); err == nil {
			if sheet1 := xlFile.GetSheet(0); sheet1 != nil {

				list := make(map[int]map[int]string)
				for i := 0; i <= (int(sheet1.MaxRow)); i++ {
					row1 := sheet1.Row(i)

					if row1 == nil {
						continue
					}
					r := make(map[int]string)
					lastcol := row1.LastCol()

					for index := 0; index < lastcol; index++ {
						r[index] = row1.Col(index)
					}
					list[i] = r
				}

				result, err := json.MarshalIndent(list, "", "  ")
				if err != nil {
					fmt.Println("err = ", err)
					return
				}
				fmt.Println("result = ", string(result))
			}
		} else {
			fmt.Print(err)
		}

	} else {
		fmt.Println("文件格式错误")
	}
}
