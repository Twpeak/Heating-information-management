package system

import (
	"errors"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"go.uber.org/zap"
	"sync"
)

/**
用户鉴权逻辑
 */

type CasbinService struct {}

var CasbinServiceApp = new(CasbinService)

//重点：若直接new了，则在外面也可以直接system.CasbinServiceApp调用，而不用service.ServiceGroupApp调用了
//在gva中，CasbinServiceApp只在同包总调用，而api层使用service.ServiceGroupApp的调用方法

var (
	syncedEnforcer *casbin.SyncedEnforcer
	once           sync.Once
)

// Casbinsync
//@function: Casbin
//@description: 持久化到数据库  引入自定义规则
//@return: *casbin.Enforcer

func (casbinService *CasbinService) Casbin() *casbin.SyncedEnforcer {
	adapter, _ := gormadapter.NewAdapterByDB(global.G_DB)		//创建适配器，策略表默认为 casbin_rule 表
	once.Do(func() {											//保证线程安全
		syncedEnforcer, _ = casbin.NewSyncedEnforcer(global.G_CONFIG.Casbin.ModelPath,adapter)	////获取一个工具对象，载入配置文件,这里载入了配适器，所以可以在数据库中调用策略(权限表)
	})
	err := syncedEnforcer.LoadPolicy()		//加载策略
	if err != nil {
		global.G_LOG.Error("加载鉴权策略失败",zap.Error(err))
		return nil
	}
	return syncedEnforcer
}

// InitCasbin
//@function: InitCasbin
//@description: 初始化自定义规则
//@return: *casbin.Enforcer

func (casbinService *CasbinService) InitCasbin()  {
	casbin := CasbinServiceApp.Casbin()
	//存储策略合集
	rules := [][]string{
		{"2","/index","GET"},
		{"2","/type/*","GET"},
		{"2","/archives","GET"},
		{"2","/music","GET"},
		{"2","/message","GET"},
		{"2","/friends","GET"},
		{"2","/pictrue","GET"},
		{"2","/about","GET"},
	}
	success, _ := casbin.AddPolicies(rules)	//添加多个策略
	if !success{
		errors.New("初始化策略失败")
		return
	}
}







//@function: GetPolicyPathByAuthorityId
//@description: 获取权限列表
//@param: authorityId string
//@return: pathMaps []request.CasbinInfo

func (casbinService *CasbinService) GetPolicyPathByAuthorityId(authorityId string) (pathMaps []request.CasbinInfo) {
	e := casbinService.Casbin()
	list := e.GetFilteredPolicy(0, authorityId)
	for _, v := range list {
		pathMaps = append(pathMaps, request.CasbinInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return pathMaps
}



//一条一条更新
//@function: UpdateCasbinApi
//@description: API更新随动
//@param: oldPath string, newPath string, oldMethod string, newMethod string
//@return: error

func (casbinService *CasbinService) UpdateCasbinApi(oldPath string, newPath string, oldMethod string, newMethod string) error {
	//当创建适配器时，就已经自动创建了cabsin_rule表。这里就选这个表去更新
	err := global.G_DB.Model(&gormadapter.CasbinRule{}).Where("v1 = ? AND v2 = ?", oldPath, oldMethod).Updates(map[string]interface{}{
		"v1": newPath,
		"v2": newMethod,
	}).Error
	return err
}

// UpdateCasbin  更新角色所有权限
//@function: UpdateCasbin
//@description: 更新casbin权限
//@param: authorityId string, casbinInfos []request.CasbinInfo
//@return: error

func (casbinService *CasbinService) UpdateCasbin(authorityId string, casbinInfos []request.CasbinInfo) error {
	//指定角色id字段删除该角色所有策略，也就是根据v0字段(存储的角色id)删除对应记录
	casbinService.ClearCasbin(0,authorityId)
	rules := [][]string{}		//存储策略合集
	for _, v := range casbinInfos{	//遍历其信息获取path和Method
		rules  =  append(rules,[]string{authorityId,v.Path,v.Method})	//拼接策略信息
	}
	e := casbinService.Casbin()	//创建工具类
	success, _ := e.AddPolicies(rules)	//添加多个策略
	if !success{
		return errors.New("存在相同api,添加失败,请联系管理员")
	}
	return nil
}


// ClearCasbin
//@function: ClearCasbin
//@description: 清除匹配的权限
//@param: v int, p ...string
//@return: bool

func (casbinService *CasbinService) ClearCasbin(v int, p ...string) bool {
	e := casbinService.Casbin()
	//指定字段删除策略
	//RemoveFilteredPolicy方法非常灵活，可以通过v参数指定 字段，
	//字段在规则中有规定，goWeb学习中v0为用户,v1角色。gva项目中，v0为角色id，v1为path路由权限 v2为方法.没有存储用户的权限
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success
}



