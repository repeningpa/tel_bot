FROM alpine
RUN apk add --update --no-cache build-base python3-dev python3 libffi-dev libressl-dev bash git gettext curl \
 && curl -O https://bootstrap.pypa.io/get-pip.py \
 && python3 get-pip.py \
 && pip install --upgrade six awscli awsebcli
RUN apk --purge -v del py-pip
CMD tail -f /dev/null
