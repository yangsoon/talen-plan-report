<p>
    <a href="http://jiyiren.github.io/"><img alt="logo" width="38" height="38" src="./_img/header.png" alt="jiyiren">
    </a>
</p>

<p align="center" style="font-size: 2em">
    ReportPaper
</p>

<p align="center" style="font-size: 16px">@<a href="mailto:csyiji@gmail.com">jiyiren</a></p>

<p align="center" style="margin: 30px 0 35px;">A LaTeX Template for Report, With Supporting Chinese Language
</p>


<img src="./_img/hot.jpg">

## 环境

需要先安装 LaTeX 环境，而支持中文的 LaTeX 为 [CTeX](http://www.ctex.org/HomePage)，CTeX 也只是一种标准定义，其通常分为两个发行版：

* TeXLive: [http://tug.org/texlive/](http://tug.org/texlive/) 
* MiKTeX: [https://miktex.org/](https://miktex.org/)

两个发行版都是全平台支持的，很多人会将 [MacTex](http://tug.org/mactex/) 也作为一种发行版，但我喜欢将其归类于 TeXLive 发行版中，这看自己的意愿。

因此，环境的话大家可以自行选择，这里为使用 TeXLive，下面为安装包地址，大概有 3 GB 左右：


* MacOS: [MacTex](http://tug.org/mactex/)
* Windows: [TeXLive.iso](http://mirrors.hust.edu.cn/CTAN/systems/texlive/Images/)

安装后将命令加入环境变量，并使之生效，通过 `which latex` 可查看是否设置成功：

```bash
$ which latex
/usr/local/texlive/2018/bin/x86_64-darwin/latex
```

## IDE

LaTeX 实际上如果 Java 语言一样，都需要先配置环境，然后选择一款自己喜欢的 IDE 进行编写“代码”。

当然所有文本编辑器都可以编写 LaTeX 或者 Java 等其他语言“代码”。这里的 IDE 指集成了一些语言本地化的功能，比如编译、特殊符号等等。

LaTeX 的发行版中会自带一款编辑器，用 TexLive 的话，MacOS 上会有个叫 **TexShop** 的编辑器，而 Windows 上则会是一个叫 **TexWorker** 的编辑器，这些是都可以胜任编写工作的。

另外，对于第三的 LaTeX 编辑器，笔者也用的不多，这里我推荐两个：

* TeXMaker: [http://www.xm1math.net/texmaker/](http://www.xm1math.net/texmaker/), 全平台，免费，强烈推荐；
* TeXStudio: [http://texstudio.sourceforge.net/](http://texstudio.sourceforge.net/) , 全平台，免费，推荐；
* WinEdt: [http://www.winedt.com/index.html](http://www.winedt.com/index.html) , 只支持 Windows，收费，自己选；

我个人推荐前两个，因为第三个收费且不跨平台，之所以写上第三个，主要是因为网络上很多博客或用户都推荐用第三个，这个我使用时也感觉不错，但每次我使用都得到 Windows 上使用，比较麻烦。当然，如果你使用 Windows 且有钱，WinEdt 确实使用体验和功能都是比较好的。


## 使用

上面两点都讲了没实际作用的环境配置，对于使用本项目实际上很简单：

* 使用 XeLaTeX 编译：

	```bash
	xelatex report_paper.tex
	```
* 用 BiBTeX 再次编译生成的 `report_paper.aux` 文件：

	```bash
	bibtex report_paper.aux
	```
* 之后再次用 XeLaTeX 编译：

	```bash
	xelatex report_paper.tex
	```

这时候生成的 `pdf` 文件是最全和完整的文档。

上面是命令行编译的，大家若使用 IDE 则是比较简单的操作：

* 编译选择 XeLaTeX 编译一次；
* 再选择 BiBTeX 编译一次；
* 再选择 XeLaTeX 编译一次；

这时产生的 `pdf` 和上面一致；通过编辑器形式不用自己指定文件名，因此更简单方便。

最终生成的 pdf 文档示例：

* github: [report_paper.pdf](https://github.com/jiyiren/ReportPaper/blob/master/report_paper.pdf)
* 七牛云：[report_paper.pdf](http://img.godjiyi.cn/report_paper.pdf)

## 参考

* [CTeX](http://www.ctex.org/HomePage)
* [TexLive](http://tug.org/texlive/)
* [MiKTeX](https://miktex.org/)
* [MacTeX](http://www.tug.org/mactex/index.html)
* [TexMaker(全平台支持编辑器)](http://www.xm1math.net/texmaker/index.html)
* [TeXstudio(全平台支持编辑器)](http://texstudio.sourceforge.net/)


