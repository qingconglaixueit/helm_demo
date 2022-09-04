FROM centos
ADD my_demo_svr /
ADD start.sh /
ENTRYPOINT ["sh", "start.sh"]
