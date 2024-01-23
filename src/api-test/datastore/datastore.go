package datastore

import "api-test/loader"

type StudentStore interface {
	Initialize()
	//SearchAuthor(author string, ratingOver, ratingBelow float64, limit, skip int) *[]*loader.StudentData
	//SearchBook(bookName string, ratingOver, ratingBelow float64, limit, skip int) *[]*loader.StudentData
	SearchExamRoll(exam_roll string) *loader.StudentData
	//CreateBook(book *loader.StudentData) bool
	//DeleteBook(isbn string) bool
	//UpdateBook(isbn string, book *loader.StudentData) bool
}
