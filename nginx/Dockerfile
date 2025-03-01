FROM nginx:latest

COPY ./nginx/default.conf /etc/nginx/conf.d/

COPY ./nginx/start.sh /
RUN chmod +x /start.sh

CMD /start.sh
