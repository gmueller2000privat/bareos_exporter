package main

import "time"

type JobType string

type JobLookup struct {
	JobName     string
	clientId    int64
	FileSetName string
}

// JobInfo models query results for static values of a job
type JobInfo struct {
	JobLookup
	JobType    JobType
	JobName    string
	ClientName string

	TotalCount int64
	TotalBytes int64
	TotalFiles int64
}

// LastJob models query results for job metrics
type LastJob struct {
	JobID        uint64
	JobStatus    string
	JobBytes     int64
	JobFiles     int64
	JobErrors    int64
	JobStartDate time.Time
	JobEndDate   time.Time
}

// PoolInfo models query result of pool information
type PoolInfo struct {
	Name     string
	Volumes  int64
	Bytes    int64
	Prunable bool
	Expired  bool
}
