# 包导入问题
自建包要整成一个文件夹，文件夹名字与包名一致
gomod文件导入时要包含根目录名字
比如：go.mod在2.6目录下，2.6目录下还有main.go和tempconv文件夹
在main.go导包时，import“2.6/tempconv”
编写mod时也要go get 2.6/tempconv
dd 困死了