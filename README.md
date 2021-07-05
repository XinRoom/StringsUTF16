# StringsUTF16

一个针对二进制文件中UTF-16编码的字符进行打印或者替换的工具

## 安装
```bash
go get github.com/XinRoom/StringsUTF16
go install github.com/XinRoom/StringsUTF16
```

## 使用

- 打印二进制文件中的UTF-16字符串  
  `StringsUTF16 filename`
- 以utf16编码替换二进制中的字符  
  `StringsUTF16 filename oldstr newstr`