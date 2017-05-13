# example-golang-oauth2
## 概要
Go言語のoauth認証 / 連携のサンプルソース
* サンプル内容
 * google,twitter,facebook,githubでoauth認証
 * twitterでAPI投稿

## 導入
### 事前準備
 * [Go言語インストール](https://golang.org/doc/install)
 * [beego quickstart](https://beego.me/quickstart)
### セットアップ
conf/app_template.conf → conf/app.confに変更して、
必要な設定を追加
```ruby:conf/app.conf
appname = example-golang-oauth2
httpport = 8080
runmode = dev

googleClientID = ""
googleClientSecret = ""

twitterConsumerKey = ""
twitterConsumerSecret = ""

facebookClientID = ""
facebookClientSecret = ""

githubClientID = ""
githubClientSecret = ""

SessionOn = true
```
インストール
```
glide install
```
start
```
bee run
# 「http://localhost:8080/」にブラウザでアクセス
```


