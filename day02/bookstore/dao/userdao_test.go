package dao

import (
	"day02/bookstore/model"
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("test funcitons in bookdao")
	m.Run()
}

func TestUser(t *testing.T) {
	// fmt.Println("test functions in userdao")
	// t.Run("check username and password：", testLogin)
	// t.Run("validate username：", testRegist)
	// t.Run("test save user：", testSave)
}

func testLogin(t *testing.T) {
	user, _ := CheckUserNameAndPassword("admin", "123456")
	fmt.Println("The test user is ：", user)
}
func testRegist(t *testing.T) {
	user, _ := CheckUserName("admin")
	fmt.Println("The user info is ：", user)
}
func testSave(t *testing.T) {
	SaveUser("admin3", "123456", "admin@atguigu.com")
}

func TestBook(t *testing.T) {
	fmt.Println("test funcitons in bookdao")
	t.Run("test getting all books", testGetBooks)
	// t.Run("test adding a book", testAddBook)
	// t.Run("test deleting a book", testDeleteBook)
	// t.Run("test getting a book", testGetBook)
	// t.Run("test updating a book", testUpdateBook)
	// t.Run("test getting paged books", testGetPageBooks)
	// t.Run("test getting paged books with price", testGetPageBooksByPrice)
}

func testGetBooks(t *testing.T) {
	books, _ := GetBooks()
	//traverse books
	for k, v := range books {
		fmt.Printf("%v book info is ：%v\n", k+1, v)
	}
}
func testAddBook(t *testing.T) {
	book := &model.Book{
		Title:        "Book11",
		Author:       "Mary",
		Publisher:    "PA_publisher",
		Publish_date: "2021-01-01",
		Rating:       2,
		Book_status:  "CheckingIn",
		Price:        88.88,
		ImgPath:      "/static/img/default.jpg",
	}

	AddBook(book)
}
func testDeleteBook(t *testing.T) {
	//test delete book
	DeleteBook("3")
}
func testGetBook(t *testing.T) {
	//调用获取图书的函数
	book, _ := GetBookByID("5")
	fmt.Println("获取的图书信息是：", book)
}
func testUpdateBook(t *testing.T) {
	book := &model.Book{
		ID:           2,
		Title:        "this is a book",
		Author:       "John M",
		Publisher:    "PA_23",
		Publish_date: "2020-09-19",
		Rating:       3,
		Book_status:  "CheckingIn",
		Price:        20.12,
		ImgPath:      "/static/img/default.jpg",
	}

	UpdateBook(book)
}

func testGetPageBooks(t *testing.T) {
	fmt.Println("start test in testGetPageBooks")
	page, err := GetPageBooks("2")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Current Page Number is：", page.PageNo)
	fmt.Println("Total Page Number is：", page.TotalPageNo)
	fmt.Println("Total Record is：", page.TotalRecord)
	fmt.Println("Books in the current page：")
	for _, v := range page.Books {
		fmt.Println("Book Information：", v)
	}
}
func testGetPageBooksByPrice(t *testing.T) {
	page, _ := GetPageBooksByPrice("3", "10", "30")
	fmt.Println("Current Page Number is：", page.PageNo)
	fmt.Println("Total Page Number is：", page.TotalPageNo)
	fmt.Println("Total Record is：", page.TotalRecord)
	fmt.Println("Books in the current page：")
	for _, v := range page.Books {
		fmt.Println("Book Information：", v)
	}
}
