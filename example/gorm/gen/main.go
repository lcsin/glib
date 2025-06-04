package main

import (
	"context"
	"fmt"

	"gen/dal/query"
	"gen/ioc"
)

func main() {
	ctx := context.Background()
	query.SetDefault(ioc.InitDB().Debug())

	book, err := query.Book.WithContext(ctx).First()
	if err != nil {
		panic(err)
	}
	fmt.Println(book)

	//books, count, err := query.Book.WithContext(ctx).FindByPage(0, 5)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("count: ", count)
	//for _, book := range books {
	//	fmt.Println(book)
	//}

	// 创建
	//for i := 0; i < 10; i++ {
	//	b1 := model.Book{
	//		Title:       fmt.Sprintf("《七米的Go语言之路%v》", i+1),
	//		Author:      "七米",
	//		PublishDate: time.Date(2023, 11, 15, 0, 0, 0, 0, time.UTC),
	//		Price:       100,
	//	}
	//	err := query.Book.WithContext(context.Background()).Create(&b1)
	//	if err != nil {
	//		fmt.Printf("create book fail, err:%v\n", err)
	//		return
	//	}
	//}

	// 更新
	//ret, err := query.Book.WithContext(context.Background()).
	//	Where(query.Book.ID.Eq(1)).
	//	Update(query.Book.Price, 500)
	//if err != nil {
	//	fmt.Printf("update book fail, err:%v\n", err)
	//	return
	//}
	//fmt.Printf("RowsAffected:%v\n", ret.RowsAffected)

	// 查询
	//book, err := query.Book.WithContext(context.Background()).First()
	//// 也可以使用全局Q对象查询
	////book, err := query.Q.Book.WithContext(context.Background()).First()
	//if err != nil {
	//	fmt.Printf("query book fail, err:%v\n", err)
	//	return
	//}
	//fmt.Printf("book:%v\n", book)
	//
	// 删除
	//ret, err := query.Book.WithContext(context.Background()).Where(query.Book.ID.Eq(1)).Delete()
	//if err != nil {
	//	fmt.Printf("delete book fail, err:%v\n", err)
	//	return
	//}
	//fmt.Printf("RowsAffected:%v\n", ret.RowsAffected)
}
