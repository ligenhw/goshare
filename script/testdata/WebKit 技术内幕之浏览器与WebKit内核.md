
# 前言
此文章是我最近在看的【WebKit 技术内幕】一书的一些理解和做的笔记。
而【WebKit 技术内幕】是基于 WebKit 的 Chromium 项目的讲解。


# 第一章 浏览器和浏览器内核
WebKit 内核是苹果2005年先开发并提出开源的，后面 Google 也以此为基础，并独立开发出 Chromium 的，2008年 Google 为 WebKit 为内核创建了一个新项目 chormium ，后来 Google 的 chrom 占领了浏览器的大部分市场。
![WebKit](https://upload-images.jianshu.io/upload_images/12890819-9f5fde9ee20076c7.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
图 1-6 显示的是该项目的大模块。图中“WebKit 嵌入式接口”就是批的狭义 WebKit，它批的是在 WebCore（包含上面提到的 HTML 解释器、CSS 解释器和布局等模块）和 JavaScript 引擎之上的一层绑定和嵌入式编程接口，可以被浏览器调用。

![WebKit2.png](https://upload-images.jianshu.io/upload_images/12890819-6e9c62a4ba6f4bb1.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

### Chromium 内核 Blink
2013年4月 gogle宣布从 WebKit中复制一份出来然后独立，并运作为Blink项目。

# 第二章 HTML网页与结构
### 1. 基本组成 html 、css、js。
### 2. html5新特性 video、canvas、2d、3d等，2012年就推出。
### 3. 框结构： iframe、frame、frameset，用于嵌入html文档。

![iframe.png](https://upload-images.jianshu.io/upload_images/12890819-9e44af0454bde435.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
![image.png](https://upload-images.jianshu.io/upload_images/12890819-e9eedb96b9afa5fa.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

上面的图说的是 iframe 的应用

### 4. 层次结构

理解层次结构非常重要，因为它可以帮忙你理解 WebKit 如何构建它来渲染，这有助于写高效的 HTML 代码。

网页的层次结构是指网页中的元素可能分布在不周的层次中，也就是说某些元素可以不同于它的父元素所在的层次，因为某些原因， WebKit 需要为该元素和它的子女建立一个新层。


![image.png](https://upload-images.jianshu.io/upload_images/12890819-bcc28404a696fdce.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

图中各层的前后关系。“ 根层 ” 在最后面，“ 层 3 ”和 “层 4 ” 在最前面。规律是需要复杂变换和处理的元素，它们需要新层，所以 WebKit 为它们构建新层其实是为了渲染引擎在处理上的方便和高效。对于不同的基于 WebKit 的浏览器，分层策略也有可能不一样，通常是有一些基本原则的，比如 video 、2d、3d 转换、canvas 等。


### 5. WebKit网页内核的渲染过程
![渲染过程.png](https://upload-images.jianshu.io/upload_images/12890819-b66a34049edc0545.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![从网页 URL 到构建 DOM 树 ](https://upload-images.jianshu.io/upload_images/12890819-a9ba23026cc6aa97.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![img.png](https://upload-images.jianshu.io/upload_images/12890819-0faed9d1a3b34e57.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


![从 CSS 和 DOM 树到绘图上下文.png](https://upload-images.jianshu.io/upload_images/12890819-0347437b42689e03.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![从绘图上下文到最终的图像.png](https://upload-images.jianshu.io/upload_images/12890819-ea6b341e71b35d8f.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![绘图过程说明.png](https://upload-images.jianshu.io/upload_images/12890819-dcbe11536c7faa8d.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

### 6. 编写高效代码注意点

![编写高效代码注意点](https://upload-images.jianshu.io/upload_images/12890819-73c9c74e1392e0de.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

# 6. 最后

希望本文对你有点帮助。

下期分享 第三章 WebKit 架构与模块 敬请期待。

关注公众号并回复 **福利** 便免费送你视频资源，绝对干货。

福利详情请点击： [免费资源分享--Python、Java、Linux、Go、node、vue、react、javaScript]

