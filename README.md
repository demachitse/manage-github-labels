# manage-github-labels

## 概要

Githubのラベルの複製を行えるCUIアプリ

## 前提

### インストール（goの環境がある場合）

```shell
$ go install github.com/demachitse/manage-github-labels
```

### インストール（goの環境がないWindowsの場合）

Releasesからmanage-github-labels.exeをダウンロード

### 設定

Githubの[アクセストークン](https://docs.github.com/en/github/authenticating-to-github/creating-a-personal-access-token)とメールアドレスを設定ファイルに設定する

```yaml
email: "test@example.com"
token: "xxxxxxxxxxxxxxxxxx"
```

設定ファイルは以下に配置する

| OS         | place                                               |
|------------|-----------------------------------------------------|
| Window     | `%AppData%¥manage-github-labels¥config.yaml`                         |
| Mac        | `$HOME/Library/Application Support/manage-github-labels/config.yaml` |
| Linux/Unix | `$HOME/.config/manage-github-labels/config.yaml`                     |

## 使い方

### 実行方法（goの環境がある場合）

以下のコマンドを入力することで実行可能  

```shell
$ go run github.com/demachitse/manage-github-labels
```

また、以下のようにコマンド引数を入れると対話せずに動作を指定可能

```shell
$ go run github.com/demachitse/manage-github-labels {owner} {repository} {command}
```

### 実行方法（goの環境がないWindowsの場合）

manage-github-labels.exeをコマンドプロンプトで実行するか、ダブルクリックして起動

### コマンド詳細

|実行引数|説明|
|--|--|
|owner|Githubのオーナー名</br>例　<https://github.com/demachitse/manage-github-labels>　の場合 `demachitse`|
|repository|Githubのリポジトリ名</br>例　<https://github.com/demachitse/manage-github-labels>　の場合 `manage-github-labels`|
|command|次の表を参考|

|コマンド|説明|
|--|--|
|r または reset|ラベルを全て削除し、作成し直す|
|g または get|ラベルを全て取得しコンソールに出力する|
|s または save|ラベルを全て取得し設定ファイルに保存する|
|c または create|ラベルを作成する|
|d または delete|ラベルを全て削除する|
