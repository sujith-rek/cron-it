package chores

import (
	"encoding/json"
	"github.com/google/uuid"
)

const (
	MaxScheduleClusterSize = 10
)

type Job struct {
	ID               string
	Name             string
	UserID           string
	ClusterID        string
	ExecString       string
	AdditionalParams json.RawMessage
	URL              string
}

type Cluster struct {
	ID              string
	Name            string
	ExecutionString string
	Jobs            []Job
	Size            int
}

type SchedulesClusterManager struct {
	Size     int
	Clusters []Cluster
}

// CreateScheduleClusterManager creates a new cluster manager
func CreateSchedulesClusterManager() *SchedulesClusterManager {
	return &SchedulesClusterManager{
		Size:     0,
		Clusters: []Cluster{},
	}
}

func (cm *SchedulesClusterManager) CreateScheduleCluster(name string, executionString string, id string) string {
	if len(cm.Clusters) >= MaxScheduleClusterSize {
		return ""
	}

	// generate a new cluster ID with a UUID generator
	newCluster := Cluster{
		ID:              id,
		Name:            name,
		ExecutionString: executionString,
		Size:            0,
	}

	cm.Clusters = append(cm.Clusters, newCluster)
	cm.Size++

	return newCluster.ID

}


func (cm *SchedulesClusterManager) DeleteScheduleCluster(clusterID string) {
	for i, cluster := range cm.Clusters {
		if cluster.ID == clusterID {
			cm.Clusters = append(cm.Clusters[:i], cm.Clusters[i+1:]...)
			cm.Size--
			return
		}
	}
}

func (cm *SchedulesClusterManager) AddJobToScheduleCluster(clusterID string, job Job) {
	for i, cluster := range cm.Clusters {
		if cluster.ID == clusterID {
			cm.Clusters[i].Jobs = append(cm.Clusters[i].Jobs, job)
			cm.Clusters[i].Size++

			if cm.Clusters[i].Size > MaxScheduleClusterSize {
				cm.splitScheduleCluster(clusterID)
			}

			// call cron job scheduler here

			return
		}
	}
}

func (cm *SchedulesClusterManager) RemoveJobFromScheduleCluster(clusterID string, job Job) {
	for i, cluster := range cm.Clusters {
		if cluster.ID == clusterID {
			for j, clusterJob := range cluster.Jobs {
				if clusterJob.ID == job.ID {
					cm.Clusters[i].Jobs = append(cm.Clusters[i].Jobs[:j], cm.Clusters[i].Jobs[j+1:]...)
					cm.Clusters[i].Size--
					return
				}
			}
		}
	}
}

func (cm *SchedulesClusterManager) splitScheduleCluster(clusterID string) {
	for _, cluster := range cm.Clusters {
		if cluster.ID == clusterID {
			newCluster := Cluster{
				ID:              uuid.New().String(),
				Name:            cluster.Name,
				ExecutionString: cluster.ExecutionString,
				Size:            0,
			}

			cm.Clusters = append(cm.Clusters, newCluster)

			// move half of the jobs to the new cluster
			for j := cluster.Size/2 - 1; j > cluster.Size/2; j-- {
				newCluster.Jobs = append(newCluster.Jobs, cluster.Jobs[j])
				cluster.Jobs = append(cluster.Jobs[:j], cluster.Jobs[j+1:]...)
				cluster.Size--
				newCluster.Size++
			}

			cm.Size++

			// call db update here

			return
		}
	}
}

func (cm *SchedulesClusterManager) FindScheduleClusterByExecString(executionString string) string {
	var clusters []Cluster
	for _, cluster := range cm.Clusters {
		if cluster.ExecutionString == executionString {
			clusters = append(clusters, cluster)
		}
	}

	// return cluster with least jobs
	if len(clusters) > 0 {
		minJobs := clusters[0].Size
		minCluster := clusters[0]
		for _, cluster := range clusters {
			if cluster.Size < minJobs {
				minJobs = cluster.Size
				minCluster = cluster
			}
		}
		return minCluster.ID
	}

	return ""
}


// RECOVERY FUNCTIONS
func (cm *SchedulesClusterManager) AddScheduleClusterForRecovery(cluster Cluster) {
	cm.Clusters = append(cm.Clusters, cluster)
	cm.Size++
}

func (cm *SchedulesClusterManager) RecoverScheduleClusterManager(clusters []Cluster) bool {
	for _, cluster := range clusters {
		cm.AddScheduleClusterForRecovery(cluster)
	}

	return true
}