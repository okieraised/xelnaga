package internal

import "time"

const (
	TerminationTimeout = 1 * time.Second
)

const (
	SOPInstanceUID    = "sop_instance_uid"
	SeriesInstanceUID = "series_instance_uid"
	StudyInstanceUID  = "study_instance_uid"
)

const (
	WellKnownPort = 104
	ISCLPort      = 2761
	TLSPort       = 2762
	OpenCommPort  = 11112
)
