package routes

import (
	"os"

	"github.com/BetterGR/api-gateway/pkg/controllers"
	gradesProtos "github.com/BetterGR/grades-microservice/protos"
	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
)

// InitiateGradesMicroservice initialize grades microservice
func InitiateGradesMicroservice(router *gin.Engine) {
	_, err := RegisterGradesRoutes(router)
	if err != nil {
		klog.Fatalf("Failed to register grades routes, %v", err)
	}
}

func RegisterGradesRoutes(router *gin.Engine) (gradesProtos.GradesServiceClient, error) {
	// Initialize the gRPC client connection.
	gradesAddress := os.Getenv("GRADES_ADDRESS")
	klog.Infof("Grades address: %s", gradesAddress)
	grpcClient, err := controllers.InitGradesGRPCClient(gradesAddress)
	if err != nil {
		klog.Fatalf("Failed to initialize gRPC client, %v", err)
	}

	// Rest endpoints.
	router.GET("/grades/course/:courseId/semester/:semesterId", func(c *gin.Context) {
		controllers.GetCourseGrades(c, grpcClient)
	})
	router.GET("/grades/student/:studentId/course/:courseId/semester/:semester", func(c *gin.Context) {
		controllers.GetStudentCourseGrades(c, grpcClient)
	})
	router.POST("/grades/create", func(c *gin.Context) {
		controllers.AddSingleGrade(c, grpcClient)
	})
	router.PUT("/grades/update/", func(c *gin.Context) {
		controllers.UpdateSingleGrade(c, grpcClient)
	})
	router.DELETE("/grades/delete", func(c *gin.Context) {
		controllers.DeleteSingleGrade(c, grpcClient)
	})
	router.GET("/grades/student/:studentId/semester/:semester", func(c *gin.Context) {
		controllers.GetStudentSemesterGrades(c, grpcClient)
	})

	return grpcClient, nil
}
