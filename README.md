# translation-tool

实现中。。。

**初衷：**

有时候在一个翻译平台翻译句子，结果可能不是很准确

在多个不同的翻译平台上进行翻译同一个句子，能够提升结果的准确性

但是来回切换页面，复制英文句子，是一件比较耗时的

所以打算做一个集成若干翻译平台的文本翻译功能的工具，在程序终端输入英文，
直接输出各个平台翻译的中文内容，提升效率、节省时间

**已实现：**
* 接通有道翻译
* 接通腾讯翻译君
* 接通百度翻译
* 接通谷歌翻译
* 接通搜狗翻译

接通有道翻译：
1. [有道翻译官网申请AppKey和SecKey，点击直达](https://ai.youdao.com/doc.s#guide)
2. 在api/youdao.go里，把AppKey和SecKey分别赋值给YouDaoAppKey，YouDaoSecKey
3. 运行，输入要翻译的英文

接通腾讯翻译君：
1. [开通免费试用版，点击直达](https://fanyi.qq.com/translateapi)
2. [到腾讯云控制台创建/获取SecretId和SecretKey，点击直达](https://console.cloud.tencent.com/cam/capi)
3. 在api/tmt.go里，把SecretId和SecretKey分别赋值给TMTAppKey，TMTSecKey
4. 运行，输入要翻译的英文

接通百度翻译：
1. [获取百度翻译App ID和密钥，点击直达](http://api.fanyi.baidu.com/api/trans/product/desktop?req=developer)
2. 在api/baidu.go里，把App ID和密钥分别赋值给BaiDuAppKey，BaiDuSecKey
3. 运行，输入要翻译的英文

接通搜狗翻译：
1. [获取搜狗翻译 PID和Key，点击直达](https://deepi.sogou.com/registered/textcognitive)
2. 在api/sogou.go里，把PID和Key分别赋值给SoGouAppKey，SoGouSecKey
3. 运行，输入要翻译的英文
