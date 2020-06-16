
#移除本地的ffmpeg
sudo  apt remove ffmpeg

export PKG_CONFIG_PATH=${HOME}/build_libs:${PATH}

#配置编译
./configure  --enable-shared --disable-yasm --prefix=/home/immm/build_libs/ffmpeg

#编译安装
make
make install

#参考网页
https://www.jianshu.com/p/59da3d350488