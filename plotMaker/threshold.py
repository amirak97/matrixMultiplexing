import pandas as pd
import matplotlib.pyplot as plt

# مسیر فایل‌ها
files = ['../t1.csv', '../t8.csv', '../t64.csv', '../t128.csv', '../t512.csv', '../t256.csv']

# لیبل‌های دلخواه برای هر فایل (مثلاً threshold=1 تا 5)
labels = ['threshold = 1', 'threshold = 8','threshold = 64', 'threshold = 128', 'threshold = 256', 'threshold = 512']

colors = ['blue', 'yellow', 'red', 'green', 'orange', 'purple']

plt.figure(figsize=(12, 7))

for i, file in enumerate(files):
    df = pd.read_csv(file)

    plt.plot(df['n'], df['cross_time_s'], label=f'{labels[i]} Cross Time',
             color=colors[i], linestyle='-', marker='o')
    plt.plot(df['n'], df['strassen_time_s'], label=f'{labels[i]} Strassen Time',
             color=colors[i], linestyle='--', marker='x')

plt.xlabel('n')
plt.ylabel('Time (seconds)')
plt.title('Cross Time and Strassen Time vs n from multiple CSV files')
plt.legend()
plt.grid(True)
plt.show()

