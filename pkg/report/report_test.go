// Copyright 2020 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package report

import (
	"context"
	"testing"

	"github.com/pingcap-incubator/tiup-cluster/pkg/telemetry"
	"github.com/pingcap/check"
)

type reportSuite struct{}

var _ = check.Suite(&reportSuite{})

func TestT(t *testing.T) { check.TestingT(t) }

func (s *reportSuite) TestNodeInfo(c *check.C) {
	info := new(telemetry.NodeInfo)
	err := telemetry.FillNodeInfo(context.Background(), info)
	c.Assert(err, check.IsNil)

	text, err := NodeInfoToText(info)
	c.Assert(err, check.IsNil)

	info2, err := NodeInfoFromText(text)
	c.Assert(err, check.IsNil)
	c.Assert(info2, check.DeepEquals, info)
}
