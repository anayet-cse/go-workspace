package datastore

import (
	"api-test/loader"
	"log"
	"os"
	"strings"
)

type Students struct {
	Store *[]*loader.StudentData `json:"store"`
}

func (b *Students) Initialize() {
	filename := "./assets/student.csv"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	b.Store = loader.LoadData(file)
}

/*func (b *Students) SearchAuthor(author string, ratingOver, ratingBelow float64, limit, skip int) *[]*loader.StudentData {
	ret := Filter(b.Store, func(v *loader.StudentData) bool {
		return strings.Contains(strings.ToLower(v.Authors), strings.ToLower(author)) && v.AverageRating > ratingOver && v.AverageRating < ratingBelow
	})
	if limit == 0 || limit > len(*ret) {
		limit = len(*ret)
	}
	data := (*ret)[skip:limit]
	return &data
}

func (b *Students) SearchBook(bookName string, ratingOver, ratingBelow float64, limit, skip int) *[]*loader.StudentData {
	ret := Filter(b.Store, func(v *loader.StudentData) bool {
		return strings.Contains(strings.ToLower(v.Title), strings.ToLower(bookName)) && v.AverageRating > ratingOver && v.AverageRating < ratingBelow
	})
	if limit == 0 || limit > len(*ret) {
		limit = len(*ret)
	}

	data := (*ret)[skip:limit]
	return &data
}*/

func (b *Students) SearchExamRoll(exam_roll string) *loader.StudentData {
	ret := Filter(b.Store, func(v *loader.StudentData) bool {
		return strings.ToLower(v.ExamRoll) == strings.ToLower(exam_roll)
	})
	if len(*ret) > 0 {
		return (*ret)[0]
	}
	return nil
}

/*func (b *Students) CreateBook(book *loader.StudentData) bool {
	*b.Store = append(*b.Store, book)
	return true
}

func (b *Students) DeleteBook(isbn string) bool {
	indexToDelete := -1
	for i, v := range *b.Store {
		if v.ISBN == isbn {
			indexToDelete = i
			break
		}
	}
	if indexToDelete >= 0 {
		(*b.Store)[indexToDelete], (*b.Store)[len(*b.Store)-1] = (*b.Store)[len(*b.Store)-1], (*b.Store)[indexToDelete]
		*b.Store = (*b.Store)[:len(*b.Store)-1]
		return true
	}
	return false
}

func (b *Students) UpdateBook(isbn string, book *loader.StudentData) bool {
	for _, v := range *b.Store {
		if v.ISBN == isbn {
			v = book
			return true
		}
	}
	return false
}*/

func Filter(vs *[]*loader.StudentData, f func(*loader.StudentData) bool) *[]*loader.StudentData {
	vsf := make([]*loader.StudentData, 0)
	for _, v := range *vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return &vsf
}
