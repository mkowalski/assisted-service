// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ClusterCreateParams cluster create params
//
// swagger:model cluster-create-params
type ClusterCreateParams struct {

	// A comma-separated list of NTP sources (name or IP) going to be added to all the hosts.
	AdditionalNtpSource *string `json:"additional_ntp_source,omitempty"`

	// The virtual IP used to reach the OpenShift cluster's API.
	// Pattern: ^(?:(?:(?:[0-9]{1,3}\.){3}[0-9]{1,3})|(?:(?:[0-9a-fA-F]*:[0-9a-fA-F]*){2,}))?$
	APIVip string `json:"api_vip,omitempty"`

	// The virtual IPs used to reach the OpenShift cluster's API.
	APIVips []*APIVip `json:"api_vips"`

	// Base domain of the cluster. All DNS records must be sub-domains of this base and include the cluster name.
	BaseDNSDomain string `json:"base_dns_domain,omitempty"`

	// IP address block from which Pod IPs are allocated. This block must not overlap with existing physical networks. These IP addresses are used for the Pod network, and if you need to access the Pods from an external network, configure load balancers and routers to manage the traffic.
	// Pattern: ^(?:(?:(?:[0-9]{1,3}\.){3}[0-9]{1,3}\/(?:(?:[0-9])|(?:[1-2][0-9])|(?:3[0-2])))|(?:(?:[0-9a-fA-F]*:[0-9a-fA-F]*){2,})/(?:(?:[0-9])|(?:[1-9][0-9])|(?:1[0-1][0-9])|(?:12[0-8])))$
	ClusterNetworkCidr *string `json:"cluster_network_cidr,omitempty"`

	// The subnet prefix length to assign to each individual node. For example, if clusterNetworkHostPrefix is set to 23, then each node is assigned a /23 subnet out of the given cidr (clusterNetworkCIDR), which allows for 510 (2^(32 - 23) - 2) pod IPs addresses. If you are required to provide access to nodes from an external network, configure load balancers and routers to manage the traffic.
	// Maximum: 128
	// Minimum: 1
	ClusterNetworkHostPrefix int64 `json:"cluster_network_host_prefix,omitempty"`

	// Cluster networks that are associated with this cluster.
	ClusterNetworks []*ClusterNetwork `json:"cluster_networks"`

	// The CPU architecture of the image (x86_64/arm64/etc).
	CPUArchitecture string `json:"cpu_architecture,omitempty"`

	// Installation disks encryption mode and host roles to be applied.
	DiskEncryption *DiskEncryption `json:"disk_encryption,omitempty" gorm:"embedded;embeddedPrefix:disk_encryption_"`

	// Guaranteed availability of the installed cluster. 'Full' installs a Highly-Available cluster
	// over multiple master nodes whereas 'None' installs a full cluster over one node.
	//
	// Enum: [Full None]
	HighAvailabilityMode *string `json:"high_availability_mode,omitempty"`

	// A proxy URL to use for creating HTTP connections outside the cluster.
	// http://\<username\>:\<pswd\>@\<ip\>:\<port\>
	//
	HTTPProxy *string `json:"http_proxy,omitempty"`

	// A proxy URL to use for creating HTTPS connections outside the cluster.
	// http://\<username\>:\<pswd\>@\<ip\>:\<port\>
	//
	HTTPSProxy *string `json:"https_proxy,omitempty"`

	// Enable/disable hyperthreading on master nodes, worker nodes, or all nodes.
	// Enum: [masters workers none all]
	Hyperthreading *string `json:"hyperthreading,omitempty"`

	// Explicit ignition endpoint overrides the default ignition endpoint.
	IgnitionEndpoint *IgnitionEndpoint `json:"ignition_endpoint,omitempty" gorm:"embedded;embeddedPrefix:ignition_endpoint_"`

	// The virtual IP used for cluster ingress traffic.
	// Pattern: ^(?:(?:(?:[0-9]{1,3}\.){3}[0-9]{1,3})|(?:(?:[0-9a-fA-F]*:[0-9a-fA-F]*){2,}))$
	IngressVip string `json:"ingress_vip,omitempty"`

	// The virtual IPs used for cluster ingress traffic.
	IngressVips []*IngressVip `json:"ingress_vips"`

	// Machine networks that are associated with this cluster.
	MachineNetworks []*MachineNetwork `json:"machine_networks"`

	// Name of the OpenShift cluster.
	// Required: true
	// Max Length: 54
	// Min Length: 1
	Name *string `json:"name"`

	// The desired network type used.
	// Enum: [OpenShiftSDN OVNKubernetes]
	NetworkType *string `json:"network_type,omitempty"`

	// An "*" or a comma-separated list of destination domain names, domains, IP addresses, or other network CIDRs to exclude from proxying.
	NoProxy *string `json:"no_proxy,omitempty"`

	// OpenShift release image URI.
	OcpReleaseImage string `json:"ocp_release_image,omitempty"`

	// List of OLM operators to be installed.
	OlmOperators []*OperatorCreateParams `json:"olm_operators"`

	// Version of the OpenShift cluster.
	// Required: true
	OpenshiftVersion *string `json:"openshift_version"`

	// platform
	Platform *Platform `json:"platform,omitempty" gorm:"embedded;embeddedPrefix:platform_"`

	// The pull secret obtained from Red Hat OpenShift Cluster Manager at console.redhat.com/openshift/install/pull-secret.
	// Required: true
	PullSecret *string `json:"pull_secret"`

	// Schedule workloads on masters
	SchedulableMasters *bool `json:"schedulable_masters,omitempty"`

	// The IP address pool to use for service IP addresses. You can enter only one IP address pool. If you need to access the services from an external network, configure load balancers and routers to manage the traffic.
	// Pattern: ^(?:(?:(?:[0-9]{1,3}\.){3}[0-9]{1,3}\/(?:(?:[0-9])|(?:[1-2][0-9])|(?:3[0-2])))|(?:(?:[0-9a-fA-F]*:[0-9a-fA-F]*){2,})/(?:(?:[0-9])|(?:[1-9][0-9])|(?:1[0-1][0-9])|(?:12[0-8])))$
	ServiceNetworkCidr *string `json:"service_network_cidr,omitempty"`

	// Service networks that are associated with this cluster.
	ServiceNetworks []*ServiceNetwork `json:"service_networks"`

	// SSH public key for debugging OpenShift nodes.
	SSHPublicKey string `json:"ssh_public_key,omitempty"`

	// A comma-separated list of tags that are associated to the cluster.
	Tags *string `json:"tags,omitempty"`

	// Indicate if the networking is managed by the user.
	UserManagedNetworking *bool `json:"user_managed_networking,omitempty"`

	// Indicate if virtual IP DHCP allocation mode is enabled.
	VipDhcpAllocation *bool `json:"vip_dhcp_allocation,omitempty"`
}

