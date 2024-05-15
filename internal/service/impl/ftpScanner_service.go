// Package impl
// @Description 封装成单个功能，每次点击循环调用
// @Author  Ymri  2024/5/14 20:13
// @Update 2024/5/14 20:13
package impl

import (
	"encoding/json"
	"fmt"
	"github.com/jlaffaye/ftp"
	"os"
	"time"
)

type APPConfig struct {
	Ip       string `json:"ip"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

var Cfg APPConfig

type FtpScannerService struct{}

var ftpScannerService *FtpScannerService

func GetFtpScannerServiceInstance() *FtpScannerService {
	return ftpScannerService
}

// LoadConfigInit 加载当前目录下的配置文件
func (f *FtpScannerService) LoadConfigInit() {
	// 检查当前目录下是否存在config.json
	data, err := os.ReadFile("config.json")
	if err != nil {
		// 自动创建一个新的config.json
		tempStr := "{\"ip\":\"192.0.2.1\",\"port\":\"8080\",\"user\":\"username\",\"password\":\"password\"}"
		_ = os.WriteFile("config.json", []byte(tempStr), 0644)
		data, err = os.ReadFile("config.json")
		if err != nil {
			fmt.Errorf("读取配置文件失败")
			return
		}
	}
	// json解析
	err = json.Unmarshal(data, &Cfg)
	fmt.Println(Cfg)
	if err != nil {
		fmt.Println("解析配置文件失败")
		return
	}
	fmt.Printf("加载配置成功！")

}

// ScanFtpFile 扫描ftp文件
func (f *FtpScannerService) ScanFtpFile(filePath string) []*ftp.Entry {
	// 读取某一文件路径下的所有文件
	//GetApp
	c, err := ftp.Dial(Cfg.Ip+":"+Cfg.Port, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		//log.Fatal(err)
		return nil
	}
	err = c.Login(Cfg.User, Cfg.Password)
	if err != nil {
		//log.Fatal(err)
		return nil
	}
	entries, err := c.List(filePath)
	if err != nil {
		//log.Fatal(err)
		return nil
	}
	return entries
}
