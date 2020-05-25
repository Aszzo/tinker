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
        目标文件所在的目录|目标文件路径
  -q int
        压缩图片的质量（0-100） (default 80)

``` 




