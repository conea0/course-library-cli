
---
## 問題

nに数値を代入して、1からnまでの偶数のリストと、奇数のリストを出力してください。

## テストケース
input = int
```
[
	{
		"input": [10],
		"output": [
			[2,4,6,8,10],
			[1,3,5,7,9]
		]
  	},
	{
		"input": [5],
		"output": [
			[2,4],
			[1,3,5]
		]
  	},
	{
		"input": [12],
		"output": [
			[2,4,6,8,10,12],
			[1,3,5,7,9,11]
		]
  	},
]
```


## 模範回答
```python
n = get_input()
even = []
odd = []

for i in range(1,n+1):
	if i % 2 == 0:
		even.append(i)
	else:
		odd.append(i)

print(even)
print(odd)
```