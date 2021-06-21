package controller

import (
	"day02/bookstore/dao"
	"day02/bookstore/model"
	"html/template"
	"net/http"
	"strconv"
)

//GetPageBooksByPrice
func GetPageBooksByPrice(w http.ResponseWriter, r *http.Request) {
	//get page num
	pageNo := r.FormValue("pageNo")
	//price range
	minPrice := r.FormValue("min")
	maxPrice := r.FormValue("max")
	if pageNo == "" {
		pageNo = "1"
	}
	var page *model.Page
	if minPrice == "" && maxPrice == "" {
		//call paged books funciton from bookdao
		page, _ = dao.GetPageBooks(pageNo)
	} else {
		//call paged books with price funciton from bookdao
		page, _ = dao.GetPageBooksByPrice(pageNo, minPrice, maxPrice)

		page.MinPrice = minPrice
		page.MaxPrice = maxPrice
	}

	t := template.Must(template.ParseFiles("views/index.html"))

	t.Execute(w, page)
}

//GetPageBooks
func GetPageBooks(w http.ResponseWriter, r *http.Request) {

	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	//call paged books funciton from bookdao
	page, _ := dao.GetPageBooks(pageNo)

	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))

	t.Execute(w, page)
}

//DeleteBook
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	//get delete book id
	bookID := r.FormValue("bookId")
	//call delete book function from bookdao
	dao.DeleteBook(bookID)

	GetPageBooks(w, r)
}

//ToUpdateBookPage
func ToUpdateBookPage(w http.ResponseWriter, r *http.Request) {
	//get book id
	bookID := r.FormValue("bookId")

	book, _ := dao.GetBookByID(bookID)
	if book.ID > 0 {
		//update book
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))

		t.Execute(w, book)
	} else {
		//add book
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))

		t.Execute(w, "")
	}
}

//UpdateOrAddBook
func UpdateOrAddBook(w http.ResponseWriter, r *http.Request) {
	//get book info
	bookID := r.PostFormValue("bookId")
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	publisher := r.PostFormValue("publisher")
	publish_date := r.PostFormValue("publish_date")
	rating := r.PostFormValue("rating")
	book_status := r.PostFormValue("book_status")

	price := r.PostFormValue("price")

	//convert type
	iRating, _ := strconv.Atoi(rating)
	fPrice, _ := strconv.ParseFloat(price, 64)
	ibookID, _ := strconv.ParseInt(bookID, 10, 0)
	//create book
	book := &model.Book{
		ID:           int(ibookID),
		Title:        title,
		Author:       author,
		Publisher:    publisher,
		Publish_date: publish_date,
		Rating:       iRating,
		Book_status:  book_status,
		Price:        fPrice,
		ImgPath:      "static/img/default.jpg",
	}
	if book.ID > 0 {
		//In update book
		//call update book from bookdao
		dao.UpdateBook(book)
	} else {
		//In add book
		//call add book from bookdao
		dao.AddBook(book)
	}

	GetPageBooks(w, r)
}
