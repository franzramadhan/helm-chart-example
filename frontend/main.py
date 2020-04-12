import os, requests
from flask import Flask

URL = os.getenv("API_URL", "http://localhost:8888")
PORT = 8080

app = Flask(__name__)
@app.route('/')
def serve():
    r = requests.get(url = URL)
    resp = r.json()
    text = resp["text"]
    return "<h5 style='font-size:20px;color:blue;'>Today's trivia: "+ text +"</h5>"

if __name__ == '__main__':
    app.run(host="0.0.0.0", port=PORT)