// Validate validates this cluster create params
func (m *ClusterCreateParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPIVip(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAPIVips(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterNetworkCidr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterNetworkHostPrefix(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterNetworks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDiskEncryption(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHighAvailabilityMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHyperthreading(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIgnitionEndpoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIngressVip(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIngressVips(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMachineNetworks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOlmOperators(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOpenshiftVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatform(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePullSecret(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceNetworkCidr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceNetworks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterCreateParams) validateAPIVip(formats strfmt.Registry) error {
	if swag.IsZero(m.APIVip) { // not required
		return nil
	}

	if err := validate.Pattern("api_vip", "body", m.APIVip, `^(?:(?:(?:[0-9]{1,3}\.){3}[0-9]{1,3})|(?:(?:[0-9a-fA-F]*:[0-9a-fA-F]*){2,}))?$`); err != nil {
		return err
	}

	return nil
}

func (m *ClusterCreateParams) validateAPIVips(formats strfmt.Registry) error {
	if swag.IsZero(m.APIVips) { // not required
		return nil
	}

	for i := 0; i < len(m.APIVips); i++ {
		if swag.IsZero(m.APIVips[i]) { // not required
			continue
		}

		if m.APIVips[i] != nil {
			if err := m.APIVips[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("api_vips" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("api_vips" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterCreateParams) validateClusterNetworkCidr(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterNetworkCidr) { // not required
		return nil
	}

	if err := validate.Pattern("cluster_network_cidr", "body", *m.ClusterNetworkCidr, `^(?:(?:(?:[0-9]{1,3}\.){3}[0-9]{1,3}\/(?:(?:[0-9])|(?:[1-2][0-9])|(?:3[0-2])))|(?:(?:[0-9a-fA-F]*:[0-9a-fA-F]*){2,})/(?:(?:[0-9])|(?:[1-9][0-9])|(?:1[0-1][0-9])|(?:12[0-8])))$`); err != nil {
		return err
	}

	return nil
}

func (m *ClusterCreateParams) validateClusterNetworkHostPrefix(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterNetworkHostPrefix) { // not required
		return nil
	}

	if err := validate.MinimumInt("cluster_network_host_prefix", "body", m.ClusterNetworkHostPrefix, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("cluster_network_host_prefix", "body", m.ClusterNetworkHostPrefix, 128, false); err != nil {
		return err
	}

	return nil
}

func (m *ClusterCreateParams) validateClusterNetworks(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterNetworks) { // not required
		return nil
	}

	for i := 0; i < len(m.ClusterNetworks); i++ {
		if swag.IsZero(m.ClusterNetworks[i]) { // not required
			continue
		}

		if m.ClusterNetworks[i] != nil {
			if err := m.ClusterNetworks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cluster_networks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cluster_networks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterCreateParams) validateDiskEncryption(formats strfmt.Registry) error {
	if swag.IsZero(m.DiskEncryption) { // not required
		return nil
	}

	if m.DiskEncryption != nil {
		if err := m.DiskEncryption.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disk_encryption")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("disk_encryption")
			}
			return err
		}
	}

	return nil
}

var clusterCreateParamsTypeHighAvailabilityModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Full","None"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterCreateParamsTypeHighAvailabilityModePropEnum = append(clusterCreateParamsTypeHighAvailabilityModePropEnum, v)
	}
}

const (

	// ClusterCreateParamsHighAvailabilityModeFull captures enum value "Full"
	ClusterCreateParamsHighAvailabilityModeFull string = "Full"

	// ClusterCreateParamsHighAvailabilityModeNone captures enum value "None"
	ClusterCreateParamsHighAvailabilityModeNone string = "None"
)

// prop value enum
func (m *ClusterCreateParams) validateHighAvailabilityModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, clusterCreateParamsTypeHighAvailabilityModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ClusterCreateParams) validateHighAvailabilityMode(formats strfmt.Registry) error {
	if swag.IsZero(m.HighAvailabilityMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateHighAvailabilityModeEnum("high_availability_mode", "body", *m.HighAvailabilityMode); err != nil {
		return err
	}

	return nil
}

var clusterCreateParamsTypeHyperthreadingPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["masters","workers","none","all"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterCreateParamsTypeHyperthreadingPropEnum = append(clusterCreateParamsTypeHyperthreadingPropEnum, v)
	}
}

const (

	// ClusterCreateParamsHyperthreadingMasters captures enum value "masters"
	ClusterCreateParamsHyperthreadingMasters string = "masters"

	// ClusterCreateParamsHyperthreadingWorkers captures enum value "workers"
	ClusterCreateParamsHyperthreadingWorkers string = "workers"

	// ClusterCreateParamsHyperthreadingNone captures enum value "none"
	ClusterCreateParamsHyperthreadingNone string = "none"

	// ClusterCreateParamsHyperthreadingAll captures enum value "all"
	ClusterCreateParamsHyperthreadingAll string = "all"
)

// prop value enum
func (m *ClusterCreateParams) validateHyperthreadingEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, clusterCreateParamsTypeHyperthreadingPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ClusterCreateParams) validateHyperthreading(formats strfmt.Registry) error {
	if swag.IsZero(m.Hyperthreading) { // not required
		return nil
	}

	// value enum
	if err := m.validateHyperthreadingEnum("hyperthreading", "body", *m.Hyperthreading); err != nil {
		return err
	}

	return nil
}

func (m *ClusterCreateParams) validateIgnitionEndpoint(formats strfmt.Registry) error {
	if swag.IsZero(m.IgnitionEndpoint) { // not required
		return nil
	}

	if m.IgnitionEndpoint != nil {
		if err := m.IgnitionEndpoint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ignition_endpoint")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ignition_endpoint")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterCreateParams) validateIngressVip(formats strfmt.Registry) error {
	if swag.IsZero(m.IngressVip) { // not required
		return nil
	}

	if err := validate.Pattern("ingress_vip", "body", m.IngressVip, `^(?:(?:(?:[0-9]{1,3}\.){3}[0-9]{1,3})|(?:(?:[0-9a-fA-F]*:[0-9a-fA-F]*){2,}))$`); err != nil {
		return err
	}

	return nil
}

func (m *ClusterCreateParams) validateIngressVips(formats strfmt.Registry) error {
	if swag.IsZero(m.IngressVips) { // not required
		return nil
	}

	for i := 0; i < len(m.IngressVips); i++ {
		if swag.IsZero(m.IngressVips[i]) { // not required
			continue
		}

		if m.IngressVips[i] != nil {
			if err := m.IngressVips[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ingress_vips" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ingress_vips" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterCreateParams) validateMachineNetworks(formats strfmt.Registry) error {
	if swag.IsZero(m.MachineNetworks) { // not required
		return nil
	}

	for i := 0; i < len(m.MachineNetworks); i++ {
		if swag.IsZero(m.MachineNetworks[i]) { // not required
			continue
		}

		if m.MachineNetworks[i] != nil {
			if err := m.MachineNetworks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("machine_networks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("machine_networks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterCreateParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 54); err != nil {
		return err
	}

	return nil
}

var clusterCreateParamsTypeNetworkTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["OpenShiftSDN","OVNKubernetes"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterCreateParamsTypeNetworkTypePropEnum = append(clusterCreateParamsTypeNetworkTypePropEnum, v)
	}
}

const (

	// ClusterCreateParamsNetworkTypeOpenShiftSDN captures enum value "OpenShiftSDN"
	ClusterCreateParamsNetworkTypeOpenShiftSDN string = "OpenShiftSDN"

	// ClusterCreateParamsNetworkTypeOVNKubernetes captures enum value "OVNKubernetes"
	ClusterCreateParamsNetworkTypeOVNKubernetes string = "OVNKubernetes"
)

// prop value enum
func (m *ClusterCreateParams) validateNetworkTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, clusterCreateParamsTypeNetworkTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ClusterCreateParams) validateNetworkType(formats strfmt.Registry) error {
	if swag.IsZero(m.NetworkType) { // not required
		return nil
	}

	// value enum
	if err := m.validateNetworkTypeEnum("network_type", "body", *m.NetworkType); err != nil {
		return err
	}

	return nil
}

func (m *ClusterCreateParams) validateOlmOperators(formats strfmt.Registry) error {
	if swag.IsZero(m.OlmOperators) { // not required
		return nil
	}

	for i := 0; i < len(m.OlmOperators); i++ {
		if swag.IsZero(m.OlmOperators[i]) { // not required
			continue
		}

		if m.OlmOperators[i] != nil {
			if err := m.OlmOperators[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("olm_operators" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("olm_operators" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterCreateParams) validateOpenshiftVersion(formats strfmt.Registry) error {

	if err := validate.Required("openshift_version", "body", m.OpenshiftVersion); err != nil {
		return err
	}

	return nil
}

func (m *ClusterCreateParams) validatePlatform(formats strfmt.Registry) error {
	if swag.IsZero(m.Platform) { // not required
		return nil
	}

	if m.Platform != nil {
		if err := m.Platform.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("platform")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("platform")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterCreateParams) validatePullSecret(formats strfmt.Registry) error {

	if err := validate.Required("pull_secret", "body", m.PullSecret); err != nil {
		return err
	}

	return nil
}

func (m *ClusterCreateParams) validateServiceNetworkCidr(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceNetworkCidr) { // not required
		return nil
	}

	if err := validate.Pattern("service_network_cidr", "body", *m.ServiceNetworkCidr, `^(?:(?:(?:[0-9]{1,3}\.){3}[0-9]{1,3}\/(?:(?:[0-9])|(?:[1-2][0-9])|(?:3[0-2])))|(?:(?:[0-9a-fA-F]*:[0-9a-fA-F]*){2,})/(?:(?:[0-9])|(?:[1-9][0-9])|(?:1[0-1][0-9])|(?:12[0-8])))$`); err != nil {
		return err
	}

	return nil
}

func (m *ClusterCreateParams) validateServiceNetworks(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceNetworks) { // not required
		return nil
	}

	for i := 0; i < len(m.ServiceNetworks); i++ {
		if swag.IsZero(m.ServiceNetworks[i]) { // not required
			continue
		}

		if m.ServiceNetworks[i] != nil {
			if err := m.ServiceNetworks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("service_networks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("service_networks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this cluster create params based on the context it is used
func (m *ClusterCreateParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAPIVips(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClusterNetworks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDiskEncryption(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIgnitionEndpoint(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIngressVips(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMachineNetworks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOlmOperators(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePlatform(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServiceNetworks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterCreateParams) contextValidateAPIVips(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.APIVips); i++ {

		if m.APIVips[i] != nil {
			if err := m.APIVips[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("api_vips" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("api_vips" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterCreateParams) contextValidateClusterNetworks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ClusterNetworks); i++ {

		if m.ClusterNetworks[i] != nil {
			if err := m.ClusterNetworks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cluster_networks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cluster_networks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterCreateParams) contextValidateDiskEncryption(ctx context.Context, formats strfmt.Registry) error {

	if m.DiskEncryption != nil {
		if err := m.DiskEncryption.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disk_encryption")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("disk_encryption")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterCreateParams) contextValidateIgnitionEndpoint(ctx context.Context, formats strfmt.Registry) error {

	if m.IgnitionEndpoint != nil {
		if err := m.IgnitionEndpoint.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ignition_endpoint")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ignition_endpoint")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterCreateParams) contextValidateIngressVips(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IngressVips); i++ {

		if m.IngressVips[i] != nil {
			if err := m.IngressVips[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ingress_vips" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ingress_vips" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterCreateParams) contextValidateMachineNetworks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MachineNetworks); i++ {

		if m.MachineNetworks[i] != nil {
			if err := m.MachineNetworks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("machine_networks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("machine_networks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterCreateParams) contextValidateOlmOperators(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OlmOperators); i++ {

		if m.OlmOperators[i] != nil {
			if err := m.OlmOperators[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("olm_operators" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("olm_operators" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterCreateParams) contextValidatePlatform(ctx context.Context, formats strfmt.Registry) error {

	if m.Platform != nil {
		if err := m.Platform.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("platform")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("platform")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterCreateParams) contextValidateServiceNetworks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ServiceNetworks); i++ {

		if m.ServiceNetworks[i] != nil {
			if err := m.ServiceNetworks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("service_networks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("service_networks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterCreateParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterCreateParams) UnmarshalBinary(b []byte) error {
	var res ClusterCreateParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
