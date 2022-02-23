# Owned Media
企業の情報発信メディアを想定。

## 制作の目的
1. WEBアプリケーションの構造を理解する。
2. GCPのCloudRunのサービスに慣れる。
3. CIとしてGithubActionsに慣れる。
    →　やってみた結果、Cloud BuildとGithubの連携でのCI&CDに落ち着いた。
4. GO言語を覚える。 ← 今ここ。アプリ作成中 
5. GO言語をさらに覚える。
6. フロントエンドもそこそこ頑張る。

## どんなWEBアプリケーション？
- 正直未定。
- ブログ＋コーポレートサイト＋掲示板＋αを模索中。
- もう一つの「om-api-server」を呼び出して何かをやりたい。

### Featuers
1. TOPページ
   - 一般的に必要なページCONTACT USとか.
2. 記事ページ
3. 投稿ページ
4. 管理者ページ
5. 記事の編集ページ

## Technologies
### Language
- GOlang

### Infrastructure
Cloud Runには標準でモニターとロギングの機能がついているので専用のサービスを利用する必要はない。
当分は費用の面でも標準の機能のみを使用していく。

#### PaaS
- Cloud Run

#### CI
- Cloud Build
- Github Actions(当初利用予定だったが変更)
- Google Buildpacks(クラウドのサービスではないが重要)

#### Container Repogitry
- Artifact Registry

#### Database
- Cloud SQL

#### Monitor
- Promehteus

#### Logging
- Cloud Logging

test
