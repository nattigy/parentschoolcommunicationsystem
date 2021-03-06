package studentServices

import "github.com/nattigy/parentschoolcommunicationsystem/models"

type StudentUsecase interface {
	AddStudent(newStudent models.Student) []error
	GetStudents() ([]models.Student, []error)
	GetStudentById(id uint) (models.Student, []error)
	DeleteStudent(id uint) []error
	UpdateStudent(newStudent models.Student) (models.Student, []error)
	ViewTasks(sectionId uint, subjectId uint) ([]models.Task, []error)
	Comment(taskId uint, studentId uint, studentName string, data string) []error
	Comments(taskId uint) ([]models.Comment, []error)
	ViewClass(classRoomId uint) ([]models.Student, []error)
	ViewResources(subjectId uint) ([]models.Resources, []error)
	ViewResult(studentId uint) (models.Student, []error)
	GetHomeRoomTeacher(studentId uint) (models.Teacher, []error)
}
