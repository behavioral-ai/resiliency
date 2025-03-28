package common

import "github.com/behavioral-ai/core/messaging"

const (
	AppHostKey    = "app-host"
	CacheHostKey  = "cache-host"
	TimeoutKey    = "timeout"
	regionKey     = "region"
	zoneKey       = "zone"
	subZoneKey    = "sub-zone"
	hostKey       = "host"
	instanceIdKey = "id"
	routeKey      = "routing"

	/*

	   WestRegion = "us-west1"
	   WestZoneA  = "w-zone-a"
	   WestZoneB  = "w-zone-b"

	   CentralRegion = "us-central1"
	   CentralZoneA  = "c-zone-a"
	   CentralZoneB  = "c-zone-b"

	   EastRegion = "us-east1"
	   EastZoneA  = "e-zone-a"
	   EastZoneB  = "e-zone-b"
	*/
)

var (
	origin Origin
	set    bool
)

type NewAgentFunc func(origin Origin, handler messaging.Agent)
