FROM python:alpine
WORKDIR /opt
ADD . /opt
RUN cd /opt && pip3 install --no-cache-dir --upgrade pip && pip3 install --no-cache-dir -r requirements.txt && \
    addgroup runner && \
    adduser \
    --disabled-password \
    --gecos "" \
    --home "$(pwd)" \
    --ingroup "runner" \
    --no-create-home \
    "runner"
USER runner
EXPOSE 8080
ENTRYPOINT [ "python" ]
CMD ["main.py"]
