<html>
<body>
    <form method="post" action="modify">
    {{ if .}}
        <p>ID: <input type="text"      value={{.ID}} name="id" readonly="true"/></p>
        <p>用户名: <input type="text"      value={{.Name}} name="name" /></p>
        <p>密码:   <input type="password" value={{.Password}} name="password" /></p>
        <p>注释:   <input type="text"     value={{.Note}} name="note" /></p>
        <p>管理员: <input type="bool"      value={{.Isadmin}} name="isadmin" /></p>
        <p><input type="submit"/> </p>
    {{end}}
    </form>
</body>
</html>
