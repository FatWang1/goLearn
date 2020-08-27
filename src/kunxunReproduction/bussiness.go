/*
 * Copyright(C) 2020 EASTCOM-BUPT Inc.
 * Author: wangpeng_1@ebupt.com
 * Date: 2020-08-21 11:16:48
 * LastEditTime: 2020-08-21 17:53:20
 * LastEditors: wangpeng_1@ebupt.com
 * Description: 
 */
package kunxunReproduction

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"fmt"
)

type conf struct {
	Disksign string `yaml:"diskSign"`
	Version string `yaml:"version"`

	Destination struct {
		Ip string `yaml:"ip"`
		Tcpport string `yaml:"tcpPort"`
		Postport string `yaml:"postPort"`
	}
}

type communicationModel struct{
	var c conf
	conf := c.getConf()
	ip = conf.Destination.Ip
	tcpPort = conf.Destination.Tcpport
	postPort = conf.Destination.Postport
}

func (c *conf) getConf() *conf {
	yamlFile ,err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarsshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}

func (httpcode *communicationModel)  {
	
}