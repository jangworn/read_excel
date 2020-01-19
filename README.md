php项目需要读取上传的xls/xlsx文件入库，所以开发了这个脚本
```  
$cmd = "./read_excel  --file=xxx.xls";
$xlsx = `$cmd`;
if (strstr($xlsx, 'not a valid zip file') || strstr($xlsx, 'not an excel file')) {
    return ['error' => '请另存为xlsx或xls文件再上传'];
}
$xlsx = json_decode(str_replace('result =  ', '', $xlsx), 1);
```

mac上打包linux运行文件  
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o read_excel main.go