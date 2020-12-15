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
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	envAppID           = "NF_APOLLO_APP_ID"
	envNamespace       = "NF_APOLLO_NAMESPACE"
	envCluster         = "NF_APOLLO_CLUSTER"
	envCacheDir        = "NF_APOLLO_CACHE_DIR"
	envMetaAdrr        = "NF_APOLLO_META_ADRR"
	envAccessKeySecret = "NF_APOLLO_APOLLO_ACCESSKEY_SECRET"
)

// Conf -
type Conf struct {
	AppID              string `yaml:"appID"`
	Namespace          string `yaml:"namespace"`
	Cluster            string `yaml:"cluster"`
	CacheDir           string `yaml:"cacheDir"`
	MetaAddr           string `yaml:"metaAddr"`
	AccessKeySecret    string `yaml:"accessKeySecret"`
	InsecureSkipVerify bool   `yaml:"insecureSkipVerify"`
}

func newConf(file string) (*Conf, error) {
	defaultConf := newConfFromEnvs()
	if file != "" {
		f, err := os.Open(file)
		defer f.Close()
		if err != nil {
			return nil, err
		}
		data, err := ioutil.ReadAll(f)
		if err != nil {
			return nil, err
		}
		conf := &Conf{}
		if err := yaml.Unmarshal(data, conf); err != nil {
			return nil, err
		}
		conf.setDefault(defaultConf)
		return conf, nil
	}
	return defaultConf, nil
}

func (conf *Conf) setDefault(defaultConf *Conf) {
	if defaultConf == nil {
		return
	}
	if conf.AppID == "" {
		conf.AppID = defaultConf.AppID
	}
	if conf.Namespace == "" {
		conf.Namespace = defaultConf.Namespace
	}
	if conf.Cluster == "" {
		conf.Cluster = defaultConf.Cluster
	}
	if conf.CacheDir == "" {
		conf.CacheDir = defaultConf.CacheDir
	}
	if conf.MetaAddr == "" {
		conf.MetaAddr = defaultConf.MetaAddr
	}
	if conf.AccessKeySecret == "" {
		conf.AccessKeySecret = defaultConf.AccessKeySecret
	}
}

func newConfFromEnvs() *Conf {
	return &Conf{
		AppID:           os.Getenv(envAppID),
		Namespace:       os.Getenv(envNamespace),
		Cluster:         os.Getenv(envCluster),
		CacheDir:        os.Getenv(envCacheDir),
		MetaAddr:        os.Getenv(envMetaAdrr),
		AccessKeySecret: os.Getenv(envAccessKeySecret),
	}
}
