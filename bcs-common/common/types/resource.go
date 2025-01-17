/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package types

type BcsClusterResource struct {
	DiskTotal float64               `json:"disktotal"`
	MemTotal  float64               `json:"memtotal"`
	CpuTotal  float64               `json:"cputotal"`
	DiskUsed  float64               `json:"diskused"`
	MemUsed   float64               `json:"memused"`
	CpuUsed   float64               `json:"cpuused"`
	Agents    []BcsClusterAgentInfo `json:"agents"`
}

type BcsClusterAgentInfo struct {
	HostName  string  `json:"hostname"`
	IP        string  `json:"ip"`
	DiskTotal float64 `json:"disktotal"`
	MemTotal  float64 `json:"memtotal"`
	CpuTotal  float64 `json:"cputotal"`
	DiskUsed  float64 `json:"diskused"`
	MemUsed   float64 `json:"memused"`
	CpuUsed   float64 `json:"cpuused"`

	Disabled       bool                 `json:"disabled"`
	HostAttributes []*BcsAgentAttribute `json:"host_attributes"`
	Attributes     []*BcsAgentAttribute `json:"attributes"`

	RegisteredTime   int64 `json:"registered_time"`
	ReRegisteredTime int64 `json:"reregistered_time"`
}

type MesosValue_Scalar struct {
	Value float64 `json:"value"`
}

type MesosValue_Ranges struct {
	Begin uint64 `json:"begin"`
	End   uint64 `json:"end"`
}

type MesosValue_Text struct {
	Value string `json:"value"`
}

type MesosValue_Set struct {
	Item []string `json:"item"`
}

type MesosValue_Type int32

const (
	MesosValueType_UNKNOW MesosValue_Type = 0
	MesosValueType_Scalar MesosValue_Type = 101
	MesosValueType_Ranges MesosValue_Type = 102
	MesosValueType_Text   MesosValue_Type = 103
	MesosValueType_Set    MesosValue_Type = 104
)

type BcsClusterAgentSetting struct {
	InnerIP     string                       `json:"innerIP"`
	Disabled    bool                         `json:"disabled"`
	AttrStrings map[string]MesosValue_Text   `json:"strings"`
	AttrScalars map[string]MesosValue_Scalar `json:"scalars"`
}

type BcsClusterAgentSettingUpdate struct {
	IPs         []string           `json:"ips"`
	SettingName string             `json:"name"`
	ValueType   MesosValue_Type    `json:"valuetype"`
	ValueText   *MesosValue_Text   `json:"text"`
	ValueScalar *MesosValue_Scalar `json:"scalar"`
}

type BcsAgentAttribute struct {
	Name   string               `json:"name,omitempty"`
	Type   MesosValue_Type      `json:"type,omitempty"`
	Scalar *MesosValue_Scalar   `json:"scalar,omitempty"`
	Ranges []*MesosValue_Ranges `json:"ranges,omitempty"`
	Set    *MesosValue_Set      `json:"set,omitempty"`
	Text   *MesosValue_Text     `json:"text,omitempty"`
}
