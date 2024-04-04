FROM python:3

WORKDIR /usr/src/twin-te/parser

COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

COPY . .
