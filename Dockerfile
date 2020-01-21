FROM debian:stable-slim
ADD vegeta-x86 /usr/bin/vegeta
ADD run.sh /usr/bin/run.sh

CMD ["run.sh"]
