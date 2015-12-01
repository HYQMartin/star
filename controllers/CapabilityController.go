package controllers

import (
	// "fmt"
	"github.com/astaxie/beego"
	"star/models"
	"strconv"
	"time"
)

var ALLCapabilityFromMysql []models.CapabilityMap
var AllUserFromMysql []models.Capabilities

func GetCapabilityFromMysqlPerHour() {
	for {

		all, err := models.ALLCapabilityMap()
		if err == nil {
			ALLCapabilityFromMysql = all
		}
		alluser, err := models.ALLCapabilities()
		if err == nil {
			AllUserFromMysql = alluser
		}
		time.Sleep(time.Hour * 1)
	}
}

type CapabilityController struct {
	beego.Controller
}

type PeopleValue struct {
	Levels map[string][]string
}
type BasicSubClass struct {
	PanelID    string
	Name       string
	Capability []string
	People     map[string]string
}

func (c *CapabilityController) Get() {
	c.TplNames = "main_capability/Capability.html"
	var basicsubclassmap map[string]BasicSubClass
	basicsubclassmap = make(map[string]BasicSubClass, 0)

	var specialsubclass map[string]BasicSubClass
	specialsubclass = make(map[string]BasicSubClass, 0)
	/*all := ALLCapabilityFromMysql
	alluser := AllUserFromMysql*/
	all, err := models.ALLCapabilityMap()
	if err != nil {
		return
	}
	alluser, err := models.ALLCapabilities()
	if err != nil {
		return
	}
	allusermap := make(map[string]string, 0)

	for _, v := range alluser {
		allusermap[v.UserName] = "1"
	}
	/*var basic1 BasicSubClass
	basic1.PanelID = "111"
	basic1.Name = "基础知识"
	basic1.Capability = []string{"汇编", "数据结构", "makefile", "CPU体系架构", "linux设备驱动", "linux进程线程"}
	basic1.People = make(map[string]string)
	basic1.People["宋文健"] = "<td>了解</td><td>了解</td><td>了解</td><td>了解</td><td>了解</td><td>了解</td>"
	basic1.People["刘明"] = "<td>了解</td><td>了解</td><td>了解</td><td>了解</td><td>了解</td><td>了解</td>"

	var basic2 BasicSubClass
	basic2.PanelID = "112"
	basic2.Name = "基础协议"
	basic2.Capability = []string{"LLDP", "TCP", "UDP", "SSH", "SNMP", "DHP"}
	basic2.People = make(map[string]string)
	basic2.People["胡杰"] = "<td>了解</td><td>了解</td><td>了解</td><td>了解</td><td>了解</td><td>了解</td>"
	basic2.People["刘盼"] = "<td>了解</td><td>了解</td><td>了解</td><td>了解</td><td>了解</td><td>了解</td>"

	basicsubclassmap["基础知识"] = basic1
	basicsubclassmap["基础协议"] = basic2*/

	for _, v := range all {
		// fmt.Println(v.Class)
		if v.Class == "业务基础" {
			tmp := basicsubclassmap[v.SubClass]
			tmp.PanelID = "10" + strconv.Itoa((int)(v.CapabilityId))
			tmp.Name = v.SubClass
			tmp.Capability = append(basicsubclassmap[v.SubClass].Capability, v.Capability)
			if tmp.People == nil {
				tmp.People = make(map[string]string)
			}

			for au, _ := range allusermap {
				level := models.GetLevel(v.CapabilityId, au, alluser)
				if err != nil {
					// fmt.Println("GetLevelByCIDAndUser err", err)
				}
				tmpstr := tmp.People[au]
				tmpstr = tmpstr + "<td>" + level + "</td>"
				tmp.People[au] = tmpstr
			}
			basicsubclassmap[v.SubClass] = tmp
			// basicsubclassmap[v.SubClass].People

		}
		if v.Class == "专项技术" {
			tmp := specialsubclass[v.SubClass]
			tmp.PanelID = "10" + strconv.Itoa((int)(v.CapabilityId))
			tmp.Name = v.SubClass
			tmp.Capability = append(specialsubclass[v.SubClass].Capability, v.Capability)
			if tmp.People == nil {
				tmp.People = make(map[string]string)
			}
			for au, _ := range allusermap {
				level := models.GetLevel(v.CapabilityId, au, alluser)
				if err != nil {
					// fmt.Println("GetLevelByCIDAndUser err", err)
				}
				tmpstr := tmp.People[au]
				tmpstr = tmpstr + "<td>" + level + "</td>"
				tmp.People[au] = tmpstr
			}

			specialsubclass[v.SubClass] = tmp
			// basicsubclassmap[v.SubClass].People
		}
	}
	// fmt.Println(basicsubclassmap)
	c.Data["basicsubclass"] = basicsubclassmap
	c.Data["specialsubclass"] = specialsubclass
	return
}
