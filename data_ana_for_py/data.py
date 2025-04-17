import pandas as pd
import numpy as np
import matplotlib.pyplot as plt
import seaborn as sns

# 设置图形显示风格
sns.set(style="whitegrid")
plt.rcParams['figure.figsize'] = (10, 6)

# 1. 数据导入
# 假设 CSV 文件位于当前目录中
data = pd.read_csv("/Users/guoziqiang/my_own_projects/leet_code/data_ana_for_py/sales-data-sample.csv")

# 查看前几行数据
print("数据样本：")
print(data.head())

# 2. 数据预处理
# 将 'Date' 列转换为 datetime 类型
data['Date'] = pd.to_datetime(data['OrderDate'])

# 检查缺失值情况
print("\n缺失值统计：")
print(data.isnull().sum())

# 处理缺失值（例如对 'Discount' 缺失值设为 0，当然也可以根据实际情况选择删除或其他策略）
data['Discount'].fillna(0, inplace=True)

# 如果数据中没有直接的销售收入字段，可以计算实际收入，公式：收入 = Sales * Quantity * (1 - Discount)
data['Revenue'] = data['Sales'] * data['Quantity'] * (1 - data['Discount'])

# 查看更新后的数据
print("\n预处理后的数据样本：")
print(data.head())
