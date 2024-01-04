# jenkins-job-loader

一个用于导入/导出，备份还原 Jenkins Jobs 的工具

## 语言

[English Doc](README.md)

## 基本知识

可以在 Jenkins Job 的访问地址后添加 `/config.xml` 来获取该 Job 的配置详细信息，例如：

Job 链接：<http://jenkins.xxx.com/job/your-job>, 

则 Job 配置详情为： <http://jenkins.xxx.com/job/your-job/config.xml>

## 编译可执行文件

可以使用以下指令编译可执行文件：

``` bash
go mod tidy

GOOS=linux GOARCH=amd64 GO111MODULE=on CGO_ENABLED=0 go build --ldflags="-s" -v
```

## 使用

- 导出 Jenkins Jobs 到指定目录

``` bash
# 模板
./jenkins-job-loader export -d "{domain}" -u "{user}" -p "{password}" -f "{folder_export_to}"

# 示例
./jenkins-job-loader export -d "http://jenkins.xxx.com" -u "admin" -p "123456" -f "./jobs/example"
```

- 从指定目录导入 Jenkins Jobs

``` bash
# 模板
./jenkins-job-loader load -d "{domain}" -u "{user}" -p "{password}" -f "{job_config_folder}" -c "{your_credentials_id}"

# 示例 1：导入 Jobs 但不指定访问代码访问凭证 Id
./jenkins-job-loader load -d "http://jenkins.xxx.com" -u "admin" -p "123456" -f "./jobs/example"

# 示例 2：导入 Jobs 但并指定访问代码访问凭证 Id
./jenkins-job-loader load -d "http://jenkins.xxx.com" -u "admin" -p "123456" -f "./jobs/example" -c "1b7574ed-56e7-4af9-b851-052df8e53b87"
```

- 从指定目录移除 Jenkins Jobs

仅移除指定目录中的 Jobs，目录外的 Jobs 将被保留

``` bash
# 模板
./jenkins-job-loader unload -d "{domain}" -u "{user}" -p "{password}" -f "{job_config_folder}"

# 示例
./jenkins-job-loader unload -d "http://jenkins.xxx.com" -u "admin" -p "123456" -f "./jobs/example"
```

## 初始化 Jenkins Jobs

本项目提供了 [Jenkinsfile](Jenkinsfile)，可以直接用来初始化 Jenkins，只需要创建一个 `multibranch pipeline job`，并用该 Job 导入其它 Jobs 即可
