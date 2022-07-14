package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"go.uber.org/zap"
)

type FeverService struct {

}

//删除
func (f *FeverService)DelFeverInfo(id string) (err error) {
	if err = global.G_DB.Model(&system.FeverInfo{}).Delete(&system.FeverInfo{},"id = ?",id).Error;
	err != nil{
		global.G_LOG.Error("删除发热记录数据失败", zap.Error(err))
		return err
	}
	return err
}

//修改
func (f *FeverService)UpdateFeverInfo(info system.FeverInfo) (err error) {
	if err = global.G_DB.Model(&system.FeverInfo{}).Save(&info).Error;err != nil{
		global.G_LOG.Error("修改发热记录数据失败", zap.Error(err))
		return err
	}
	return err
}

//新增
func (f *FeverService)AddFeverInfo(info system.FeverInfo) (err error) {
	if err = global.G_DB.Model(&system.FeverInfo{}).Create(&info).Error;err != nil{
		global.G_LOG.Error("新增发热记录数据失败", zap.Error(err))
		return err
	}
	return err
}

