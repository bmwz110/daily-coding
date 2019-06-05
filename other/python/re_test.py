import re

# abc*的意思是：
# 先在字符串s中找到ab，然后看ab后面有没有c，c可以有0次或者不限多次，直到遇见新的ab或其他字符。
s = 'aabccbcc'
r1 = re.findall('abc*', s)
print(r1)

# abc+同上，只是c可以有1次或者不限多次
r2 = re.findall('abc+', s)
print(r2)

# abc？同上，只是c可以有0次或者1次
r3 = re.findall('abc?', s)
print(r3)



# 利用函数替换所匹配的字符
s1 = 'aaaaC#bbbbC#'

def convert(value):
  matched = value.group()
  return '!!' + matched + '!!'

r = re.sub('C#', convert, s1)
print(r)
