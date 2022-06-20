package mypackages

import (
	"encoding/json"
	"fmt"
)

type jsondata struct {
	Id                   	 	string 		`json:"id"`
	TimeStamp             		int64 		`json:"timeStamp"`
	FileCreateRate        		int64 		`json:"fileCreateRate"`
	UtilizationChangeRate 		float64		`json:"utilizationChangeRate"`
	UtilizationChangeRateUnit 	string 		`json:"utilizationChangeRateUnit"`
	GarbageCollection     		float64		`json:"garbageCollection"`
	GarbageCollectionUnit 		string 		`json:"garbageCollectionUnit"`
	ClusterUsage 				float64 	`json:"clusterUsage"`
	ClusterUsageUnits 			string 		`json:"clusterUsageUnits"`
	MetaDataUtilization 		float64		`json:"metaDataUtilization"`
	MetaDataUtilizationUnits 	string 		`json:"metaDataUtilizationUnits"`
	MetaDataSpacePercentage 	float64 	`json:"meta_data_space_percentage"`
	ProtectionJobsInfo    		[]protectiongroup 	`json:"protectionGroups"`
}

type protectiongroup struct {
	ProtectionGroupName 		string `json:"protectionGroupName"`
	ProtectionGroupType 		string   `json:"protectionGroupType"`
	JobBackuptime       		string  `json:"jobBackuptime"`
	JobReplicationTime  		string  `json:"jobReplicationTime"`
	JobArchivalTime     		string  `json:"jobArchivalTime"`
	SlaTime             		int    `json:"slaTimes"`
	Runs                		int    `json:"runs"`
}

var Response_for_elastic jsondata

func GenerateJson()(data []byte){
	FillProtectionJobKeys()
	total_jobs := len(ProtectionJobsList)
	var myProtectionGroups []protectiongroup
	for i:=0;i<total_jobs;i++{	
		job_environment,job_backup_time,job_replication_time,job_archival_time,sla_times,runs := ProtectionJobInfo(ProtectionJobsList[i])
		mygroup := protectiongroup{
			ProtectionGroupName: ProtectionJobsList[i],
			ProtectionGroupType: ProtectiongroupEnvironment[job_environment],
			JobBackuptime:       job_backup_time,
			JobReplicationTime:  job_replication_time,
			JobArchivalTime:     job_archival_time,
			SlaTime:             sla_times,
			Runs:                runs,
			}
			myProtectionGroups = append(myProtectionGroups, mygroup)
	}
	
	file_create_rate := Filecreaterate()
	utilization_change_rate := UtilizationChangeRate()
	garbage_collection := GarbageCollection()
	Present_time_stamp := PresentTimeStamp()
	cluster_usage := ClusterUsage()
	meta_data_utilization,meta_data_space_percentage := MetaDataUtilization()

	myjsondata := jsondata{
		Id:                    "longivitycluster",
		TimeStamp:             Present_time_stamp,
		FileCreateRate:        file_create_rate,
		UtilizationChangeRate: utilization_change_rate,
		UtilizationChangeRateUnit: "GiB/sec",
		GarbageCollection:     garbage_collection,
		GarbageCollectionUnit: "GiB",
		ClusterUsage: cluster_usage,
		ClusterUsageUnits: "TiB",
		MetaDataUtilization: meta_data_utilization,
		MetaDataUtilizationUnits: "TiB",
		MetaDataSpacePercentage: meta_data_space_percentage,
		ProtectionJobsInfo:    myProtectionGroups,
	}

	data, _ = json.Marshal(myjsondata)
	fmt.Printf("%s", data)
		return 
}