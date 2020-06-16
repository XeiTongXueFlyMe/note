# 下载x264源码
git clone https://github.com/mirror/x264.git

#创建安装目录
export PATH=${HOME}/build_libs/bin:${PATH}

#安装汇编器
sudo apt install assembler
sudo apt-get install nasm
> 可以不安装，配置参数需要填写 --disable-asm

#配置编译
./configure --prefix=${HOME}/build_libs --enable-shared --enable-pic --enable-static

#编译安装
make
make install