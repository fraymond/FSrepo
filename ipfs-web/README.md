```bash
# 克隆项目
git clone 

# 安装依赖(优先使用)
npm install
# 安装依赖(下载较慢时使用)
npm install --registry=https://registry.npm.taobao.org

#上述2种【安装依赖】无法正常时，请尝试删除node_modules文件夹后，使用cnpm安装
# 第一步
npm install -g cnpm --registry=https://registry.npm.taobao.org
# 第二步
cnpm install

# 启动服务
npm run dev
```

- 开发时，如何连接后台项目api接口？
> 修改ipfs/static/config/index.js目录文件中window.SITE_CONFIG.baseUrl = '本地api接口请求地址'

- 开发时，如何解决跨域？
> 1. 修改ipfs/config/dev.env.js目录文件中OPEN_PROXY: true开启代理
> 2. 修改ipfs/config/index.js目录文件中proxyTable对象target: '代理api接口请求地址'
> 3. 重启本地服务

- 开发时，如何提前配置CDN静态资源？
> 修改ipfs/static/config/index-[qa/uat/prod].js目录文件中window.SITE_CONFIG.cdnUrl = '静态资源cdn地址' + window.SITE_CONFIG.staticFileName

## 发布
> 构建生成的资源文件保存在ipfs/dist目录下，可通过config/index.js目录文件修改相关配置信息

```bash
# 构建生产环境(默认)
npm run build

# 构建测试环境
npm run build --qa

# 构建验收环境
npm run build --uat

# 构建生产环境
npm run build --prod
```

- 构建生成后，发布需要上传哪些文件？
> 修改ipfs/dist目录下：181101（静态资源，由当前日期动态生成文件夹名）、config（配置文件）、index.html

- 构建生成后，如何动态配置CDN静态资源？
> 修改ipfs/dist/config/index.js目录文件中window.SITE_CONFIG.cdnUrl = '静态资源cdn地址' + window.SITE_CONFIG.staticFileName

- 构建生成后，如何动态切换新旧版本？
> 修改ipfs/dist/config/index.js目录文件中window.SITE_CONFIG.staticFileName = '181101（静态资源文件夹名称）'


## 其他
``` bash
# build for production and view the bundle analyzer report
npm run build --report

# run unit tests
npm run unit

# run e2e tests
npm run e2e

# run all tests
npm test
```