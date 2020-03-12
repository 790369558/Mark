package Mark

import (
	"errors"
	"github.com/imroc/biu"
	"strconv"
)

//分割ip段位以.为分割标志
func SplitPoint(IP string) []string {
	var IPS []string
	j := 0
	for i, s := range IP {
		if s == 46 {
			IPS = append(IPS, IP[j:i])
			j = i + 1
		}
	}
	IPS = append(IPS, IP[j:])
	return IPS
}

//获得掩码位数
func GetMarkNum(mark string) (int, error) {
	var num = 0
	var err error
	marks := SplitPoint(mark)
	var isEnd = false
	for _, m := range marks {
		int_m, err := strconv.Atoi(m)
		if err != nil {
			return num, err
		}
		Bin := biu.ToBinaryString(int8(int_m))
		for _, bin := range Bin {
			if bin == 49 {
				if isEnd == true {
					return num, errors.New("mark error:mask format error")
				}
				num++
			} else if bin == 48 {
				isEnd = true
			}
		}
	}
	return num, err
}

//获得IP经过掩码后的网络地址
func GetWebsiteAddress(ip string, mark string) (string, error) {
	var websiteAddress string
	IPS := SplitPoint(ip)
	MarkNum, err := GetMarkNum(mark)
	if err != nil {
		return websiteAddress, err
	}
	Num_res := MarkNum / 8
	Num_rem := MarkNum % 8
	for i, ip := range IPS {
		if i < Num_res {
			websiteAddress = websiteAddress + "." + IPS[i]
		} else if i == Num_res {
			int_ip, err := strconv.Atoi(ip)
			if err != nil {
				return websiteAddress, err
			}
			Bin := biu.ToBinaryString(int8(int_ip))
			var CBin string
			for j := 0; j < 8; j++ {
				if j < Num_rem {
					CBin += string(Bin[j])
				}
				if j >= Num_rem {
					CBin += "0"
				}
			}
			websiteAddress = websiteAddress + "." + strconv.Itoa(BinToInt(CBin))
		} else {
			websiteAddress = websiteAddress + ".0"
		}
	}
	return websiteAddress[1:], err
}

//二进制字符串转换int类型
func BinToInt(Bin string) int {
	var result = 0
	Bit_Num := len(Bin) - 1
	N := 1
	for i := 0; i <= Bit_Num; i++ {
		if Bin[Bit_Num-i] == 49 {
			result += N
		}
		N *= 2
	}
	return result
}
