package utils

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
	"os"
)
//关于文件夹的操作

// PathExists 判断文件目录是否存在
func PathExists(path string) (bool,error)  {
	fi ,err := os.Stat(path)	//获取该文件信息，如同ll命令一样
	if err == nil {		//若成功获取信息
		if fi.IsDir(){	//若是目录
			return true,nil
		}
		return false,errors.New("存在同名文件")
	}
	if os.IsNotExist(err){	//最后判断一下是不是因为文件不存在爆的错
		return false,nil
	}
	return false,err
}

// CreateDir 批量创建文件夹
func CreateDir(dirs ...string)  (err error){
	for _,v := range dirs{
		exist, err := PathExists(v)	//查看文件目录是否存在
		if err != nil {
			return err
		}
		if !exist{	//若不存在
			global.G_LOG.Debug("create directory" + v)
			if err := os.MkdirAll(v,os.ModePerm);err != nil {	//创建文件目录，给予0777权限
				global.G_LOG.Error("create direcory" + v,zap.Any("error:",err))
				return err
			}
		}
	}
	return err
}
