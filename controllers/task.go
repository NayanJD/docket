package controllers

import (
	"nayanjd/docket/models"
	"nayanjd/docket/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type TaskForm interface {
	getDescription() *string
	getScheduledFor() *time.Time
}

type TaskInputForm struct {
	Description   *string    `form:"description"   binding:"required"`
	Scheduled_for *time.Time `form:"scheduled_for" binding:"required"`
}

func (f *TaskInputForm) getDescription() *string {
	return f.Description
}

func (f *TaskInputForm) getScheduledFor() *time.Time {
	return f.Scheduled_for
}

type PatchTaskInputForm struct {
	Description   *string    `form:"description"`
	Scheduled_for *time.Time `form:"scheduled_for"`
}

func (f *PatchTaskInputForm) getDescription() *string {
	return f.Description
}

func (f *PatchTaskInputForm) getScheduledFor() *time.Time {
	return f.Scheduled_for
}

type TaskController struct{}

func (ctrl *TaskController) Create(c *gin.Context) {
	taskForm := c.MustGet(gin.BindKey).(*TaskInputForm)

	user := c.MustGet(gin.AuthUserKey).(models.User)

	newTask := models.Task{
		Description:   taskForm.Description,
		Scheduled_for: taskForm.Scheduled_for,
		UserID:        user.ID,
	}

	err := models.GetDB().Create(&newTask).Error

	if err != nil {
		c.Error(err).SetType(utils.ErrorTypeDB)
		return
	}

	utils.AbortWithGenericJson(c, utils.CreateOKResponse(newTask, nil), nil)
}

func (ctl *TaskController) GetUserTasks(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(models.User)

	tasks := []models.Task{}

	log.Error().Msg("Starting query")
	if err := models.GetDB().Preload("Tags").Where("user_id = ?", user.ID).Find(&tasks).Error; err != nil {
		c.Error(err).SetType(utils.ErrorTypeDB)
		return
	}

	utils.AbortWithGenericJson(c, utils.CreateOKResponse(tasks, nil), nil)
}

func (ctl *TaskController) UpdateUserTask(c *gin.Context) {
	var taskForm TaskForm

	taskForm, ok := c.MustGet(gin.BindKey).(*TaskInputForm)

	if !ok {
		taskForm = c.MustGet(gin.BindKey).(*PatchTaskInputForm)
	}

	taskId := c.Param("id")

	user := c.MustGet(gin.AuthUserKey).(models.User)

	task := models.Task{}

	if err := models.GetDB().Where("user_id = ? and id = ?", user.ID, taskId).First(&task).Error; err != nil {
		c.Error(err).SetType(utils.ErrorTypeDB)
		return
	}

	if taskForm.getDescription() != nil {
		task.Description = taskForm.getDescription()
	}

	if taskForm.getScheduledFor() != nil {
		task.Scheduled_for = taskForm.getScheduledFor()
	}

	if err := models.GetDB().Save(&task).Error; err != nil {
		c.Error(err).SetType(utils.ErrorTypeDB)
		return
	}

	utils.AbortWithGenericJson(c, utils.CreateOKResponse(task, nil), nil)
}
