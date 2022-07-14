package core

import (
	"flag"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

// Viper 优先级: 命令行 > 环境变量 > 默认值
//我们会先从命令行中获取参数，之后进行判断，命令行是否有参数？环境变量是否有参数？最后是程序调用传入的参数
func Viper(path ...string) *viper.Viper  {
	var config string	//参数变量，存储的是配置文件路径。从命令行/环境变量/默认值中获取到，根据其值选择哪个配置文件

	//判断参数
	if len(path) == 0 {
		//通过命令行获取其输入的参数，详情请看flag标准库 flag.TypeVar(Type指针, flag名, 默认值, 帮助信息)
		flag.StringVar(&config,"c","","choose config file.")
		flag.Parse() //解析命令行参数
		if config == ""{	// 判断命令行参数是否为空
			if configEnv := os.Getenv(ConfigEnv); configEnv == ""{// 判断 internal.ConfigEnv 常量存储的环境变量是否为空
				//注意，判断环境变量可以用os.Getenv，也可以用env库去获取,还可以用Viper.bindEnv去获取。我们现在用的是os,去读取变量中存储的环境变量key是否存在
				switch gin.Mode() {	//判断当前gin运行模式，一般默认是debug，开发环境中要用Release.根据不同的环境选择不同的配置文件
				case gin.DebugMode:
					config = ConfigDefaultFile
					fmt.Printf("您正在使用gin模式的%s环境,config的路径为%s\n",gin.EnvGinMode,ConfigDefaultFile)
				case gin.ReleaseMode:
					config = ConfigReleaseFile
					fmt.Printf("您正在使用gin模式的%s环境,config的路径为%s\n",gin.EnvGinMode,ConfigReleaseFile)
				case gin.TestMode:
					config = ConfigTestFile
					fmt.Printf("您正在使用gin模式的%s环境,config的路径为%s\n",gin.EnvGinMode,ConfigTestFile)
				}
			}else{	//存储的环境变量不为空，则赋值给config
				config = configEnv
				fmt.Printf("您正在使用%s环境变量,config的路径为%s\n", ConfigEnv, config)
			}
		}else{	//命令行参数不为空，则StringVar已将值赋值给config
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%s\n", config)
		}
	} else {	//函数传递的可变参数的第一个值赋值给config
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%s\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)	//重点：上面做了那么多，就是为了这句
	v.SetConfigType("yaml")
	err := v.ReadInConfig()		//读取并解析配置
	if err != nil {
		panic(fmt.Errorf("viper读取配置失败，%s\n",err))	//若失败了，则这个viper初始化就失败了，所以引发panic
	}
	v.WatchConfig() 		//开启监听变化，若有变化则热加载

	v.OnConfigChange(func(e fsnotify.Event) {	//回调函数,监听配置文件变化，触发事件
		fmt.Println("配置文件发生变化：",e.Name)
		if err = v.Unmarshal(&global.G_CONFIG);err !=nil {		//若文件发生改变，则再次初始化配置
			fmt.Println(err)
		}
	})

	//重点：将配置文件的所有配置初始化在G_CONFIG中
	if err = v.Unmarshal(&global.G_CONFIG); err != nil {
		fmt.Println(err)
	}


	/*// root 适配性 根据root位置去找到对应迁移位置,保证root路径有效	//问题：这里看不懂，原理是什么
	global.GVA_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	global.BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(time.Second * time.Duration(global.GVA_CONFIG.JWT.ExpiresTime)),
	)*/

	return v
}