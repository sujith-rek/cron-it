package controller

import (
	"cronbackend/chores"
	"cronbackend/models"
	"cronbackend/utils"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ScheduleController struct {
	DB *gorm.DB
	CM *chores.ClusterManager
}

func NewScheduleController(db *gorm.DB) *ScheduleController {

	cm := chores.CreateClusterManager()

	sc := &ScheduleController{
		DB: db,
		CM: cm,
	}

	sc.recoverCluster()

	return sc
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

	if !utils.ValidateCronString(inputJob.ExecString) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cron string"})
		return
	}

	clusterID = sc.CM.FindClusterByExecString(inputJob.ExecString)
	clusterName := utils.ExtractNameFromExecString(inputJob.ExecString)

	if clusterID == "" {

		// generate a new cluster ID with a UUID generator
		clusterID = uuid.New().String()

		clusterID = sc.CM.CreateCluster(clusterName, inputJob.ExecString, clusterID)

		if clusterID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Max cluster size reached"})
			return
		}

		clusterDB = models.Cluster{
			ID:              clusterID,
			Name:            clusterName,
			ExecutionString: inputJob.ExecString,
			Size:            0,
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

	// update cluster size in db
	res = sc.DB.Model(&models.Cluster{}).Where("id = ?", clusterID).Update("size", gorm.Expr("size + 1"))

	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": res.Error.Error()})
		return
	}

	// update user limit in db done through trigger
	jobChore := chores.Job{
		ID:               jobDB.ID,
		Name:             jobDB.Name,
		UserID:           jobDB.UserID,
		ClusterID:        jobDB.ClusterID,
		ExecString:       jobDB.ExecString,
		AdditionalParams: jobDB.AdditionalParams,
		URL:              jobDB.URL,
	}

	sc.CM.AddJobToCluster(clusterID, jobChore)

	c.JSON(http.StatusOK, gin.H{"job": jobDB})

}

func (sc *ScheduleController) PrintCluster(c *gin.Context) {

	for _, cluster := range sc.CM.Clusters {

		for _, job := range cluster.Jobs {
			fmt.Println("Cluster: ", cluster.Name)
			fmt.Println("Job: ", job.Name)
			fmt.Println("ExecString: ", job.ExecString)
			fmt.Println("URL: ", job.URL)
			fmt.Println("AdditionalParams: ", string(job.AdditionalParams))
			fmt.Println("-------------------------------------------------")
		}

		fmt.Println("=====================================================")

	}

	c.JSON(http.StatusOK, gin.H{"clusters": sc.CM.Clusters})

}

func (sc *ScheduleController) recoverCluster() {

	// fetch all clusters from db
	var clusters []models.Cluster

	// fetch clusters along with jobs
	res := sc.DB.Preload("Jobs").Find(&clusters)

	if res.Error != nil {
		fmt.Println("Error fetching clusters: ", res.Error)
		return
	}

	for _, cluster := range clusters {

		jobChore := chores.Cluster{
			ID:              cluster.ID,
			Name:            cluster.Name,
			ExecutionString: cluster.ExecutionString,
			Size:            cluster.Size,
			Jobs:            []chores.Job{},
		}

		for _, job := range cluster.Jobs {
			jobChore.Jobs = append(jobChore.Jobs, chores.Job{
				ID:               job.ID,
				Name:             job.Name,
				UserID:           job.UserID,
				ClusterID:        job.ClusterID,
				ExecString:       job.ExecString,
				AdditionalParams: job.AdditionalParams,
				URL:              job.URL,
			})
		}

		sc.CM.AddClusterForRecovery(jobChore)
	}

	fmt.Println("CLusterManager Recovered")

}
