class Student():
  sum = 0

  # 构造函数
  def __init__(self, name, age):
    self.name = name
    self.age = age
    self.__class__.sum += 1
    print('当前班级学生总数为：' + str(self.__class__.sum))

  # 实例方法
  def do_homework(self):
    print('homework')

  # 类方法
  @classmethod
  def plus_sum(cls):
    print(cls.sum)

  # 静态方法
  @staticmethod
  def add(x, y):
    print(Student.sum)
    print('This is static method')

student1 = Student('AAA', 18)
student2 = Student('BBB', 19)
student3 = Student('CCC', 20)
print(student1.name)
print(Student.sum) 