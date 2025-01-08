# hotfix

热更新文档

1,前提需要被热更的公共属性和方法(首字母大写)

2,热更所import的包必须提前在服务启动之前生成符号

3，具体生成在symbols包下面执行[symbols.go](example%2Fsymbols%2Fsymbols.go)

go:generate yaegi extract xxx路径 包名（包名跟路径一样可以省去包名）

4,脚本[test_go.patch](example%2Fpatch%2Ftest_go.patch) 文件任何位置下

## 示例
[example](example)