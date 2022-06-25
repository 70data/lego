package ipQuery

import (
	"bufio"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"k8s.io/klog/v2"
)

const (
	lineNum = 192922
)

type ipInfo struct {
	PR  string
	ISP string
}

var ISPArray []int

var IPData []ipInfo

var ipInfoArray [2]string

func LoadIPdata(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		klog.Infoln(err)
	}
	buf := bufio.NewReader(f)
	for i := 0; i < lineNum; i++ {
		line, _ := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		ipBegin := strings.Split(line, "\t")[0]
		ipPR := strings.Split(line, "\t")[3]
		ipISP := strings.Split(line, "\t")[7]
		ipISPInt, _ := strconv.Atoi(ipBegin)
		ISPArray = append(ISPArray, ipISPInt)
		IPData = append(IPData, ipInfo{ipPR, ipISP})
	}
	return nil
}

func ip2Long(ipstr string) (ip uint32) {
	r := `^(\d{1,3})\.(\d{1,3})\.(\d{1,3})\.(\d{1,3})`
	reg, err := regexp.Compile(r)
	if err != nil {
		klog.Infoln(err)
	}
	ips := reg.FindStringSubmatch(ipstr)
	if ips == nil {
		klog.Infoln(err)
	}
	ip1, _ := strconv.Atoi(ips[1])
	ip2, _ := strconv.Atoi(ips[2])
	ip3, _ := strconv.Atoi(ips[3])
	ip4, _ := strconv.Atoi(ips[4])
	if ip1 > 255 || ip2 > 255 || ip3 > 255 || ip4 > 255 {
		return
	}
	ip += uint32(ip1 * 0x1000000)
	ip += uint32(ip2 * 0x10000)
	ip += uint32(ip3 * 0x100)
	ip += uint32(ip4)
	return
}

func binarySearch(ISPArray []int, k int) int {
	left, right, mid := 1, len(ISPArray), 0
	for {
		mid = int(math.Floor(float64((left + right) / 2)))
		if ISPArray[mid] > k {
			right = mid - 1
		} else if ISPArray[mid] < k {
			left = mid + 1
		} else {
			break
		}
		if left > right {
			mid = right
			break
		}
	}
	return mid
}

func ProvinceIsp(ipAddr string) [2]string {
	ipMid := ip2Long(ipAddr)
	ipInfoMid := binarySearch(ISPArray, int(ipMid))
	province := IPData[ipInfoMid].PR
	isp := IPData[ipInfoMid].ISP
	ipInfoArray[0] = province
	ipInfoArray[1] = isp
	return ipInfoArray
}
