
1. 按照官方步骤下载arm-linux-gnueabihf-gcc
> 建议版本： gcc-linaro-6.3.1-2017.02-x86_64_arm-linux-gnueabihf
2. $ sudo apt-get install device-tree-compiler


## 小知识：
1. 交叉编译工具链的命名规则为：arch [-vendor] [-os] [-(gnu)eabi]

* arch - 体系架构，如ARM，MIPS
* vendor - 工具链提供商
* os - 目标操作系统
* eabi - 嵌入式应用二进制接口（Embedded Application Binary Interface）

2. 根据对操作系统的支持与否，ARM GCC可分为支持和不支持操作系统，如
* arm-none-eabi：这个是没有操作系统的，自然不可能支持那些跟操作系统关系密切的函数，比如fork(2)。他使用的是newlib这个专用于嵌入式系统的C库。
* arm-none-linux-eabi：用于Linux的，使用Glibc

3. ABI 和 EABI
* ABI：二进制应用程序接口(Application Binary Interface (ABI) for the ARM Architecture)。在计算机中，应用二进制接口描述了应用程序（或者其他类型）和操作系统之间或其他应用程序的低级接口。

* EABI：嵌入式ABI。嵌入式应用二进制接口指定了文件格式、数据类型、寄存器使用、堆积组织优化和在一个嵌入式软件中的参数的标准约定。开发者使用自己的汇编语言也可以使用 EABI 作为与兼容的编译器生成的汇编语言的接口。

> 两者主要区别是，ABI是计算机上的，EABI是嵌入式平台上（如ARM，MIPS等）。

4. arm-linux-gnueabi-gcc 和 arm-linux-gnueabihf-gcc
* 两个交叉编译器分别适用于 armel 和 armhf 两个不同的架构，armel 和 armhf 这两种架构在对待浮点运算采取了不同的策略（有 fpu 的 arm 才能支持这两种浮点运算策略）。

* 其实这两个交叉编译器只不过是 gcc 的选项 -mfloat-abi 的默认值不同。gcc 的选项 -mfloat-abi 有三种值 soft、softfp、hard（其中后两者都要求 arm 里有 fpu 浮点运算单元，soft 与后两者是兼容的，但 softfp 和 hard 两种模式互不兼容）：
* soft： 不用fpu进行浮点计算，即使有fpu浮点运算单元也不用，而是使用软件模式。
* softfp： armel架构（对应的编译器为 arm-linux-gnueabi-gcc ）采用的默认值，用fpu计算，但是传参数用普通寄存器传，这样中断的时候，只需要保存普通寄存器，中断负荷小，但是参数需要转换成浮点的再计算。
* hard： armhf架构（对应的编译器 arm-linux-gnueabihf-gcc ）采用的默认值，用fpu计算，传参数也用fpu中的浮点寄存器传，省去了转换，性能最好，但是中断负荷高。