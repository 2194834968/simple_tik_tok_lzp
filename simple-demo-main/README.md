# simple-demo

## 抖音项目服务端简单实现

安卓apk下载以及apk文档：https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7#
接口文档https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18345145

工程需要创建数据库才能运行，数据库名simple_tik_tok，定义见测试数据

```shell
go build && ./simple-demo
```

### 功能说明

* 用户数据库中尚未包含用户头像链接avatar，个性签名signature，背景图链接background_image
* 用户数据库中新注册用户，用户名为账号名
* 如此处理以上数据的原因，是应用并未添加编辑相应数据的接口，以上数据均需要在数据库手动更改

* 用户数据库额外支持了total_favorited（被赞的总数量）和Favorite_Count（喜欢的视频总数量）两个int

* 视频”熊“链接时好时坏，暂未更换
* 部分视频数据是功能不完善时编写的，会有封面和视频不匹配的问题

* 视频上传后会保存到本地 public 目录中，访问时用 127.0.0.1:8080/static/video_name 即可
* 会自动截取视频第一帧保存为视频同名文件

### 测试数据

测试数据全部采用数据库中数据，数据库定义已在service文件中给出，例如user的数据库定义是user加后缀DataBase
