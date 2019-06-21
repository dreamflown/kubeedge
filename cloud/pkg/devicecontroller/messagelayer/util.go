package messagelayer

import (
	"errors"
	"fmt"
	"strings"

	deviceconstants "github.com/kubeedge/kubeedge/cloud/pkg/devicecontroller/constants"
	constants "github.com/kubeedge/kubeedge/common/constants"
)

// BuildResource return a string as "beehive/pkg/core/model".Message.Router.Resource
func BuildResource(nodeID, resourceType, resourceID string) (resource string, err error) {
	if nodeID == "" || resourceType == "" {
		err = fmt.Errorf("required parameter are not set (node id, namespace or resource type)")
		return
	}
	resource = fmt.Sprintf("%s%s%s%s%s", deviceconstants.ResourceNode, constants.ResourceSep, nodeID, constants.ResourceSep, resourceType)
	if resourceID != "" {
		resource += fmt.Sprintf("%s%s", constants.ResourceSep, resourceID)
	}
	return
}

// GetDeviceID returns the ID of the device
func GetDeviceID(resource string) (string, error) {
	res := strings.Split(resource, "/")
	if len(res) >= deviceconstants.ResourceDeviceIDIndex+1 && res[deviceconstants.ResourceDeviceIndex] == deviceconstants.ResourceDevice {
		return res[deviceconstants.ResourceDeviceIDIndex], nil
	}
	return "", errors.New("failed to get device id")
}

// GetResourceType returns the resourceType of message received from edge
func GetResourceType(resource string) (string, error) {
	// resource as below
	// $hw/event/device/+/updated
	// $hw/event/device/+/state/(update,get)
	// $hw/event/device/+/node/+/membership/updated
	// $hw/event/device/+/twin/edge_updated
	if strings.Contains(resource, deviceconstants.ResourceTypeTwinEdgeUpdated) {
		return deviceconstants.ResourceTypeTwinEdgeUpdated, nil
	} else if strings.Contains(resource, deviceconstants.ResourceNode) {
		return deviceconstants.ResourceNode, nil
	} else if strings.Contains(resource, deviceconstants.ResourceDevice) {
		return deviceconstants.ResourceDevice, nil
	}
	return "", errors.New("unknown resource")
}

// GetNodeID returns the nodeID of device create message received from edge
func GetNodeID(resource string) (string, error) {
	res := strings.Split(resource, "/")
	if len(res) > deviceconstants.ResourceDeviceIDIndex+3 && res[deviceconstants.ResourceDeviceIndex] == deviceconstants.ResourceDevice && res[deviceconstants.ResourceDeviceIndex+2] == deviceconstants.ResourceDevice {
		return res[deviceconstants.ResourceDeviceIndex+3], nil
	}
	return "", errors.New("failed to get device id and node id")
}
