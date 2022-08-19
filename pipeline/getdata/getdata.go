package getdata

import (
	"encoding/csv"
	"fmt"
	"go++/pipeline/AppInit"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

type Book struct {
	BookId int `gorm:"column:book_id"`
	BookName string `gorm:"column:book_name"`
}

type BookList struct {
	Data  []*Book
	Page int
}

type Result struct {
	Page int
	Err error
} 
type InChan chan *BookList
type OutChan chan *Result

type DataCmd func() InChan
type DataPipeCmd func(in InChan) OutChan

const sql  = "select * from books order by book_id limit ? offset ?"

// 使用channel改造
func ReadData() InChan {
	page := 1
	pageSize := 1000
	in := make(InChan)
	go func() {
		defer  close(in)
		for {
			bookList := &BookList{make([]*Book,0),page}
			db := AppInit.GetDB().Raw(sql,pageSize,(page-1)*pageSize).Find(&bookList.Data)
			if db.Error!= nil || db.RowsAffected==0 {  // 有错误，或者查询到的数据为0行
				break
			}
			in <- bookList
			page ++
		}
	}()
	return  in
}

func WriteData(in InChan) OutChan {
	out := make(OutChan)
	go func() {
		defer close(out)
		for d := range in {
			out <- &Result{
				Page: d.Page,
				Err:  SaveData(d),
			}
		}
	}()
	return out
}

func Pipe(c1 func() InChan, cs ...DataPipeCmd) OutChan  {
	in := c1()
	out := make(OutChan)
	wg := sync.WaitGroup{}
	for _,c:= range cs {
		getChan := c(in)
		wg.Add(1)
		go func(input OutChan) {
			defer wg.Done()
			for v := range input {
				out <- v
			}
		}(getChan)
	}
	go func() {
		defer close(out)
		wg.Wait()
	}()
	return  out
}

// 传统方法
func ReadData1()  {
	page := 1
	pageSize := 1000
	for {
		bookList := &BookList{make([]*Book,0),page}
		db := AppInit.GetDB().Raw(sql,pageSize,(page-1)*pageSize).Find(&bookList.Data)
		if db.Error!= nil || db.RowsAffected==0 {  // 有错误，或者查询到的数据为0行
			break
		}
		err := SaveData(bookList)
		if err != nil {
			log.Println(err)
		}
		page ++
	}
}


// 写入到csv
func SaveData(data *BookList) error   {
	time.Sleep(time.Millisecond*500)
	file:=fmt.Sprintf("./pipeline/csv/%d.csv",data.Page)
	csvFile,err:= os.OpenFile(file,os.O_RDWR|os.O_CREATE|os.O_TRUNC,0666)
	if err!=nil{
		return err
	}
	defer csvFile.Close()
	w := csv.NewWriter(csvFile)//创建一个新的写入文件流
	header := []string{"book_id", "book_name"}
	export := [][]string{   // 值得学习，二维切片
		header,
	}
	for _,d:=range data.Data{
		cnt:=[]string{
			strconv.Itoa(d.BookId),
			d.BookName,
		}
		export=append(export,cnt)
	}
	err=w.WriteAll(export)
	if err!=nil{
		return err
	}
	w.Flush()  // 刷新，不刷新是无法写入的
	return nil
}

func Test()  {
	out := Pipe(ReadData,WriteData,WriteData)
	for o := range out {
		fmt.Printf("%d文件执行完成，结果:%v\n", o.Page,o.Err)
	}

}


