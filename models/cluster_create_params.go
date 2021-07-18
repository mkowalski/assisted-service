// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
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

	// Base domain of the cluster. All DNS records must be sub-domains of this base and include the cluster name.
	BaseDNSDomain string `json:"base_dns_domain,omitempty"`

	// IP address block from which Pod IPs are allocated. This block must not overlap with existing physical networks. These IP addresses are used for the Pod network, and if you need to access the Pods from an external network, configure load balancers and routers to manage the traffic.
	// Pattern: ^(?:(?:(?:[0-9]{1,3}\.){3}[0-9]{1,3}\/(?:(?:[0-9])|(?:[1-2][0-9])|(?:3[0-2])))|(?:(?:[0-9a-fA-F]*:[0-9a-fA-F]*){2,})/(?:(?:[0-9])|(?:[1-9][0-9])|(?:1[0-1][0-9])|(?:12[0-8])))$
	ClusterNetworkCidr *string `json:"cluster_network_cidr,omitempty"`

	// The subnet prefix length to assign to each individual node. For example, if clusterNetworkHostPrefix is set to 23, then each node is assigned a /23 subnet out of the given cidr (clusterNetworkCIDR), which allows for 510 (2^(32 - 23) - 2) pod IPs addresses. If you are required to provide access to nodes from an external network, configure load balancers and routers to manage the traffic.
	// Maximum: 128
	// Minimum: 1
	ClusterNetworkHostPrefix int64 `json:"cluster_network_host_prefix,omitempty"`

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

	// The virtual IP used for cluster ingress traffic.
	// Pattern: ^(?:(?:(?:[0-9]{1,3}\.){3}[0-9]{1,3})|(?:(?:[0-9a-fA-F]*:[0-9a-fA-F]*){2,}))$
	IngressVip string `json:"ingress_vip,omitempty"`

	// Name of the OpenShift cluster.
	// Required: true
	// Max Length: 54
	// Min Length: 1
	Name *string `json:"name"`

	// The desired network type used.
	// Enum: [OpenShiftSDN OVNKubernetes auto-assign]
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

	// The pull secret obtained from Red Hat OpenShift Cluster Manager at cloud.redhat.com/openshift/install/pull-secret.
	// Required: true
	PullSecret *string `json:"pull_secret"`

	// The IP address pool to use for service IP addresses. You can enter only one IP address pool. If you need to access the services from an external network, configure load balancers and routers to manage the traffic.
	// Pattern: ^(?:(?:(?:[0-9]{1,3}\.){3}[0-9]{1,3}\/(?:(?:[0-9])|(?:[1-2][0-9])|(?:3[0-2])))|(?:(?:[0-9a-fA-F]*:[0-9a-fA-F]*){2,})/(?:(?:[0-9])|(?:[1-9][0-9])|(?:1[0-1][0-9])|(?:12[0-8])))$
	ServiceNetworkCidr *string `json:"service_network_cidr,omitempty"`

	// SSH public key for debugging OpenShift nodes.
	SSHPublicKey string `json:"ssh_public_key,omitempty"`

	// Indicate if the networking is managed by the user.
	UserManagedNetworking *bool `json:"user_managed_networking,omitempty"`

	// Indicate if virtual IP DHCP allocation mode is enabled.
	VipDhcpAllocation *bool `json:"vip_dhcp_allocation,omitempty"`
}

// Validate validates this cluster create params
func (m *ClusterCreateParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterNetworkCidr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterNetworkHostPrefix(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHighAvailabilityMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHyperthreading(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIngressVip(formats); err != nil {
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

	if err := m.validatePullSecret(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceNetworkCidr(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterCreateParams) validateClusterNetworkCidr(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterNetworkCidr) { // not required
		return nil
	}

	if err := validate.Pattern("cluster_network_cidr", "body", string(*m.ClusterNetworkCidr), `^(?:(?:(?:[0-9]{1,3}\.){3}[0-9]{1,3}\/(?:(?:[0-9])|(?:[1-2][0-9])|(?:3[0-2])))|(?:(?:[0-9a-fA-F]*:[0-9a-fA-F]*){2,})/(?:(?:[0-9])|(?:[1-9][0-9])|(?:1[0-1][0-9])|(?:12[0-8])))$`); err != nil {
		return err
	}

	return nil
}

func (m *ClusterCreateParams) validateClusterNetworkHostPrefix(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterNetworkHostPrefix) { // not required
		return nil
	}

	if err := validate.MinimumInt("cluster_network_host_prefix", "body", int64(m.ClusterNetworkHostPrefix), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("cluster_network_host_prefix", "body", int64(m.ClusterNetworkHostPrefix), 128, false); err != nil {
		return err
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

func (m *ClusterCreateParams) validateIngressVip(formats strfmt.Registry) error {

	if swag.IsZero(m.IngressVip) { // not required
		return nil
	}

	if err := validate.Pattern("ingress_vip", "body", string(m.IngressVip), `^(?:(?:(?:[0-9]{1,3}\.){3}[0-9]{1,3})|(?:(?:[0-9a-fA-F]*:[0-9a-fA-F]*){2,}))$`); err != nil {
		return err
	}

	return nil
}

func (m *ClusterCreateParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 54); err != nil {
		return err
	}

	return nil
}

var clusterCreateParamsTypeNetworkTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["OpenShiftSDN","OVNKubernetes","auto-assign"]`), &res); err != nil {
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

	// ClusterCreateParamsNetworkTypeAutoAssign captures enum value "auto-assign"
	ClusterCreateParamsNetworkTypeAutoAssign string = "auto-assign"
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

	if err := validate.Pattern("service_network_cidr", "body", string(*m.ServiceNetworkCidr), `^(?:(?:(?:[0-9]{1,3}\.){3}[0-9]{1,3}\/(?:(?:[0-9])|(?:[1-2][0-9])|(?:3[0-2])))|(?:(?:[0-9a-fA-F]*:[0-9a-fA-F]*){2,})/(?:(?:[0-9])|(?:[1-9][0-9])|(?:1[0-1][0-9])|(?:12[0-8])))$`); err != nil {
		return err
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
