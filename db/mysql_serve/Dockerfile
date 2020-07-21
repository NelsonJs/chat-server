FROM mysql:8.0

#将所需文件放到容器中
RUN mkdir mysql
COPY setup.sh /mysql/setup.sh
COPY schema.sql /mysql/schema.sql

#设置容器启动时执行的命令
CMD ["sh", "/mysql/setup.sh"]