package controller

import (
	"cronbackend/chores"
	"cronbackend/models"
	"cronbackend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type ScheduleController struct {
	DB *gorm.DB
	CM *chores.ClusterManager
}

func NewScheduleController(db *gorm.DB) *ScheduleController {

	cm := chores.CreateClusterManager()

	return &ScheduleController{
		DB: db,
		CM: cm,
	}
}

func (sc *ScheduleController) CreateJobSchedule(c *gin.Context) {

	user, _ := c.Get("user")
	var inputJob models.JobInput
	var clusterDB models.Cluster
	var jobDB models.Job
	var clusterID string
	var userDB models.User

	if err := c.ShouldBindJSON(&inputJob); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// fetch user from db
	res := sc.DB.Where("id = ?", user.(models.User).ID).First(&userDB)

	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": res.Error.Error()})
		return
	}

	// if user.limit == 0 then abort
	if userDB.Limit == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User limit reached"})
		return
	}

	clusterID = sc.CM.FindClusterByExecString(inputJob.ExecString)
	clusterName := utils.ExtractNameFromExecString(inputJob.ExecString)

	if clusterID == "" {
		clusterID = sc.CM.CreateCluster(clusterName, inputJob.ExecString)

		if clusterID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Max cluster size reached"})
			return
		}

		clusterDB = models.Cluster{
			ID:              clusterID,
			Name:            clusterName,
			ExecutionString: inputJob.ExecString,
			Size:            1,
		}

		res := sc.DB.Create(&clusterDB)

		if res.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": res.Error.Error()})
			return
		}

	}

	jobDB = models.Job{
		Name:             inputJob.Name,
		ExecString:       inputJob.ExecString,
		UserID:           user.(models.User).ID,
		ClusterID:        clusterID,
		URL:              inputJob.URL,
		AdditionalParams: inputJob.AdditionalParams,
	}

	res = sc.DB.Create(&jobDB)

	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": res.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"job": jobDB})

}
