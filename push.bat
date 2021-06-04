@echo off
git add .
echo "请添加更新描述"
set /p msg="123" 
git commit -m %msg%
echo "正在提交"
git push