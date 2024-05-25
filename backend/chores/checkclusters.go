package chores

import (
	"encoding/json"
	"time"
)

const (
	MaxCheckClusterSize = 4
)

type CheckJob struct {
	ID               string
	UserID           string
	UserMail         string
	ClusterID        string
	Name             string
	Body             json.RawMessage
	PingString       string
	ExecString       string
	LastCheck        time.Time
	LastCheckSuccess bool
	NextCheck        time.Time
}

type CheckCluster struct {
	ID              string
	Name            string
	ExecutionString string
	Jobs            []CheckJob
	Size            int
}

type CheckClusterManager struct {
	Size     int
	Clusters []CheckCluster
}

func CreateCheckClusterManager() *CheckClusterManager {
	return &CheckClusterManager{
		Size:     0,
		Clusters: []CheckCluster{},
	}
}

