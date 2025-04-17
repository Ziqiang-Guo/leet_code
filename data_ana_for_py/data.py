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

# # 按 'Region' 汇总总销售收入
# region_sales = data.groupby("Region")['Sales'].sum().reset_index()
# print("\n各区域总销售收入：")
# print(region_sales)

# # 可视化：绘制各区域销售收入的条形图
# plt.figure()
# plt.rcParams['font.family'] = ['Arial Unicode MS']  # 设置字体为支持中文的字体
# sns.barplot(x="Region", y="Sales", data=region_sales, hue="Region", palette="muted", legend=False)
# plt.title("各区域总销售收入")
# plt.xlabel("区域")
# plt.ylabel("销售数量")
# plt.show()

# # 设置 Date 为索引便于重采样
# data.set_index('Date', inplace=True)

# # 按月重采样，并统计当月的总销售收入
# monthly_sales = data['Revenue'].resample('M').sum()

# # 绘制时间序列趋势图
# plt.rcParams['font.family'] = ['Arial Unicode MS']  # 设置字体为支持中文的字体
# plt.figure()
# monthly_sales.plot(marker='o')
# plt.title("按月汇总的销售收入趋势")
# plt.xlabel("时间")
# plt.ylabel("销售收入")
# plt.grid(True)
# plt.show()

# 如有需要，还可以查看不同区域或不同产品随时间的趋势，使用 pivot_table 或 groupby 进一步拆分数据

# 按产品汇总销售收入
product_sales = data.groupby("Segment")['Revenue'].sum().reset_index()
product_sales = product_sales.sort_values(by='Revenue', ascending=False)
print("\n不同产品的销售收入排序：")
print(product_sales)

# 设置字体以支持中文显示
plt.rcParams['font.family'] = ['Arial Unicode MS']  # macOS 上常用的中文字体

# 绘制水平条形图展示各产品销售收入
plt.figure()
sns.barplot(x="Segment", y="Revenue", data=product_sales, palette="deep")
plt.title("不同产品销售收入排名")
plt.xlabel("销售收入")
plt.ylabel("产品")
plt.show()
