FROM python:3.12-slim

RUN apt-get update && \
    apt-get upgrade -y

RUN pip install --upgrade pip

WORKDIR /code

COPY ./README.md /code/README.md
COPY ./requirements.txt /code/requirements.txt
RUN pip install -r requirements.txt

COPY ./app /code/app

EXPOSE 8000
CMD fastapi dev --host 0.0.0.0