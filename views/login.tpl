<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>

<body>
  <header>
    <h1 class="logo">SSO LOGIN SERVER</h1>
  </header>
  <div>
    <form  method="post">
        <label for="username" >用户名</label>
        <input type="text" name="username"/>
        <label for="password" >密码</label>
        <input type="text" name="password" />
        <input type="submit" value="登陆" />
        <input type="reset" value="重置" />
    </form>
  </div>
  <div>{{.msg}}</div>
</body>
</html>
