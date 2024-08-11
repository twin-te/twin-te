FROM python:3.12

RUN git clone --depth 1 -b feature/staging https://github.com/twin-te/twinte-back /twinte-back
RUN bash -c 'echo -e "\nopenpyxl" >> /twinte-back/twinte-parser/requirements.txt'
RUN pip install -r /twinte-back/twinte-parser/requirements.txt
