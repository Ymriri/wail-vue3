# 商品系统

效果图
![img.png](img/img.png)

![img.png](img.png)

##  技术栈
    
    本系统使用 wails + vue3 + elementplus + gorm
    数据库使用 sqlite 
    本地系统没有上云和账户能力，可以通过系统备份和导入来同步数据库
   


## 如何使用
    
    打包：wails build --platform windows darwin
    生成目录：goods-system/build/bin
    本地使用需要sqlite，程序打开时自动创建表,同步备份根据主键id判断时创建还是更新
    wails build -nsis  // 免安装版本
    wails build --platform windows/amd64 -nsis
    wails build --platform windows/amd64  -webview2 Embed
