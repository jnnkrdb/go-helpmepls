package synchronization

type HostList struct {
	Hosts               []HostListItem
	PeriodSeconds       int
	FailureThresholdMax int
}

type HostListItem struct {
	Host             Host
	failureThreshold int
}

// init a new HostList
func InitHostList(self Host, failureThresholdMax int, periodSeconds int) *HostList {

	var hostlist = &HostList{
		Hosts: []HostListItem{
			{Host: self, failureThreshold: 0},
		},
		PeriodSeconds:       periodSeconds,
		FailureThresholdMax: failureThresholdMax,
	}

	logWriter.Write([]byte("created new hostlist and appended self"))

	return hostlist
}

// 