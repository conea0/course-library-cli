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

### ローカルでテスト

作成した資料と問題が、course-libraryのプラットフォーム上に正しく反映され、動作するのかをローカルで検証することができます。

```sh
clc test ./
```
