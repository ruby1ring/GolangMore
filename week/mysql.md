docker 使用 mysql教程


建立mysql容器

```
docker run -itd --name mysql-test -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql
```

进入mysql容器

``docker exec -it 84d bash``
