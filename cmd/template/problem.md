# 1
---
## 問題

配列を二つ入力させ、その配列の要素数を足した結果を出力しなさい。

### inputの方法
配列のinputは`1 2 3 4`という形式で行われる。
これをpythonの配列にするために、以下のコードを利用してよい。
```python
a = list(map(int, input().split()))
```
　
## テストケース
```json
[
	{
		"input":
		[
			"1 2 3",
			"3 4 5"
		]
	},
	{
		"input":
		[
			"a b c",
			"aa bb cc"
		]
	}
]
```
## 模範回答
```python
A = list(map(int, input().split()))
B = list(map(int, input().split()))
lenA, lenB = len(A), len(B)
print(lenA + lenB)
```
