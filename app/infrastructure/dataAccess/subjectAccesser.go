package dataAccess

import (
	"exam-preparation-app/app/domain/model"
	"fmt"
	"log"
)

// SubjectAccesser subjectテーブルへのアクセッサーです。
type SubjectAccesser struct {
	DBAgent         *DBAgent
	FacultyAccesser *FacultyAccesser
}

// FindAll 学科一覧を取得します。
func (a *SubjectAccesser) FindAll() []model.Subject {
	rows, err := a.DBAgent.Conn.Query("SELECT * FROM subject")
	if err != nil {
		log.Printf(failedToGetData.value, err)
		return nil
	}
	defer rows.Close()
	var subjectResult []model.Subject
	var facultyID int
	for rows.Next() {
		subject := model.Subject{}
		if err := rows.Scan(&subject.ID, &subject.Name, &facultyID); err != nil {
			log.Printf(failedToGetData.value, err)
		}
		subject.Faculty = a.FacultyAccesser.FindByID(facultyID)
		subjectResult = append(subjectResult, subject)
	}
	return subjectResult
}

// FindByFacultyID 指定したFacultyIDの学科一覧を取得します。
func (a *SubjectAccesser) FindByFacultyID(facultyID int) []model.Subject {
	rows, err := a.DBAgent.Conn.Query(fmt.Sprintf("SELECT * FROM subject WHERE faculty_id = %d", facultyID))
	if err != nil {
		log.Printf(failedToGetData.value, err)
		return nil
	}
	defer rows.Close()
	var subjectResult []model.Subject
	var trash int
	for rows.Next() {
		subject := model.Subject{}
		if err := rows.Scan(&subject.ID, &subject.Name, &trash); err != nil {
			log.Printf(failedToGetData.value, err)
		}
		subject.Faculty = a.FacultyAccesser.FindByID(facultyID)
		subjectResult = append(subjectResult, subject)
	}
	return subjectResult
}

// FindByID 指定したIDの学科を取得します。
func (a *SubjectAccesser) FindByID(ID int) *model.Subject {
	rows, err := a.DBAgent.Conn.Query(fmt.Sprintf("SELECT * FROM subject WHERE id = %d", ID))
	if err != nil {
		log.Printf(failedToGetData.value, err)
		return nil
	}
	defer rows.Close()
	var facultyID int
	subject := model.Subject{}

	for rows.Next() {
		if err := rows.Scan(&subject.ID, &subject.Name, &facultyID); err != nil {
			log.Printf(failedToGetData.value, err)
		}
		subject.Faculty = a.FacultyAccesser.FindByID(facultyID)
	}
	return &subject
}
