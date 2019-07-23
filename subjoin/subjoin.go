package subjoin

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type Excel struct {
	Start       string
	End         string
	Content     string
	fileDir     string
	Auth        string
	FileName    string
	NextContent string
}

func (E *Excel) GetFileDir() string {
	return E.fileDir
}

func (E *Excel) SetFileName(filename string) {
	E.FileName = filename
}

func (E *Excel) GetFileName() string {
	return E.FileName
}

func (E *Excel) SetFileDir(name string) {
	E.fileDir = name
}

func GetAppPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	dirstring := strings.Replace(dir, "\\", "/", -1)
	if err != nil {
		log.Fatal(err)
	}
	return dirstring
}

func (E *Excel) NewExcel() error {
	start := E.Start
	end := E.End
	content := E.Content

	s := strings.Split(start, "-")
	e := strings.Split(end, "-")
	if len(s) != 3 || len(e) != 3 {
		return errors.New("检查你开始时间和结束时间的格式!")
	}
	titTime := fmt.Sprint(s[0] + "年" + s[1] + "月" + s[2] + "日" + " - " + e[1] + "月" + e[2] + "日")
	title := "软件部" + E.Auth + "工作周报" + strings.ReplaceAll(end, "-", "")
	filename := title + ".xlsx"
	E.SetFileName(title)
	fileDir := GetAppPath() + "/" + filename
	E.SetFileDir(fileDir)

	f := excelize.NewFile()
	// 创建一个工作表
	index := f.NewSheet("Sheet1")
	// 合并单元格
	f.MergeCell("Sheet1", "A1", "I1")
	f.MergeCell("Sheet1", "A2", "E2")
	f.MergeCell("Sheet1", "G2", "I2")
	f.SetCellValue("Sheet1", "B1", "                                                                                   农担项目组周报表                                "+titTime)

	// 设置列宽度
	_ = f.SetColWidth("Sheet1", "B", "B", 50)
	_ = f.SetColWidth("Sheet1", "C", "C", 30)
	_ = f.SetColWidth("Sheet1", "D", "D", 40)
	_ = f.SetColWidth("Sheet1", "H", "H", 20)

	// 设置全局边框样式
	borderStyle, err := f.NewStyle(`{"border":[
		       {"type":"left","color":"#000000","style":1},
		       {"type":"right","color":"#000000","style":1},
		       {"type":"top","color":"#000000","style":1},
		       {"type":"bottom","color":"#000000","style":1}
		       ],
		      "font":{
			 "size":10
			  },
		      "alignment":{
			  "horizontal":"center",
			  "vertical":"center"
			}
		       }`)

	if err != nil {
		return err
	}
	if err := f.SetCellStyle("Sheet1", "A2", "E4", borderStyle); err != nil {
		return err
	}
	if err := f.SetCellStyle("Sheet1", "G2", "I4", borderStyle); err != nil {
		return err
	}
	if err := f.SetCellStyle("Sheet1", "A1", "I1", borderStyle); err != nil {
		return err
	}

	// 表头样式
	tableHeadStyle, headErr := f.NewStyle(`{"font":
		       {"bold":true,
			"size":11
		       },
		      "alignment":{
			  "horizontal":"center",
			  "vertical":"center"
			},
		       "border":[
			       {"type":"left","color":"#000000","style":1},
			       {"type":"right","color":"#000000","style":1},
			       {"type":"top","color":"#000000","style":1},
			       {"type":"bottom","color":"#000000","style":1}
		       ]
}`)
	if headErr != nil {
		return headErr
	}

	if err := f.SetCellStyle("Sheet1", "A3", "I3", tableHeadStyle); err != nil {
		return err
	}

	// 标题样式
	titleStyle, titleErr := f.NewStyle(`{"border":[
		       {"type":"left","color":"#000000","style":1},
		       {"type":"right","color":"#000000","style":1},
		       {"type":"top","color":"#000000","style":1},
		       {"type":"bottom","color":"#000000","style":1}
		       ],
		      "font":{
			 "size":20,
			 "bold":true
			  },
		      "alignment":{
			  "horizontal":"center",
			  "vertical":"center"
			}
		       }`)

	if titleErr != nil {
		return titleErr
	}
	f.SetCellStyle("Sheet1", "A1", "A1", titleStyle)

	f.SetCellValue("Sheet1", "G2", "下周计划")

	f.SetCellValue("Sheet1", "A3", "序号")
	f.SetCellValue("Sheet1", "B3", "工作（项目）内容")
	f.SetCellValue("Sheet1", "C3", "完成情况")
	f.SetCellValue("Sheet1", "C4", "负责人")
	f.SetCellValue("Sheet1", "D3", "是否异常/需要协调")
	f.SetCellValue("Sheet1", "G3", "序号")
	f.SetCellValue("Sheet1", "H3", "工作（项目）内容")
	f.SetCellValue("Sheet1", "I3", "负责人")

	f.SetCellValue("Sheet1", "A4", "1")
	f.SetCellValue("Sheet1", "B4", content)
	f.SetCellValue("Sheet1", "C4", "100%")
	f.SetCellValue("Sheet1", "D4", "暂无异常")
	f.SetCellValue("Sheet1", "E4", E.Auth)
	f.SetCellValue("Sheet1", "G4", "1")
	f.SetCellValue("Sheet1", "H4", E.NextContent)
	f.SetCellValue("Sheet1", "I4", E.Auth)

	_ = f.SetRowHeight("Sheet1", 1, 50)
	_ = f.SetRowHeight("Sheet1", 2, 30)
	_ = f.SetRowHeight("Sheet1", 3, 30)
	_ = f.SetRowHeight("Sheet1", 4, 70)

	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)
	// 根据指定路径保存文件
	error := f.SaveAs(fileDir)
	if error != nil {
		fmt.Println(error)
	}
	return nil
}
