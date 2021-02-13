# Meetrecord

- ## login
- controller

1. 获取用户密码
2. 根据用户、密码查询db，如果存在则设置session，并回报登陆成功，否则回报登陆失败

- ## register
- controller

1. 获取用户、密码（以邮箱作为用户名，repassword在前端做检验）
2. 判断该usrname是否存在
3. 若不存在，将usrname，MD5（passwd）status写入数据库，并返回注册成功
4. 若存在，则返回用户已被注册

- ### base-controller
- 用于判断用户登陆状态

1. 尝试获取session
2. 若session为nil，则返回false
3. 若session不为nil，则返回true