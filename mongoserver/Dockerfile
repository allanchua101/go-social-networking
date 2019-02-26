FROM mongo:3.4.19-jessie

ADD boot.sh ./boot-projections.sh
RUN ["chmod", "+x", "/boot-projections.sh"]
CMD ["/boot-projections.sh"]