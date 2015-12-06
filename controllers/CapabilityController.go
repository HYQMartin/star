package controllers

import (
	"fmt"
	// "github.com/astaxie/beego"
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
	BaseController
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

type CapabilityUpdateController struct {
	BaseController
}

func (c *CapabilityUpdateController) Post() {
	all, err := models.ALLCapabilityMap()
	if err != nil {
		c.Ctx.WriteString("failed")
		return
	}
	domain := c.GetString("domain")
	name := c.GetString("name")
	level := c.GetString("level")
	desc := c.GetString("desc")
	capability := c.GetString("capability")
	var capabilityid int64
	for _, v := range all {
		if v.Capability == capability {
			capabilityid = v.CapabilityId
		}
	}
	ret := models.AddCapabilityData(capabilityid, domain, name, level, desc)
	c.Ctx.WriteString(ret)
	return
}

type CapabilityAddController struct {
	BaseController
}

func (c *CapabilityAddController) Post() {
	all, err := models.ALLCapabilityMap()
	if err != nil {
		fmt.Println("all capabilitymaap err", err)
		c.Ctx.WriteString("failed")
		return
	}
	capability := c.GetString("capability")
	bigclass := c.GetString("bigclass")
	subclass := c.GetString("subclass")
	desc := c.GetString("desc")
	pool := c.GetString("pool")
	if capability == "" {
		c.Ctx.WriteString("能力名称不能为空")
		return
	}
	if bigclass == "" {
		c.Ctx.WriteString("大类不能为空")
		return
	}
	if subclass == "" {
		c.Ctx.WriteString("小类不能为空")
		return
	}
	var capabilityid int64
	capabilityid = all[len(all)-1].CapabilityId + 1
	ret := models.AddCapability(bigclass, subclass, capabilityid, capability, desc, pool)
	c.Ctx.WriteString(ret)
	return
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
	// ALLCapabilityFromMysql = all
	alluser, err := models.ALLCapabilities()
	if err != nil {
		return
	}
	// AllUserFromMysql = alluser
	allusermap := make(map[string]string, 0)

	for _, v := range alluser {
		allusermap[v.UserName] = "1"
	}

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
