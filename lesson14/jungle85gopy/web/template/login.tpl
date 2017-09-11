<html>
<body>
    <form method="post" action="checkLogin">
        <p>username:<input type="text" name="user" /></p>
        <p>password:<input type="password" name="password" /></p>
        <p><input type="submit"/> </p>
    </form>
    
    {{ if . }}
    <p style="color:red">{{.}}</p>
    {{end}}
</body>
</html>