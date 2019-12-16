# -*- coding: utf-8 -*-
import os
from flask import Flask, jsonify
app = Flask(__name__)


@app.route("/")
def index():
    TOKEN = os.getenv('TOKEN', '')
    LANGUAGE = os.getenv('LANGUAGE', '')
    return jsonify(token=TOKEN, lang=LANGUAGE)


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)