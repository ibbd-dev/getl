# getl
Golang ETL

借鉴类似shell的管道概念实现ETL功能，使用上，如：

```sh
# 将age字段值+12，并替换原有的age值，输出也是csv格式
getl input --type=csv --filename=test.csv --fields-type="age:int" \
  | getl operation --expression="{age}+12" --result-field="age" \
  | getl output --type=csv --filename=output.csv
```

## 三个组件

- input: 输入
  - csv
- transform: 转换
- output: 输出
  - csv


