package models

import (
	"github.com/astaxie/beego/orm"
)

type CapabilityMap struct {
	Id           int64
	CapabilityId int64
	Capability   string
	Class        string
	SubClass     string
	Description  string
	ExpertPool   string
}

type Capabilities struct {
	Id           int64
	CapabilityId int64
	UserId       int64
	UserName     string
	Level        string
}

func AddCapabilityMap() {
	AddCapability("业务基础", "基础知识", 1, "汇编", "", "")
	AddCapability("业务基础", "基础知识", 2, "数据结构", "", "")
	AddCapability("业务基础", "基础知识", 3, "makefile", "", "")
	AddCapability("业务基础", "基础知识", 4, "CPU体系架构", "", "")
	AddCapability("业务基础", "基础知识", 5, "linux设备驱动", "", "")
	AddCapability("业务基础", "基础知识", 6, "linux进程线程", "", "")
	AddCapability("业务基础", "基础知识", 7, "linux并发控制", "", "")
	AddCapability("业务基础", "基础知识", 8, "linux中断时钟", "", "")
	AddCapability("业务基础", "基础知识", 9, "linux裁剪与打包", "", "")
	AddCapability("业务基础", "基础知识", 10, "表项管理", "", "")
	AddCapability("业务基础", "基础知识", 11, "V5和V8架构", "", "")
	AddCapability("业务基础", "基础知识", 12, "设备管理", "", "")
	AddCapability("业务基础", "基础知识", 13, "启动流程", "", "")
	AddCapability("业务基础", "基础知识", 14, "C", "", "")
	AddCapability("业务基础", "基础知识", 15, "python", "", "")
	AddCapability("业务基础", "基础知识", 16, "java/spring/osgi", "", "")

	AddCapability("业务基础", "基础协议", 17, "LLDP", "", "")
	AddCapability("业务基础", "基础协议", 18, "STP", "", "")
	AddCapability("业务基础", "基础协议", 19, "ARP", "", "")
	AddCapability("业务基础", "基础协议", 20, "ISIS", "", "")
	AddCapability("业务基础", "基础协议", 21, "路由协议基础", "", "")
	AddCapability("业务基础", "基础协议", 22, "OSPF", "", "")
	AddCapability("业务基础", "基础协议", 23, "RIP", "", "")
	AddCapability("业务基础", "基础协议", 24, "BGP", "", "")

	AddCapability("业务基础", "基础业务", 25, "VLAN", "", "")
	AddCapability("业务基础", "基础业务", 26, "MAC管理", "", "")
	AddCapability("业务基础", "基础业务", 27, "QinQ", "", "")
	AddCapability("业务基础", "基础业务", 28, "堆叠", "", "")
	AddCapability("业务基础", "基础业务", 29, "隧道管理", "", "")
	AddCapability("业务基础", "基础业务", 30, "IPV4基础", "", "")
	AddCapability("业务基础", "基础业务", 31, "IPV6基础", "", "")
	AddCapability("业务基础", "基础业务", 32, "MPLS基础", "", "")
	AddCapability("业务基础", "基础业务", 33, "VPLS", "", "")

	AddCapability("业务基础", "QoE", 34, "ACL", "", "")
	AddCapability("业务基础", "QoE", 35, "Qos", "", "")
	AddCapability("业务基础", "QoE", 36, "MQC", "", "")
	AddCapability("业务基础", "QoE", 37, "合法监听", "", "")
	AddCapability("业务基础", "QoE", 38, "端口镜像", "", "")

	AddCapability("业务基础", "业务应用", 39, "NAT/FW", "", "")
	AddCapability("业务基础", "业务应用", 40, "防火墙", "", "")
	AddCapability("业务基础", "业务应用", 41, "IPSec VPN", "", "")
	AddCapability("业务基础", "业务应用", 42, "SSL VPN", "", "")
	AddCapability("业务基础", "业务应用", 43, "DHCP", "", "")
	AddCapability("业务基础", "业务应用", 44, "DNS", "", "")
	AddCapability("业务基础", "业务应用", 45, "IDS/IPS", "", "")
	AddCapability("业务基础", "业务应用", 46, "Cache优化", "", "")
	AddCapability("业务基础", "业务应用", 47, "TCP优化", "", "")
	AddCapability("业务基础", "业务应用", 48, "报文压缩", "", "")
	AddCapability("业务基础", "业务应用", 49, "负载分担", "", "")

	AddCapability("专项技术", "系统管理", 50, "启动流程", "", "")
	AddCapability("专项技术", "系统管理", 51, "NTP", "", "")
	AddCapability("专项技术", "系统管理", 52, "进程管理", "", "")
	AddCapability("专项技术", "系统管理", 53, "asn", "", "")
	AddCapability("专项技术", "系统管理", 54, "linux诊断", "", "")
	AddCapability("专项技术", "系统管理", 55, "组件管理", "", "")

	AddCapability("专项技术", "应用管理", 56, "openstack", "", "")
	AddCapability("专项技术", "应用管理", 57, "opendaylight", "", "")
	AddCapability("专项技术", "应用管理", 58, "nova", "", "")
	AddCapability("专项技术", "应用管理", 59, "ceilometer", "", "")

	AddCapability("专项技术", "配置管理", 60, "linux pam", "", "")
	AddCapability("专项技术", "配置管理", 61, "netconf", "", "")
	AddCapability("专项技术", "配置管理", 62, "restconf", "", "")
	AddCapability("专项技术", "配置管理", 63, "snmp", "", "")
	AddCapability("专项技术", "配置管理", 64, "shell", "", "")
	AddCapability("专项技术", "配置管理", 65, "json", "", "")
	AddCapability("专项技术", "配置管理", 66, "key-value database", "", "")
	AddCapability("专项技术", "配置管理", 67, "database adapter", "", "")
	AddCapability("专项技术", "配置管理", 68, "broker/subscriber", "", "")
	AddCapability("专项技术", "配置管理", 69, "notification", "", "")

	AddCapability("专项技术", "设备管理", 70, "电子标签", "", "")
	AddCapability("专项技术", "设备管理", 71, "License", "", "")
	AddCapability("专项技术", "设备管理", 72, "告警管理", "", "")
	AddCapability("专项技术", "设备管理", 73, "监控管理", "", "")
	AddCapability("专项技术", "设备管理", 74, "虚拟化管理", "", "")
	AddCapability("专项技术", "设备管理", 75, "HA", "", "")
	AddCapability("专项技术", "设备管理", 76, "能力管理", "", "")
	AddCapability("专项技术", "设备管理", 77, "时钟管理", "", "")
	AddCapability("专项技术", "设备管理", 78, "整机硬件管理", "", "")
	AddCapability("专项技术", "设备管理", 79, "固件管理", "", "")

	AddCapability("专项技术", "协议", 80, "Nginx", "", "")
	AddCapability("专项技术", "协议", 81, "Squid", "", "")
	AddCapability("专项技术", "协议", 82, "dnsmasq", "", "")
	AddCapability("专项技术", "协议", 83, "ISC DHCP", "", "")
	AddCapability("专项技术", "协议", 84, "OpenSSH", "", "")
	AddCapability("专项技术", "协议", 85, "OpenSSL", "", "")
	AddCapability("专项技术", "协议", 86, "NGE", "", "")
	AddCapability("专项技术", "协议", 87, "brctl", "", "")
	AddCapability("专项技术", "协议", 88, "nftables", "", "")
	AddCapability("专项技术", "协议", 89, "ACL", "", "")
	AddCapability("专项技术", "协议", 90, "ACL-Adp", "", "")
	AddCapability("专项技术", "协议", 91, "AAA", "", "")
	AddCapability("专项技术", "协议", 92, "DAA", "", "")
	AddCapability("专项技术", "协议", 93, "NAC", "", "")
	AddCapability("专项技术", "协议", 94, "UCM", "", "")
	AddCapability("专项技术", "协议", 95, "ethtool", "", "")
	AddCapability("专项技术", "协议", 96, "WIFIdog", "", "")
	AddCapability("专项技术", "协议", 97, "rp-pppoe", "", "")
	AddCapability("专项技术", "协议", 98, "pppd", "", "")
	AddCapability("专项技术", "协议", 99, "hostapd", "", "")
	AddCapability("专项技术", "协议", 100, "iproute2", "", "")
	AddCapability("专项技术", "协议", 101, "iw", "", "")
	AddCapability("专项技术", "协议", 102, "GPS", "", "")
	AddCapability("专项技术", "协议", 103, "PPPoE", "", "")
	AddCapability("专项技术", "协议", 104, "PPP", "", "")
	AddCapability("专项技术", "协议", 105, "ARP", "", "")
	AddCapability("专项技术", "协议", 106, "VLAN", "", "")
	AddCapability("专项技术", "协议", 107, "bridge", "", "")
	AddCapability("专项技术", "协议", 108, "fib", "", "")
	AddCapability("专项技术", "协议", 109, "WPA_Supplicant", "", "")

	AddCapability("专项技术", "基础OS", 110, "netlink", "", "")
	AddCapability("专项技术", "基础OS", 111, "ioctl", "", "")
	AddCapability("专项技术", "基础OS", 112, "overlayfs", "", "")
	AddCapability("专项技术", "基础OS", 113, "nftables", "", "")
	AddCapability("专项技术", "基础OS", 114, "socket", "", "")
	AddCapability("专项技术", "基础OS", 115, "netmap", "", "")
	AddCapability("专项技术", "基础OS", 116, "dbus", "", "")
	AddCapability("专项技术", "基础OS", 117, "aufs", "", "")
	AddCapability("专项技术", "基础OS", 118, "工具链", "", "")
	AddCapability("专项技术", "基础OS", 119, "应用仓库", "", "")
	AddCapability("专项技术", "基础OS", 120, "文件系统", "", "")

	AddCapability("专项技术", "虚拟化", 121, "libvirt", "", "")
	AddCapability("专项技术", "虚拟化", 122, "lxc", "", "")
	AddCapability("专项技术", "虚拟化", 123, "namespace", "", "")
	AddCapability("专项技术", "虚拟化", 124, "cgroup", "", "")
	AddCapability("专项技术", "虚拟化", 125, "docker", "", "")
	AddCapability("专项技术", "虚拟化", 126, "kvm", "", "")
	AddCapability("专项技术", "虚拟化", 127, "qemu", "", "")
	AddCapability("专项技术", "虚拟化", 128, "iov", "", "")
	AddCapability("专项技术", "虚拟化", 129, "gpu", "", "")
	AddCapability("专项技术", "虚拟化", 130, "包管理（apt/dpkg）", "", "")
	AddCapability("专项技术", "虚拟化", 131, "管理接口", "", "")

}

