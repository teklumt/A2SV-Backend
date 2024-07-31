package controllers

import (
	"net/http"
	"task_manager/model"
	"task_manager/services"

	"github.com/gin-gonic/gin"
)


func GetAllTasks(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, gin.H{"tasks": services.FindAllData()})
}

func GetSpecificTask(context *gin.Context) {
    id := context.Param("id")

    task := services.FindByIndex(id)
    
    if task == (model.Task{}) {
        context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Resource Not Found"})
        return
    }

    context.IndentedJSON(http.StatusOK, task)
}



func UpdateSpecificTask(context *gin.Context){
	id := context.Param("id")

	var updatedTask  model.Task

	if err := context.ShouldBindJSON(&updatedTask); err != nil{
		context.IndentedJSON(http.StatusBadRequest,  err.Error())
	}  


	var Updated  bool = services.UpdateTask(id , updatedTask)
	
	if  Updated{
		context.IndentedJSON(http.StatusOK, gin.H{"message":"Successfully Updated"})
		return
	}
	
	context.IndentedJSON(http.StatusBadGateway ,gin.H{"message":"Task not found"})
	
}



func DeleteSpecificTask(context *gin.Context) {
	id := context.Param("id")

	var Deleted bool = services.DeleteTask(id)

	if Deleted{
		context.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully Deleted"})
		return
	}

	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Resourse Not Found"})
}



func AddSpecificTask(context *gin.Context) {
	var newTask model.Task

	if err := context.ShouldBindJSON(&newTask); err != nil {
		context.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	var Added bool = services.AddTask(newTask)

	if Added{
		context.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully Added"})
		return 
	}
	context.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "Not Added Successfully"})

}