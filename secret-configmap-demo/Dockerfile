FROM python:3.6.4

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app
RUN pip install flask
ADD . /usr/src/app

ENV TOKEN abcdefg0000
ENV LANGUAGE  English

EXPOSE 5000
CMD python read-env-app.py