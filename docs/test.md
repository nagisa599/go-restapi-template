# go のテスト戦略

目標(ゴール):github copilot を使った AI 駆動による go テスト開発

1. repository のテスト
2. usecase のテスト
3. handler のテスト

## 利用ライブラリー

| ライブラリー名 | github                                   | 目的                                     |
| -------------- | ---------------------------------------- | ---------------------------------------- |
| gomockhandler  | ithub.com/sanposhiho/gomockhandler@lates | mock を一言管理する                      |
| sqlmock        | github.com/DATA-DOG/go-sqlmoc            | repository のテストする                  |
| mockgen        | github.com/golang/mock/mockgen@latest    | repository と usecase の mock を作成する |

## 手順

## 1, repository のテスト

sqlmock を利用する。

## 2 usecase のテスト

gomock を利用する

## 3 usecase のテスト

gomock を利用する

[参考文献](https://engineering.mercari.com/blog/entry/20210406-gomockhandler/)
