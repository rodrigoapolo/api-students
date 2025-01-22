package db

import (
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"github.com/rodrigoapolo/api-students/schemas"
)

type StudentHandler struct{
	DB *gorm.DB
}

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("student.db"), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to initialize SQLite: %s", err.Error())
	}

	db.AutoMigrate(&schemas.Student{})

	return db
}

func NewStudentHanler(db *gorm.DB) *StudentHandler{
	return &StudentHandler{DB: db}
}

func (s *StudentHandler) AddStudent(student schemas.Student) error{
	if result := s.DB.Create(&student); result.Error != nil{
		log.Error().Msg("Error to create student")
		return result.Error
	}

	log.Info().Msg("Create student")
	return nil
}

func (s *StudentHandler) GetStudents() ([]schemas.Student, error) {
	students := []schemas.Student{}

	err := s.DB.Find(&students).Error

	return students, err
}

func (s *StudentHandler) GetStudent(id int) (schemas.Student, error) {
	student := schemas.Student{}

	err := s.DB.First(&student, id)
	
	return student, err.Error
}

func (s *StudentHandler) UpdateStudent(updateStudent schemas.Student) error {
		return s.DB.Save(&updateStudent).Error
}


func (s *StudentHandler) DeleteStudent(student schemas.Student) error {
	return s.DB.Delete(&student).Error
}

func (s *StudentHandler) GetFilteredStudent(active bool) ([]schemas.Student, error) {
	filteredStudents := []schemas.Student{}
	err := s.DB.Where("active = ?", active).Find(&filteredStudents)

	return filteredStudents, err.Error
}
