buf := make([]byte, 4096)
for {
	n , err := f.Read(buf)
	if err == io.EOF {
		break 
	}
	conn.Write(buf[:n])
}

//-------------------------------
os.MkdirAll(filepath.Dir(), 0755)
f  , err := os.Create(name)
if err != nil {
	log.Print(err)
	return
}

io.Copy(f, r)
f.Close()
//---------------------------------
