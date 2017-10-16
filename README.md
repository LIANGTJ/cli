## 简介

 selpg 是从文本输入选择页范围的实用程序，由go语言实现。该输入可以来自作为最后一个命令行参数指定的文件，在没有给出文件名参数时也可以来自标准输入。

## 用法

```shell
./selpg -s startPage -e endPage [-f] -l linesPerPages [-d des] [filename]
```
```shell
Usage of ./selpg:
  -d string
    	the target destiny
  -e int
    	the last page selected (default 1)
  -f bool
  		whether the page is separeted by \f, confilt with -l
  -l int
    	the capacity of a page [default: 72 lines] (default 72)
  -s int
    	the first page selected (default 1)
```

## 实现原理

本程序主要用到的flag库以及io，bufio库。通过flag来解析命令行参数，然后根据对应的参数，实现不同的标准输入输出。由于shell本身会帮我们处理管道和重定向，所以我们只要适配其接口即可，因此本次程序只要理解原理，就变得比较简单，



## 测试样例

> 相关图片可在./picture下找到

本次测试文件有2个，分别是 

```shell
file test1：

1
this is a test file
this is a test file
this is a test file
this is a test file
this is a test file
this is a test file
....(以上重复4次，序号递增)

file test2：

total 1
drwxrwxr-x 5 ltj ltj    4096 Oct 16 22:35 ./
drwxrwxr-x 5 ltj ltj    4096 Oct 14 21:36 ../
-rw-rw-r-- 1 ltj ltj    1256 Oct 16 18:26 a
-rw-rw-r-- 1 ltj ltj     644 Oct 16 21:01 b
-rw-rw-r-- 1 ltj ltj     114 Oct 16 19:05 c.go
....(以上重复7次，序号递增)

```

从指定文件读取指定页范围(支持多文件)

```shell
ltj@ubuntu:~/Desktop/workSpace/src$ ./selpg -s 1 -e 2 -l 3  test test2
this is file 1 :
1
this is a test file
this is a test file
this is a test file
this is a test file
this is a test file

this is file 2 :
total 1
drwxrwxr-x 5 ltj ltj    4096 Oct 16 22:35 ./
drwxrwxr-x 5 ltj ltj    4096 Oct 14 21:36 ../
-rw-rw-r-- 1 ltj ltj    1256 Oct 16 18:26 a
-rw-rw-r-- 1 ltj ltj     644 Oct 16 21:01 b
-rw-rw-r-- 1 ltj ltj     114 Oct 16 19:05 c.go

```

selpg 读取标准输入，而标准输入已被 shell／内核重定向为来自“input_file”而不是显式命名的文件名参数

``` shell
ltj@ubuntu:~/Desktop/workSpace/src$ ./selpg -s 1 -e 2 -l 3 < test
1
this is a test file
this is a test file
this is a test file
this is a test file
this is a test file

```

selpg 标准输出被 shell／内核通过管道作为 grep的标准输入

```shell
ltj@ubuntu:~/Desktop/workSpace/src$ ./selpg -s 1 -e 2 -l 3 test | grep "test"
this is a test file
this is a test file
this is a test file
this is a test file
this is a test file

```

标准输出被 shell／内核重定向至“test3”文件。

```shell
ltj@ubuntu:~/Desktop/workSpace/src$ ./selpg -s 1 -e 2 -l 3 -f test 2> test3
ltj@ubuntu:~/Desktop/workSpace/src$ cat test3
you should not specify -f and -l at the same time
  -d string
    	the target destiny
  -e int
    	the last page selected (default 1)
  -f	whether the page is separeted by \f?
  -l int
    	the capacity of a page [default: 72 lines] (default 72)
  -s int

```

selpg 支持后台运行进程

```shell
ltj@ubuntu:~/Desktop/workSpace/src$ ./selpg -s 1 -e 2 -l 3 test &
[1] 72077
ltj@ubuntu:~/Desktop/workSpace/src$ this is file 1 :
1
this is a test file
this is a test file
this is a test file
this is a test file
this is a test file
ll
total 1976
drwxrwxr-x 3 ltj ltj    4096 Oct 16 23:43 ./
drwxrwxr-x 5 ltj ltj    4096 Oct 14 21:36 ../
drwxrwxr-x 8 ltj ltj    4096 Oct 16 23:17 .git/
-rwxrwxr-x 1 ltj ltj 1991108 Oct 16 23:21 selpg*
-rw-rw-r-- 1 ltj ltj    2245 Oct 16 23:10 selpg.go
-rw-rw-r-- 1 ltj ltj     488 Oct 16 22:57 test
-rw-rw-r-- 1 ltj ltj    1638 Oct 16 23:08 test2
-rw-rw-r-- 1 ltj ltj     297 Oct 16 23:45 test3
[1]+  Done                    ./selpg -s 1 -e 2 -l 3 test

```

