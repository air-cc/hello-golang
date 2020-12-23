# GO 入门

## 入门资料

- [Go 指南（在线演示）](https://tour.go-zh.org/)：【初体验】在线使用 GO 语言
- [Go 简明教程](https://geektutu.com/post/quick-golang.html)：【懒人福音】go 技术栈全知道（皮毛），适合有一点语言基础，快速上手 go 技术栈
- [Go 入门指南](https://www.kancloud.cn/kancloud/the-way-to-go/72432)：【推荐】全面介绍基础语法，相较于「Go语言圣经」描述更详尽，用例更丰富，需耐心学习
- [Go Web 编程](https://astaxie.gitbooks.io/build-web-application-with-golang/zh/)：【概览】基础语法涉及不多，更偏重go语言在常规 Web 技术方面的使用
- [Go语言圣经（中文版）](https://yar999.gitbooks.io/gopl-zh/content/)：go 基础语法完整表述，内容详实

## 一些注意点

- new 为预定义的函数，非关键字
- 局部变量的存储空间是分配在栈上还是堆上由编译器自动选择
- 自增、自减是语句 不是表达式 不可以 j = i++
- map 非线程安全，多线程存取需要使用互斥锁
- 不要通过共享来通信，而要通过通信来共享

## hello-golang

├── basic          # go 基础语法
├── module         # go 模块
└── web-framework  # 流行 web 框架
    └── hello-gin  # gin