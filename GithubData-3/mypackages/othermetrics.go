package mypackages

import (
	"encoding/json"
	"math"
)

type clustersummary struct {
	Stats                       stats   `json:"stats"`
	AvailableMetadataSpace      int64   `json:"availableMetadataSpace"`
	UsedMetadataSpacePercentage float64 `json:"usedMetadataSpacePct"`
}

type stats struct {
	LocalUsageStats localusagestats `json:"localUsagePerfStats"`
}
type localusagestats struct {
	TotalPhysicalUsageBytes int64 `json:"totalPhysicalUsageBytes"`
}

func UrlData() (responsedata clustersummary) {

	url := "https://10.14.19.226/irisservices/api/v1/public/cluster?fetchStats=true"
	response := PostRequestForAccessToken()
	data := GetRequestForJsonData(response, url)

	json.Unmarshal(data, &responsedata)
	return
}

func ClusterUsage() (cluster_usage float64) {
	responsedata := UrlData()

	total_physical_usage_bytes := responsedata.Stats.LocalUsageStats.TotalPhysicalUsageBytes

	cluster_usage = ContertToTeraBytes(total_physical_usage_bytes)

	return
}

func MetaDataUtilization() (metadata float64, used_meta_space_percentage float64) {
	responsedata := UrlData()
	available_meta_space := responsedata.AvailableMetadataSpace

	used_meta_space_percentage = responsedata.UsedMetadataSpacePercentage

	metadata = ContertToTeraBytes(available_meta_space)

	return
}

func ContertToTeraBytes(space int64) (space_converted float64) {

	terabyte := 1024 * 1024 * 1024 * 1024

	space_converted = float64(space) / float64(terabyte)

	space_converted = math.Round(space_converted*100) / 100

	return
}
