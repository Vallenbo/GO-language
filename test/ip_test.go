package test

import (
	"fmt"
	"log"
	"net"
	"testing"
)

// 验证当前网段是否处于指定子网内
// 查看特定网段下不同子网掩码时的 网段划分情况
func Test_ip(t *testing.T) {
	// 初始化子网掩码映射
	subnetMask := map[int]int{
		16: 256, 17: 128, 18: 64, 19: 32,
		20: 16, 21: 8, 22: 4, 23: 2, 24: 1,
	}

	// 创建嵌套的映射
	ipSegments := make(map[string]map[string]map[string][]string)
	for ipSeg1 := 1; ipSeg1 < 256; ipSeg1++ {
		if ipSegments[fmt.Sprintf("%d", ipSeg1)] == nil {
			ipSegments[fmt.Sprintf("%d", ipSeg1)] = make(map[string]map[string][]string)
		}
		for ipSeg2 := 0; ipSeg2 < 256; ipSeg2++ {
			if ipSegments[fmt.Sprintf("%d", ipSeg1)][fmt.Sprintf("%d", ipSeg2)] == nil {
				ipSegments[fmt.Sprintf("%d", ipSeg1)][fmt.Sprintf("%d", ipSeg2)] = make(map[string][]string)
			}
			for mask := 16; mask <= 24; mask++ {
				if ipSegments[fmt.Sprintf("%d", ipSeg1)][fmt.Sprintf("%d", ipSeg2)][fmt.Sprintf("%d", mask)] == nil {
					// 【1】.【0】.【16】
					ipSegments[fmt.Sprintf("%d", ipSeg1)][fmt.Sprintf("%d", ipSeg2)][fmt.Sprintf("%d", mask)] = make([]string, 0)
				}
				// 使用子网掩码计算IP段
				for ipSeg3 := 0; ipSeg3 < 256; ipSeg3 += subnetMask[mask] {
					// 1.0.0.0
					ipSegment := fmt.Sprintf("%d.%d.%d.0", ipSeg1, ipSeg2, ipSeg3)
					// 【1】.【0】.【16】.【1.0.0.0】 --  【1】.【0】.【16】.【1.0.255.0】
					ipSegments[fmt.Sprintf("%d", ipSeg1)][fmt.Sprintf("%d", ipSeg2)][fmt.Sprintf("%d", mask)] = append(
						ipSegments[fmt.Sprintf("%d", ipSeg1)][fmt.Sprintf("%d", ipSeg2)][fmt.Sprintf("%d", mask)],
						ipSegment,
					)
				}
			}
		}
	}

	// 【1】.【0】.【16】.【1.0.0.0】 --  【1】.【0】.【16】.【1.0.255.0】
	// 打印部分结果以验证（可选）
	for ipSeg1 := 1; ipSeg1 < 3; ipSeg1++ { // 仅打印前两个以简化输出
		for ipSeg2 := 0; ipSeg2 < 3; ipSeg2++ { // 仅打印前三个以简化输出
			for mask := 16; mask <= 24; mask++ {
				fmt.Printf("IP Segments for %d.%d.x.0 with mask /%d:\n", ipSeg1, ipSeg2, mask)
				for _, segment := range ipSegments[fmt.Sprintf("%d", ipSeg1)][fmt.Sprintf("%d", ipSeg2)][fmt.Sprintf("%d", mask)] {
					fmt.Println(segment)
				}
			}
		}
	}
}

// Test_IPMaskSegment 按照 Mask 划分指定网段
func Test_IPMaskSegment(t *testing.T) {
	ipStr := "192.168.1.1" // 指定的网段
	//maskStr := "255.255.255.0" // 指定的子网掩码

	// 第一种方式 根据子网掩码 解析IP网段
	ipv4Mask := net.CIDRMask(24, 32) // 指定的子网掩码
	ipSegment := net.ParseIP(ipStr).Mask(ipv4Mask)
	fmt.Println(ipSegment) // 192.168.1.0

	// 第二种方式 计算网络地址
	ipv4Addr, ipv4Net, err := net.ParseCIDR("192.0.2.1/24")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ipv4Addr) // 192.0.2.1
	fmt.Println(ipv4Net)  // 192.0.2.0/24

	// 打印结果
	fmt.Printf("Network Address: %v\n", ipSegment)
	//fmt.Printf("Broadcast Address: %s\n", broadcast)
	//fmt.Printf("First Usable IP: %s\n", firstIP)
	//fmt.Printf("Last Usable IP: %s\n", lastIP)
}
