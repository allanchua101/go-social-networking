FROM rabbitmq:3.6-management

ADD boot.sh ./bootmq.sh
EXPOSE 15672
RUN ["chmod", "+x", "/bootmq.sh"]
# Define default command
CMD ["/bootmq.sh"]