// Copyright 2020 The nfgo Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package napollo

import (
	"context"
	"log"

	"github.com/philchia/agollo/v4"
	"nfgo.ga/nfgo/nconf"
	"nfgo.ga/nfgo/nlog"
)

// Apollo -
type Apollo struct {
	conf   *Conf
	client agollo.Client
}

// NewApollo -
func NewApollo(confFile string) (*Apollo, error) {
	conf, err := newConf(confFile)
	if err != nil {
		return nil, err
	}
	logger := nlog.Logger(context.Background())
	opts := agollo.WithLogger(logger)
	client := agollo.NewClient(&agollo.Conf{
		AppID:              conf.AppID,
		NameSpaceNames:     []string{conf.Namespace},
		Cluster:            conf.Cluster,
		CacheDir:           conf.CacheDir,
		MetaAddr:           conf.MetaAddr,
		AccesskeySecret:    conf.AccessKeySecret,
		InsecureSkipVerify: conf.InsecureSkipVerify,
	}, opts)
	if err := client.Start(); err != nil {
		logger.WithError(err).Warn("apollo client start with error.")
	}
	return &Apollo{
		conf:   conf,
		client: client,
	}, nil
}

// MustNewApollo -
func MustNewApollo(confFile string) *Apollo {
	apollo, err := NewApollo(confFile)
	if err != nil {
		log.Fatal("fail to new apolllo: ", err)
	}
	return apollo
}

// OnUpdate -
func (a *Apollo) OnUpdate(handler func(newContent, oldContent string)) {
	a.client.OnUpdate(func(event *agollo.ChangeEvent) {
		change := event.Changes["content"]
		handler(change.NewValue, change.OldValue)
	})
}

// MustNewConfig -
func (a *Apollo) MustNewConfig() *nconf.Config {
	return a.MustNewConfigCustom(nil)
}

// MustNewConfigCustom  -
func (a *Apollo) MustNewConfigCustom(customConfig interface{ SetConfig(config *nconf.Config) }) *nconf.Config {
	content := a.client.GetContent(agollo.WithNamespace(a.conf.Namespace))
	return nconf.MustNewConfigCustom([]byte(content), customConfig)
}
