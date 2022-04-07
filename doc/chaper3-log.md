## 日志

一个应用日志模块必不可少，尤其编译型语言，在定位线上问题的时候显得尤其重要。
有个段子：线上怎么定位问题，二分log法。

### [logrus]("https://github.com/sirupsen/logrus") 

使用一个包的时候，直接看 README.md是最快的方式。下载这个包，然后直接在本地查看，非常方便。

既然说到本地开发，再另外说一个，之前写php的时候，遇到难复现的问题，直接预发布环境debug，简单快捷。go也可以有类似的做法，
方法就是再本地打个二进制包，scp、ftp等方式上传到服务器即可。但是要使用交叉编辑，打成 linux 可以使用的包:

```php
GOOS=linux GOARCH=amd64 go build -o $out/gin_web_layout main.go 
```

#### logrus初始化

```php

	rotate := &lumberjack.Logger{
		Filename:   log.Filename,
		MaxSize:    500, // 最大的文件2000M
		MaxBackups: 30,  // 最多保留10个文件
		MaxAge:     10,  // 最长保留7天
	}
	logrus.SetOutput(rotate)
	
```

