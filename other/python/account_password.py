ACCOUNT = 'bmwz110'
PASSWORD = 'hello123'

print('please input account:')
user_account = input()

print('please input password:')
user_password = input()

if user_account == ACCOUNT and user_password == PASSWORD :
  print('successful')
else :
  print('fail')