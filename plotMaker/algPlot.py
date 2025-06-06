import pandas as pd
import matplotlib.pyplot as plt

# مسیر فایل CSV (یک فایل واحد)
file = '../results.csv'  # مسیر فایل را بر اساس مکان واقعی تغییر بده

# خواندن فایل
df = pd.read_csv(file)

# رسم نمودار
plt.figure(figsize=(10, 6))
plt.plot(df['n'], df['cross_time_s'], label='Classic Algorithm', color='blue', linestyle='-', marker='o')
plt.plot(df['n'], df['strassen_time_s'], label='Strassen Algorithm', color='red', linestyle='--', marker='x')

plt.xlabel('Matrix Size (n)')
plt.ylabel('Time (seconds)')
plt.title('Comparison of Classic and Strassen Matrix Multiplication')
plt.legend()
plt.grid(True)
plt.tight_layout()
plt.show()

