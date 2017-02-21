// Copyright 2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package types

import (
	"encoding/json"
	"fmt"

	log "github.com/cihub/seelog"
	"github.com/containernetworking/cni/pkg/skel"
	"github.com/containernetworking/cni/pkg/types"
	"github.com/pkg/errors"
)

// NetConf defines the parameters required to configure a contaner's namespace
// with an ENI
type NetConf struct {
	types.NetConf
	ENIID       string `json:"eni"`
	IPV4Address string `json:"ipv4-address"`
}

func NewConf(args *skel.CmdArgs) (*NetConf, error) {
	var conf NetConf
	if err := json.Unmarshal(args.StdinData, &conf); err != nil {
		return nil, errors.Wrapf(err, "Failed to parse config")
	}
	if conf.ENIID == "" {
		return nil, fmt.Errorf("Missing required parameter in config: '%s'", "eni")
	}
	if conf.IPV4Address == "" {
		return nil, fmt.Errorf("Missing required parameter in config: '%s'", "ipv4-address")
	}

	log.Infof("Loaded config: %v", conf)
	return &conf, nil
}
