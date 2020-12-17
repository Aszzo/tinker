# tinker
> 修补匠（Tinker）是Valve公司开发的战略游戏《DOTA2》（Defense of the Ancients）中的一个智力型英雄
### golang 图片压缩工具
- png压缩

    使用[pngquant](https://github.com/kornelski/pngquant)。通过cargo构建成可执行文件
    
    相关链接：https://pngquant.org/install.html
- jpp压缩

    使用golang的[resize](https://github.com/nfnt/resize)包来实现jpg文件的压缩

```bazaar
version/1.0.0
Usage: tinker [-i input] [-q quality]
Options:
    -h    this help
    -i string
          目标文件夹|目标文件
    -q int
          压缩图片的质量（0-100） (default 80)
``` 
#### node环境使用
```bazaar
只做了mac系统的支持

1、git clone 
2、cd npm/tinker
3、npm install
4、sudo npm link
```
[node环境使用说明](npm/tinker/README.md)



