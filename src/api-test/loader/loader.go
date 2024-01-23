package loader

import (
	"encoding/csv"
	"io"
	"log"
)

type StudentData struct {
	StudentName  string `json:"student_name"`
	Registration string `json:"registration"`
	ExamRoll     string `json:"exam_roll"`
	Year         string `json:"year"`
	Result       string `json:"result"`
}

func LoadData(r io.Reader) *[]*StudentData {
	reader := csv.NewReader(r)

	ret := make([]*StudentData, 0)

	for {
		row, err := reader.Read()

		if err == io.EOF {
			log.Println("End of File")
			break
		} else if err != nil {
			log.Println(err)
			break
		}
		/*averageRating, _ := strconv.ParseFloat(row[3], 64)
		numPages, _ := strconv.Atoi(row[7])
		ratings, _ := strconv.Atoi(row[8])
		reviews, _ := strconv.Atoi(row[9])*/

		if err != nil {
			log.Println(err)
		}
		student := &StudentData{
			StudentName:  row[0],
			Registration: row[1],
			ExamRoll:     row[2],
			//Year:         row[3],
			Result: row[3],
		}

		if err != nil {
			log.Fatalln(err)
		}

		ret = append(ret, student)
	}
	return &ret
}
