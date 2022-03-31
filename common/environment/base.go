package environment

import (
	"github.com/v2fly/v2ray-core/v5/common/environment/filesystemcap"
	"github.com/v2fly/v2ray-core/v5/common/log"
	"github.com/v2fly/v2ray-core/v5/transport/internet"
	"github.com/v2fly/v2ray-core/v5/transport/internet/tagged"
)

type BaseEnvironmentCapabilitySet interface {
	FeaturesLookupCapabilitySet
	LogCapabilitySet
}

type BaseEnvironment interface {
	BaseEnvironmentCapabilitySet
	doNotImpl()
}

type SystemNetworkCapabilitySet interface {
	Dialer() internet.SystemDialer
	Listener() internet.SystemListener
}

type InstanceNetworkCapabilitySet interface {
	OutboundDialer() tagged.DialFunc
}

type FeaturesLookupCapabilitySet interface {
	RequireFeatures(callback interface{}) error
}

type LogCapabilitySet interface {
	RecordLog(msg log.Message)
}

type FileSystemCapabilitySet interface {
	filesystemcap.FileSystemCapabilitySet
}
