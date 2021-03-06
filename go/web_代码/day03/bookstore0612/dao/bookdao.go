package dao

import (
	"bookstore0612/model"
	"bookstore0612/utils"
)

//GetBooks 获取数据库中所有的图书
func GetBooks() ([]*model.Book, error) {
	//写sql语句
	sqlStr := "select id,title,author,price,sales,stock,img_path from books"
	//执行
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		//给book中的字段赋值
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		//将book添加到books中
		books = append(books, book)
	}
	return books, nil
}

//AddBook 向数据库中添加一本图书
func AddBook(b *model.Book) error {
	//写sql语句
	slqStr := "insert into books(title,author,price,sales,stock,img_path) values(?,?,?,?,?,?)"
	//执行
	_, err := utils.Db.Exec(slqStr, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ImgPath)
	if err != nil {
		return err
	}
	return nil
}

//DeleteBook 根据图书的id从数据库中删除一本图书
func DeleteBook(bookID string) error {
	//写sql语句
	sqlStr := "delete from books where id = ?"
	//执行
	_, err := utils.Db.Exec(sqlStr, bookID)
	if err != nil {
		return err
	}
	return nil
}

//GetBookByID 根据图书的id从数据库中查询出一本图书
func GetBookByID(bookID string) (*model.Book, error) {
	//写sql语句
	sqlStr := "select id,title,author,price,sales,stock,img_path from books where id = ?"
	//执行
	row := utils.Db.QueryRow(sqlStr, bookID)
	//创建Book
	book := &model.Book{}
	//为book中的字段赋值
	row.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
	return book, nil
}

//UpdateBook 根据图书的id更新图书信息
func UpdateBook(b *model.Book) error {
	//写sql语句
	sqlStr := "update books set title=?,author=?,price=?,sales=?,stock=? where id=?"
	//执行
	_, err := utils.Db.Exec(sqlStr, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ID)
	if err != nil {
		return err
	}
	return nil
}
