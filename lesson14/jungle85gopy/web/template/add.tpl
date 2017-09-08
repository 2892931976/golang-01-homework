<html>
<body>
    <form method="post" action="create">
        <p>用户名:<input type="text" name="name" /></p>
        <p>密码:<input type="password" name="password" /></p>
        <p>注释:<input type="text" name="note" /></p>
        <p>管理员:<input type="bool" name="isadmin" /></p>
        <p><input type="submit"/> </p>
    </form>
    
    {{ if . }}
    <p style="color:red">{{.}}</p>
    {{end}}
</body>
</html>
