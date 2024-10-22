### 说明
* 简易的文件存储,通过系统的文件管理实现
* 方便小项目的快速部署

### 如何使用
1. 通过配置文件,配置文件的位置在` ./config/config.yaml `
    ```
    #服务的端口
    port: 8080
    #存储的目录
    dir: './resource/'
    #启用下载
    downloadEnable: true
    #下载的Token,启用下载时有效
    downloadToken: ''
    #启用上传
    uploadEnable: false
    #上传的Token,启用上传时有效
    uploadToken: ''
    #启用删除
    deleteEnable: false
    #删除的Token,启用删除时有效
    deleteToken: ''
    ```
2. 通过命令行
    ```
   filestore --port=8080 --dir==./upload --uploadEnable=true
   ```
    
3. 启用下载后可以使用query的show进行预览
    ```
   http://localhost:8080/img.png?show
   ```