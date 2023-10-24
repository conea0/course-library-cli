# course-library-cli

## 概要

[course-library](https://github.com/conea0/course-library)に登録するコース作成を支援するcliツールです。

## できること

### テンプレートの作成

#### initialize
以下のコマンドで、正しく動作するpart1のフォルダと、そのプロジェクトファイルが作成されます。
```sh
clc init
```

##### 構成
```sh
├── part1
│   └── 001
│       ├── 001.md
│       └── problems
│           ├── 1.md
│           └── 2.md
```


#### partごとの作成
以下のコマンドで、./part1/002に資料と問題20問を作成するためのファイルが作成されます。
```sh
clc g ./part1/002 -n 20
```

## 実装予定

### outputの自動生成
テストケースにinputを書けば、outputを自動生成することができます。
[output自動生成の仕様](https://github.com/conea0/course-library-cli/wiki/output%E8%87%AA%E5%8B%95%E7%94%9F%E6%88%90%E3%81%AE%E4%BB%95%E6%A7%98)を参照してください。

#### partごとの出力
```sh
clc output ./part1/001/problems
```

#### 問題ごとの出力
```sh
clc output -p ./part1/001/problems/1.md
```

### outputの削除
テストケースのoutputのみを削除します。
#### partごとの削除
```sh
clc clean ./part1/001/problems
```

#### 問題ごとの削除
```sh
clc clean -p ./part1/001/problems/1.md
```

## 未実装

### ローカルでテスト

作成した資料と問題が、course-libraryのプラットフォーム上に正しく反映され、動作するのかをローカルで検証することができます。

```sh
clc test ./
```

### デプロイ

course-library のアカウントがあれば、ログインしてコースを登録することができます。

#### ログイン
以下のコマンドを実行すると、cli上またはブラウザで認証することができます。
```sh
clc login
```

#### デプロイ

以下のコマンドを実行すると、作成したコースをcourse-libraryに反映させることができます。この状態では、まだパブリックに公開されていません。

```sh
clc deploy --name name 
```

#### 公開

登録したコースを、パブリックに公開することができます。
新規に公開する場合、料金を設定するには、course-libraryのWebサービス上で設定をする必要があります。

以下のコマンドでコース一覧を取得できます。

```sh
clc list
```

以下のコマンドで公開することができます。

```sh
clc publish name