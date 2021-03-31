# translation-tool

实现中。。。

**初衷：**

有时候在一个翻译平台翻译句子，结果可能不是很准确

在多个不同的翻译平台上进行翻译同一个句子，能够提升结果的准确性

但是来回切换页面，复制英文句子，是一件比较耗时的

所以打算做一个集成若干翻译平台的文本翻译功能的工具，在程序终端输入英文，
直接输出各个平台翻译的中文内容，提升效率、节省时间

**已实现：**
* 有道翻译

接入：
1. [有道翻译官网申请AppKey和SecKey，点击直达](https://ai.youdao.com/doc.s#guide)
2. 在api/youdao.go里，把AppKey和SecKey分别赋值给YouDaoAppKey，YouDaoSecKey
3. 运行，输入要翻译的英文

**待实现：**
* 腾讯翻译君Api
* 金山词典Api
* 百度翻译Api
* 搜狗翻译Api
* 谷歌翻译Api