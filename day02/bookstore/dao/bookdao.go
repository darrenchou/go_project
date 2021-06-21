package dao

import (
	"day02/bookstore/model"
	"day02/bookstore/utils"
	"fmt"
	"strconv"
)

//GetBooks
func GetBooks() ([]*model.Book, error) {

	sqlStr := "select id,title,author,publisher, publish_date, rating, book_status, price,img_path from books"

	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		//assign value to book
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Publisher, &book.Publish_date, &book.Rating,
			&book.Book_status, &book.Price, &book.ImgPath)
		//add book to books
		books = append(books, book)
	}
	return books, nil
}

//AddBook 向数据库中添加一本图书
func AddBook(b *model.Book) error {
	//写sql语句
	slqStr := "insert into books(title,author,publisher, publish_date, rating, book_status, price,img_path) values(?,?,?,?,?,?,?,?)"
	//执行
	_, err := utils.Db.Exec(slqStr, b.Title, b.Author, b.Publisher, b.Publish_date, b.Rating, b.Book_status, b.Price, b.ImgPath)
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
	sqlStr := "select id,title,author,publisher, publish_date, rating, book_status, price,img_path from books where id = ?"
	//执行
	row := utils.Db.QueryRow(sqlStr, bookID)
	//创建Book
	book := &model.Book{}
	//为book中的字段赋值
	row.Scan(&book.ID, &book.Title, &book.Author, &book.Publisher, &book.Publish_date, &book.Rating,
		&book.Book_status, &book.Price, &book.ImgPath)
	return book, nil
}

//UpdateBook 根据图书的id更新图书信息
func UpdateBook(b *model.Book) error {
	//写sql语句
	sqlStr := "update books set title=?,author=?,publisher=?, publish_date=?, rating=?, book_status=?, price=? where id=?"
	//执行
	_, err := utils.Db.Exec(sqlStr, b.Title, b.Author, b.Publisher, b.Publish_date, b.Rating, b.Book_status,
		b.Price, b.ID)
	if err != nil {
		return err
	}
	return nil
}

//GetPageBooks 获取带分页的图书信息
func GetPageBooks(pageNo string) (*model.Page, error) {
	//将页码转换为int64类型
	iPageNo, _ := strconv.ParseInt(pageNo, 10, 64)
	//获取数据库中图书的总记录数
	sqlStr := "select count(*) from books"
	//设置一个变量接收总记录数
	var totalRecord int64
	//执行
	row := utils.Db.QueryRow(sqlStr)
	row.Scan(&totalRecord)
	//设置每页只显示4条记录
	var pageSize int64 = 4
	//设置一个变量接收总页数
	var totalPageNo int64
	if totalRecord%pageSize == 0 {
		totalPageNo = totalRecord / pageSize
	} else {
		totalPageNo = totalRecord/pageSize + 1
	}
	/*
		math.Ceil( totalRecord/PageSize )
		(totalRecord - 1)/pageSize + 1
	*/
	//获取当前页中的图书
	sqlStr2 := "select id,title,author,publisher, publish_date, rating, book_status, price,img_path from books limit ?,?"
	//执行
	rows, err := utils.Db.Query(sqlStr2, (iPageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Publisher, &book.Publish_date, &book.Rating,
			&book.Book_status, &book.Price, &book.ImgPath)
		//将book添加到books中
		books = append(books, book)
	}
	//创建page
	page := &model.Page{
		Books:       books,
		PageNo:      iPageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
	}
	return page, nil
}

//GetPageBooksByPrice 获取带分页和价格范围的图书信息
func GetPageBooksByPrice(pageNo string, minPrice string, maxPrice string) (*model.Page, error) {
	//将页码转换为int64类型
	iPageNo, _ := strconv.ParseInt(pageNo, 10, 64)
	//获取数据库中图书的总记录数
	sqlStr := "select count(*) from books where price between ? and ?"
	//设置一个变量接收总记录数
	var totalRecord int64
	//执行
	row := utils.Db.QueryRow(sqlStr, minPrice, maxPrice)
	row.Scan(&totalRecord)
	//设置每页只显示4条记录
	var pageSize int64 = 4
	//设置一个变量接收总页数
	var totalPageNo int64
	if totalRecord%pageSize == 0 {
		totalPageNo = totalRecord / pageSize
	} else {
		totalPageNo = totalRecord/pageSize + 1
	}
	/*
		math.Ceil( totalRecord/PageSize )
		(totalRecord - 1)/pageSize + 1
	*/
	//获取当前页中的图书
	sqlStr2 := "select id,title,author,publisher, publish_date, rating, book_status, price,img_path from books where price between ? and ? limit ?,?"
	//执行
	rows, err := utils.Db.Query(sqlStr2, minPrice, maxPrice, (iPageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Publisher, &book.Publish_date, &book.Rating,
			&book.Book_status, &book.Price, &book.ImgPath)
		//将book添加到books中
		books = append(books, book)
	}
	//创建page
	page := &model.Page{
		Books:       books,
		PageNo:      iPageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
	}
	return page, nil
}