func AddCapability(class string, subclass string, capabilityid int64, capability string, description string, expertpool string) string {
	o := orm.NewOrm()

	//验证合法性
	qs := o.QueryTable("capability_map")

	var c CapabilityMap

	err := qs.Filter("Class", class).Filter("SubClass", subclass).Filter("Capability", capability).One(&c)
	if err != orm.ErrNoRows {
		return "exist such capability"
	}

	err = qs.Filter("CapabilityId", capabilityid).One(&c)
	if err != orm.ErrNoRows {
		return "exist such capabilityid"
	}

	ca := new(CapabilityMap)
	ca.CapabilityId = capabilityid
	ca.Capability = capability
	ca.Class = class
	ca.SubClass = subclass
	ca.Description = description
	ca.ExpertPool = expertpool
	o.Insert(ca)
	return "success"
}

func AddCapabilitiesData() {
	AddCapabilityData(1, 1, "黄业钦", "了解")
	AddCapabilityData(2, 1, "黄业钦", "了解")
	AddCapabilityData(3, 1, "黄业钦", "了解")
	AddCapabilityData(4, 1, "黄业钦", "了解")
	AddCapabilityData(5, 1, "黄业钦", "了解")
	AddCapabilityData(6, 1, "黄业钦", "了解")
	AddCapabilityData(7, 1, "黄业钦", "了解")
	AddCapabilityData(8, 1, "黄业钦", "了解")

}

// add user in capabilities
func AddCapabilityData(capabilityid int64, userid int64, username string, level string) string {
	o := orm.NewOrm()

	//验证合法性
	qs := o.QueryTable("capabilities")

	var c Capabilities

	err := qs.Filter("UserId", userid).Filter("UserName", username).Filter("CapabilityId", capabilityid).One(&c)
	if err != orm.ErrNoRows {
		return "exist such userid"
	}

	ca := new(Capabilities)
	ca.CapabilityId = capabilityid
	ca.Level = level
	ca.UserId = userid
	ca.UserName = username
	o.Insert(ca)
	return "success"
}
