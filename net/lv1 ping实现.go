package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"net"
	"time"
)

type ICMP struct {
	Type        uint8
	Code        uint8
	Checksum    uint16
	Identifier  uint16
	SequenceNum uint16
}

var (
	laddr         = net.IPAddr{IP: net.ParseIP("ip")}
	num     int   = 30
	timeout int64 = 1000 // 超时时间，单位为毫秒
	size    int   = 32   // 发送数据包的大小，单位为字节
	stop    bool
)

func main() {
	fmt.Printf("ping ")
	var webUrl string
	_, err := fmt.Scanln(&webUrl)
	if err != nil {
		log.Println(err)
		return
	}
	ip := net.ParseIP(webUrl)
	if ip == nil || ip.To4() == nil {
		ipAddrs, err := net.LookupIP(webUrl)
		if err != nil {
			log.Println(err)
		} else {
			ip = ipAddrs[0]
		}
	}
	icmp := ICMP{}
	//icmp头部填充
	icmp.Type = 8
	icmp.Code = 0
	icmp.Checksum = 0
	icmp.Identifier = 1
	icmp.SequenceNum = 1
	//建立与目标主机的连接，并检查连接是否成功
	conn, err := net.DialTimeout("ip:icmp", webUrl, time.Duration(timeout)*time.Millisecond)
	if err != nil {
		log.Fatal(err)
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}(conn)
	fmt.Printf("\n正在 ping %s 具有 %d 字节的数据:\n", webUrl, size)
	//构造 ICMP 报文的数据字段（date）
	var buffer bytes.Buffer
	err = binary.Write(&buffer, binary.BigEndian, icmp) //按大端模式将icmp写入buffer
	if err != nil {
		log.Println(err)
		return
	}
	data := make([]byte, size)
	buffer.Write(data)
	data = buffer.Bytes()
	var successTimes int
	var failTimes int
	var minTime int = int(math.MaxInt32)
	var maxTime int
	var totalTime int
	for i := 0; i < num; i++ {
		icmp.SequenceNum = uint16(1)
		// 检验和设为0
		data[2] = byte(0)
		data[3] = byte(0)
		data[6] = byte(icmp.SequenceNum >> 8)
		data[7] = byte(icmp.SequenceNum)
		icmp.Checksum = CheckSum(data)
		data[2] = byte(icmp.Checksum >> 8)
		data[3] = byte(icmp.Checksum)
		// 开始时间
		nowTime := time.Now()
		err := conn.SetDeadline(nowTime.Add(time.Duration(timeout) * time.Millisecond))
		if err != nil {
			log.Println(err)
		}
		n, err := conn.Write(data)
		if err != nil {
			log.Fatal(err)
		}
		buf := make([]byte, 65535)
		n, err = conn.Read(buf)
		if err != nil {
			fmt.Println("请求超时。")
			failTimes++
			continue
		}
		et := int(time.Since(nowTime) / 1000000)
		if minTime > et {
			minTime = et
		}
		if maxTime < et {
			maxTime = et
		}
		totalTime += et
		fmt.Printf("来自 %s 的回复: 字节=%d 时间=%dms TTL=%d\n", webUrl, len(buf[28:n]), et, buf[8])
		successTimes++
		if successTimes == 5 {
			break
		}
		time.Sleep(1 * time.Second)
	}
	fmt.Printf("\n%s 的 Ping 统计信息:\n", webUrl)
	fmt.Printf("    数据包: 已发送 = %d，已接收 = %d，丢失 = %d (%.2f%% 丢失)，\n", successTimes+failTimes, successTimes, failTimes, float64(failTimes*100)/float64(successTimes+failTimes))
	if maxTime != 0 && minTime != int(math.MaxInt32) {
		fmt.Printf("往返行程的估计时间(以毫秒为单位):\n")
		fmt.Printf("    最短 = %dms，最长 = %dms，平均 = %dms\n", minTime, maxTime, totalTime/successTimes)
	}
}

func CheckSum(data []byte) uint16 {
	var sum uint32
	var length = len(data)
	var index int
	for length > 1 {
		sum += uint32(data[index])<<8 | uint32(data[index+1])
		index += 2
		length -= 2
	}
	if length == 1 {
		sum += uint32(data[index]) << 8
	}
	for sum>>16 != 0 {
		sum = (sum & 0xFFFF) + (sum >> 16)
	}
	return uint16(^sum)
}
