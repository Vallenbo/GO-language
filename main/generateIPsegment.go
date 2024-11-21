package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type IpSubnetMaskSegment struct {
	//IpSubnetMaskSegmentID int64  `gorm:"column:ip_subnet_mask_segment_id"`
	Segment1         int64  `gorm:"column:segment_1"`
	Segment2         int64  `gorm:"column:segment_2"`
	SubnetMaskLength int64  `gorm:"column:subnet_mask_length"`
	IpSegment        string `gorm:"column:ip_segment"`
	Created          int64  `gorm:"column:created"`
	Updated          int64  `gorm:"column:updated"`
	Status           int    `gorm:"column:status"`
}

func (IpSubnetMaskSegment) TableName() string {
	return "ip_subnet_mask_segment"
}

func generateData() (data []IpSubnetMaskSegment) {
	subnetMask := map[int64]int64{
		16: 256, 17: 128, 18: 64, 19: 32,
		20: 16, 21: 8, 22: 4, 23: 2, 24: 1,
	}
	for ipSeg1 := int64(1); ipSeg1 < 191; ipSeg1++ {
		for ipSeg2 := int64(0); ipSeg2 < 255; ipSeg2++ {
			for mask := int64(16); mask <= 24; mask++ {
				fmt.Printf("IP Segments for %d.%d.x.0 with mask /%d:\n", ipSeg1, ipSeg2, mask)
				for ipSeg3 := int64(0); ipSeg3 < 256; {
					fmt.Printf("%d.%d.%d.0 \n", ipSeg1, ipSeg2, ipSeg3)
					data = append(data, IpSubnetMaskSegment{
						Segment1:         ipSeg1,
						Segment2:         ipSeg2,
						SubnetMaskLength: mask,
						IpSegment:        fmt.Sprintf("%d.%d.%d.0", ipSeg1, ipSeg2, ipSeg3),
						Created:          time.Now().Unix(),
						Updated:          0,
						Status:           1,
					})
					ipSeg3 += subnetMask[mask]
				}
			}
		}
	}
	return data
}

func main() {
	dsn := "root:123456@tcp(10.15.0.202:3306)/qfwa?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	IpSubnetMaskSegmentData := generateData()

	for i := 1; i < len(IpSubnetMaskSegmentData); i++ {
		endIndex := i + 1000
		if endIndex > len(IpSubnetMaskSegmentData) {
			endIndex = len(IpSubnetMaskSegmentData)
		}
		batchData := IpSubnetMaskSegmentData[i:endIndex]
		err = db.Save(&batchData).Error
		if err != nil {
			panic(err)
		}
	}
}

//func main() {
//	fmt.Println("请输入子网掩码（tag : 24）")
//	var mask string
//	_, err := fmt.Scanln(&mask)
//	if err != nil {
//		panic(err)
//	}
//	for ipSeg1 := 1; ipSeg1 < 3; ipSeg1++ {
//		for ipSeg2 := 0; ipSeg2 < 3; ipSeg2++ {
//			ipStr := fmt.Sprintf("%d.%d.x.0/%s", ipSeg1, ipSeg2, mask)
//			ipv4Addr, ipv4Net, err := net.ParseCIDR(ipStr)
//			if err != nil {
//				log.Fatal(err)
//			} else {
//				fmt.Printf("ipv4Addr: %v, ipv4Net: %v\n", ipv4Addr, ipv4Net)
//			}
//		}
//	}
//}
