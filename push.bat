@echo off
git add .
echo "����Ӹ�������"
set /p msg="123" 
git commit -m %msg%
echo "�����ύ"
git push