name="filestore"

echo "开始编译linux_amd64..."
GOOS=linux GOARCH=amd64 go build -v -ldflags="-w -s" -o ./bin/linux_amd64/$name
echo "开始压缩..."
upx -9 -k "./bin/linux_amd64/$name"
if [ -f "./bin/linux_amd64/$name.~" ]; then
  rm "./bin/linux_amd64/$name.~"
fi
if [ -f "./bin/linux_amd64/$name.000" ]; then
  rm "./bin/linux_amd64/$name.000"
fi

echo "开始编译linux_arm7..."
GOOS=linux GOARCH=arm GOARM=7 go build -v -ldflags="-w -s" -o ./bin/linux_arm7/$name
echo "开始压缩..."
upx -9 -k "./bin/linux_arm7/$name"
if [ -f "./bin/linux_arm7/$name.~" ]; then
  rm "./bin/linux_arm7/$name.~"
fi
if [ -f "./bin/linux_arm7/$name.000" ]; then
  rm "./bin/linux_arm7/$name.000"
fi

echo "开始编译windows_amd64..."
GOOS=windows go build -v -ldflags="-w -s" -o ./bin/windows_amd64/$name.exe
echo "开始压缩..."
upx -9 -k "./bin/windows_amd64/$name.exe"
if [ -f "./bin/windows_amd64/$name.ex~" ]; then
  rm "./bin/windows_amd64/$name.ex~"
fi
if [ -f "./bin/windows_amd64/$name.000" ]; then
  rm "./bin/windows_amd64/$name.000"
fi

echo "编译完成,等待结束"
sleep 8