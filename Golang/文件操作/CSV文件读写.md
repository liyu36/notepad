## CSV 文件操作
### 读取CSV文件
``` golang
func Reader(filename string) (err error) {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	f := csv.NewReader(file)
	for {
		row, err := f.Read()
		if err != nil && err != io.EOF {
			return err
		} else if err == io.EOF {
			break
		}
		fmt.Println(row) // []string{}
	}
	return nil
}

func main() {
	if err := Reader(os.Args[1]); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
```

### 写入CSV文件
``` golang
func Writer(filename string) (err error) {
	row := []string{"marshmallow", "liy"}
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	f := csv.NewWriter(file)
	if err = f.Write(row); err != nil {
		return err
	}
	f.Flush()
	return nil
}
func main() {
	if err := Writer(os.Args[1]); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

```