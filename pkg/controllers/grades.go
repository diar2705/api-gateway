package controllers

import (
	"context"
	"net/http"

	gradesProtos "github.com/BetterGR/grades-microservice/protos"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// InitGradesGRPCClient initializes the grades-microservice gRPC client connection.
func InitGradesGRPCClient(address string) (gradesProtos.GradesServiceClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return gradesProtos.NewGradesServiceClient(conn), nil
}

// GetCourseGrades handles the request to get grades for a specific course in a specific semester.
func GetCourseGrades(c *gin.Context, client gradesProtos.GradesServiceClient) {
	courseID := c.Param("courseID")
	semester := c.Param("semester")

	// Call the gRPC method
	resp, err := client.GetCourseGrades(context.Background(), &gradesProtos.GetCourseGradesRequest{
		CourseID:   courseID,
		Semester: semester,

	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetStudentCourseGrades handles the request to get grades for a specific student in a specific course and semester.
func GetStudentCourseGrades(c *gin.Context, client gradesProtos.GradesServiceClient) {
	studentID := c.Param("studentID")
	courseID := c.Param("courseID")
	semester := c.Param("semester")

	// Call the gRPC method
	resp, err := client.GetStudentCourseGrades(context.Background(), &gradesProtos.GetStudentCourseGradesRequest{
		StudentID:  studentID,
		CourseID:   courseID,
		Semester: semester,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// AddSingleGrade handles the request to add a single grade.
func AddSingleGrade(c *gin.Context, client gradesProtos.GradesServiceClient) {
	var grade gradesProtos.SingleGrade
	if err := c.ShouldBindJSON(&grade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the gRPC method
	resp, err := client.AddSingleGrade(context.Background(), &gradesProtos.AddSingleGradeRequest{
		Grade: &grade,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateSingleGrade handles the request to update a single grade.
func UpdateSingleGrade(c *gin.Context, client gradesProtos.GradesServiceClient) {
	var grade gradesProtos.SingleGrade
	if err := c.ShouldBindJSON(&grade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the gRPC method
	resp, err := client.UpdateSingleGrade(context.Background(), &gradesProtos.UpdateSingleGradeRequest{
		Grade: &grade,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteSingleGrade handles the request to delete a single grade.
func DeleteSingleGrade(c *gin.Context, client gradesProtos.GradesServiceClient) {
	var grade gradesProtos.SingleGrade
	if err := c.ShouldBindJSON(&grade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the gRPC method
	resp, err := client.RemoveSingleGrade(context.Background(), &gradesProtos.RemoveSingleGradeRequest{
		GradeID: grade.GetGradeID(),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetStudentSemesterGrades handles the request to get grades for a specific student in a specific semester.
func GetStudentSemesterGrades(c *gin.Context, client gradesProtos.GradesServiceClient) {
	studentId := c.Param("studentId")
	semester := c.Param("semester")

	// Call the gRPC method
	resp, err := client.GetStudentSemesterGrades(context.Background(), &gradesProtos.GetStudentSemesterGradesRequest{
		StudentID:  studentId,
		Semester: semester,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
