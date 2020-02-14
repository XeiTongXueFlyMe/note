
> 以下开发环境，任选其一即可

# 环境1：在Linux下使用STM32Cube工具集

> 工具集都可以在ST官网下载

## 安装 STM32CubeMX
> 这个是ST公司官方开发的用于初始化STM32外设驱动的软件。这款软件可以非常方便地生成高质量的硬件配置代码，使得STM32的开发可以专注于功能的实现

//安装java环境
1. sudo apt-get install openjdk-8-jdk openjdk-8-source
//添加运行权限
2. chmod 777 SetupSTM32CubeMX-5.4.0.linux
//安装
3. ./SetupSTM32CubeMX-5.4.0.linux 
> 安装完成后，程序启动图标，需要在安装地址找 STM32CubeMX

## 安装 STM32CubeIde
1. chmod +x st-stm32cubeide_1.1.0_4551_20191014_1140_amd64.deb_bundle.sh
2. sudo ./st-stm32cubeide_1.1.0_4551_20191014_1140_amd64.deb_bundle.sh  


# 环境2：安装makefile版环境
## 安装编译器
sudo apt install gcc-arm-none-eabi
sudo apt install gdb-arm-none-eabi
sudo apt install binutils-arm-none-eabi

##　安装 cmake 和 libusb
sudo apt update
sudo apt install cmake
sudo apt install libusb-1.0-0 libusb-1.0-0-dev libusb-1.0-0-dbg

## 安装　jlink 工具
sudo  apt install  openocd 

##　安装 stlink
git clone https://github.com/texane/stlink.git
cd stlink/
make
make install